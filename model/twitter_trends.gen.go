// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTwitterTrend = "twitter_trends"

// TwitterTrend mapped from table <twitter_trends>
type TwitterTrend struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Rank      int32     `gorm:"column:rank" json:"rank"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName TwitterTrend's table name
func (*TwitterTrend) TableName() string {
	return TableNameTwitterTrend
}