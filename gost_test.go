package gost_test

import (
	"reflect"
	"testing"

	gc "gopkg.in/check.v1"

	"github.com/synapse-garden/gost"
)

func Test(t *testing.T) { gc.TestingT(t) }

type GostSuite struct{}

var _ = gc.Suite(&GostSuite{})

func (s *GostSuite) TestPutGet(c *gc.C) {
	tests := []struct {
		description string
		toSet       interface{}
		setKey      string
		fetchKey    string
		errMsg      string
	}{{
		description: "set and get normally",
		toSet: struct {
			S string
			I int
		}{"foo", 5},
		setKey:   "foo",
		fetchKey: "foo",
	}, {
		description: "fail get with bad key",
		toSet:       struct{ I int }{5},
		setKey:      "foo",
		fetchKey:    "bar",
		errMsg:      "no value for key \"bar\"",
	}}

	for i, t := range tests {
		c.Logf("test %v: %s", i, t.description)
		g, err := gost.GetGost(gost.DefaultKind)
		c.Assert(err, gc.IsNil)
		err = g.Put(t.toSet, t.setKey)
		c.Assert(err, gc.IsNil)
		putType := reflect.TypeOf(t.toSet)
		result := reflect.New(putType)
		resultRef := result.Interface()
		err = g.Get(resultRef, t.fetchKey)
		if t.errMsg != "" {
			c.Assert(err.Error(), gc.Equals, t.errMsg)
			c.Check(result.Elem().Interface(), gc.DeepEquals, reflect.Zero(putType).Interface())
		} else {
			c.Assert(err, gc.IsNil)
			c.Check(result.Elem().Interface(), gc.DeepEquals, t.toSet)
		}
	}
}
