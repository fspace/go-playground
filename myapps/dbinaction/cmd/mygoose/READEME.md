- 看状态：
> go run main.go  sqlite3 ./foo.db status
- 创建
> go run main.go  sqlite3 ./foo.db  create fetch_user_data go
运行
> go run main.go  sqlite3 ./foo.db up
回滚一个版本
> go run main.go  sqlite3 ./foo.db down
