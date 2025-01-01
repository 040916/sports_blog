package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "./models" // 使用相对路径引用models包
)

var DB *gorm.DB

func Connect() {
    var err error
    // 替换下面的username和password为您的MySQL数据库的用户名和密码
    // 同时更新数据库名称为sykgotest
    DB, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/sykgotest?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the schema
    err = DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
    if err != nil {
        log.Fatal("Error migrating the schema:", err)
    }
}
  