package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	// Mengaktifkan middleware CORS, untuk sementara
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Mengizinkan semua origin (bisa disesuaikan)
		AllowMethods: "GET,POST,PUT,DELETE", // Metode yang diizinkan
		AllowHeaders: "Origin, Content-Type, Accept", // Header yang diizinkan
	}))

	// data waktu 
	Tm := time.Now()

	app.Get("/users", func (c *fiber.Ctx)  error {
		DateLog := fmt.Sprintf("%d-%s-%d", Tm.Day(), Tm.Month(), Tm.Year())

		D := map[string]any{
			"Name":   "naim",
			"Id":     "1255",
			"Role":   "users",
			"Email":  "naimmmmab@gmail.com",
			"Year":   "14-09-2004",
			"Nik":    "18393782598472",
			"Pass":   "123",
			"Bis_Loc": "paniki",
			"Date_Log": DateLog,
		}
		return c.JSON(D)
    })

	app.Static("/", "./build")
	log.Fatal(app.Listen(":3000"))
}