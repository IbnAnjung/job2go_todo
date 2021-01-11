package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct models
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:25;not null;unique" json:"username"`
	Name      string    `gorm:"size:25;not null" json:"name"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Hash method for hashing password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword method for matching password hash
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//BeforeSave IS FOR
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

//Prepare method
func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// FindUser methos for cek user
func (u *User) FindUser(db *gorm.DB, username string) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User not Found")
	}

	return u, err
}

//Validate method for register, login
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "register":
		if u.Username == "" {
			return errors.New("Username is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Name == "" {
			return errors.New("Name is required")
		}
		return nil
	case "login":
		if u.Username == "" {
			return errors.New("Username is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		return nil
	default:
		return errors.New("choose you validate for")
	}
}

//SaveUser method
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
