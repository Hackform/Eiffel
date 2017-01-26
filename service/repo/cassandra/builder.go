package cassandra

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type (
	// Fields is a map of db model fields to cassandra data types
	Fields map[string]string
)

// BuilderTable creates a new cql string to create a table
func BuilderTable(sector string, fields Fields, partitionKey, clusterKey []string) (string, error) {
	if len(sector) < 1 {
		return "", errors.New("sector must be defined")
	}
	if len(fields) < 1 {
		return "", errors.New("there must be at least one field")
	}
	if len(partitionKey) < 1 {
		return "", errors.New("there must be at least one partition key")
	}

	b := []string{}
	for k, v := range fields {
		b = append(b, fmt.Sprintf("%s %s", k, v))
	}
	sort.Strings(b)

	var k string
	if len(partitionKey) == 1 {
		k = partitionKey[0]
	} else {
		k = fmt.Sprintf("(%s)", strings.Join(partitionKey, ", "))
	}
	keys := []string{k}
	if clusterKey != nil && len(clusterKey) > 0 {
		keys = append(keys, clusterKey...)
	}

	return fmt.Sprintf("CREATE TABLE %s (%s, PRIMARY KEY (%s))", sector, strings.Join(b, ", "), strings.Join(keys, ", ")), nil
}
