package models

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"go-db-labs/util"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	USER_ENABLED  = "ENABLED"
	USER_DISABLED = "DISABLED"
)

type User struct {
	ID        string
	Name      string
	Username  string
	Email     string
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser() *User {
	return &User{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewDummyUser() *User {
	dummy := NewUser()
	random := util.Random().Int63n(60)
	dummy.Status = USER_ENABLED
	if random%2 == 0 {
		dummy.Status = USER_DISABLED
	}
	dummy.Name = faker.Username()
	dummy.Email = faker.Email()
	dummy.Username = faker.Username()
	dummy.Name = faker.Name()
	password, _ := bcrypt.GenerateFromPassword([]byte(faker.Password()), bcrypt.DefaultCost)
	dummy.Password = string(password)
	dummy.CreatedAt = time.Now().Add(time.Hour * 24 * time.Duration(-random)).UTC()
	dummy.UpdatedAt = dummy.CreatedAt.Add(time.Hour * 24 * time.Duration(random)).UTC()
	return dummy
}

