package q

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Q_NewQOne(t *testing.T) {
	assert := assert.New(t)
	query := NewQOne("test_sector", Props{"prop_1", "prop_2", "another_prop"}, Constraints{NewCon("key_1", EQUAL, "value_1"), NewCon("key_2", EQUAL, "value_2")})
	assert.Equal(ACTION_QUERY_ONE, query.Action, "property Action should be instantiated")
	assert.Equal("test_sector", query.Sector, "property Sector should be instantiated")
	assert.Equal(3, len(query.RProps), "property RProps should be instantiated")
	assert.Equal(2, len(query.Cons), "property Cons should be instantiated")
}

func Test_Q_NewQMulti(t *testing.T) {
	assert := assert.New(t)
	query := NewQMulti("test_sector", Props{"prop_1", "prop_2", "another_prop"}, Constraints{NewCon("key_1", EQUAL, "value_1"), NewCon("key_2", EQUAL, "value_2")}, 5, Constraints{NewOrd("prop_1", ASC)})
	assert.Equal(ACTION_QUERY_MULTI, query.Action, "property Action should be instantiated")
	assert.Equal("test_sector", query.Sector, "property Sector should be instantiated")
	assert.Equal(3, len(query.RProps), "property RProps should be instantiated")
	assert.Equal(2, len(query.Cons), "property Cons should be instantiated")
	assert.Equal(5, query.Limit, "property Limit should be instantiated")
	assert.Equal(1, len(query.Order), "property Order should be instantiated")
}

func Test_Q_NewI(t *testing.T) {
	assert := assert.New(t)
	query := NewI("test_sector", Props{"prop_1", "prop_2", "another_prop"}, Props{"val_1", "val_2", "val_3"})
	assert.Equal(ACTION_INSERT, query.Action, "property Action should be instantiated")
	assert.Equal("test_sector", query.Sector, "property Sector should be instantiated")
	assert.Equal(3, len(query.RProps), "property RProps should be instantiated")
	assert.Equal(len(query.RProps), len(query.Vals), "property Vals should be instantiated")
}

func Test_Q_NewU(t *testing.T) {
	assert := assert.New(t)
	query := NewU("test_sector", Constraints{NewCon("key_1", EQUAL, "value_1"), NewCon("key_2", EQUAL, "value_2")}, Constraints{NewCon("key_1", EQUAL, "value_1"), NewCon("key_2", EQUAL, "value_2")})
	assert.Equal(ACTION_UPDATE, query.Action, "property Action should be instantiated")
	assert.Equal("test_sector", query.Sector, "property Sector should be instantiated")
	assert.Equal(2, len(query.Mods), "property Mods should be instantiated")
	assert.Equal(2, len(query.Cons), "property Cons should be instantiated")
}

func Test_Q_NewD(t *testing.T) {
	assert := assert.New(t)
	query := NewD("test_sector", Constraints{NewCon("key_1", EQUAL, "value_1")})
	assert.Equal(ACTION_DELETE, query.Action, "property Action should be instantiated")
	assert.Equal("test_sector", query.Sector, "property Sector should be instantiated")
	assert.Equal(1, len(query.Cons), "property Cons should be instantiated")
}

func Test_Q_NewT(t *testing.T) {
	assert := assert.New(t)
	query := NewT("test_sector", Constraints{NewType("key_1", INT, NONE, 0), NewType("key_2", BIGINT, UNIQUE, 0)})
	assert.Equal(ACTION_CREATE_TABLE, query.Action, "property Action should be instantiated")
	assert.Equal("test_sector", query.Sector, "property Sector should be instantiated")
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

func Test_Q_Constraint_NewEq(t *testing.T) {
	assert := assert.New(t)
	a := NewEq("key_1", "value_1")
	assert.Equal(EQUAL, a.Condition, "property Condition should be instantiated")
	assert.Equal("key_1", a.Key, "property Key should be instantiated")
	assert.Equal("value_1", a.Value, "property Value should be instantiated")
	b := NewEq("key_2", "$")
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

func Test_Q_Constraint_NewOrd(t *testing.T) {
	assert := assert.New(t)
	a := NewOrd("key_1", ASC)
	b := NewOrd("key_2", DESC)

	assert.Equal("key_1", a.Key, "property Key should be instantiated")
	assert.Equal(ASC, a.Condition, "property Condition should be instantiated")
	assert.Equal("key_2", b.Key, "property Key should be instantiated")
	assert.Equal(DESC, b.Condition, "property Condition should be instantiated")
}

func Test_Q_Constraint_NewType(t *testing.T) {
	assert := assert.New(t)
	a := NewType("key_1", INT, NOT_NULL_UNIQUE, 0)
	assert.Equal("key_1", a.Key, "property Key should be instantiated")
	assert.Equal(INT, a.Condition, "property Condition should be instantiated")
	assert.Equal(NOT_NULL_UNIQUE, a.ColCon, "property ColCon should be instantiated")
}
