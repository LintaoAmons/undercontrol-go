//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type AccountHistory struct {
	ID           int32 `sql:"primary_key"`
	Name         string
	Amount       *string
	CurrencyCode *string
	CreatedAt    *time.Time
	CreatedBy    *string
	UpdatedAt    *time.Time
	UpdatedBy    *string
}