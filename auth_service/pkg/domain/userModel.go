package domain

import "time"

type User struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username    *string   `json:"username"`
	Pid         uint      `json:"pid" gorm:"unique"`
	Password    string    `json:"password" validate:"required,min=8,max=16"`
	Phone       string    `json:"phone" validate:"required,len=10"`
	Email       string    `json:"email" validate:"email,required"`
	Otp         string    `json:"otp"`
	Provider    *string   `json:"provider"`
	Status      *string   `json:"status"`
	VipLevel    uint      `json:"vip_level" gorm:"default:0"`
	IsVerified  bool      `json:"isverified" gorm:"default:false"`
	IsAdmin     bool      `json:"isadmin" gorm:"default:false"`
	CanTrade    bool      `json:"can_trade" gorm:"default:true"`
	CanWithdraw bool      `json:"can_withdraw" gorm:"default:true"`
	CanDeposit  bool      `json:"can_deposit" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
