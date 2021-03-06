package main

import (
	"context"
	"fmt"
	"log"

	"github.com/B6102647/app/controllers"
	"github.com/B6102647/app/ent"
	"github.com/B6102647/app/ent/role"
	"github.com/B6102647/app/ent/status"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	Name  string
	Email string
	Role  int
}

type Books struct {
	Book []Book
}

type Book struct {
	Name     string
	Username string
	Author   string
	Category string
	Status   int
}

type Purposes struct {
	Puspose []Purpose
}

type Purpose struct {
	Name string
}

type Roles struct {
	Role []Role
}

type Role struct {
	Name string
}

type Statuss struct {
	Status []Status
}

type Status struct {
	Name string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewBookController(v1, client)
	controllers.NewPurposeController(v1, client)
	controllers.NewBookBorrowController(v1, client)
	controllers.NewRoleController(v1, client)
	controllers.NewStatusController(v1, client)

	// Set Status Data
	statuss := []string{"Available", "in use"}

	for _, st := range statuss {

		client.Status.
			Create().
			SetSTATUSNAME(st).
			Save(context.Background())
	}

	// Set Books Data
	books := Books{
		Book: []Book{
			Book{"Microprocessor for Beginner", "Somsri Tongdee", "Mr.VC", "Study", 1},
			Book{"System Analysis", "Somsri Tongdee", "Mr.CW", "Study", 1},
			Book{"Solo Leveling Novel", "Makmee Deedee", "JinWoo", "Relax", 1},
			Book{"Basic Math for Kids", "MakMee Deedee", "QG.", "Study", 1},
			Book{"How to Sleep 5 hrs. in 3 hrs.", "Somsri Tongdee", "Okay", "Other", 1},
		},
	}

	for _, b := range books.Book {

		st, err := client.Status.
			Query().
			Where(status.IDEQ(int(b.Status))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Book.
			Create().
			SetBOOKNAME(b.Name).
			SetUSERNAME(b.Username).
			SetAuthor(b.Author).
			SetCATEGORY(b.Category).
			SetStatus(st).
			Save(context.Background())
	}

	// Set Purposes Data
	purposes := []string{"เพื่อการศึกษาค้นคว้าเพิ่มเติม", "เพื่อการทำงานวิจัย", "เพื่อความบันเทิง", "เพื่อวัตถุประสงค์อื่น ๆ"}

	for _, pp := range purposes {

		client.Purpose.
			Create().
			SetPURPOSENAME(pp).
			Save(context.Background())
	}

	// Set Roles Data
	roles := []string{"สมาชิกห้องสมุด", "บรรณารักษ์", "บุคคลภายนอก"}

	for _, ro := range roles {

		client.Role.
			Create().
			SetROLENAME(ro).
			Save(context.Background())
	}

	// Set Users Data
	users := Users{
		User: []User{
			User{"Manuschanok Srikhrueadong", "Manuschanok@gmail.com", 1},
			User{"Sridong Manus", "B6102647@example.com", 1},
			User{"Dang Green", "DangDang@Yahoo.com", 1},
			User{"Sorn Dee", "SornZaa55@gmail.com", 1},
		},
	}

	for _, u := range users.User {

		ro, err := client.Role.
			Query().
			Where(role.IDEQ(int(u.Role))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.User.
			Create().
			SetUSERNAME(u.Name).
			SetUSEREMAIL(u.Email).
			SetRolePlay(ro).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
