package wotoTranslate

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TrHandler(message *tg.Message, args pTools.Arg) {
	// set the first element of args to empty,
	// because pTools.Arg also contains the command itself,
	// however we don't want that here.
	args[wv.BaseIndex] = wv.EMPTY

	send_pv := args.HasFlag(PV_FLAG, PRIVATE_FLAG)
	flags := args.GetFlags()
	is_reply := message.ReplyToMessage != nil
	gnu := args.HasFlag("gnu")

	if len(flags) == wv.BaseIndex {
		return
	}

	toLang := flags[wv.BaseIndex]

	var full string
	if is_reply {
		if !ws.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			if !ws.IsEmpty(&message.ReplyToMessage.Caption) {
				full = message.ReplyToMessage.Caption
			} else {
				// we should check if the replied message has text or not,
				// and since it doesn't have any text value, then we should not
				// do any operations on it.
				return
			}
		}
	} else {
		// do not convert the flags to the morse code.
		full = args.JoinNoneFlags(false)
	}

	l1 := DetectLanguage(full)
	if l1 == nil {
		return
	}

	cl := l1.Data.Detections
	if len(cl) == wv.BaseIndex {
		return
	}

	var str string

	if gnu {
		str = TrGnuTxt(l1.Data.Detections[0].TheLang, toLang, full)
	} else {
		trl := Translate(l1, toLang, full)
		if trl.HasWrongNess {
			str += "Did you mean \"" + trl.CorrectedValue + "\"?\n"
		}
		str += trl.TranslatedText
	}

	sendTr(message, &str, is_reply, send_pv)
}

// please before using this function, ensure that the `morse` value is
// converted to markdown (using wotoMarkdown) functions.
func sendTr(message *tg.Message, morse *string, reply, pv bool) {
	// always check before sending:
	// Bad Request: message text is empty
	if ws.IsEmpty(morse) {
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

	var msg tg.MessageConfig
	if pv {
		//settings.SendSudo(strconv.Itoa(int(message.From.ID)))
		msg = tg.NewMessage(message.From.ID, (*morse))
	} else {
		//settings.SendSudo(strconv.Itoa(int(message.Chat.ID)))
		msg = tg.NewMessage(message.Chat.ID, (*morse))
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

	//msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}
