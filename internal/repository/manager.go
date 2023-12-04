package repository

import (
	"context"
	"log"
	"time"

	"github.com/lalizita/go-crud-boilerplate/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Task struct {
		Id             primitive.ObjectID `bson:"_id,omitempty"`
		Title          string             `bson:"title"`
		Description    string             `bson:"descrition"`
		Points         int                `bson:"points"`
		CreationTime   time.Time          `bson:"creation_time"`
		ExpirationTime time.Time          `bson:"expiration_time"`
		Status         string             `bson:"status"`
	}

	TaskRepository struct {
		mongoCollection *mongo.Collection
	}
)

type ITaskRepository interface {
	Insert(context.Context, entity.TaskDTOInput) (primitive.ObjectID, error)
	FindOneByID(context.Context, primitive.ObjectID) (entity.TaskDTOOutput, error)
}

func NewTaskRepository(collection *mongo.Collection) ITaskRepository {
	return &TaskRepository{
		mongoCollection: collection,
	}
}

func (m *TaskRepository) Insert(ctx context.Context, consent entity.TaskDTOInput) (primitive.ObjectID, error) {
	consentToInput := Task{
		CreationTime: time.Now(),
		Status:       consent.Status,
	}

	result, err := m.mongoCollection.InsertOne(ctx, consentToInput)
	if err != nil {
		log.Println("Error insert task", err.Error())
		return primitive.ObjectID{}, err
	}

	ok, _ := result.InsertedID.(primitive.ObjectID)
	return ok, nil
}

func (m *TaskRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (entity.TaskDTOOutput, error) {
	var output entity.TaskDTOOutput
	err := m.mongoCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&output)
	if err != nil {
		log.Println("Error find one task")
		return entity.TaskDTOOutput{}, err
	}

	return output, nil
}
