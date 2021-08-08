package listeners

import (
	"fmt"

	"github.com/abetobing/go-eventlistener-example/ext"
)

var UserCreated = make(chan interface{})

func init() {
	go func() {
		for {
			user := <-UserCreated
			fmt.Println("user event ", user)
			ext.SubmitElastic(user)
		}
	}()
}
