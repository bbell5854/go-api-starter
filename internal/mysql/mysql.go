package mysql

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
