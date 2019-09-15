package main

import (
	"fmt"
	qb "github.com/didi/gendry/builder"
)

func main() {

	mp := map[string]interface{}{
		"country":  "China",
		"role":     "driver",
		"age >":    45,
		"_groupby": "name",
		"_having": map[string]interface{}{
			"total >":  1000,
			"total <=": 50000,
		},
		"_orderby": "age desc",
	}
	cond, vals, err := qb.BuildSelect("tableName", mp, []string{"name", "count(price) as total", "age"})

	//cond: SELECT name,count(price) as total,age FROM tableName WHERE (age>? AND country=? AND role=?) GROUP BY name HAVING (total>? AND total<=?) ORDER BY age DESC
	//vals: []interface{}{45, "China", "driver", 1000, 50000}

	if nil != err {
		panic(err)
	}
	fmt.Println(cond, vals)

	//have fun !!
}
