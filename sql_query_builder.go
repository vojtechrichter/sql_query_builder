package main

import (
	"log"
	"strings"
)

type QueryBuilder struct {
	query map[int]string
}

func InitQueryBuilder() QueryBuilder {
	var qb QueryBuilder
	qb.query = make(map[int]string)

	return qb
}

func (qb QueryBuilder) Select(colNames ...string) QueryBuilder {
	qb.query[STATEMENT_SELECT] = strings.Join(colNames, ", ")

	return qb
}

func (qb QueryBuilder) From(tableName string) QueryBuilder {
	qb.query[STATEMENT_FROM] = tableName

	return qb
}

func (qb QueryBuilder) Where(condition string) QueryBuilder {

	return qb
}

func (qb QueryBuilder) GetFinal() string {
	var finalQuery strings.Builder

	finalQuery.WriteString("SELECT ")
	finalQuery.WriteString(qb.query[STATEMENT_SELECT])
	finalQuery.WriteString(" ")

	finalQuery.WriteString("FROM ")
	finalQuery.WriteString(qb.query[STATEMENT_FROM])

	finalQuery.WriteString(";")

	return finalQuery.String()
}

func main() {
	builder := InitQueryBuilder()
	builder.Select("user", "email", "is_admin").From("administration")

	log.Println(builder.GetFinal())
}
