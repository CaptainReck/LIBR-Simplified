package db

import (
	"context"
	"fmt"
	model "libr-simplified/model"
	"log"
)

func InsertMessage(message model.Message) (string, error) {
	query := "INSERT INTO messages(id,content,timestamp,status) VALUES ($1,$2,$3,$4)"
	_, err := Pool.Exec(context.Background(), query, message.Id, message.Content, message.Timestamp, message.Status)
	if err != nil {
		fmt.Printf("error inserting message: %v", err)
		return "Error", err
	}
	return "Message Successfully Inserted", nil
}

func GetMessages(ts int64) []model.Message {
	query := "SELECT * FROM messages WHERE timestamp = $1"
	rows, err := Pool.Query(context.Background(), query, ts)
	if err != nil {
		fmt.Printf("error inserting message: %v", err)
		return nil
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var message model.Message
		if err := rows.Scan(&message.Id, &message.Content, &message.Timestamp, &message.Status); err != nil {
			log.Fatalf("Error scanning row: %v", err)
			continue
		}
		messages = append(messages, message)
	}
	return messages
}
