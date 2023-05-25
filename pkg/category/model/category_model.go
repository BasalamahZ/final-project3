package model

import (
	"final-project3/pkg/task/model"
	"time"
)

// Items        []modelItem.Item `gorm:"Foreignkey:OrderId;association_foreignkey:OrderId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// OrderId      int64            `json:"order_id" gorm:"primaryKey;UNIQUE"`
// CreatedAt    time.Time        `json:"created_at" gorm:"autoCreateTime:true"`
// UpdatedAt    time.Time        `json:"updated_at" gorm:"autoUpdateTime:true"`

type Category struct {
	Id        int64        `json:"id" gorm:"primaryKey;UNIQUE"`
	Type      string       `json:"type" binding:"required"`
	Tasks     []model.Task `json:"tasks" gorm:"foreignKey:CategoryId"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime:true"`
}
