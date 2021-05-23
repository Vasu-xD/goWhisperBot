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

package whispers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Whispers Whispers_

type Whisper struct {
	Sender   int64  `json:"sender_uid"`
	Receiver string `json:"receiver_uname"`
	Text     string `json:"text"`
}

type Whispers_ struct {
	Whispers map[string]Whisper `json:"whispers"`
}

func InitWhispers() {
	file, err := os.Open("whispers.json")

	if err != nil {
		panic(err)
	}

	defer ioutil.WriteFile("whispers.json", []byte("{\"whispers\": {}}"), 0644)
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &Whispers)
	file.Close()
}

func SaveWhispers() {
	bytes, err := json.Marshal(&Whispers)

	if err != nil {
		fmt.Println("Cannot save Whispers:", err.Error())
	}

	ioutil.WriteFile("whispers.json", bytes, 0644)
}
