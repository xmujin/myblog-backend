package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xmujin/myblog-backend/internal/controller"
	"github.com/xmujin/myblog-backend/internal/middlewares"
	"github.com/xmujin/myblog-backend/internal/model"
	"github.com/xmujin/myblog-backend/internal/repository"
	"github.com/xmujin/myblog-backend/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	{

		dsn := "host=192.168.1.100 user=postgres password=123456 dbname=myblog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		db.AutoMigrate(&model.Post{})
		db.AutoMigrate(&model.User{})
		postRepository := repository.NewPostRepository(db)
		postService := service.NewPostService(postRepository)
		postController := controller.NewPostController(postService)

		userRepository := repository.NewUserRepository(db)
		userService := service.NewUserService(userRepository)
		userController := controller.NewUserController(userService)

		api := r.Group("/api/v1")

		posts := api.Group("/posts").Use(middlewares.JWTAuth())
		posts.GET("/", postController.GetPosts())
		posts.GET("/:id", postController.GetPostById())
		posts.POST("/", postController.CreatePost())
		posts.DELETE("/:id", postController.DeletePostById())
		posts.PUT("/:id", postController.UpdatePost())

		api.POST("/register", userController.Register())
		api.POST("/login", userController.Login())

	}
	return r
}
