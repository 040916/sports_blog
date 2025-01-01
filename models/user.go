package models

import (
"golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
    Email    string `gorm:"unique"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
    if u.Password != "" {
        bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        u.Password = string(bytes)
    }
 return nil
}