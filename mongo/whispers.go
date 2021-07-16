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

package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
)

type Whisper struct {
	Id       string `bson:"id"`
	Sender   int64  `bson:"sender"`
	Receiver string `bson:"receiver"`
	Text     string `bson:"text"`
}

func getCollection() *mongo.Collection {
	if collection != nil {
		return collection
	}

	collection = GetDatabase().Collection("whispers")
	return collection
}

func SaveWhisper(whisper Whisper) error {
	_, err := getCollection().InsertOne(Ctx, whisper)
	return err
}

func GetWhisper(id string) (Whisper, error) {
	result := Whisper{}
	err := getCollection().FindOne(Ctx, bson.M{"id": id}).Decode(&result)
	return result, err
}

func GetWhispersCount(sender int64) (int64, error) {
	count, err := getCollection().CountDocuments(Ctx, bson.M{"sender": sender})
	return count, err
}

func DeleteWhisper(id string) error {
	_, err := getCollection().DeleteOne(Ctx, bson.M{"id": id})
	return err
}

func DeleteWhispers(sender int64) (int64, error) {
	result, err := getCollection().DeleteMany(Ctx, bson.M{"sender": sender})
	return result.DeletedCount, err
}
