//go:generate goagen bootstrap -d github.com/cboecking/gorma-sample/design

package main

import (
	"fmt"

	"github.com/cboecking/gorma-sample/app"
	"github.com/cboecking/gorma-sample/models"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var accountModel *models.AccountdddDB

func main() {
	// create db service
	var err error
	url := fmt.Sprintf("dbname=gorma user=gorma password=gorma sslmode=disable port=%d host=%s", 5432, "localhost")
	fmt.Println(url)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DropTable(&models.Accountddd{})
	db.AutoMigrate(&models.Accountddd{})

	accountModel = models.NewAccountdddDB(db)

	// Create service
	service := goa.New("cellar-aaa")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "accountbbb" controller
	c := NewAccountbbbController(service)
	app.MountAccountbbbController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}
}
