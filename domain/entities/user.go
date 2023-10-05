package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id                   primitive.ObjectID `json:"id" bson:"_id"`
	Name                 string             `json:"name" bson:"name"`
	Mobile               string             `json:"mobile" bson:"mobile"`
	Email                string             `json:"email" bson:"email"`
	Role                 string             `json:"role" bson:"role"`
	Avatar               string             `json:"avatar" bson:"avatar,omitempty"`
	AvatarBucketKey      string             `json:"avatarBucketKey" bson:"avatarBucketKey,omitempty"`
	BankDataName         string             `json:"bankDataName" bson:"bankDataName,omitempty"`
	BankDataAgency       string             `json:"bankDataAgency" bson:"bankDataAgency,omitempty"`
	BankDataAccount      string             `json:"bankDataAccount" bson:"bankDataAccount,omitempty"`
	BankDataType         string             `json:"bankDataType" bson:"bankDataType,omitempty"`
	BankDataPix          string             `json:"bankDataPix" bson:"bankDataPix,omitempty"`
	TenantCompanyId      string             `json:"tenantCompanyId" bson:"tenantCompanyId,omitempty"`
	TenantCompanyName    string             `json:"tenantCompanyName" bson:"tenantCompanyName,omitempty"`
	TenantCompanyMobile  string             `json:"tenantCompanyMobile" bson:"tenantCompanyMobile,omitempty"`
	TenantCompanyEmail   string             `json:"tenantCompanyEmail" bson:"tenantCompanyEmail,omitempty"`
	IsAuthorized         bool               `json:"isAuthorized" bson:"isAuthorized"`
	Password             string             `json:"password" bson:"password"`
	PasswordResetToken   string             `json:"passwordResetToken" bson:"passwordResetToken,omitempty"`
	PasswordResetExpires string             `json:"passwordResetExpires" bson:"passwordResetExpires,omitempty"`
	CreatedDate          string             `json:"createdDate" bson:"createdDate"`
	CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
}

type UserDto struct {
	Name                 string `json:"name" bson:"name"`
	Mobile               string `json:"mobile" bson:"mobile"`
	Email                string `json:"email" bson:"email"`
	Role                 string `json:"role" bson:"role"`
	Avatar               string `json:"avatar" bson:"avatar"`
	AvatarBucketKey      string `json:"avatarBucketKey" bson:"avatarBucketKey"`
	BankDataName         string `json:"bankDataName" bson:"bankDataName"`
	BankDataAgency       string `json:"bankDataAgency" bson:"bankDataAgency"`
	BankDataAccount      string `json:"bankDataAccount" bson:"bankDataAccount"`
	BankDataType         string `json:"bankDataType" bson:"bankDataType"`
	BankDataPix          string `json:"bankDataPix" bson:"bankDataPix"`
	TenantCompanyId      string `json:"tenantCompanyId" bson:"tenantCompanyId"`
	TenantCompanyName    string `json:"tenantCompanyName" bson:"tenantCompanyName"`
	TenantCompanyMobile  string `json:"tenantCompanyMobile" bson:"tenantCompanyMobile"`
	TenantCompanyEmail   string `json:"tenantCompanyEmail" bson:"tenantCompanyEmail"`
	IsAuthorized         bool   `json:"isAuthorized" bson:"isAuthorized"`
	Password             string `json:"password" bson:"password"`
	PasswordResetToken   string `json:"passwordResetToken" bson:"passwordResetToken"`
	PasswordResetExpires string `json:"passwordResetExpires" bson:"passwordResetExpires"`
}

func NewUser(user UserDto) *User {
	return &User{
		Id:                   primitive.NewObjectID(),
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
		CreatedDate:          time.Now().Format("02/01/2006 15:04:05"),
		CreatedAt:            time.Now(),
	}
}

type UserRepository interface {
	Create(student *User) error
	List() ([]User, error)
	FindById(id primitive.ObjectID) (User, error)
	Update(user *User) error
	Delete(id primitive.ObjectID) error
}
