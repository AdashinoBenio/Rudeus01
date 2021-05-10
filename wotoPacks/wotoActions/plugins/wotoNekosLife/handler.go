package wotoNekosLife

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoNeko"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TickleHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomTickle()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func SlapHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomSlap()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func PokeHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomPoke()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func NekoHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomNeko()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func MeowHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomMeow()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func LizardHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomLizard()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func KissHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomKiss()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func HugHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomHug()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func FoxHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomFoxGirl()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func FeedHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomFeed()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func CuddleHandler(message *tg.Message, args pTools.Arg) {
	base, err := wotoNeko.GetRandomCuddle()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func sendNekoBase(message *tg.Message, neko *wotoNeko.NekoBase, args pTools.Arg) {
	if message == nil || neko == nil {
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
	reply := message.ReplyToMessage != nil
	del := reply && args.HasFlag(DEL_FLAG, DELETE_FLAG)
	var sendIt tg.Chattable

	if neko.IsPhoto() {
		var photo tg.PhotoConfig
		if pv {
			photo = tg.NewPhoto(message.From.ID, tg.FileURL(neko.Url))
			if message.From.ID == message.Chat.ID {
				photo.ReplyToMessageID = message.MessageID
			}
		} else {
			photo = tg.NewPhoto(message.Chat.ID, tg.FileURL(neko.Url))
			if reply && message.ReplyToMessage != nil {
				photo.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				photo.ReplyToMessageID = message.MessageID
			}
		}

		sendIt = photo
	} else if neko.IsGif() {
		var gif tg.DocumentConfig
		if pv {
			gif = tg.NewDocument(message.From.ID, tg.FileURL(neko.Url))
			if message.From.ID == message.Chat.ID {
				gif.ReplyToMessageID = message.MessageID
			}
		} else {
			gif = tg.NewDocument(message.Chat.ID, tg.FileURL(neko.Url))
			if reply && message.ReplyToMessage != nil {
				gif.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				gif.ReplyToMessageID = message.MessageID
			}
		}

		sendIt = gif
	} else {
		log.Println(neko.Url)
		return
	}

	if del {
		req := tg.NewDeleteMessage(message.Chat.ID, message.MessageID)
		// don't check error or response, we have
		// more important things to do
		go settings.GetAPI().Request(req)
	}

	if _, err := api.Send(sendIt); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}
