/**
 * ezWhisperBot - A Telegram bot for sending whisper messages
 * Copyright (C) 2021  Roj Serbest
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

/*
 * A simple custom Chosen Inline Result Handler,
 * as gotgbot isn't currently (May 23, 2021) supporting it.
 */

package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Response func(b *gotgbot.Bot, ctx *ext.Context) error

type ChosenInlineResult struct {
	Response Response
}

func NewChosenInlineResult(r Response) ChosenInlineResult {
	return ChosenInlineResult{
		Response: r,
	}
}

func (i ChosenInlineResult) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	return i.Response(b, ctx)
}

func (i ChosenInlineResult) CheckUpdate(b *gotgbot.Bot, u *gotgbot.Update) bool {
	return u.ChosenInlineResult != nil
}

func (i ChosenInlineResult) Name() string {
	return fmt.Sprintf("choseninlineresult_%p", i.Response)
}
