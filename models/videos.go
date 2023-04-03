package models

import (
	"gorm.io/gorm"
	"time"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	gorm.Model
	ID       int64      `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Video    *[]byte    `gorm:"column:video;type:blob" json:"video"`
	CutFirst *time.Time `gorm:"column:CutFirst;type:datetime(6)" json:"CutFirst"`
	CutLate  *time.Time `gorm:"column:CutLate;type:datetime(6)" json:"CutLate"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
