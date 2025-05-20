package db

import (
	"LesyaBack/poll/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetAnswers(l string, sessionId string) (*[]model.Answer, error) {
	var a model.Answer
	err := PollCol.FindOne(context.TODO(), &bson.M{
		"link": l, "sessionId": sessionId,
	}).Decode(&a)

	if err == mongo.ErrNoDocuments {
		return nil,
			errors.New("Not found")
	}

	return &a, nil
}
