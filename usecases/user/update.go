package user

import (
	"awesomeProject/domain/entities"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *UserUseCase) Update(id primitive.ObjectID, ur entities.User) (entities.User, error) {
	user, err := s.Database.UserRepository.FindById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("Não foi possível encontrar o usuário!")
		}
		return user, err
	}

	user.Name = ur.Name
	user.Mobile = ur.Mobile
	user.Email = ur.Email
	user.Role = ur.Role
	user.Avatar = ur.Avatar
	user.AvatarBucketKey = ur.AvatarBucketKey
	user.BankDataName = ur.BankDataName
	user.BankDataAgency = ur.BankDataAgency
	user.BankDataAccount = ur.BankDataAccount
	user.BankDataType = ur.BankDataType
	user.BankDataPix = ur.BankDataPix
	user.TenantCompanyId = ur.TenantCompanyId
	user.TenantCompanyName = ur.TenantCompanyName
	user.TenantCompanyMobile = ur.TenantCompanyMobile
	user.TenantCompanyEmail = ur.TenantCompanyEmail
	user.IsAuthorized = ur.IsAuthorized

	err = s.Database.UserRepository.Update(&user)

	return user, err
}
