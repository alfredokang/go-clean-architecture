package repositories

import (
	"awesomeProject/domain/entities"
	mongoinfra "awesomeProject/infra/database/mongo"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mongodrive "go.mongodb.org/mongo-driver/mongo"
)

const (
	StudentCollection = "students"
)

type StudentRepository struct {
	Ctx        context.Context
	Collection *mongodrive.Collection
}

func NewStudentRepository(ctx context.Context, client *mongodrive.Client) *StudentRepository {
	return &StudentRepository{
		Ctx:        ctx,
		Collection: mongoinfra.GetCollection(client, StudentCollection),
	}
}

func (sr StudentRepository) Create(student *entities.Student) error {
	_, err := sr.Collection.InsertOne(sr.Ctx, student)
	if err != nil {
		return err
	}
	return nil
}

func (sr StudentRepository) List() (students []entities.Student, err error) {
	cur, err := sr.Collection.Find(sr.Ctx, bson.M{})
	if err != nil {
		return students, err
	}

	// O cursor esta como bytes, o cursor é como se fosse um ponteiro da agulha de um leitor de disco de vinil
	for cur.Next(sr.Ctx) {
		var student entities.Student

		if err = cur.Decode(&student); err != nil {
			return students, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (sr StudentRepository) FindById(id uuid.UUID) (student entities.Student, err error) {
	err = sr.Collection.FindOne(sr.Ctx, bson.M{"_id": id}).Decode(&student)

	return student, err
}

func (sr StudentRepository) Update(student *entities.Student) error {
	_, err := sr.Collection.UpdateOne(sr.Ctx, bson.M{"_id": student.ID}, bson.M{"$set": student})

	return err
}

func (sr StudentRepository) Delete(id uuid.UUID) error {
	_, err := sr.Collection.DeleteOne(sr.Ctx, bson.M{"_id": id})

	return err
}
