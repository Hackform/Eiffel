package q

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Q_New(t *testing.T) {
	assert := assert.New(t)
	query := New(ACTION_QUERY_ONE, "test_sector", Props{"prop_1", "prop_2", "another_prop"}, Constraints{NewCon("key_1", EQUAL, "value_1"), NewCon("key_2", EQUAL, "value_2")})
	assert.Equal(query.Action, ACTION_QUERY_ONE, "property Action should be instantiated")
	assert.Equal(query.Sector, "test_sector", "property Sector should be instantiated")
	assert.Equal(len(query.RProps), 3, "property RProps should be instantiated")
	assert.Equal(len(query.Cons), 2, "property Cons should be instantiated")
}

func Test_Q_Constraint_NewCon(t *testing.T) {
	assert := assert.New(t)
	a := NewCon("key_1", EQUAL, "value_1")
	assert.Equal(EQUAL, a.Condition, "property Condition should be instantiated")
	assert.Equal("key_1", a.Key, "property Key should be instantiated")
	assert.Equal("value_1", a.Value, "property Value should be instantiated")
	b := NewCon("key_2", EQUAL, "$")
	assert.Equal("$", b.Value, "property Value should be instantiated")
}

func Test_Q_Constraint_NewOp(t *testing.T) {
	a := NewCon("key_1", EQUAL, "value_1")
	b := NewCon("key_2", EQUAL, "$")
	assert := assert.New(t)
	c := NewOp(a, OR, b)
	assert.Equal(a, c.Con1, "property Con1 should be instantiated")
	assert.Equal(b, c.Con2, "property Con2 should be instantiated")
	assert.Equal(OR, c.Condition, "property Condition should be instantiated")
}
