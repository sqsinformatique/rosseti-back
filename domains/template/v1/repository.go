package templatev1

import (
	"context"
	"time"

	"github.com/sqsinformatique/rosseti-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (tt *TemplateV1) templateDB() *mongo.Collection {
	mongoconn := *tt.mongodb
	return mongoconn.Database(tt.cfg.Mongo.TemplateDB).Collection("templates")
}

func (t *TemplateV1) GetTemplateByID(id string) (data *models.Template, err error) {
	filter := bson.D{{"id", id}}

	var result models.Template
	err = t.templateDB().FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *TemplateV1) CreateTemplate(request *models.Template) (*models.Template, error) {
	data := *request

	data.CreateTimestamp()

	_, err := t.templateDB().InsertOne(context.Background(), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (t *TemplateV1) UpdateTemplateByID(id string, request *models.Template) (*models.Template, error) {
	data := *request
	data.UpdateTimestamp()

	filter := bson.D{{"name", id}}

	_, err := t.templateDB().UpdateOne(context.Background(), filter, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (t *TemplateV1) SoftDeleteTemplateByID(id string) (err error) {
	data, err := t.GetTemplateByID(id)
	if err != nil {
		return
	}

	if data.DeletedAt.Valid {
		return
	}

	data.DeletedAt.Time = time.Now()
	data.DeletedAt.Valid = true
	data.UpdatedAt.Time = time.Now()
	data.UpdatedAt.Valid = true

	filter := bson.D{{"name", data.Name}}

	_, err = t.templateDB().UpdateOne(context.Background(), filter, data)
	if err != nil {
		return err
	}

	return
}

func (t TemplateV1) HardDeleteTemplateByID(id string) (err error) {
	filter := bson.D{{"name", id}}

	_, err = t.templateDB().DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
