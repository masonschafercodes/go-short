package worker

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func updateCountForShortId(client *pgxpool.Pool, shortId string) error {
	_, err := client.Exec(context.Background(), "UPDATE links SET access_count = access_count + 1 WHERE short_url = ($1)", shortId)

	if err != nil {
		log.Println("Error updating access count", err)
		return err
	}

	return nil
}

type UpdateTask struct {
	ShortId string
}

var UpdateQueue = make(chan UpdateTask, 1000)

func UpdateWorker(dbPool *pgxpool.Pool) {
	for task := range UpdateQueue {
		err := updateCountForShortId(dbPool, task.ShortId)
		if err != nil {
			log.Println("Error updating access count", err)
		}
	}
}
