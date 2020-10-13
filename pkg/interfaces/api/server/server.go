package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Serve サーバー起動処理
func Serve(addr string) {
	e := echo.New()
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Content-Type", "Accept", "Origin", "x-token"},
		}),
		middleware.Logger(),
		middleware.Recover(),
	)

	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
