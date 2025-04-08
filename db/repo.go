package db

import (
	. "LesyaBack/poll/model"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
)

var col *mongo.Collection

func init() {
	c, err := mongo.Connect(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	db := c.Database("LETI")
	if db.Collection("Polls") == nil {
		db.CreateCollection(context.TODO(), "Polls")
	}
	col = db.Collection("Polls")

}

func GetAllPolls() ([]Poll, error) {
	c, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var ps []Poll
	err = c.All(context.TODO(), &ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func CreateNewPoll(c *CreatePollRequest) (*string, error) {
	p := Poll{
		Link:      uuid.New().String(),
		Course:    c.Course,
		Group:     c.Group,
		Professor: c.Professor,
		Questions: []Question{
			{
				Text:     "Оцените качество программы",
				Type:     Grade,
				MaxGrade: 10,
			},
			{
				Text:     "Оцените сложность обучения",
				Type:     Grade,
				MaxGrade: 10,
			},
			{
				Text:     "Оцените уровень материала",
				Type:     Grade,
				MaxGrade: 10,
			},
		},
	}

	_, err := col.InsertOne(context.TODO(), p)
	if err != nil {
		return nil, err
	}
	l := p.Link
	return &l, nil
}

func GetPoll(l string) (*Poll, error) {
	var p Poll
	res := col.FindOne(context.TODO(), &bson.M{
		"link": l,
	})
	if err := res.Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func UpdatePoll(l string, new Poll) (*Poll, error) {

	_, err := col.UpdateOne(context.TODO(), &bson.M{
		"link": l,
	}, new)

	if err != nil {
		return nil, err
	}
	upd, err := GetPoll(l)

	if err != nil {
		return nil, err
	}

	return upd, nil
}

func DeletePoll(l string) error {
	var p Poll
	err := col.FindOne(context.TODO(), &bson.M{
		"link": l,
	}).Decode(&p)

	if err == mongo.ErrNoDocuments {
		return errors.New("Not found")
	}
	_, err = col.DeleteOne(context.TODO(), &bson.M{
		"link": l,
	})

	return err
}
