package rest_api

import (
	"github.com/jmoiron/sqlx"
)

//Data ...
type repository struct {
	DB *sqlx.DB
}

// Create Repository
func NewRepository(db *sqlx.DB) Repository {
	return repository{DB: db}
}

func (r repository) Get(ID string) (rest_apiResponse, error) {

	//you can add your logic for database
	//r.DB.Get(&your_db_object,"your_sql_query",your_params1, ...)

	return rest_apiResponse{ ID }, nil
}
