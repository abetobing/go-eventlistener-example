package events

type ApplicationEvent interface {
	Subscribe(string, chan interface{})
	Unsubscribe(string, chan interface{})
	Publish(string, interface{})
}

type UserEvent struct {
	listeners map[string][]chan interface{}
}

func (d *UserEvent) Subscribe(eventType string, ch chan interface{}) {
	if d.listeners == nil {
		d.listeners = make(map[string][]chan interface{})
	}
	d.listeners[eventType] = append(d.listeners[eventType], ch)
}

func (d *UserEvent) Unsubscribe(eventType string, ch chan interface{}) {
	if _, ok := d.listeners[eventType]; ok {
		for i := range d.listeners[eventType] {
			if d.listeners[eventType][i] == ch {
				d.listeners[eventType] = append(d.listeners[eventType][:i], d.listeners[eventType][i+1:]...)
				break
			}
		}
	}
}

func (d *UserEvent) Publish(eventType string, content interface{}) {
	if _, ok := d.listeners[eventType]; ok {
		for _, handler := range d.listeners[eventType] {
			go func(handler chan interface{}) {
				handler <- content
			}(handler)
		}
	}
}
