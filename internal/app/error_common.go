package app

import (
	"errors"
	"fmt"
	"go-template/internal/config"
	"go-template/utils"
	"gorm.io/gorm"
	"strings"
)

func GormErr(tx *gorm.DB, data ErrData) error {
	if tx.Error == nil {
		return nil
	}

	var err *Error
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		err = NotFoundError(tx.Statement.Model, data)
	} else if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
		err = &Error{Code: ErrDuplicateKey, Err: tx.Error, Data: data}
	} else {
		err = &Error{Code: ErrDatabase, Err: tx.Error, Data: data}
	}

	val, _ := tx.Get(config.GormSQLKey)
	str, ok := val.(string)
	if ok {
		if err.Data == nil {
			err.Data = make(map[string]interface{})
			err.Data[config.GormSQLKey] = str
		}
	}

	return err
}

func NotFoundError(typ interface{}, data ErrData) *Error {
	resourceType := fmt.Sprintf("%s", strings.ToLower(utils.TypeName(typ)))

	return &Error{
		Code:    ErrNotFound,
		Message: fmt.Sprintf("%s not found", resourceType),
		Data:    data,
	}
}
