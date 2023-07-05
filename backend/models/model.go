package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id"`
	Task_name   string             `json:"task_name"`
	Task_detail string             `json:"task_detail"`
	Date        time.Time          `json:"date"`
}
