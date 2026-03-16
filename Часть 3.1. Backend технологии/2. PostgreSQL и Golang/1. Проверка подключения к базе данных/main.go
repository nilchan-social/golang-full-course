package main

import (
	"context"
	"fmt"
	simple_connection "study/feature_postgres"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CheckConnection() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключение к базе данных произошло успешно!")
}

func main() {
	fmt.Println("Hello, Git!")
	simple_connection.CheckConnection()
}
