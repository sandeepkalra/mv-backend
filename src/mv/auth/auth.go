package main

import (
	"database/sql"
	"mv/utils"
	"time"
)

/* -------- Signup -------- */
type SignupReq struct {
	Email     string    `json:"email"`
	FName     string    `json:"first_name"`
	LName     string    `json:"last_name"`
	DOB       time.Time `json:"date_of_birth"`
	Password  string    `json:"password"`
	DigitLock int       `json:"four_digit_lock"`
	IsBlocked bool      `json:"is_blocked"`
}

type SignupResp struct {
}

/* -------- Login -------- */
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
}

/* -------- ValidateSingup -------- */
type ValidateSingupReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ValidateSingupResp struct {
}

/* -------- Logout -------- */
type LogoutReq struct {
}

type LogoutResp struct {
}

/* -------- ForgotPassword -------- */
type ForgotPasswordReq struct {
	Email     string `json:"email"`
	DigitLock int    `json:"four_digit_lock"`
}

type ForgotPasswordResp struct {
	Message string `json:"message"`
}

/* -------- ForgotDigitLock -------- */
type ForgotDigitLockReq struct {
	Email string `json:"email"`
}

type ForgotDigitLockResp struct {
}

/* -------- ChangeOldPassword -------- */
type ChangeOldPasswordReq struct {
	Email       string `json:"email"`
	DigitLock   int    `json:"four_digit_lock"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangeOldPasswordResp struct {
}

/* -------- ChangeOldDigitLock -------- */
type ChangeOldDigitLockReq struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	OldDigitLock int    `json:"old_four_digit_lock"`
	NewDigitLock int    `json:"new_four_digit_lock"`
}

type ChangeOldDigitLockResp struct {
}

type AuthModule struct {
	DataBase *sql.DB
	RedisDB  *utils.RedisDb
}
