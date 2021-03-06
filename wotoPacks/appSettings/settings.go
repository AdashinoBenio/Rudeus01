// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _alreadyRunning bool
var _settings *AppSettings

type AppSettings struct {
	index      int
	totalIndex int
	isGlobal   bool
	tObt       string
	url        string
	nextUrl    string
	sudoList   []int64
	mainSudo   int64
	wClient    interfaces.WClient
	patClient  interfaces.WClient
	botAPI     *tgbotapi.BotAPI
	_apiObtain func(interfaces.WSettings) bool
	_nextChild func(string, interfaces.WSettings)
}

func GetSettings() interfaces.WSettings {
	if _settings != nil {
		return GetExisting()
	}
	settings := AppSettings{}
	_settings = &settings
	return _settings
}

func GetExisting() interfaces.WSettings {
	return _settings
}

func App_enter() {
	_alreadyRunning = true
}

func App_exit() {
	_alreadyRunning = false
}

func IsRunning() bool {
	return _alreadyRunning
}
