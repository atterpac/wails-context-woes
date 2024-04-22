package backend

import "fmt"

func Deploy(channel chan string, event string) {
	// Count to 25
	for i := 0; i < 25; i++ {
		channel <- fmt.Sprintf("Event %d", i)
	}
}
