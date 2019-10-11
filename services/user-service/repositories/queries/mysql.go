package queries

import (
	"strings"
)

type query int

const (
	QueryGetAllUser query = iota + 1
)

func replace(s string, v map[string]string) string {
	for k, v := range v {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

// Q is function to get Query dinamic
func Q(key query, opt map[string]string) string {
	tables := map[string]string{
		"users": "rpc.users",
	}

	queries := map[query]string{
		QueryGetAllUser: `SELECT 
		id, name, last_name 
		FROM users`,
	}

	val, ok := queries[key]
	if ok {
		val = replace(val, opt)
		return replace(val, tables)
	}
	return ""
}
