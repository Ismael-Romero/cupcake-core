package mysql

import (
	"cupcake-core/src/settings"
	"database/sql"
)

type Engine struct {
	database *sql.DB
}

func New(mse settings.MySQLEngine) *Engine {

	return &Engine{}
}

func (mse *Engine) Ping() {

}
