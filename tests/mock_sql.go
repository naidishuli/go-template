package tests

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
)

var EverythingMatcher sqlmock.QueryMatcher = sqlmock.QueryMatcherFunc(
	func(expectedSQL, actualSQL string) error {
		return nil
	},
)

type MockSqlSpecs struct {
	IsTransaction bool
	IsExec        bool

	QueryResult *sqlmock.Rows
	ExecResult  driver.Result

	MatchSQL string
	MockDB   sqlmock.Sqlmock
	SubSpecs []MockSqlSpecs
	Err      error
}

func SetSqlSpecsExpectation(specs []MockSqlSpecs, mockSqlDB sqlmock.Sqlmock) {
	for _, spec := range specs {
		spec.SetExpectation(mockSqlDB)
	}
}

func (m MockSqlSpecs) SetExpectation(sqlDB sqlmock.Sqlmock) {
	if m.MockDB == nil {
		m.MockDB = sqlDB
	}

	if m.IsTransaction {
		m.MockDB.ExpectBegin()
	}

	for i, subSpec := range m.SubSpecs {
		if subSpec.IsTransaction {
			subSpec.MockDB.ExpectBegin()
		}

		if len(subSpec.SubSpecs) > 0 {
			subSpec.SetExpectation(m.MockDB)
			continue
		}

		last := i == len(subSpec.SubSpecs)-1

		if subSpec.IsExec {
			setExec(subSpec, last)
		} else {
			setQuery(subSpec, last)
		}
	}

	if len(m.SubSpecs) == 0 {
		if m.IsExec {
			setExec(m, m.IsTransaction)
		} else {
			setQuery(m, m.IsTransaction)
		}
	}
}

func setQuery(spec MockSqlSpecs, last bool) {
	if spec.QueryResult != nil {
		spec.MockDB.ExpectQuery(``).WillReturnRows(spec.QueryResult)
		if last {
			spec.MockDB.ExpectCommit()
		}
		return
	}

	if spec.Err != nil {
		spec.MockDB.ExpectQuery(``).WillReturnError(spec.Err)
		if last {
			spec.MockDB.ExpectRollback()
		}
		return
	}

	spec.MockDB.ExpectCommit()
}

func setExec(spec MockSqlSpecs, last bool) {
	if spec.ExecResult != nil {
		spec.MockDB.ExpectExec(``).WillReturnResult(spec.ExecResult)
		if last {
			spec.MockDB.ExpectCommit()
		}
		return
	}

	if spec.Err != nil {
		spec.MockDB.ExpectExec(``).WillReturnError(spec.Err)
		if last {
			spec.MockDB.ExpectRollback()
		}
		return
	}

	spec.MockDB.ExpectCommit()
}
