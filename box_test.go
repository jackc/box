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
	c.Check(b.Status(), Equals, byte(box.Undefined))
}

func (s *MySuite) TestNewTime(c *C) {
	val := time.Now()
	b := box.NewTime(val)

	val2, present := b.Get()
	c.Check(val2, Equals, val)
	c.Check(present, Equals, true)
}

func (s *MySuite) TestSetAndGet(c *C) {
	var b box.Time
	val := time.Now()

	b.Set(val)

	val2, present := b.Get()
	c.Check(val2, Equals, val)
	c.Check(present, Equals, true)

	c.Check(b.MustGet(), Equals, val)

	b.SetUndefined()
	_, present = b.Get()
	c.Check(present, Equals, false)

	b.SetUnknown()
	_, present = b.Get()
	c.Check(present, Equals, false)

	b.SetEmpty()
	_, present = b.Get()
	c.Check(present, Equals, false)
}

func (s *MySuite) TestGetCoerceNil(c *C) {
	var b box.Time

	b.SetUndefined()
	c.Check(b.GetCoerceNil(), Equals, nil)

	b.SetUnknown()
	c.Check(b.GetCoerceNil(), Equals, nil)

	b.SetEmpty()
	c.Check(b.GetCoerceNil(), Equals, nil)

	val := time.Now()
	b.Set(val)
	c.Check(b.GetCoerceNil(), Equals, val)
}

func (s *MySuite) TestSetCoerceNil(c *C) {
	var b box.Time

	b.SetCoerceNil(nil, box.Empty)

	c.Check(b.Status(), Equals, byte(box.Empty))
}

func (s *MySuite) TestGetCoerceZero(c *C) {
	var b box.Time
	var zero time.Time

	b.SetUndefined()
	c.Check(b.GetCoerceZero(), Equals, zero)

	b.SetUnknown()
	c.Check(b.GetCoerceZero(), Equals, zero)

	b.SetEmpty()
	c.Check(b.GetCoerceZero(), Equals, zero)

	val := time.Now()
	b.Set(val)
	c.Check(b.GetCoerceNil(), Equals, val)
}

func (s *MySuite) TestSetCoerceZero(c *C) {
	var b box.Time
	var zero time.Time

	b.SetCoerceZero(zero, box.Empty)
	c.Check(b.Status(), Equals, byte(box.Empty))
}

func (s *MySuite) TestMustGetPanicsWhenNotFull(c *C) {
	var b box.Time

	b.SetUndefined()
	c.Check(b.MustGet, Panics, "called MustGet on a box that was not full")

	b.SetUnknown()
	c.Check(b.MustGet, Panics, "called MustGet on a box that was not full")

	b.SetEmpty()
	c.Check(b.MustGet, Panics, "called MustGet on a box that was not full")
}
