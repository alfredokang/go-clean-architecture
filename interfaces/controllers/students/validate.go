package students

import (
	"awesomeProject/domain/entities"
	"awesomeProject/domain/entities/shared"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gookit/validate"
	"io"
	"net/http"
)

func getInputBody(r *http.Request) (input InputBody, err error) {
	// Lê o corpo da solicitação HTTP
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return input, err
	}

	// Converte o corpo da solicitação para a struct Person
	err = json.Unmarshal(body, &input)
	if err != nil {
		return input, err
	}

	validation := validate.Struct(input)
	if !validation.Validate() {
		return input, validation.Errors
	}

	return input, err
}

func getInputID(r *http.Request) (id uuid.UUID, err error) {
	inputID := chi.URLParam(r, "id")

	id, err = shared.GetUuidByString(inputID)
	if err != nil {
		return id, err
	}

	return id, nil
}

func getOutputListStudents(students []entities.Student) (output OutputStudents, err error) {
	for _, s := range students {
		outputStudent, err := getOutputStudent(s)
		if err != nil {
			return output, err
		}

		output.Students = append(output.Students, outputStudent)
	}

	return output, err
}

func getOutputStudent(student entities.Student) (output OutputStudent, err error) {
	return OutputStudent{
		ID:   student.ID,
		Name: student.Name,
		Age:  student.Age,
	}, err
}
