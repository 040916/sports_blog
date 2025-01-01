package models

import (
    "gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    Content   string
    ArticleID uint
    User      User `gorm:"foreignKey:UserID"`
}