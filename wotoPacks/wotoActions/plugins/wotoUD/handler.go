package wotoUD

import (
	"log"
	"reflect"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	wq "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/wotoQuery"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UdHandler(message *tg.Message, args pTools.Arg) {
	if message == nil {
		return
	}

	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	pv := args.HasFlag(PRIVATE_FLAG, PV_FLAG)
	//hasMsg := args.HasFlag(MSG_FLAG, MESSAGE_FLAG)
	reply := message.ReplyToMessage != nil
	//del := reply && args.HasFlag(DEL_FLAG, DELETE_FLAG)
	//single := args.HasFlag(SINGLE_FLAG)
	text := args.JoinNoneFlags(false)
	if wotoStrings.IsEmpty(&text) {
		return
	}

	finalText, ud := defineText(text, wv.BaseIndex)

	var msg tg.MessageConfig

	if pv {
		msg = tg.NewMessage(message.From.ID, finalText)
	} else {
		msg = tg.NewMessage(message.Chat.ID, finalText)
	}

	// !pv => for fixing: Bad Request: replied message not found
	if reply && !pv {
		r := message.ReplyToMessage
		if r != nil {
			msg.ReplyToMessageID = r.MessageID
		} else {
			msg.ReplyToMessageID = message.MessageID
		}
	} else {
		// for fixing: Bad Request: replied message not found
		if !pv {
			msg.ReplyToMessageID = message.MessageID
		}
	}
	msg.ReplyMarkup = getButton(text, wv.BaseOneIndex, message.MessageID, ud)

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}

}

func defineText(word string, page int) (str string, c *UrbanCollection) {
	ud, err := Define(word)
	if err != nil {
		log.Println(err)
		return wv.EMPTY, nil
	}

	if len(ud.List) == wv.BaseIndex {
		// TODO: return not found or something.
		return "not found", ud
	}

	//text = wotoMD.GetNormal(ud.List[0].Def).ToString()
	normal := wotoMD.GetNormal(ud.List[page].Def)
	ex := wotoMD.GetItalic(wv.N_ESCAPE + wv.N_ESCAPE + ud.List[0].Example)
	return normal.Append(ex).ToString(), ud
}

func getButton(word string, page uint8, id int,
	c *UrbanCollection) *tg.InlineKeyboardMarkup {
	next := getNextUdQuery(word, page, id)
	pre := getPreviousUdQuery(word, page, id)
	next.setData(c)
	pre.setData(c)
	nextID := wq.SetInMap(next)
	preID := wq.SetInMap(pre)
	q1 := wq.GetQuery(wq.UdQueryPlugin, nextID)
	q2 := wq.GetQuery(wq.UdQueryPlugin, preID)
	b01 := tg.NewInlineKeyboardButtonData(preText, q1.ToString())
	b02 := tg.NewInlineKeyboardButtonData(nextText, q2.ToString())
	buttons := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			b01,
			b02,
		),
	)

	return &buttons
}

func QUdHanler(query *tg.CallbackQuery, q *wq.QueryBase) {
	data := q.GetDataQuery()
	if data == nil {
		return
	}

	udData := &udQuery{}
	if reflect.TypeOf(udData) != reflect.TypeOf(data) {
		return
	}

	udData = data.(*udQuery)

	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	if udData.next {
		udData.currentPage += wv.BaseOneIndex
	} else if udData.previous {
		udData.currentPage -= wv.BaseOneIndex
	}

	text := udData.collection.List[udData.currentPage].Def
	chat := query.Message.Chat.ID
	msgId := query.Message.MessageID

	eConfig := tg.NewEditMessageText(chat, msgId, text)
	eConfig.ParseMode = tg.ModeMarkdownV2

	api.Request(eConfig)
	wq.RemoveFromMap(q.Data)
}
