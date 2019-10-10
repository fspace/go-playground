package main

import (
	"fmt"
	"github.com/doug-martin/goqu"
)

func main() {
	sql, _, _ := goqu.From("test").Where(goqu.Ex{
		"d": []string{"a", "b", "c"},
	}).ToSql()
	fmt.Println(sql)

	// 主要考察下where子句的构造能力
	w1 := goqu.Ex{"a": goqu.Op{"gt": 10}}
	fmt.Println(w1.Expression())
}
