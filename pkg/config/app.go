package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    db *gorm.DB
)

func Connect() {
    connectionString := "root:GfyWSXYfFXeoxtlsTRjYTQhlIQMWiLuY@tcp(monorail.proxy.rlwy.net:40118)/railway?charset=utf8&parseTime=True&loc=Local"

    d, err := gorm.Open("mysql", connectionString)
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB {
    return db
}
