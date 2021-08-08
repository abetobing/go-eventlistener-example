package listeners

import "fmt"

var SlackNotifier = make(chan interface{})

func init() {
	go func() {
		for {
			user := <-SlackNotifier
			fmt.Println("Notify about user to slack ", user)
		}
	}()
}
