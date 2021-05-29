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
	"goWhisperBot/mongo"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func chosenInlineResult(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.ChosenInlineResult.Query
	if len(strings.Fields(query)) == 0 || len(query) > 200 {
		return nil
	}
	_username := strings.Fields(query)[0]
	username := strings.TrimPrefix(_username, "@")
	text := strings.Trim(query[len(_username)+1:], " ")
	inlineMessageId := ctx.ChosenInlineResult.InlineMessageId
	mongo.SaveWhisper(mongo.Whisper{
		Id:       inlineMessageId,
		Sender:   ctx.ChosenInlineResult.From.Id,
		Receiver: username,
		Text:     text,
	})
	return nil
}
