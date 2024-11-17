package main

import (
	"go_limiter_rate/internal/api"
	"go_limiter_rate/internal/database"
)

func main() {
	sqlite := *database.ConnectSqlite()
	rdb := *database.ConnectRedis()
	api.StartApp(&sqlite, &rdb)
}
