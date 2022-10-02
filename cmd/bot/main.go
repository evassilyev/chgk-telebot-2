package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/evassilyev/chgk-telebot-2/internal/core"
	"github.com/evassilyev/chgk-telebot-2/internal/messages"
	storage2 "github.com/evassilyev/chgk-telebot-2/internal/storage"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var c core.Configuration
	err := envconfig.Process("chgkbot", &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Configuation:\n%+v\n", c)

	db, err := storage2.NewPostgresStorage(c.DbUser, c.DbPass, c.DbHost, c.DbName, c.DbPort)
	if err != nil {
		fmt.Printf("can not init storage: %s", err.Error())
		return
	}

	mh, err := messages.NewHandler(db, c.BotToken, c.Domain, c.Url, c.CertFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Handler url: ", mh.GetHandlerUrl())

	r := mux.NewRouter()
	r.HandleFunc(mh.GetHandlerUrl(), mh.HandleMessages).Methods(http.MethodPost)
	r.HandleFunc("/healthcheck", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
	http.Handle("/", r)
	fmt.Println("started on port", c.Port)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", c.Port), c.CertFile, c.KeyFile, nil))
}
