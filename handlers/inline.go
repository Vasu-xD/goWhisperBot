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
	"strings"

	uuid "github.com/satori/go.uuid"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func inline(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.InlineQuery.Query
	if query == "" || len(query) > 200 || len(strings.Fields(query)) < 2 {
		ctx.InlineQuery.Answer(
			b,
			[]gotgbot.InlineQueryResult{
				gotgbot.InlineQueryResultArticle{
					Id:    uuid.NewV4().String(),
					Title: "🔥 Write a whisper message",
					InputMessageContent: gotgbot.InputTextMessageContent{
						MessageText: fmt.Sprintf("<b>Send whisper messages through inline mode</b>\n\nUsage: <code>@%s [@username] text</code>", b.User.Username),
						ParseMode:   "HTML",
					},
					Description: fmt.Sprintf("Usage: @%s [@username] text", b.User.Username),
					ThumbUrl:    "https://www.freeiconspng.com/uploads/whisper-icon-0.png",
					ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
						InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
							{
								gotgbot.InlineKeyboardButton{
									Text: "Learn more...",
									Url:  fmt.Sprintf("https://t.me/%s?start=learn", b.User.Username),
								},
							},
						},
					},
				},
			},
			&gotgbot.AnswerInlineQueryOpts{
				SwitchPmText:      "ℹ️ Learn how to send whispers",
				SwitchPmParameter: "learn",
				CacheTime:         0,
				IsPersonal:        true,
			},
		)
		return nil
	}
	_username := strings.Fields(query)[0]
	username := strings.TrimPrefix(_username, "@")
	text := strings.Trim(query[len(_username)+1:], " ")
	if username == "all" {
		ctx.InlineQuery.Answer(
			b,
			[]gotgbot.InlineQueryResult{
				gotgbot.InlineQueryResultArticle{
					Id:    uuid.NewV4().String(),
					Title: "👁️ Whisper once to the first one who open it",
					InputMessageContent: gotgbot.InputTextMessageContent{
						MessageText: "👁️ The first one who open the whisper can read it",
					},
					Description: "🤫 " + text,
					ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
						InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
							{
								gotgbot.InlineKeyboardButton{
									Text:         "👁️ show message",
									CallbackData: "show_whisper",
								},
							},
						},
					},
				},
			},
			&gotgbot.AnswerInlineQueryOpts{CacheTime: 0, IsPersonal: true},
		)
	} else {
		ctx.InlineQuery.Answer(
			b,
			[]gotgbot.InlineQueryResult{
				gotgbot.InlineQueryResultArticle{
					Id:    uuid.NewV4().String(),
					Title: fmt.Sprintf("🔒 A whisper message to @%s", username),
					InputMessageContent: gotgbot.InputTextMessageContent{
						MessageText: fmt.Sprintf("🔒 A whisper message to @%s", username),
					},
					Description: text,
					ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
						InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
							{
								gotgbot.InlineKeyboardButton{
									Text:         "👁️ show message",
									CallbackData: "show_whisper",
								},
							},
						},
					},
				},
			},
			&gotgbot.AnswerInlineQueryOpts{CacheTime: 0, IsPersonal: true},
		)
	}
	return nil
}
