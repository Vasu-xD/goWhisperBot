/**
 * goWhisperBot - A Telegram bot for sending whisper messages
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
	"goWhisperBot/mongo"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func listWhispers(b *gotgbot.Bot, ctx *ext.Context) error {
	result, err := mongo.GetWhisper(ctx.CallbackQuery.InlineMessageId)

	if err != nil {
		return err
	}

	if result == (mongo.Whisper{}) {
		ctx.CallbackQuery.Answer(
			b,
			&gotgbot.AnswerCallbackQueryOpts{
				Text:      "Can't find the whisper text",
				ShowAlert: true,
			},
		)

		b.EditMessageText(
			"⛔ invalid whisper",
			&gotgbot.EditMessageTextOpts{
				InlineMessageId: ctx.CallbackQuery.InlineMessageId,
			},
		)

		return nil
	}

	sender := result.Sender
	receiver := result.Receiver
	text := result.Text

	if ctx.EffectiveUser.Id == sender {
		ctx.CallbackQuery.Answer(
			b,
			&gotgbot.AnswerCallbackQueryOpts{
				Text:      text,
				ShowAlert: true,
			})

	} else if receiver == "all" || strings.EqualFold(ctx.EffectiveUser.Username, receiver) {
		ctx.CallbackQuery.Answer(
			b,
			&gotgbot.AnswerCallbackQueryOpts{
				Text:      text,
				ShowAlert: true,
			})

		b.EditMessageText(
			fmt.Sprintf("🔓 %s read the message", ctx.EffectiveUser.FirstName),
			&gotgbot.EditMessageTextOpts{
				InlineMessageId: ctx.CallbackQuery.InlineMessageId,
			},
		)

		return mongo.DeleteWhisper(ctx.CallbackQuery.InlineMessageId)
	} else {
		ctx.CallbackQuery.Answer(
			b,
			&gotgbot.AnswerCallbackQueryOpts{
				Text:      "This is not for you",
				ShowAlert: true,
			})
	}

	return err
}
