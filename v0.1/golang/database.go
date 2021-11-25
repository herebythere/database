package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseDetails struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int64  `json:"port"`
	Username string `json:"username"`
}

type SQLStatement struct {
	SQL    string
	Values []interface{}
}

type DatabaseInterface struct {
	pool *pgxpool.Pool
}

const (
	connectionString = "postgresql://%s:%s@%s:%d/%s?sslmode=disable"
)

var (
	errNilDetails = errors.New("pgsqlx.Create() - nil details provided")
)

func (dbi *DatabaseInterface) Query(statement *SQLStatement, err error) (*[][]interface{}, error) {
	if err != nil {
		return nil, err
	}

	results, errResults := dbi.pool.Query(context.Background(), statement.SQL, statement.Values...)
	if errResults != nil {
		return nil, errResults
	}

	defer results.Close()

	var parsedRows [][]interface{}
	for results.Next() {
		values, errValues := results.Values()
		if errValues != nil {
			return nil, errValues
		}

		parsedRows = append(parsedRows, values)
	}

	return &parsedRows, nil
}

func getConnectionStr(details *DatabaseDetails) string {
	return fmt.Sprintf(
		connectionString,
		details.Username,
		details.Password,
		details.Host,
		details.Port,
		details.Name,
	)
}

func NewInterface(details *DatabaseDetails) (*DatabaseInterface, error) {
	if details == nil {
		return nil, errNilDetails
	}

	connStr := getConnectionStr(details)
	fmt.Println(connStr)
	pool, errPool := pgxpool.Connect(context.Background(), connStr)
	if errPool != nil {
		return nil, errPool
	}

	database := DatabaseInterface{
		pool: pool,
	}

	return &database, nil
}
