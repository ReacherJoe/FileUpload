package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type File struct {
	Images string `json:"images"`
}

var db *gorm.DB
var err error

func main() {
	e := echo.New()
	dsn := "user=postgres password=password dbname=filesever host=localhost port=5433 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	e.Use(middleware.CORS())
	db.AutoMigrate(&File{})

	e.Static("/", "./")

	e.POST("/upload", Upload)
	// e.POST("/upload", Insertfoeupload)
	e.GET("/data", ReadAll)

	e.Logger.Fatal(e.Start(":8080"))
}

func Upload(c echo.Context) error {

	file, err := c.FormFile("file")

	if err != nil {
		log.Println(err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		log.Println(err)
		return err
	}
	defer src.Close()

	rootFileName := "assets"
	subFileName := time.Now().Format("02-Jan-2006")

	if _, err := os.Stat(rootFileName); os.IsNotExist(err) {
		err := os.Mkdir(rootFileName, 0755)
		if err != nil {
			log.Println(err)
			return err
		}
		if _, err := os.Stat(subFileName); os.IsNotExist(err) {
			err := os.Mkdir(fmt.Sprintf("%s/%s", rootFileName, subFileName), 0755)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s/%s", rootFileName, subFileName, file.Filename))

	if err != nil {
		log.Println(err)
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
		return err
	}
	files := File{
		Images: fmt.Sprintf("http://localhost:8080/%s/%s/%s", rootFileName, subFileName, file.Filename),
	}

	err = db.Create(&files).Error
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"images": ("http://localhost:8080/" + rootFileName + "/" + subFileName + "/" + file.Filename),
	})
}

func ReadAll(c echo.Context) error {

	var files []File

	err := db.Find(&files).Error
	if err != nil {
		return nil
	}
	fmt.Println(files)

	return c.JSON(http.StatusOK, files)

}
