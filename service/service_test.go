package service

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }
type ServiceSuite struct{}
var _ = Suite(&ServiceSuite{})

func init() {
}

func (s *ServiceSuite) SetUpSuite(c *C) {
}

func (s *ServiceSuite) TearDownSuite(c *C) {
}
