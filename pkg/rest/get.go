package rest

import "gorm.io/gorm"

type GetArgs struct {
    ID     uint `params:"id" validate:"required"`
    Joins  []string
    Wheres []WhereCondition
}

func (g *GetArgs) DbDeps(db *gorm.DB) *gorm.DB {
    for _, join := range g.Joins {
        db = db.Joins(join)
    }

    for _, where := range g.Wheres {
        db = db.Where(where.Query, where.Values...)
    }

    return db
}

type GetAllArgs struct {
}
