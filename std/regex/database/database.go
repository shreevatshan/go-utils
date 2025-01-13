package database

import (
	"regexp"
	"strings"
)

const (
	dbOperationPattern1 = `^(?i)(set.*)?(\s)*(select).*(\s)+(from)(\s)+(\x60)?(['"\[]?\w+['"\]]?(?:\.[\[]?\w+[\]]?)?)(\x60)?.*`
	dbOperationPattern2 = `^(?i)(set.*)?(\s)*(insert|delete)(\s)+(from|into)(\s)+(\x60)?(['"\[]?\w+['"\]]?(?:\.[\[]?\w+[\]]?)?)(\x60)?.*`
	dbOperationPattern3 = `^(?i)(set.*)?(\s)*(update)(\s)+(\x60)?(['"\[]?\w+['"\]]?(?:\.[\[]?\w+[\]]?)?)(\x60)?.*`
	dbOperationPattern4 = `^(?i)(set.*)?(\s)*(create|drop)(\s)+(procedure|database|table|keyspace)(\s)+((IF NOT EXISTS|IF EXISTS)(\s)+)?(\x60)?(['"\[]?\w+['"\]]?(?:\.[\[]?\w+[\]]?)?)(\x60)?.*`
	dbOperationPattern5 = `^(?i)(set.*)?(\s)*(alter)+(\s)+(table)(\s)+(\x60)?(['"\[]?\w+['"\]]?(?:\.[\[]?\w+[\]]?)?)(\x60)?.*`
	dbOperationPattern6 = `^(?i)(set.*)?(\s)*(call|execute)(\s)+(\x60)?(['"\[]?\w+['"\]]?(?:\.[\[]?\w+[\]]?)?)(\x60)?.*`
	dbOperationPattern7 = `^(?i)NoSQL:\s(\w+)(\s)(from|into|in|key)(\s)+(\x60)?(['"\[]?\w+['"\]]?(?:\.[\[]?\w+[\]]?)?)(\x60)?.*`

	sqlParametrizePattern1 = "(\\d+)"
	sqlParametrizePattern2 = "'(.*?)'"
)

var (
	databaseOperationRegexes map[string]*regexp.Regexp
	sqlParametrizeRegexes    map[string]*regexp.Regexp
)

func init() {

	//DATABASE REGEX
	databaseOperationRegexes = make(map[string]*regexp.Regexp)

	databaseOperationRegexes["select"], _ = regexp.Compile(dbOperationPattern1)

	databaseOperationRegexes["insert-delete"], _ = regexp.Compile(dbOperationPattern2)

	databaseOperationRegexes["update"], _ = regexp.Compile(dbOperationPattern3)

	databaseOperationRegexes["create-drop"], _ = regexp.Compile(dbOperationPattern4)

	databaseOperationRegexes["alter"], _ = regexp.Compile(dbOperationPattern5)

	databaseOperationRegexes["call-execute"], _ = regexp.Compile(dbOperationPattern6)

	databaseOperationRegexes["nosql"], _ = regexp.Compile(dbOperationPattern7)

	//SQLPARAMETRIZE REGEX
	sqlParametrizeRegexes = make(map[string]*regexp.Regexp)

	sqlParametrizeRegexes["parametrize-integer"], _ = regexp.Compile(sqlParametrizePattern1)

	sqlParametrizeRegexes["parametrize-withinquote"], _ = regexp.Compile(sqlParametrizePattern2)

}

func GetDatabaseOperationRegexes() map[string]*regexp.Regexp {
	return databaseOperationRegexes
}

func GetSQLParametrizeRegexes() map[string]*regexp.Regexp {
	return sqlParametrizeRegexes
}

type Operation struct {
	Name  string
	Table string
}

type Operations []Operation

func ParseStatement(query string) Operations {
	var dbOperations Operations
	queries := strings.Split(query, ";")
	for _, query := range queries {

		var db Operation
		var regex *regexp.Regexp

		// the order is chosen to match the most common queries first

		regex = databaseOperationRegexes["select"]
		s := regex.FindAllStringSubmatch(query, -1)
		if len(s) != 0 {
			db.Name = s[0][3]
			db.Table = s[0][8]
		}

		regex = databaseOperationRegexes["insert-delete"]
		s = regex.FindAllStringSubmatch(query, -1)
		if len(s) != 0 {
			db.Name = s[0][3]
			db.Table = s[0][8]
		}

		regex = databaseOperationRegexes["update"]
		s = regex.FindAllStringSubmatch(query, -1)
		if len(s) != 0 {
			db.Name = s[0][3]
			db.Table = s[0][6]
		}

		regex = databaseOperationRegexes["create-drop"]
		s = regex.FindAllStringSubmatch(query, -1)
		if len(s) != 0 {
			db.Name = s[0][3]
			db.Table = s[0][11]
		}

		regex = databaseOperationRegexes["alter"]
		s = regex.FindAllStringSubmatch(query, -1)
		if len(s) != 0 {
			db.Name = s[0][3]
			db.Table = s[0][8]
		}

		regex = databaseOperationRegexes["call-execute"]
		s = regex.FindAllStringSubmatch(query, -1)
		if len(s) != 0 {
			db.Name = s[0][3]
			db.Table = s[0][6]
		}

		regex = databaseOperationRegexes["nosql"]
		s = regex.FindAllStringSubmatch(query, -1)
		if len(s) != 0 {
			db.Name = s[0][1]
			db.Table = s[0][6]
		}

		if (db.Name != "") && (db.Table != "") {
			dbOperations = append(dbOperations, db)
		}

	}
	return dbOperations
}

func ParametrizeStatement(query string) string {
	var parametrizedQuery string

	for _, regex := range sqlParametrizeRegexes {
		parametrizedQuery = regex.ReplaceAllString(query, "?")
	}

	return parametrizedQuery
}

func (dbOperations Operations) ToString() string {
	var operationname string
	var tablename string
	for _, dboperation := range dbOperations {
		if operationname != "" && tablename != "" {
			operationname += "," + dboperation.Name
			tablename += "," + dboperation.Table
		} else {
			operationname = dboperation.Name
			tablename = dboperation.Table
		}
	}
	return operationname + "/" + tablename
}
