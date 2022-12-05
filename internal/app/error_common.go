package app

import (
	"errors"
	"fmt"
	"go-template/internal/config"
	"go-template/utils"
	"gorm.io/gorm"
)

func NotFoundError(typ interface{}, err error, data ErrData) *Error {
	if err == nil {
		err = errors.New("not found")
	}

	e := NewError(NotFoundErr, err, data)

	resourceType := fmt.Sprintf(" %s", utils.TypeName(typ))
	e.Message = fmt.Sprintf("Not found%s", resourceType)

	return e
}

func GormErr(tx *gorm.DB, data ErrData) error {
	if tx.Error == nil {
		return nil
	}

	var err *Error
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		err = NotFoundError(tx.Statement.Model, tx.Error, data)
	} else {
		err = NewError(DatabaseErr, tx.Error, data)
	}

	val, _ := tx.Get(config.GormSQLKey)
	str, ok := val.(string)
	if ok {
		if err.ErrData == nil {
			err.ErrData = make(map[string]interface{})
			err.ErrData[config.GormSQLKey] = str
		}
	}

	return err
}

func IsErr(err interface{}, code ErrCode) bool {
	appErr, ok := err.(*Error)
	if !ok {
		return false
	}

	return appErr.Code == code
}
