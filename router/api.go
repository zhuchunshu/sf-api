package router

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/zhuchunshu/sf-api/client/controllers/auth"
)

func SetupAPIV1Routes(app *fiber.App) {
	api := app.Group("/api")

	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// 添加 API 路由
	api.Post("/login", auth.Login)
}
