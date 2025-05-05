package main

import (
	"context"
	"log"
	"os"
	"task-scheduler-api/internal/create"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load env.")
	}
	ctx := context.Background()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("failed to connect db:", err)
	}
	defer conn.Close()

	createRepo := create.NewCreateTaskRepository(conn)
	createUsecase := create.NewCreateTaskUsecase(createRepo)
	createTaskHandler := create.CreateTaskHandler(createUsecase)

	app := fiber.New()

	app.Post("/tasks", createTaskHandler)

	app.Listen(":8080")
}
