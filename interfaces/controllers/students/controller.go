package students

import (
	"awesomeProject/domain/repositories"
	"awesomeProject/interfaces/controllers"
	"encoding/json"
	"net/http"
)

type StudentController struct {
	StudentUseCase repositories.StudentUsecaseContract
}

func NewStudentController(su repositories.StudentUsecaseContract) *StudentController {
	return &StudentController{
		StudentUseCase: su,
	}
}

func (sc *StudentController) Create(w http.ResponseWriter, r *http.Request) {
	input, err := getInputBody(r)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	student, err := sc.StudentUseCase.Create(input.Name, input.Age)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	js, err := json.Marshal(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(js)
}

func (sc *StudentController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getInputID(r)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	if err = sc.StudentUseCase.Delete(id); err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	_, err = w.Write([]byte("Estudante exluido com sucesso!"))
	if err != nil {
		return
	}
}

func (sc *StudentController) Details(w http.ResponseWriter, r *http.Request) {
	id, err := getInputID(r)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError("Erro com o seu ID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	studentFound, err := sc.StudentUseCase.SearchByID(id)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	output, err := getOutputStudent(studentFound)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}

func (sc *StudentController) List(w http.ResponseWriter, r *http.Request) {
	// Exemplo para acessar as chaves dos tokens enviados
	//claims, ok := r.Context().Value("claims").(jwt.MapClaims)
	//if !ok {
	//	http.Error(w, "Invalid claims", http.StatusInternalServerError)
	//	return
	//}
	//
	//// Acessar as chaves do token JWT
	//nome := claims["nome"].(string)
	//email := claims["email"].(string)

	students, err := sc.StudentUseCase.List()
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	output, err := getOutputListStudents(students)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}

func (sc *StudentController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getInputID(r)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	input, err := getInputBody(r)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	student, err := sc.StudentUseCase.Update(id, input.Name, input.Age)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	output, err := getOutputStudent(student)
	if err != nil {
		js, err := json.Marshal(controllers.NewResponseMessageError(err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}
