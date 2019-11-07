package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct
type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}
