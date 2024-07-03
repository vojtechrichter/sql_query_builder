package main

import (
	"log"
	"strings"
)

type QueryBuilder struct {
	query map[int]string
}

type WhereInterface struct {
	prevBuilderInstance *QueryBuilder
	conditions          []string
}

func (wi WhereInterface) Equals(col string, value string) WhereInterface {
	cond := col + " = " + value
	wi.conditions = append(wi.conditions, cond)

	return wi
}

func (wi WhereInterface) NotEquals(col string, value string) WhereInterface {
	cond := col + " != " + value
	wi.conditions = append(wi.conditions, cond)

	return wi
}

func (wi WhereInterface) EndWhere() QueryBuilder {
	wi.prevBuilderInstance.query[STATEMENT_WHERE] = strings.Join(wi.conditions, " AND ")

	return *wi.prevBuilderInstance
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

func (qb QueryBuilder) StartWhere() WhereInterface {
	wi := WhereInterface{
		prevBuilderInstance: &qb,
		conditions:          make([]string, 0),
	}

	return wi
}

func (qb QueryBuilder) GetFinal() string {
	var finalQuery strings.Builder

	finalQuery.WriteString("SELECT ")
	finalQuery.WriteString(qb.query[STATEMENT_SELECT])
	finalQuery.WriteString(" ")

	finalQuery.WriteString("FROM ")
	finalQuery.WriteString(qb.query[STATEMENT_FROM])

	finalQuery.WriteString(" WHERE ")
	finalQuery.WriteString("(" + qb.query[STATEMENT_WHERE] + ")")

	finalQuery.WriteString(";")

	return finalQuery.String()
}

func main() {
	builder := InitQueryBuilder()
	builder.Select("user", "email", "is_admin").From("administration").StartWhere().Equals("user", "admin").EndWhere()

	log.Println(builder.GetFinal())
}
