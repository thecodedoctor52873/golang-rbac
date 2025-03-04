package main

import (
    "log"
    "rbac-app/config"
    "rbac-app/controllers"
    "rbac-app/middleware"
    "rbac-app/storage"

    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()
    storage.ConnectRedis()

    r := gin.Default()

    public := r.Group("/")
    public.POST("register", controllers.RegisterUser)
    public.POST("login", controllers.LoginUser)

    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware())
    protected.Use(middleware.RateLimiter())

    protected.GET("admin/data", middleware.RequireRole("admin"), controllers.GetAdminData)
    protected.GET("user/data", middleware.RequireRole("user"), controllers.GetUserData)

    log.Println("Server running on :8080")
    r.Run(":8080")
}