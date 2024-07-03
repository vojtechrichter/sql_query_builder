package main

const (
	STATEMENT_SELECT      = iota
	STATEMENT_FROM        = iota
	STATEMENT_WHERE       = iota
	STATEMENT_JOIN        = iota
	STATEMENT_INSERT_INTO = iota
	STATEMENT_VALUES      = iota
	STATEMENT_GROUP_BY    = iota
	STATEMENT_ORDER_BY    = iota
	STATEMENT_LIMIT       = iota
)
