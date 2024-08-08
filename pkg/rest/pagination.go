package rest

import "gorm.io/gorm"

type PaginatedArgs struct {
	Model           any // should be a pointer to a mode struct
	ListArgs        ListArgs
	BeforeCountHook func(db *gorm.DB) *gorm.DB
	AfterCountHook  func(db *gorm.DB) *gorm.DB
}
