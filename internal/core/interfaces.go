package core

import "github.com/evassilyev/chgk-telebot-2/internal/core/gen/models"

type StorageHandler interface {
	AddGroup(chatId int64) error
	GetGroup(groupId int64) (*models.Group, error)
}
