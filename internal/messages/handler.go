package messages

import (
	"encoding/json"
	"fmt"
	"github.com/evassilyev/chgk-telebot-2/internal/core"
	"github.com/evassilyev/chgk-telebot-2/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Handler struct {
	inputMessages   chan *tgbotapi.Message
	callbackQueries chan *tgbotapi.CallbackQuery
	/*
		messageHandlers  map[MessageAction]func(message *tgbotapi.Message) error
		callbackHandlers map[ButtonAction]callbackHandlerFunc

		expectedAnswers expectedAnswersStorage
	*/

	db         core.StorageHandler
	handlerUrl string
	bot        *telegram.Tbot
}

func (mh *Handler) GetHandlerUrl() string {
	return mh.handlerUrl
}

func NewHandler(db core.StorageHandler, token, domain, url, cert string) (*Handler, error) {
	// generate webhook URL
	uuidToken := uuid.New().String()
	handlerUrl := fmt.Sprintf("%s/%s", url, uuidToken)
	bot, err := telegram.NewBot(token, cert, domain+handlerUrl)
	if err != nil {
		return nil, err
	}

	var m Handler
	m.db = db
	m.bot = bot
	m.handlerUrl = handlerUrl
	m.inputMessages = make(chan *tgbotapi.Message)
	m.callbackQueries = make(chan *tgbotapi.CallbackQuery)

	/*
		m.expectedAnswers = NewExpectedAnswersStorage(m.cacheExpirationCallback, time.Minute*1)

		err = m.db.ClearTemporaryEvents()
		if err != nil {
			return nil, err
		}

		m.initHandlers()

	*/
	go m.handleMessages()
	go m.handleCallbacks()

	return &m, nil
}

func (mh *Handler) handleCallbacks() {
	for callback := range mh.callbackQueries {
		if callback == nil {
			continue
		}
		if callback.Data == "" {
			fmt.Println("get empty callback data")
		}
		/*

			var (
				handler callbackHandlerFunc
				exists  bool
			)

			callbackExpectedAnswer := mh.expectedAnswers.Get(callbackId(callback))
			// checking in callback.Data = 0:TOPIC_YES || 0:TOPIC_NO
			if strings.HasSuffix(callback.Data, string(chooseTopicYes)) {
				handler, exists = mh.callbackHandlers[chooseTopicYes]
			} else if strings.HasSuffix(callback.Data, string(chooseTopicNo)) {
				handler, exists = mh.callbackHandlers[chooseTopicNo]
			} else if strings.HasSuffix(callback.Data, string(topicEdited)) {
				handler, exists = mh.callbackHandlers[topicEdited]
			} else if strings.HasSuffix(callback.Data, string(newEventTopic)) && callbackExpectedAnswer == AskEventTopic {
				handler, exists = mh.callbackHandlers[newEventTopic]
			} else if callback.Data == string(setEventComplexityEasy) || callback.Data == string(setEventComplexityMid) || callback.Data == string(setEventComplexityHard) {
				if callbackExpectedAnswer == AskEventComplexity {
					handler, exists = mh.callbackHandlers[ButtonAction(callback.Data)]
				} else {
					return
				}
			} else if strings.HasSuffix(callback.Data, string(showEventDetails)) {
				handler, exists = mh.callbackHandlers[showEventDetails]
			} else {
				handler, exists = mh.callbackHandlers[ButtonAction(callback.Data)]
			}
			if exists {
				err := handler(callback)
				if err != nil {
					fmt.Printf("error in callback %s handler: %v\n", callback.Data, err)
				}
			} else {

		*/
		fmt.Printf("unknown callback %s\n", callback.Data)
		/*
			}

		*/
	}
}

func (mh *Handler) handleMessages() {
	for message := range mh.inputMessages {
		if message == nil {
			continue
		}

		/*
			var action MessageAction
			action = mh.expectedAnswers.Get(messageId(message))
			if action == Undefined {
				action = determine(message)
			}

			handler, exist := mh.messageHandlers[action]
			if exist {
				err := handler(message)
				if err != nil {
					fmt.Printf("error in handler %d: %v\n", action, err)
				}
			} else {

		*/
		fmt.Printf("no handler for the message: %+v", message)
		/*
			}
		*/
	}
}

func (mh *Handler) HandleMessages(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(string(bytes))
	var message tgbotapi.Update
	err = json.Unmarshal(bytes, &message)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	go func() { mh.inputMessages <- message.Message }()
	go func() { mh.callbackQueries <- message.CallbackQuery }()
}
