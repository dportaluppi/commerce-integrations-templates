package mongodb

import (
	"context"
	"time"

	"github.com/dportaluppi/commerce-integrations-templates/pkg/template"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "commerce_integrations_templates"
	collectionName = "templates"
	timeout        = 10 * time.Second
)

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepository(mongoURI string) (*Repository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	collection := client.Database(databaseName).Collection(collectionName)
	return &Repository{
		client:     client,
		collection: collection,
	}, nil
}

func (r *Repository) Save(t *template.Template) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"name": t.Name}, bson.M{"$set": t}, options.Update().SetUpsert(true))
	return err
}

func (r *Repository) FindByName(name string) (*template.Template, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var tmpl template.Template
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&tmpl)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, template.ErrTemplateNotFound
		}
		return nil, err
	}

	return &tmpl, nil
}

func (r *Repository) Delete(t *template.Template) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"name": t.Name})
	return err
}
