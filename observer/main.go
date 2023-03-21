package main

import (
	"fmt"
	"math"
)

type Observer interface {
	OnNotify(Event)
}

type Notifier interface {
	Register(Observer)
	Deregister(Observer)
	Notify(Event)
}

type Event struct {
	Data int64
}

type eventObserver struct {
	id int
}

type eventNotifier struct {
	observers map[Observer]struct{}
}

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("%d:%d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (o *eventNotifier) Notify(e Event) {
	for p := range o.observers {
		p.OnNotify(e)
	}
}

func main() {
	notifier := eventNotifier{
		observers: make(map[Observer]struct{}),
	}

	notifier.Register(&eventObserver{1})
	notifier.Register(&eventObserver{2})
	notifier.Register(&eventObserver{3})

	notifier.Notify(Event{Data: math.MaxInt64})
}
