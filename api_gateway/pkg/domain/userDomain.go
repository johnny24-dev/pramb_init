package domain

import "time"

type EmailRequest struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type UpdatePasswordRequest struct {
	Otp             string `json:"otp"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type User struct {
	Id           uint      `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username     *string   `json:"username"`
	PId          uint      `json:"pid" gorm:"unique"`
	Password     string    `json:"password" validate:"required,min=8,max=16"`
	Phone        string    `json:"phone" validate:"required,len=10"`
	Email        string    `json:"email" validate:"email,required"`
	Otp          string    `json:"otp"`
	Provider     *string   `json:"provider"`
	Status       *string   `json:"status"`
	Vip_Level    uint      `json:"vip_level" gorm:"default:0"`
	Isverified   bool      `json:"isverified" gorm:"default:false"`
	Isadmin      bool      `json:"isadmin" gorm:"default:false"`
	Can_Trade    bool      `json:"can_trade" gorm:"default:true"`
	Can_Withdraw bool      `json:"can_withdraw" gorm:"default:true"`
	Can_Deposit  bool      `json:"can_deposit" gorm:"default:true"`
	Created_At   time.Time `json:"created_at"`
	Updated_At   time.Time `json:"updated_at"`
}
