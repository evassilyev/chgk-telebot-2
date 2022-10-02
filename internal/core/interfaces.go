package core

type StorageHandler interface {
	AddGroup(chatId int64) error
}
