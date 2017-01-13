package q

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Q_New(t *testing.T) {
	assert := assert.New(t)
	query := New(ACTION_QUERY_ONE, "test_sector", Props{"prop_1", "prop_2", "another_prop"}, Constraints{NewCon("key_1", EQUAL, "value_1"), NewCon("key_2", EQUAL, "value_2")})
	assert.Equal(ACTION_QUERY_ONE, query.Action, "property Action should be instantiated")
	assert.Equal("test_sector", query.Sector, "property Sector should be instantiated")
	assert.Equal(3, len(query.RProps), "property RProps should be instantiated")
	assert.Equal(2, len(query.Cons), "property Cons should be instantiated")
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
	assert := assert.New(t)
	a := NewCon("key_1", EQUAL, "value_1")
	b := NewCon("key_2", EQUAL, "$")
	c := NewOp(a, OR, b)
	assert.Equal(a, c.Con1, "property Con1 should be instantiated")
	assert.Equal(b, c.Con2, "property Con2 should be instantiated")
	assert.Equal(OR, c.Condition, "property Condition should be instantiated")
}