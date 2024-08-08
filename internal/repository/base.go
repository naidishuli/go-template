package repository

import (
    "fmt"

    "bets/internal/app"
    "bets/pkg/rest"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type Base struct {
    db *gorm.DB
}

func NewBase(dep app.App) *Base {
    return &Base{db: dep.DB().Session(&gorm.Session{})}
}

func (b *Base) WithCtx(ctx app.BaseRepoCtx) app.BaseRepo {
    return &Base{
        db: ctx.DB(b.db),
    }
}

func (b *Base) DB() *gorm.DB {
    return b.db
}

// ==================== Builder methods ==============================

func (b *Base) Clauses(conds ...clause.Expression) app.BaseRepo {
    return &Base{
        db: b.db.Clauses(conds...),
    }
}

func (b *Base) Joins(query string, args ...any) app.BaseRepo {
    return &Base{
        db: b.db.Joins(query, args...),
    }
}

func (b *Base) Model(value any) app.BaseRepo {
    return &Base{
        db: b.db.Model(value),
    }
}

func (b *Base) Select(query any, args ...any) app.BaseRepo {
    return &Base{
        db: b.db.Select(query, args...),
    }
}

func (b *Base) Omit(columns ...string) app.BaseRepo {
    return &Base{
        db: b.db.Omit(columns...),
    }
}

func (b *Base) Scopes(funcs ...func(*gorm.DB) *gorm.DB) app.BaseRepo {
    return &Base{
        db: b.db.Scopes(funcs...),
    }
}

func (b *Base) Where(query any, args ...any) app.BaseRepo {
    return &Base{
        db: b.db.Where(query, args...),
    }
}

func (b *Base) Unscoped() app.BaseRepo {
    return &Base{
        db: b.db.Unscoped(),
    }
}

// ==================== Finisher methods ==============================

func (b *Base) Create(value any) error {
    rtx := b.db.Create(value)
    return app.GormErr(rtx)
}

func (b *Base) Delete(value any, cond ...any) error {
    rtx := b.db.Delete(value, cond...)
    return app.GormErr(rtx)
}

func (b *Base) Find(dest any, cond ...any) error {
    rtx := b.db.Find(dest, cond...)
    return app.GormErr(rtx)
}

func (b *Base) First(dest any, cond ...any) error {
    rtx := b.db.First(dest, cond...)
    return app.GormErr(rtx)
}

func (b *Base) FirstOrCreate(dest any, cond ...any) error {
    rtx := b.db.FirstOrCreate(dest, cond...)
    return app.GormErr(rtx)
}
func (b *Base) Update(column string, value any) (int64, error) {
    rtx := b.db.Update(column, value)
    return rtx.RowsAffected, app.GormErr(rtx)
}

func (b *Base) Updates(values any) error {
    rtx := b.db.Updates(values)
    return app.GormErr(rtx)
}

func (b *Base) Save(value any) error {
    rtx := b.db.Save(value)
    return app.GormErr(rtx)
}

func (b *Base) Transaction(fun func(repo app.BaseRepo) error) error {
    return b.db.Transaction(func(tx *gorm.DB) error {
        return fun(&Base{db: tx})
    })
}

func (b *Base) TransactionCtx(ctx app.Context) app.Context {
    return ctx.WithDB(b.db)
}

func (b *Base) Paginated(output any, args rest.PaginatedArgs) (int64, error) {
    var totalCount int64

    query := b.db.Model(args.Model)

    if args.ListArgs.Unscoped {
        query = query.Unscoped()
    }

    if args.BeforeCountHook != nil {
        query = args.BeforeCountHook(query)
    }

    // count results
    countQuery := query.Model(args.Model).Session(&gorm.Session{})
    rtx := countQuery.Session(&gorm.Session{}).Count(&totalCount)
    if rtx.Error != nil {
        return 0, app.GormErr(rtx)
    }

    // add orders
    for field, order := range args.ListArgs.Orders {
        query = query.Order(fmt.Sprintf("%s %s", field, order))
    }

    if args.AfterCountHook != nil {
        query = args.AfterCountHook(query)
    }

    // query results
    rtx = query.Offset(args.ListArgs.Offset()).Limit(args.ListArgs.Limit()).Find(output)
    return totalCount, app.GormErr(rtx)
}
