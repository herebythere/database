package database

import (
	"fmt"
	"testing"
)

const (
	defaultDBAddress	= "127.0.0.1"
	defaultDBPort		= "3015"
	defaultDBUsername	= "superdb"
	defaultDBName		= "superdb"
	defaultDBPassword	= "suberdb_is_awesome"
)

// Env variables for tests
var (
	dbAddress  = os.Getenv("TEST_DATABASE_HOST_ADDRESS")
	dbPort     = os.Getenv("TEST_DATABASE_HOST_PORT")
	dbUsername = os.Getenv("TEST_DATABASE_USERNAME")
	dbName     = os.Getenv("TEST_DATABASE_NAME")
	dbPassword = os.Getenv("TEST_DATABASE_PASSWORD")

	details, errDetails   = getDetails()
	database, errDatabase = NewInterface(&details)
)

func getDetails() (*DatabaseDetails, error) {
	if dbAddress == "" {
		dbAddress = defaultDBAddress
	}
	if dbPort == "" {
		dbPort = defaultDBPort
	}
	if dbName == "" {
		dbAddress = defaultDBName
	}
	if dbUsername == "" {
		dbPort = defaultDBUsername
	}
	if dbPassword == "" {
		dbPort = defaultDBPassword
	}

	dbPortInt64, errDbPortInt64 := strconv.ParseInt(dbPort, 10, 64)
	if errDbPortInt64 != nil {
		return nil, errDbPortInt64
	}

	details := CacheDetails{
		Host:     dbAddress,
		Port:     dbPortInt64,
		Name:     dbName,
		Username: dbUsername,
		Password: dbPassword,
	}

	return &details, nil
}

func TestDetailsExists(t *testing.T) {
	if details == nil {
		t.Error("nil parameters should return nil")
	}
	if errDetails != nil {
		t.Error(errDetails.Error())
	}
}

func TestDatabaseInterfaceExists(t *testing.T) {
	if d == nil {
		t.Error("nil parameters should return nil")
	}
	if errCache != nil {
		t.Error(errCache.Error())
	}
}

func TestSetterQueries(t *testing.T) {
	expected := "hello world!"
	statement := &Statement{
		Sql:    "SELECT $1",
		Values: []interface{}{expected},
	}
	results, errResults := Query(statement, nil)

	if results == nil {
		t.Fail()
		t.Logf("there should be sql results!")
		return
	}

	if errResults != nil {
		t.Fail()
		t.Logf(errResults.Error())
		return
	}

	result := (*results)[0][0]
	if result != expected {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", expected, ", found: ", result))
	}
}
