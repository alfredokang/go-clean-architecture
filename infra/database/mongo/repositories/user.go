package repositories

import (
	"awesomeProject/domain/entities"
	mongoinfra "awesomeProject/infra/database/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodrive "go.mongodb.org/mongo-driver/mongo"
)

const (
	UserCollection = "users"
)

type UserRepository struct {
	Ctx        context.Context
	Collection *mongodrive.Collection
}

func NewUserRepository(ctx context.Context, client *mongodrive.Client) *UserRepository {
	return &UserRepository{
		Ctx:        ctx,
		Collection: mongoinfra.GetCollection(client, UserCollection),
	}
}

func (ur UserRepository) Create(user *entities.User) error {
	_, err := ur.Collection.InsertOne(ur.Ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) List() (users []entities.User, err error) {
	cur, err := ur.Collection.Find(ur.Ctx, bson.M{})
	if err != nil {
		return users, err
	}

	// O cursor esta como bytes, o cursor é como se fosse um ponteiro da agulha de um leitor de disco de vinil
	for cur.Next(ur.Ctx) {
		var user entities.User

		if err = cur.Decode(&user); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur UserRepository) FindById(id primitive.ObjectID) (user entities.User, err error) {
	err = ur.Collection.FindOne(ur.Ctx, bson.M{"_id": id}).Decode(&user)

	return user, err
}

func (ur UserRepository) Update(user *entities.User) error {
	_, err := ur.Collection.UpdateOne(ur.Ctx, bson.M{"_id": user.Id}, bson.M{"$set": user})

	return err
}

func (ur UserRepository) Delete(id primitive.ObjectID) error {
	_, err := ur.Collection.DeleteOne(ur.Ctx, bson.M{"_id": id})

	return err
}
