package wotoUD

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type udQuery struct {
	next        bool   // true if it's the next button
	previous    bool   // false if it's the previous button
	currentPage uint8  // must be between 1 and 10.
	word        string // the word used in urban dictionary
	id          int
	collection  *UrbanCollection // the collection
	keyboard    *tgbotapi.InlineKeyboardMarkup
}

func getNextUdQuery(word string, page uint8, id int) *udQuery {
	return getUdQuery(word, true, page, id)
}

func getPreviousUdQuery(word string, page uint8, id int) *udQuery {
	return getUdQuery(word, false, page, id)
}

func getUdQuery(word string, next bool, page uint8, mId int) *udQuery {
	return &udQuery{
		next:        next,
		previous:    !next,
		currentPage: page,
		word:        word,
		id:          mId,
	}
}
