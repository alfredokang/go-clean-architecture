package users

import (
	"awesomeProject/domain/entities"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gookit/validate"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

func getInputBody(r *http.Request) (input InputBody, err error) {
	// Lê o corpo da solicitação HTTP
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return input, err
	}

	// Converte o corpo da solicitação para a struct
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

func getUpdateBody(r *http.Request) (update UpdateBody, err error) {
	// Lê o corpo da solicitação HTTP
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return update, err
	}

	// Converte o corpo da solicitação para a struct
	err = json.Unmarshal(body, &update)
	if err != nil {
		return update, err
	}

	validation := validate.Struct(update)
	if !validation.Validate() {
		return update, validation.Errors
	}

	return update, err
}

func getInputID(r *http.Request) (id primitive.ObjectID, err error) {
	inputID := chi.URLParam(r, "id")

	id, err = primitive.ObjectIDFromHex(inputID)
	if err != nil {
		return id, err
	}

	return id, nil
}

func getOutputListUsers(users []entities.User) (output OutputUsers, err error) {
	for _, s := range users {
		outputUsers, err := getOutputDetailedUser(s)
		if err != nil {
			return output, err
		}

		output.Users = append(output.Users, outputUsers)
	}

	return output, err
}

func getOutputDetailedUser(user entities.User) (output OutputDetailedUser, err error) {
	return OutputDetailedUser{
		Id:                   user.Id,
		Name:                 user.Name,
		Mobile:               user.Mobile,
		Email:                user.Email,
		Role:                 user.Role,
		Avatar:               user.Avatar,
		AvatarBucketKey:      user.AvatarBucketKey,
		BankDataName:         user.BankDataName,
		BankDataAgency:       user.BankDataAgency,
		BankDataAccount:      user.BankDataAccount,
		BankDataType:         user.BankDataType,
		BankDataPix:          user.BankDataPix,
		TenantCompanyId:      user.TenantCompanyId,
		TenantCompanyName:    user.TenantCompanyName,
		TenantCompanyMobile:  user.TenantCompanyMobile,
		TenantCompanyEmail:   user.TenantCompanyEmail,
		IsAuthorized:         user.IsAuthorized,
		Password:             user.Password,
		PasswordResetToken:   user.PasswordResetToken,
		PasswordResetExpires: user.PasswordResetExpires,
		CreatedDate:          user.CreatedDate,
		CreatedAt:            user.CreatedAt,
	}, err
}
