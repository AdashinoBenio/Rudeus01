// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoMorse"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoNekosLife"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoPat"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func toMorseHandler(message *tg.Message, args pTools.Arg) {
	wotoMorse.ToMorseHandler(message, args)
}

func fromMorseHandler(message *tg.Message, args pTools.Arg) {
	wotoMorse.FromMorseHandler(message, args)
}

func patHandler(message *tg.Message, args pTools.Arg) {
	wotoPat.PatHandler(message, args)
}

func tickleNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.TickleHandler(message, args)
}

func slapNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.SlapHandler(message, args)
}

func pokeNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.PokeHandler(message, args)
}

func nekoNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.NekoHandler(message, args)
}

func meowNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.MeowHandler(message, args)
}

func lizardNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.LizardHandler(message, args)
}

func kissNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.KissHandler(message, args)
}

func hugNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.HugHandler(message, args)
}

func foxNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.FoxHandler(message, args)
}

func feedNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.FeedHandler(message, args)
}

func cuddleNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNekosLife.CuddleHandler(message, args)
}
