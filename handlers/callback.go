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

package handlers

import (
	"fmt"
	"goWhisperBot/whispers"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func showWhisper(b *gotgbot.Bot, ctx *ext.Context) error {
	inlineMessageId := ctx.CallbackQuery.InlineMessageId

	if whispers.Whispers.Whispers[inlineMessageId].Text == "" {
		ctx.CallbackQuery.Answer(
			b,
			&gotgbot.AnswerCallbackQueryOpts{
				Text:      "Can't find the whisper text",
				ShowAlert: true,
			},
		)
		ctx.EffectiveMessage.EditText(
			b,
			"⛔ invalid whisper",
			&gotgbot.EditMessageTextOpts{},
		)
	} else {
		whisper := whispers.Whispers.Whispers[inlineMessageId]
		sender := whisper.Sender
		receiver := whisper.Receiver
		text := whisper.Text

		if receiver == "all" || strings.EqualFold(ctx.EffectiveUser.Username, receiver) {
			ctx.CallbackQuery.Answer(
				b,
				&gotgbot.AnswerCallbackQueryOpts{
					Text:      text,
					ShowAlert: true,
				})
			ctx.EffectiveMessage.EditText(
				b,
				fmt.Sprintf("🔓 %s read the message", ctx.EffectiveUser.FirstName),
				&gotgbot.EditMessageTextOpts{},
			)
			delete(whispers.Whispers.Whispers, inlineMessageId)
		} else if ctx.EffectiveUser.Id == sender {
			ctx.CallbackQuery.Answer(
				b,
				&gotgbot.AnswerCallbackQueryOpts{
					Text:      text,
					ShowAlert: true,
				})
		} else {
			ctx.CallbackQuery.Answer(
				b,
				&gotgbot.AnswerCallbackQueryOpts{
					Text:      "This is not for you",
					ShowAlert: true,
				})
		}
	}

	return nil
}
