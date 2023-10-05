package users

import (
	"awesomeProject/domain/entities"
	"awesomeProject/domain/repositories"
	"awesomeProject/interfaces/controllers"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
	UserUseCase repositories.UserUsecaseContract
}

func NewUserController(uu repositories.UserUsecaseContract) *UserController {
	return &UserController{
		UserUseCase: uu,
	}
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	input, err := getInputBody(r)
	fmt.Println(input)

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

	userDto := entities.UserDto{
		Name:         input.Name,
		Mobile:       input.Mobile,
		Email:        input.Email,
		Role:         input.Role,
		IsAuthorized: input.IsAuthorized,
		Password:     input.Password,
	}

	user, err := uc.UserUseCase.Create(userDto)
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

	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(js)
}

func (uc *UserController) List(w http.ResponseWriter, r *http.Request) {
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

	users, err := uc.UserUseCase.List()
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

	output, err := getOutputListUsers(users)
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

	response := NewResponseUsersMessage(output.Users, "Usuários enviados com sucesso!", "success")

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}

func (uc *UserController) Details(w http.ResponseWriter, r *http.Request) {
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

	studentFound, err := uc.UserUseCase.FindById(id)
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

	output, err := getOutputDetailedUser(studentFound)
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

	response := NewResponseUserMessage(output, "Usuário enviado com sucesso!", "success")

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
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

	updateBody, err := getUpdateBody(r)
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

	ur := entities.User{
		Name:                updateBody.Name,
		Mobile:              updateBody.Mobile,
		Email:               updateBody.Email,
		Role:                updateBody.Role,
		Avatar:              updateBody.Avatar,
		AvatarBucketKey:     updateBody.AvatarBucketKey,
		BankDataName:        updateBody.BankDataName,
		BankDataAgency:      updateBody.BankDataAgency,
		BankDataAccount:     updateBody.BankDataAccount,
		BankDataType:        updateBody.BankDataType,
		BankDataPix:         updateBody.BankDataPix,
		TenantCompanyId:     updateBody.TenantCompanyId,
		TenantCompanyName:   updateBody.TenantCompanyName,
		TenantCompanyMobile: updateBody.TenantCompanyMobile,
		TenantCompanyEmail:  updateBody.TenantCompanyEmail,
		IsAuthorized:        updateBody.IsAuthorized,
	}

	user, err := uc.UserUseCase.Update(id, ur)
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

	output, err := getOutputDetailedUser(user)
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

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
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

	if err = uc.UserUseCase.Delete(id); err != nil {
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

	response := NewResponseMessage("Usuário exluído com sucesso!", "success")

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(js)
	if err != nil {
		return
	}
}
