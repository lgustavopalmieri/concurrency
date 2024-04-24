package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/configuration/logger"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/entity/user_entity"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	UserId string `bson:"_user_id"`
	Name   string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) FindUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_user_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with id %s", userId), err)
			return nil, internal_error.NewNotFoundError(fmt.Sprintf("User not found with id %s", userId))
		}
		logger.Error("Error trying to find user by id", err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by id")
	}

	userEntity := &user_entity.User{
		UserId: userEntityMongo.UserId,
		Name:   userEntityMongo.Name,
	}
	return userEntity, nil
}
