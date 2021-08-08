# go-eventlistener-example
Go Lang REST API example with event listeners

This project is a demo of a very basic implementation of ``Observer Pattern`` with one event publisher and many listeners.

Each commit in `master` branch represent development progressions from a basic rest api server (before implementing event listeners) to the very end of the commit (after implements event listeners). You can compare the difference between each commits.

## Running

To run the example:
```bash
go run .
```


## User Event

In this source code, we have one event publisher `UserEvent` which implement `ApplicationEvent` interface.

Also we have two listeners:
* 1st listening on `UserCreated` channel
* 2nd listening on `SlackNotifier` channel

All channel accept arbitraty `interface{}` as object that we can pass everytime an event is published.

Example below showing 2 listener subscibed to `user_created` event:
```go
var userEvent *events.UserEvent = &events.UserEvent{}
userEvent.Subscribe("user_created", listeners.UserCreated)
userEvent.Subscribe("user_created", listeners.SlackNotifier)
```

Then, whenever `user_created` event is published, it will be broadcasted to all its listeners
```go
userEvent.Publish("user_created", user)
```
in the code above, the `user` is the arbitraty object passed to all the listeners, so then they can do further logic to it.


Example code showing a listener doing its job:
```go
var SlackNotifier = make(chan interface{})

func init() {
	go func() {
		for {
			user := <-SlackNotifier
			fmt.Println("Notify about user to slack ", user)
		}
	}()
}
```
