package main

import (
	"context"
	"log"
	"os"
	"task-scheduler-api/db"
	"task-scheduler-api/internal/create"
	"task-scheduler-api/internal/get"
	"task-scheduler-api/internal/getall"
	"task-scheduler-api/jobs"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
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

	// Create Task
	createRepo := create.NewCreateTaskRepository(conn)
	createUsecase := create.NewCreateTaskUsecase(createRepo)
	createHandler := create.CreateTaskHandler(createUsecase)

	// Get Task By Code
	getByCodeRepo := get.NewGetTaskByCodeRepository(conn)
	getByCodeUsecase := get.NewGetTaskByCodeUsecase(getByCodeRepo)
	getByCodeHandler := get.GetTaskByCodeHandler(getByCodeUsecase)

	// Get All Tasks
	getAllRepo := getall.NewGetAllTaskRepository(conn)
	getAllUsecase := getall.NewGetAllTaskUsecase(getAllRepo)
	getAllHandler := getall.GetAllTaskHandler(getAllUsecase)

	// Fiber setup
	app := fiber.New()
	app.Post("/tasks", createHandler)
	app.Get("/tasks/:code", getByCodeHandler)
	app.Get("/tasksList", getAllHandler)

	// Cron setup
	queries := db.New(conn)
	c := cron.New()
	jobs.ScheduleAllJobs(c, queries, ctx)
	c.Start()

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
