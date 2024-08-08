package rest

import (
    "strings"

    "go-template/utils"
)

type OrderType string

const (
    OrderASC  OrderType = "ASC"
    OrderDESC OrderType = "DESC"
)

func ParseOrderParam(str string) OrderType {
    switch strings.ToUpper(str) {
    case string(OrderASC):
        return OrderASC
    case string(OrderDESC):
        return OrderDESC
    default:
        return OrderASC
    }
}

type ListArgs struct {
    Page    int `query:"page"`
    PerPage int `query:"per_page"`

    Order

    FiltersQuery string `query:"filters"`

    Unscoped bool `query:"unscoped"` // for querying over all rows or only those that are not deleted
}

func (g *ListArgs) Limit() int {
    return g.PerPage
}

func (g *ListArgs) Offset() int {
    return (g.PerPage * g.Page) - g.PerPage
}

func (g *ListArgs) ProcessParams() {
    g.defaults()
    g.Order.Process()
}

func (g *ListArgs) defaults() {
    if g.Page < 1 {
        g.Page = 1
    }

    if g.PerPage == 0 || g.PerPage > 15 {
        g.PerPage = 15
    }
}

func (g *ListArgs) GetOrderFields() []string {
    fields := make([]string, 0, len(g.Orders))
    for key, _ := range g.Orders {
        fields = append(fields, key)
    }

    return fields
}

type Order struct {
    Orders      map[string]OrderType `swaggerignore:"true"` // holds the extracted query params value from OrdersQuery
    OrdersQuery []string             `query:"orders"`
}

func (o *Order) Process() {
    o.Orders = make(map[string]OrderType)
    for _, query := range o.OrdersQuery {
        for _, orderStr := range strings.Split(query, ",") {
            parts := strings.Split(orderStr, "|")
            field := strings.ToLower(parts[0])
            if utils.Strip(field) == "" {
                continue
            }

            orderType := OrderASC
            if len(parts) > 1 {
                orderType = ParseOrderParam(parts[1])
            }

            o.Orders[field] = orderType
        }
    }
}
