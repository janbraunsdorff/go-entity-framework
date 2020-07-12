package test

import (
	"errors"
	"fmt"
	"reflect"
)

type Comparator interface {
	IsEqualTo(i interface{}, msg ...string)
	GotError(i string, msg ...string)
	IsFalse(msg ...string)
	IsNotNil(msg ...string)
	IsNil(msg ...string)
	IsTypeOf(typ interface{}, msg ...string)
	IsTrue(msg ...string)
}

type comparator struct {
	tester *tester
	a      interface{}
}

func (c *comparator) IsTypeOf(t interface{}, msg ...string) {
	if reflect.TypeOf(t) != reflect.TypeOf(c.a) {
		c.tester.tester.Errorf(getErrorMessage(reflect.TypeOf(c.a), reflect.TypeOf(t), msg[0]))
	}

}

func (c *comparator) IsFalse(msg ...string) {
	if c.a != false {
		c.tester.tester.Errorf(getErrorMessage(c.a, false, msg[0]))

	}
}

func (c *comparator) IsTrue(msg ...string) {
	if c.a != true {
		c.tester.tester.Errorf(getErrorMessage(c.a, true, msg[0]))

	}
}

func (c *comparator) IsNotNil(msg ...string) {
	if c.a == nil {
		c.tester.tester.Errorf(getErrorMessage(c.a, nil, msg[0]))

	}
}

func (c *comparator) IsNil(msg ...string) {
	if c.a != nil {
		c.tester.tester.Errorf(getErrorMessage(c.a, nil, msg[0]))

	}
}

func (c *comparator) IsEqualTo(b interface{}, msg ...string) {
	if !reflect.DeepEqual(c.a, b) {
		c.tester.tester.Errorf(getErrorMessage(c.a, b, msg[0]))
	}
}

func (c *comparator) GotError(i string, msg ...string) {
	c.IsNotNil(msg...)

	if reflect.TypeOf(c.a).Name() == "error" {
		c.tester.tester.Errorf(getErrorMessage(c.a, errors.New(""), msg[0]))
	}

	if c.a.(error).Error() != i {
		c.tester.tester.Errorf(getErrorMessage(c.a, i, msg[0]))
	}
}

func getErrorMessage(a, b interface{}, msg string) string {
	return fmt.Sprintf("\n%-12s: %v (%T)\n%-12s: %v (%T)\n%-12s: %s", "expected", b, b, "got", a, a, "message", msg)
}
