package main

type QueryBuilder interface {
	Select(table string, columns ...string) QueryBuilder
	Where(conditions ...string) QueryBuilder
	GetRawSQL() string
}

type MySQLQueryBuilder struct {
	Opt        string
	Table      string
	Columns    []string
	Conditions []string
}

func (m *MySQLQueryBuilder) Select(table string, columns ...string) QueryBuilder {
	m.Opt = "SELECT "
	m.Table = table
	m.Columns = columns
	return m
}

func (m *MySQLQueryBuilder) Where(conditions ...string) QueryBuilder {
	m.Conditions = append(m.Conditions, conditions...)
	return m
}

func (m *MySQLQueryBuilder) GetRawSQL() string {
	sql := ""

	// todo 区分 select/update/delete etc.
	sql = m.Opt
	lenColumns := len(m.Columns)
	for i, v := range m.Columns {
		sql += v
		if i != lenColumns-1 {
			sql += ", "
		}
	}
	sql += " FROM " + m.Table
	lenConditions := len(m.Conditions)
	if lenConditions != 0 {
		sql += " WHERE "
		for i, v := range m.Conditions {
			sql += v
			if i != lenConditions-1 {
				sql += " AND "
			}
		}
	}

	return sql
}

func main() {
	m := &MySQLQueryBuilder{}

	sql := m.Select("table", "id", "name", "age").
		Where("age = 10").
		Where("name = zs").
		GetRawSQL()
	println(sql)
}
