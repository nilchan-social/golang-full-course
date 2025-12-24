package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("tasks:", tasks)

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Покормить кошку"
			task.Description = "Отсыпать кошке 30 грамм корма"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}

			break
		}
	}

	fmt.Println("succeed!")
}
