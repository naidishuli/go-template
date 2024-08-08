package app

import (
    "errors"
    "fmt"
    "strings"

    "go-template/utils"
    "gorm.io/gorm"
)

func GormErr(tx *gorm.DB, errs ...error) error {
    var dbError error

    if tx != nil && tx.Error != nil {
        dbError = tx.Error
    } else if len(errs) > 0 {
        dbError = errs[0]
    } else {
        return nil
    }

    var err error
    if errors.Is(dbError, gorm.ErrRecordNotFound) {
        err = DBNotFoundErr(tx.Statement.Model)
    } else if errors.Is(dbError, gorm.ErrDuplicatedKey) {
        err = &Error{Code: "database.duplicated_key", Err: dbError}
    } else {
        err = &Error{Code: "database.undefined", Err: dbError}
    }

    // this is to get extra key from the db context
    // val, _ := tx.Get(config.GormSQLKey)
    // str, ok := val.(string)
    // if ok {
    // 	if err.Data == nil {
    // 		err.Data = make(map[string]interface{})
    // 		err.Data[config.GormSQLKey] = str
    // 	}
    // }

    return err
}

func DBNotFoundErr(typ interface{}) error {
    resourceType := fmt.Sprintf("%s", strings.ToLower(utils.TypeName(typ)))

    return &Error{
        Code:    "database.not_found",
        Message: fmt.Sprintf("%s not found", resourceType),
    }
}
