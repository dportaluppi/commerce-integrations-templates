package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/dportaluppi/commerce-integrations-templates/pkg/template"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "commerce-integrations-templates"
	collectionName = "templates"
	timeout        = 10 * time.Second
)

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepository(mongoURI string) (*Repository, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Add a timeout to the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping MongoDB to check if the connection is valid
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
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
