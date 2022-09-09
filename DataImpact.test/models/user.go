package models


type Friends struct {
	Id_friend string `json:"id_friend" bson:"id_friend"`
	Name_friend string `json:"name_friend" bson:"name_friend"`
}

type User struct {
	Id string `json:"id" bson:"id"`
	Password string `json:"password" bson:"password"`
	IsActive bool `json:"IsActive" bson:"IsActive"`
	Balance string `json:"balance" bson:"balance"`
	Age int `json:"age" bson:"age"`
	Name string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Company string `json:"company" bson:"company"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
	About string `json:"about" bson:"about"`
	Registered string `json:"registered" bson:"registered"`
	Latitude float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
	Tags []string `json:"tags" bson:"tags"`
	Data string `json:"data" bson:"data"`
}