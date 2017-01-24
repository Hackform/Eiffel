package cassandra

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type (
	Fields map[string]string
)

func BuilderTable(name string, fields Fields, partitionKey, clusterKey []string) (string, error) {
	if len(name) < 1 {
		return "", errors.New("name must be defined")
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

	return fmt.Sprintf("CREATE TABLE %s (%s, PRIMARY KEY (%s))", name, strings.Join(b, ", "), strings.Join(keys, ", ")), nil
}
