package moderator

import (
	"context"
	"libr-simplified/model"
	"math/rand"
	"sync"
	"time"
)

func SendMod(message model.Message) string {
	var wg sync.WaitGroup
	modChan := make(chan model.ModResponse, 3)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			response := modSimulate(message, id)
			select {
			case modChan <- response:
			case <-ctx.Done():
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(modChan)
	}()

	approved := 0
	for res := range modChan {
		if res.Status == "approved" {
			approved++
		}
	}

	if approved >= 2 {
		return "approved"
	} else {
		return "rejected"
	}

}

func modSimulate(message model.Message, id int) model.ModResponse {
	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)

	i := rand.Intn(2)
	var modresponse model.ModResponse
	modresponse.ModeratorId = id
	modresponse.ResponseTime = int(delay)
	modresponse.MessageId = message.Id

	if i == 0 {
		modresponse.Status = "approved"
	} else {
		modresponse.Status = "rejected"
	}

	return modresponse
}
