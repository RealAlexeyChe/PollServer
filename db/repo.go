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

var PollCol *mongo.Collection

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
	PollCol = db.Collection("Polls")

	if db.Collection("Admins") == nil {
		db.CreateCollection(context.TODO(), "Admins")
	}
	adminCol = db.Collection("Admins")

	if db.Collection("Answers") == nil {
		db.CreateCollection(context.TODO(), "Answers")
	}
	answersCol = db.Collection("Answers")

}

func GetAllPolls(sessionId string) ([]Link, error) {

	filter := bson.D{{"sessionid", sessionId}}

	pipeline := mongo.Pipeline{
		{{"$match", filter}},
		{{"$group", bson.D{
			{"_id", "$link"},
		}}},
	}

	c, err := PollCol.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	l := []Link{}

	err = c.All(context.TODO(), &l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func CreateNewPoll(c *CreatePollRequest, sessionId string) (*Poll, error) {
	p := Poll{}

	if c.Template {
		p = Poll{
			SessionId: sessionId,
			Link:      uuid.New().String(),
			Questions: []Question{
				{
					IsRequired: true,
					Text:       "1. Оцените дисциплину предыдущего семестра:",
					Type:       Grid,
					Options: []string{
						"Отлично",
						"Хорошо",
						"Удовл.",
						"Плохо",
						"Очень плохо",
					},
					Rows: []string{
						"Содержание дисциплины (программа)",
						"Соответствие лекций практике",
						"Методическое обеспечение",
						"Материально-техническое обеспечение",
					},
				},
				{
					IsRequired: true,
					Text:       "",
					Type:       Grid,
					Options: []string{
						"Слишком большой",
						"Достаточный",
						"Средний",
						"Недостаточный",
						"Слишком малый",
					},
					Rows: []string{
						"Объем теоретической подготовки",
						"Объем практической подготовки",
					},
				},
				{
					Text: "2. Ваши замечания и пожелания по улучшению структуры, содержания и ресурсного обеспечения дисциплины:",
					Type: Text,
				},
				{
					IsRequired: true,
					Text:       "Оцените преподавателя теоретической части (лекции) дисциплины",
					Type:       Grid,
					Options: []string{
						"Отлично",
						"Хорошо",
						"Удовл.",
						"Плохо",
						"Очень плохо",
					},
					Rows: []string{
						"Ясность и последовательность изложения материала",
						"Контакт с аудиторией",
						"Современность преподаваемого материала",
						"Объективность оценивания студентов",
					},
				},
				{
					IsRequired: true,
					Text:       "",
					Type:       Grid,
					Options: []string{
						"Слишком высокий",
						"Высокий",
						"Средний",
						"Низкий",
						"Слишком низкий",
					},
					Rows: []string{
						"Уровень требовательности преподавателя",
					},
				},
				{
					Text: "4. Ваши пожелания, замечания, предложения в адрес преподавателя теоретической части:",
					Type: Text,
				},
				{
					IsRequired: true,
					Text:       "Оцените преподавателя практической части дисциплины",
					Type:       Grid,
					Options: []string{
						"Отлично",
						"Хорошо",
						"Удовл.",
						"Плохо",
						"Очень плохо",
					},
					Rows: []string{
						"Ясность и последовательность изложения материала",
						"Контакт с аудиторией",
						"Современность преподаваемого материала",
						"Объективность оценивания студентов",
					},
				},
				{
					IsRequired: true,
					Text:       "",
					Type:       Grid,
					Options: []string{
						"Слишком большой",
						"Достаточный",
						"Средний",
						"Недостаточный",
						"Слишком малый",
					},
					Rows: []string{
						"Уровень требовательности преподавателя",
					},
				},
				{
					Text: "6. Ваши пожелания, замечания, предложения в адрес преподавателя практической части:",
					Type: Text,
				},

				{
					IsRequired: true,
					Text:       "7. Оцените полезность полученных знаний и умений по дисциплине:",
					Type:       Grid,
					Options: []string{
						"Полезны",
						"Скорее полезны",
						"Скорее полезны",
						"Скорее бесполезны",
						"Бесполезны",
					},
					Rows: []string{
						"Полезность для будущей карьеры",
						"Полезность для расширения кругозора",
					},
				},
				{
					IsRequired: true,
					Text:       "8. Укажите, какую часть занятий по дисциплине Вы посетили в предыдущем семестре:\n",
					Type:       Grid,
					Options: []string{
						"100-70%",
						"70-50%",
						"Менее 50%",
					},
					Rows: []string{
						"Теоретическая часть",
						"Практическая часть\t\n",
					},
				},
				{
					Text: "9. Укажите причины непосещаемости Вами занятий по дисциплине в предыдущем семестре:",
					Type: Text,
				},
			},
		}
	} else {
		p = Poll{
			SessionId: sessionId,
			Link:      uuid.New().String(),
			Questions: make([]Question, 0),
		}
	}

	_, err := PollCol.InsertOne(context.TODO(), p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func GetPoll(l string) (*Poll, error) {
	var p Poll
	res := PollCol.FindOne(context.TODO(), &bson.M{
		"link": l,
	})
	if err := res.Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func UpdatePoll(l string, new Poll, sessionId string) (*Poll, error) {

	var p Poll
	err := PollCol.FindOne(context.TODO(), &bson.M{
		"link": l, "sessionId": sessionId,
	}).Decode(&p)

	if p.SessionId != sessionId {
		return nil, errors.New("Unauthorized")
	}
	_, err = PollCol.UpdateOne(context.TODO(), &bson.M{
		"link": l}, &bson.M{"$set": new})

	if err != nil {
		return nil, err
	}
	upd, err := GetPoll(l)

	if err != nil {
		return nil, err
	}

	return upd, nil
}

func DeletePoll(l string, sessionId string) error {
	var p Poll
	err := PollCol.FindOne(context.TODO(), &bson.M{
		"link": l, "sessionId": sessionId,
	}).Decode(&p)

	if err == mongo.ErrNoDocuments {
		return errors.New("Not found")
	}
	_, err = PollCol.DeleteOne(context.TODO(), &bson.M{
		"link": l, "sessionId": sessionId,
	})

	return err
}
