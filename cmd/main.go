package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "fjh:1@(10.0.1.238:3306)/mysql_study?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("connect db error: ", err)
	}
	defer db.Close()

	viper.SetConfigName("config")            // 设置配置名称
	viper.SetConfigFile("config/common.ini") // 设置配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	summary := viper.GetString("mysql")
	fmt.Println(summary)
}
