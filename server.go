package main

import (
	"programming/golang/rest-api/config"
	"programming/golang/rest-api/controller"
	"programming/golang/rest-api/middleware"
	"programming/golang/rest-api/repository"
	"programming/golang/rest-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository   repository.UserRepository   = repository.NewUserRepository(db)
	pirateRepository repository.PirateRepository = repository.NewPirateRepository(db)
	jwtService       service.JWTService          = service.NewJWTService()
	userService      service.UserService         = service.NewUserService(userRepository)
	pirateService    service.PirateService       = service.NewPirateService(pirateRepository)
	authService      service.AuthService         = service.NewAuthService(userRepository)
	authController   controller.AuthController   = controller.NewAuthController(authService, jwtService)
	userController   controller.UserController   = controller.NewUserController(userService, jwtService)
	pirateController controller.PirateController = controller.NewPirateController(pirateService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	pirateRoutes := r.Group("api/pirates", middleware.AuthorizeJWT(jwtService))
	{
		pirateRoutes.GET("/", pirateController.All)
		pirateRoutes.POST("/", pirateController.Insert)
		pirateRoutes.GET("/:id", pirateController.FindByID)
		pirateRoutes.PUT("/:id", pirateController.Update)
		pirateRoutes.DELETE("/:id", pirateController.Delete)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
