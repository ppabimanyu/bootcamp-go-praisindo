package database

import (
	"fmt"

	"gorm.io/gorm"
)

func Order(order, orderBy string, query *gorm.DB) *gorm.DB {
	if order != "" && orderBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", orderBy, order))
	}
	return query
}
