package main

import (
	"../utils"
	"database/sql"
	"time"
)

// SignupReq  incoming signup request
type SignupReq struct {
	Email     string    `json:"email"`
	FName     string    `json:"first_name"`
	LName     string    `json:"last_name"`
	DOB       time.Time `json:"date_of_birth"`
	Password  string    `json:"password"`
	DigitLock string    `json:"four_digit_lock"`
	IsBlocked bool      `json:"is_blocked"`
}

// SignupResp  outgoing response
type SignupResp struct {
}

// LoginReq incoming request
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResp  outgoing response
type LoginResp struct {
}

// ValidateSingupReq incoming request
type ValidateSingupReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// ValidateSingupResp  outgoing response
type ValidateSingupResp struct {
}

// LogoutReq incoming request
type LogoutReq struct {
}

// LogoutResp outgoing response
type LogoutResp struct {
}

// ForgotPasswordReq incoming request of forgot password.
type ForgotPasswordReq struct {
	Email     string `json:"email"`
	DigitLock int    `json:"four_digit_lock"`
}

// ForgotPasswordResp outgoing response
type ForgotPasswordResp struct {
	Message string `json:"message"`
}

// ForgotDigitLockReq incoming req
type ForgotDigitLockReq struct {
	Email string `json:"email"`
}

// ForgotDigitLockResp outgoing resp
type ForgotDigitLockResp struct {
}

// ChangeOldPasswordReq incoming requst
type ChangeOldPasswordReq struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

//ChangeOldPasswordResp outgoing resp
type ChangeOldPasswordResp struct {
}

// ChangeOldDigitLockReq incoming req
type ChangeOldDigitLockReq struct {
	Email        string `json:"email"`
	OldDigitLock string `json:"old_digit_lock"`
	NewDigitLock string `json:"new_digit_lock"`
}

//ChangeOldDigitLockResp outgoing resp
type ChangeOldDigitLockResp struct {
}

//ResetDigitLockReq  incoming req
type ResetDigitLockReq struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	NewDigitLock string `json:"new_digit_lock"`
}

// ResetDigitLockResp outgoing resp
type ResetDigitLockResp struct {
}

// ResetPasswordReq incoming req.
type ResetPasswordReq struct {
	Email       string `json:"email"`
	DigitLock   string `json:"digit_lock"`
	NewPassword string `json:"new_password"`
}

// ResetPasswordResp outgoing resp
type ResetPasswordResp struct {
}

// AuthModule auth module
type AuthModule struct {
	DataBase *sql.DB
	RedisDB  *utils.RedisDb
}
