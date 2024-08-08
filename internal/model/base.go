package model

import (
    "database/sql/driver"
    "fmt"
    "reflect"
    "strings"
    "time"

    "bets/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type Base struct {
    isNew         bool
    isUpdated     bool
    updateColumns []string
}

func (b *Base) IsNew() bool {
    return b.isNew
}

func (b *Base) IsUpdated() bool {
    return b.isUpdated
}

func (b *Base) AfterCreate(tx *gorm.DB) (err error) {
    b.isNew = true
    return nil
}

func (b *Base) AfterUpdate(tx *gorm.DB) (err error) {
    b.isUpdated = true
    return nil
}

func (b *Base) AddUpdateColumns(columns ...string) {
    b.updateColumns = append(b.updateColumns, columns...)
}

func (b *Base) GetUpdateColumns() []string {
    return b.updateColumns
}

func (b *Base) CleanUpdateColumns() {
    b.updateColumns = nil
}

func (p *Base) SelectColumns(db *gorm.DB) *gorm.DB {
    if len(p.updateColumns) == 0 {
        return db.Select("")
    }

    if len(p.updateColumns) == 1 {
        return db.Select(p.updateColumns[0])
    }

    return db.Select(p.updateColumns[0], p.updateColumns[1:])
}

func (b *Base) SetUpdateColumns(columns []string) {
    b.updateColumns = columns
}

func (b *Base) UpdateColumnExist(column string) bool {
    for _, uc := range b.updateColumns {
        if uc == column {
            return true
        }
    }

    return false
}

// ==========================================================================================

type Timestamp struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

type TimestampWithDelete struct {
    DeletedAt gorm.DeletedAt
    CreatedAt time.Time
    UpdatedAt time.Time
}

// ==========================================================================================

func Unscoped(db *gorm.DB) *gorm.DB {
    return db.Unscoped()
}

func OmitAssociations(db *gorm.DB) *gorm.DB {
    return db.Omit(clause.Associations)
}

func Omit(assocs ...string) func(*gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Omit(assocs...)
    }
}

func WithJoins(assoc string, args ...any) func(*gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Joins(assoc, args...)
    }
}

func WithPreloads(assoc string, args ...any) func(*gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Preload(assoc, args...)
    }
}

func OnOrganization(orgID uint, table ...string) func(db *gorm.DB) *gorm.DB {
    query := "organization_id = ?"
    if len(table) > 0 {
        query = fmt.Sprintf("%s.%s", table[0], query)
    }

    return func(db *gorm.DB) *gorm.DB {
        return db.Where(query, orgID)
    }
}

// ==========================================================================================

// StringArray save an array with strings to a varchar database column by joining them with ',' separator
type StringArray []string

func (f *StringArray) Scan(value any) error {
    if value == nil {
        return nil
    }

    data, ok := value.(string)
    if !ok {
        return fmt.Errorf("failed to extract StringArray data value: %+v", value)
    }

    if data == "" {
        return nil
    }

    *f = strings.Split(data, ",")
    return nil
}

func (f StringArray) Value() (driver.Value, error) {
    if len(f) == 0 {
        return nil, nil
    }

    return strings.Join(f, ","), nil
}

// ==========================================================================================

func NonZeroFields(value any) []string {
    val := reflect.ValueOf(value)
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    var nonZeroFields []string
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)

        if !field.IsZero() {
            nonZeroFields = append(nonZeroFields, utils.ToSnakeCase(val.Type().Field(i).Name))
        }
    }

    return nonZeroFields
}
