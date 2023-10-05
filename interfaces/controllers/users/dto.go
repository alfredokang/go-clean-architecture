package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InputBody struct {
	Name                 string `json:"name" validate:"required|string"`
	Mobile               string `json:"mobile" validate:"required|string"`
	Email                string `json:"email" validate:"required|string"`
	Role                 string `json:"role" validate:"required|string"`
	Avatar               string `json:"avatar"`
	AvatarBucketKey      string `json:"avatarBucketKey"`
	BankDataName         string `json:"bankDataName"`
	BankDataAgency       string `json:"bankDataAgency"`
	BankDataAccount      string `json:"bankDataAccount"`
	BankDataType         string `json:"bankDataType"`
	BankDataPix          string `json:"bankDataPix"`
	TenantCompanyId      string `json:"tenantCompanyId"`
	TenantCompanyName    string `json:"tenantCompanyName"`
	TenantCompanyMobile  string `json:"tenantCompanyMobile"`
	TenantCompanyEmail   string `json:"tenantCompanyEmail"`
	IsAuthorized         bool   `json:"isAuthorized"`
	Password             string `json:"password" validate:"required|string"`
	PasswordResetToken   string `json:"passwordResetToken"`
	PasswordResetExpires string `json:"passwordResetExpires"`
}

type UpdateBody struct {
	Name                 string `json:"name" validate:"required|string"`
	Mobile               string `json:"mobile" validate:"required|string"`
	Email                string `json:"email" validate:"required|string"`
	Role                 string `json:"role" validate:"required|string"`
	Avatar               string `json:"avatar"`
	AvatarBucketKey      string `json:"avatarBucketKey"`
	BankDataName         string `json:"bankDataName"`
	BankDataAgency       string `json:"bankDataAgency"`
	BankDataAccount      string `json:"bankDataAccount"`
	BankDataType         string `json:"bankDataType"`
	BankDataPix          string `json:"bankDataPix"`
	TenantCompanyId      string `json:"tenantCompanyId"`
	TenantCompanyName    string `json:"tenantCompanyName"`
	TenantCompanyMobile  string `json:"tenantCompanyMobile"`
	TenantCompanyEmail   string `json:"tenantCompanyEmail"`
	IsAuthorized         bool   `json:"isAuthorized"`
	Password             string `json:"password"`
	PasswordResetToken   string `json:"passwordResetToken"`
	PasswordResetExpires string `json:"passwordResetExpires"`
}

type OutputDetailedUser struct {
	Id                   primitive.ObjectID `json:"_id"`
	Name                 string             `json:"name"`
	Mobile               string             `json:"mobile"`
	Email                string             `json:"email"`
	Role                 string             `json:"role"`
	Avatar               string             `json:"avatar"`
	AvatarBucketKey      string             `json:"avatarBucketKey"`
	BankDataName         string             `json:"bankDataName"`
	BankDataAgency       string             `json:"bankDataAgency"`
	BankDataAccount      string             `json:"bankDataAccount"`
	BankDataType         string             `json:"bankDataType"`
	BankDataPix          string             `json:"bankDataPix"`
	TenantCompanyId      string             `json:"tenantCompanyId"`
	TenantCompanyName    string             `json:"tenantCompanyName"`
	TenantCompanyMobile  string             `json:"tenantCompanyMobile"`
	TenantCompanyEmail   string             `json:"tenantCompanyEmail"`
	IsAuthorized         bool               `json:"isAuthorized"`
	Password             string             `json:"password"`
	PasswordResetToken   string             `json:"passwordResetToken"`
	PasswordResetExpires string             `json:"passwordResetExpires"`
	CreatedDate          string             `json:"createdDate"`
	CreatedAt            time.Time          `json:"createdAt"`
}

type OutputUser struct {
	User OutputDetailedUser `json:"user"`
	Msg  string             `json:"msg"`
	Flag string             `json:"flag"`
}

type OutputUsers struct {
	Users []OutputDetailedUser `json:"users"`
	Msg   string               `json:"msg"`
	Flag  string               `json:"flag"`
}

type OutputRaw struct {
	Msg  string `json:"msg"`
	Flag string `json:"flag"`
}
