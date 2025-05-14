package db

import (
	"LesyaBack/poll/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var answersCol *mongo.Collection

func CreateNewAnswers(a *model.AnswersForm) error {
	_, err := answersCol.InsertOne(context.TODO(), a)
	if err != nil {
		return err
	}
	return nil
}

func GetAnswersForPoll(l string) ([]model.AnswersForm, error) {
	a := []model.AnswersForm{}

	c, err := answersCol.Find(context.TODO(), &bson.M{
		"link": l,
	})
	if err != nil {
		return nil, err
	}

	err = c.All(context.TODO(), &a)
	if err != nil {
		return nil, err
	}
	return a, nil

}

func GetAllAnswers() ([]model.AnswersForm, error) {
	c, err := answersCol.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var a []model.AnswersForm
	err = c.All(context.TODO(), &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
