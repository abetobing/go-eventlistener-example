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
			fmt.Println("User event listener, about to submit data to elasticsearch ", user)
			ext.SubmitElastic(user)
		}
	}()
}
