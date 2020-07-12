package test

import (
	"testing"
)

type Tester interface {
	That(interface{}) Comparator
}

type tester struct {
	tester *testing.T
}

func (t *tester) That(a interface{}) Comparator {
	return &comparator{tester: t, a: a}
}

func Assert(t *testing.T) Tester {
	return &tester{tester: t}
}
