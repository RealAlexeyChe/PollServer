package db

import (
	"LesyaBack/poll/model"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var adminCol *mongo.Collection

func CreateAdmin(a *model.AdminRequest) error {

	admin := &model.Admin{
		Email:    a.Email,
		Password: a.Password,
		Id:       uuid.NewString(),
	}

	_, err := adminCol.InsertOne(context.TODO(), admin)
	if err != nil {
		return err
	}
	return nil
}

func GetAdminId(r *model.AdminRequest) *string {
	admin := model.Admin{}
	res := adminCol.FindOne(context.TODO(), &bson.M{
		"email": r.Email, "password": r.Password,
	})

	if res.Decode(&admin) != nil {
		return nil
	}
	return &admin.Id
}
