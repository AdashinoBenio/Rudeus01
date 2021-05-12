package wotoUD

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (q *udQuery) setData(collection *UrbanCollection) {
	q.collection = collection
}

func (q *udQuery) setButtons(b *tgbotapi.InlineKeyboardMarkup) {
	q.keyboard = b
}
