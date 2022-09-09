package services

import (
	"context"
	"errors"
	//"time"
	

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"dataimpact/test/golang/models"

)
    
type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(users *[]interface{}) error {
	_, err := u.usercollection.InsertMany(u.ctx, *users)
	return err
}


func (u *UserServiceImpl) LoginUser(id string, password string) (error, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	query1 := bson.D{bson.E{Key: "password", Value: password}}
	err1 := u.usercollection.FindOne(u.ctx, query1).Decode(&user)
	return err, err1
}


func (u *UserServiceImpl) GetUser(id *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}






func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{primitive.E{Key: "id", Value: user.Id}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "id", Value: user.Id}, 
		primitive.E{Key: "password", Value: user.Password}, 
		primitive.E{Key: "isActive", Value: user.IsActive},
		primitive.E{Key: "balance", Value: user.Balance},
		primitive.E{Key: "age", Value: user.Age}, 
		primitive.E{Key: "name", Value: user.Name}, 
		primitive.E{Key: "gender", Value: user.Gender},
		primitive.E{Key: "company", Value: user.Company}, 
		primitive.E{Key: "email", Value: user.Email}, 
		primitive.E{Key: "phone", Value: user.Phone},
		primitive.E{Key: "address", Value: user.Address}, 
		primitive.E{Key: "about", Value: user.About}, 
		primitive.E{Key: "registred", Value: user.Registered},
		primitive.E{Key: "latitude", Value: user.Latitude}, 
		primitive.E{Key: "longitude", Value: user.Longitude}, 
		primitive.E{Key: "tags", Value: user.Tags},
		primitive.E{Key: "data", Value: user.Data},
	}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}