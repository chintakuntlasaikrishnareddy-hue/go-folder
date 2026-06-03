package main

import (
	"context"
	"fmt"
	"time"
)

type Item struct {
	ID        int
	value     string
	Timestamp time.Time
}

func DataGenarator(ctx context.Context, limit int) <-chan Item {
	// create unbuffered channel.

	out := make(chan Item)
	//  create a ananymous function&call as goroutine.

	go func() {
		defer close(out)

		for i := 1; i <= limit; i++ {
			time.Sleep(200 * time.Millisecond)

			Item := Item{
				ID:        i,
				value:     fmt.Sprintf("Datapayload %d", i),
				Timestamp: time.Now(),
			}
			select {

			case <-ctx.Done():
				fmt.Println("cleanup,Exiting")
			case out <- Item:
				fmt.Println("")
			}
		}
		fmt.Println("[Generator] finished generating all Items")
	}()
	return out
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)

	defer cancel()
	fmt.Println("[Main]starting Generator.........")
	Itemschan := DataGenarator(ctx, 7)

	for Item := range Itemschan {
		fmt.Printf("[main]Recived: ID=%d|value=%s|Time=%s\n",
			Item.ID,
			Item.value,
			Item.Timestamp.Format("24:00:00:000"),
		)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("[main]Execution finished")
}
