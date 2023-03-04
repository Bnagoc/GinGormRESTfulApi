package main

import (
	"log"

	"github.com/bnagos/GinGormRESTfulApi/models"
	"github.com/bnagos/GinGormRESTfulApi/routers"
	"github.com/bnagos/GinGormRESTfulApi/storage"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "host= , user= , password= , dbname= ")
	if err != nil {
		log.Println("error while accessing to DB", err)
	}
	defer storage.DB.Close()
	storage.DB.AutoMigrate(&models.Article{})

	r := routers.SetupRouter()
	r.Run()
}
