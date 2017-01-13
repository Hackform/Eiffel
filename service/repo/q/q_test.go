package q

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Q_New(t *testing.T) {
	assert := assert.New(t)
	query := New(ACTION_QUERY_ONE, "test_sector", Props{"prop_1", "prop_2", "another_prop"}, Constraints{NewConstraint("key_1", EQUAL, "value_1"), NewConstraint("key_2", EQUAL, "value_2")})
	assert.Equal(query.Action, ACTION_QUERY_ONE, "property Action should be instantiated")
	assert.Equal(query.Sector, "test_sector", "property Sector should be instantiated")
	assert.Equal(len(query.RProps), 3, "property RProps should be instantiated")
	assert.Equal(len(query.Cons), 2, "property Cons should be instantiated")
}
