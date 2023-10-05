package database

import (
	"awesomeProject/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Conn *mongo.Client

	StudentRepository entities.StudentRepository
	UserRepository    entities.UserRepository
}

func NewDatabase(conn *mongo.Client, sr entities.StudentRepository, ur entities.UserRepository) *Database {
	return &Database{
		Conn:              conn,
		StudentRepository: sr,
		UserRepository:    ur,
	}
}
