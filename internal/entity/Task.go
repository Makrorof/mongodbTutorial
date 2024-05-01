package entity

import (
	"fmt"
	"time"

	"github.com/Makrorof/mongodbTutorial/tools"
	"go.mongodb.org/mongo-driver/bson"
)

type Task struct {
	ID        TaskID    `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Completed bool      `bson:"completed"`
}

type TaskID struct {
	//ID   primitive.ObjectID `bson:"id" json:"id"`
	Text string `bson:"text" json:"text"`
}

func ParseTaskID(data interface{}) TaskID {
	taskID := TaskID{}
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &taskID)

	return taskID
}

func (t TaskID) String() string {
	//return fmt.Sprintf("ID: %s, Text[:30]: %s", t.ID.String(), tools.StrLimit(t.Text, 30))
	return fmt.Sprintf("Text[:30]: %s", tools.StrLimit(t.Text, 30))
}

func (t TaskID) IsZero() bool {
	return len(t.Text) == 0
}
