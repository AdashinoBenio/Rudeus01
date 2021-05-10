// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import wn "github.com/ALiwoto/rudeus01/wotoPacks/wotoNeko"

// non-sudo commands
const (
	TMorseCmd    = "tmorse"    // wotoMorse plugin
	FMorseCmd    = "fmorse"    // wotoMorse plugin
	TrCmd        = "tr"        // wotoTranslate plugin
	TlCmd        = "tl"        // wotoTranslate plugin
	TranslateCmd = "translate" // wotoTranslate plugin
	PatCmd       = "pat"       // wotoPat plugin
	PattuCmd     = "pattu"     // wotoPat plugin
	FoxCmd       = "fox"       // wotoNekosLife plugin
	FoxyCmd      = "foxy"      // wotoNekosLife plugin
)

var cmdList map[string]CmdHandler

func cmdListInit() {
	if cmdList != nil {
		return
	}

	cmdList = make(map[string]CmdHandler)

	addMorseCmdList()
	addTrCmdList()
	addPatCmdList()
	addNekosLifeCmdList()
}

func addMorseCmdList() {
	if cmdList != nil {
		if cmdList[TMorseCmd] == nil {
			cmdList[TMorseCmd] = toMorseHandler
		}
		if cmdList[FMorseCmd] == nil {
			cmdList[FMorseCmd] = fromMorseHandler
		}
	}
}

func addTrCmdList() {
	if cmdList != nil {
		if cmdList[TrCmd] == nil {
			cmdList[TrCmd] = fromMorseHandler
		}
		if cmdList[TranslateCmd] == nil {
			cmdList[TranslateCmd] = fromMorseHandler
		}
	}
}

func addPatCmdList() {
	if cmdList != nil {
		if cmdList[PatCmd] == nil {
			cmdList[PatCmd] = patHandler
		}
		if cmdList[PattuCmd] == nil {
			cmdList[PattuCmd] = patHandler
		}
	}
}

func addNekosLifeCmdList() {
	//first: sfw
	if cmdList != nil {
		if cmdList[string(wn.Tickle)] == nil {
			cmdList[string(wn.Tickle)] = tickleNekoHandler
		}
		if cmdList[string(wn.Slap)] == nil {
			cmdList[string(wn.Slap)] = slapNekoHandler
		}
		if cmdList[string(wn.Poke)] == nil {
			cmdList[string(wn.Poke)] = pokeNekoHandler
		}
		// don't set pat here, we have wotoPat for handling this.
		//if cmdList[string(wn.Pat)] == nil {
		//	cmdList[string(wn.Pat)] = patHandler
		//}
		if cmdList[string(wn.Neko)] == nil {
			cmdList[string(wn.Neko)] = nekoNekoHandler
		}
		if cmdList[string(wn.Meow)] == nil {
			cmdList[string(wn.Meow)] = meowNekoHandler
		}
		if cmdList[string(wn.Lizard)] == nil {
			cmdList[string(wn.Lizard)] = lizardNekoHandler
		}
		if cmdList[string(wn.Kiss)] == nil {
			cmdList[string(wn.Kiss)] = kissNekoHandler
		}
		if cmdList[string(wn.Hug)] == nil {
			cmdList[string(wn.Hug)] = hugNekoHandler
		}
		if cmdList[string(wn.Fox_Girl)] == nil {
			cmdList[string(wn.Fox_Girl)] = foxNekoHandler
		}
		if cmdList[FoxCmd] == nil {
			cmdList[FoxCmd] = foxNekoHandler
		}
		if cmdList[FoxyCmd] == nil {
			cmdList[FoxyCmd] = foxNekoHandler
		}
		if cmdList[string(wn.Feed)] == nil {
			cmdList[string(wn.Feed)] = feedNekoHandler
		}
		if cmdList[string(wn.Cuddle)] == nil {
			cmdList[string(wn.Cuddle)] = cuddleNekoHandler
		}
	}
}
