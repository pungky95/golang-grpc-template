package entities

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string
	Email     string
	Phone     string
	Password  string
	Status    string         `json:"status" gorm:"type:user_status;not null;default:'ACTIVE'"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;type:timestamp;autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;type:timestamp;autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt ` gorm:"index"`
}

func (*User) TableName() string {
	return "users"
}

func (u *User) hashPassword(password string) error {
	bytePassword := []byte(password)
	argon := argon2.DefaultConfig()
	passwordHash, err := argon.HashEncoded(bytePassword)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)
	return nil
}

func (u *User) ComparePassword(password string) (bool, error) {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	match, err := argon2.VerifyEncoded(bytePassword, byteHashedPassword)
	if err != nil {
		return false, err
	}
	return match, nil
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	err = u.hashPassword(u.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
	}
	return err
}

func (u *User) BeforeUpdate(*gorm.DB) (err error) {
	err = u.hashPassword(u.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
	}
	return err
}
