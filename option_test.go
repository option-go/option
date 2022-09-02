package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {

	o1 := Some[int](2)
	o2 := None[int]()
	assert.Equal(t, true, o1.And(&o2).None)

	o3 := None[int]()
	o4 := Some[string]("foo")
	assert.Equal(t, true, OptAndOpt(&o3, &o4).None)

	o5 := Some[int](2)
	o6 := Some[string]("foo")
	assert.Equal(t, true, OptAndOpt(&o5, &o6).IsSome())
	assert.Equal(t, "foo", OptAndOpt(&o5, &o6).Some)

}
