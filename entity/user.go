package entity

type User struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Age       uint   `bson:"age" json:"age"`
	Email     string `bson:"email" json:"email"`
}
