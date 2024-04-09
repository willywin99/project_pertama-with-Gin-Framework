package main

import (
	"project_pertama/controller"
	"project_pertama/lib"
	"project_pertama/middleware"
	"project_pertama/model"
	"project_pertama/repository"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "project_pertama/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {

	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Person{}, &model.CreditCard{}, &model.User{})
	if err != nil {
		panic(err)
	}

	personRepository := repository.NewPersonRepository(db)
	personController := controller.NewPersonController(personRepository)

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	ginEngine := gin.Default()

	// personGroup := ginEngine.Group("/person", gin.BasicAuth(gin.Accounts{
	// 	"admin": "12345",
	// 	"willy": "golang_mantap",
	// }))

	// personGroup.GET("/", personController.GetAll)
	// personGroup.POST("/", personController.Create)

	// ginEngine.Use(middleware.LogMiddleware)
	// personGroup := ginEngine.Group("/person")
	// personGroup.Use(middleware.LogMiddleware)
	// personGroup.GET("/", personController.GetAll)
	// personGroup.POST("/", personController.Create)

	ginEngine.POST("/users/register", userController.Register)
	ginEngine.POST("/users/login", userController.Login)

	personGroup := ginEngine.Group("/person", middleware.AuthMiddleware, middleware.AdminMiddleware)
	personGroup.GET("", personController.GetAll)
	personGroup.POST("", personController.Create)
	// ginEngine.PUT("/person/:id", personController.Update)
	personGroup.DELETE("/:id", personController.Delete)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = ginEngine.Run("localhost:8082")
	if err != nil {
		panic(err)
	}
}
