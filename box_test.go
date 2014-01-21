package box_test

import (
	"github.com/JackC/box"
	. "launchpad.net/gocheck"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestZeroValueTimeIsUndefined(c *C) {
	var b box.Time
	c.Check(b.Status(), Equals, box.Undefined)
}

func (s *MySuite) TestNewTime(c *C) {
	val := time.Now()
	b := box.NewTime(val)

	c.Assert(b.IsFull(), Equals, true)
	c.Check(b.Get(), Equals, val)
}

func (s *MySuite) TestSetAndGet(c *C) {
	var b box.Time
	val := time.Now()

	b.Set(val)

	c.Assert(b.IsFull(), Equals, true)
	c.Check(b.Get(), Equals, val)
}

func (s *MySuite) TestGetPanicsWhenNotFull(c *C) {
	var b box.Time

	b.SetUndefined()
	c.Check(b.Get, Panics, "Tried to get from box that was not full")

	b.SetUnknown()
	c.Check(b.Get, Panics, "Tried to get from box that was not full")

	b.SetEmpty()
	c.Check(b.Get, Panics, "Tried to get from box that was not full")
}
