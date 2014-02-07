package box

import (
	"time"
)

const (
	Undefined = iota
	Unknown   = iota
	Empty     = iota
	Full      = iota
)

type Time struct {
	value  time.Time
	status int
}

func NewTime(v time.Time) (box Time) {
	box.Set(v)
	return box
}

func (box *Time) Set(v time.Time) {
	box.value = v
	box.status = Full
}

func (box *Time) SetUndefined() {
	box.status = Undefined
}

func (box *Time) SetUnknown() {
	box.status = Unknown
}

func (box *Time) SetEmpty() {
	box.status = Empty
}

// Panics unless box is Full
func (box *Time) Get() time.Time {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Time) Status() int {
	return box.status
}

func (box *Time) IsFull() bool {
	return box.status == Full
}
