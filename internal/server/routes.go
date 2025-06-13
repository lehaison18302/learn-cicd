package server

import (
	"net/http"

	"my_project/internal/middleware"
	"my_project/internal/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	// Khởi tạo repository, service, handler cho users
	repo := users.NewRepository(s.db)
	service := users.NewService(repo)
	handler := users.NewHandler(service)

	r.POST("/login", handler.LoginHandler)

	// Bảo vệ các route cần xác thực
	api := r.Group("/users", middleware.JWTAuthRequired())
	api.POST("/create", handler.CreateUserHandle)
	api.PUT(":id", handler.UpdateUserHandler)
	api.DELETE(":id", handler.DeleteUserHandler)

	r.GET("/users", handler.GetUsersHandler)

	// Serve static files from /web
	r.Static("/web", "./web")

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
