package main

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"reflect"
)

func main() {
	cond1 := squirrel.Eq{"id": 1}
	fmt.Println(cond1.ToSql())

	cond2 := squirrel.Like{"name": "%1irrel"}
	fmt.Println(cond2.ToSql())

	cond3 := squirrel.NotLike{"name": "%irrel"}
	fmt.Println(cond3.ToSql())

	orCond1 := squirrel.Or{cond1, cond2, cond3}
	fmt.Println(orCond1.ToSql())

	andCond1 := squirrel.And{
		FilterCond("", cond1),
		FilterCond(0, cond2),
		cond3,
		FilterCond(true, cond2),
		FilterCond("some input exist!", cond1),
	}
	fmt.Println(andCond1.ToSql())
}

func FilterCond(by interface{}, cond squirrel.Sqlizer) squirrel.Sqlizer {
	v := reflect.ValueOf(by)
	if isBlank(v) {
		return squirrel.Eq{} // 等价1=1
	}
	return cond
}

// borrow from gorm/utils.go
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}

	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
