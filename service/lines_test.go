package service

import (
	"strings"
	. "gopkg.in/check.v1"
)

func (s *ServiceSuite) TestLinesSize(c *C) {
	for i := 0; i < 1000; i++ {
		l := Lines(i)
		c.Assert(len(l), Equals, i)
	}
}

func (s *ServiceSuite) TestLoad(c *C) {
	values := Load(strings.NewReader("one\ntwo\nthree"))
	c.Assert(len(values), Equals, 3)
	c.Assert(values[0], Equals, "one")
	c.Assert(values[1], Equals, "two")
	c.Assert(values[2], Equals, "three")
}
