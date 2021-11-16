package route

import (
	"admin-ais/controller"
	"admin-ais/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)


//func MainHandler(resp http.ResponseWriter, _ *http.Request) {
//	resp.Write([]byte("Hi there! i AM TEST"))
//}

func SetupRoutes(db *gorm.DB) {
	r := gin.Default()

	// ссылку на бд
	roleRepository := repository.NewRoleRepository(db)
	rightRepository := repository.NewRightRepository(db)
	userRepository := repository.NewUserRepository(db)
	err := roleRepository.Migrate()
	if err != nil {
		log.Println("Ошибка при создании Миграции", err)

	}
	err = rightRepository.Migrate()
	if err != nil {
		log.Println("Ошибка при создании Миграции", err)
	}
	err = userRepository.Migrate()
	if err != nil {
		log.Println("Ошибка при создании Миграции", err)
	}

	// ROUTES PART
	roleController := controller.NewRoleController(roleRepository)
	rightController := controller.NewRightController(rightRepository)
	userController := controller.NewUserController(userRepository)

	roleRoutes := r.Group("/api")

	{
		roleRoutes.POST("/role", roleController.AddRole)
		roleRoutes.GET("/roles", roleController.GetAllRoles)
		roleRoutes.PUT("/role/:id", roleController.ChangeRole)
		roleRoutes.DELETE("/roles/:id", roleController.DeleteRole)

	}
	rightRoutes := r.Group("/api")

	{
		rightRoutes.POST("/right",rightController.AddRight)
		rightRoutes.GET("/rights",rightController.GetAllRights)
		rightRoutes.PUT("/right/:id",rightController.ChangeRight)
		rightRoutes.DELETE("/right/:id",rightController.DeleteRight)

	}

	userRoutes := r.Group("/api")

	{
		userRoutes.POST("/user", userController.AddUser)
		userRoutes.GET("/users", userController.GetAllUsers)
		userRoutes.PUT(	"/user/:id", userController.ChangeUser)
		userRoutes.DELETE(	"/user/:id", userController.DeleteUser)
	}

	r.Run(":9995")
}
