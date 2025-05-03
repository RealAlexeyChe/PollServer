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
	p := Poll{}
	if c.Template {
		p = Poll{
			Link:      uuid.New().String(),
			Course:    c.Course,
			Group:     c.Group,
			Professor: c.Professor,
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
		p = Poll{}
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
