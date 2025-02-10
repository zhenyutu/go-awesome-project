package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancal := context.WithCancel(context.Background())

	go reqTask(ctx, "work1")
	go reqTask(ctx, "work2")

	time.Sleep(3 * time.Second)
	cancal()
	time.Sleep(3 * time.Second)

}
