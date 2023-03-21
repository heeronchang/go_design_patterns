package main

import "fmt"

type SQL struct {
	Table   string
	Columns []string
	Where   []string
}

type Option func(s *SQL)

func Table(t string) Option {
	return func(s *SQL) {
		s.Table = t
	}
}

func Columns(cs ...string) Option {
	return func(s *SQL) {
		s.Columns = append(s.Columns, cs...)
	}
}

func Where(conditions ...string) Option {
	return func(s *SQL) {
		s.Where = append(s.Where, conditions...)
	}
}

func NewSQL(options ...Option) *SQL {
	sql := &SQL{}

	for _, option := range options {
		option(sql)
	}

	return sql
}

func main() {
	sql := NewSQL(Table("user"),
		Columns("name", "password"),
		Where("id = 100"),
	)

	fmt.Printf("%v", sql)
}
