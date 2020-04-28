package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	DBMigrate(DBConnect())
	r:=gin.Default()
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "stay alive",
		})
	})
	r.Run() // 0.0.0.0:8080
}

type User struct {
	gorm.Model
	Name string `json:"Name"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}

func DBConnect() *gorm.DB {
    DBMS := "mysql"
    USER := "root"
    PROTOCOL := "tcp(db:3306)"
    DBNAME := "test_app"
    CONNECT := USER + ":" + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
  		panic(err.Error())
    }
    return db
}
