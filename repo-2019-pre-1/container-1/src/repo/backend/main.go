package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	db := DBMigrate(DBConnect())
	app := App{DB: db}
	app.DBInit()
	r:=gin.Default()
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "stay alive",
		})
	})
	r.GET("/api/team", app.GetData)
	r.Run() // 0.0.0.0:8080
}
type App struct {
	DB *gorm.DB
}
// Handler
func (app App) GetData(c *gin.Context) {
	teams := []Team{}
	if err := app.DB.Find(&teams).Error; err != nil{
		c.JSON(500, "DB error")
	}
	c.JSON(200, teams)
}

// DB
type Team struct {
	gorm.Model
	Name string
	Score int64
	Description string
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.DropTable(&Team{})
	db.AutoMigrate(&Team{})
	return db
}

func (app App) DBInit() {
	team1 := Team{Name: "usagi-san", Score: 4000, Description: "hop-jump"}
	team2 := Team{Name: "yagi-san", Score: 3000, Description: "meee"}
	team3 := Team{Name: "kame-san", Score: 2000, Description: "byu-nnn"}
	team4 := Team{Name: "rakuda-san", Score: 1000, Description: "tokotoko"}
	team5 := Team{Name: "ookami-san", Score: 0, Description: "gao-gao-"}

	db := app.DB
	db.NewRecord(&team1)
	app.DB.Create(&team1)
	app.DB.Create(&team2)
	app.DB.Create(&team3)
	app.DB.Create(&team4)
	app.DB.Create(&team5)
}

func DBConnect() *gorm.DB {
    DBMS := "mysql"
    USER := "useradmin"
		PASS := "pass"
    PROTOCOL := "tcp(db:3306)"
    DBNAME := "test_app"
    CONNECT := USER + ":" + PASS  + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
			panic(err.Error())
    }
    return db
}
