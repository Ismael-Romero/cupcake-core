package repository

import (
	"cupcake-core/src/repository/mysql"
	"cupcake-core/src/settings"
	"fmt"
)

type Repository interface {
	Ping()
}

/*
	This code defines a function called New that takes two arguments: engine and dbe.
	engine is a string representing the type of database engine you want to use,
	and dbe is a struct containing database settings.
	The function returns a value of type Repository and an error.
	The function has a switch statement that checks the value of engine.
	If engine is equal to "mysql", the function creates a new MySQL repository by calling the mysql.New function.
	Otherwise, the function returns an error with a message indicating that the specified engine is not supported.
*/

func New(engine string, dbe settings.DatabaseEngine) (repo Repository, err error) {
	switch engine {
	case "mysql":
		repo = mysql.New(dbe.MySQL)

	default:
		err = fmt.Errorf("bad database engine %s", engine)
	}

	return repo, err
}
