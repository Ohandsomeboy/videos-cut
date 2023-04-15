package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	gorm.Model
	ID       int64     `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Video    []byte    `gorm:"column:video;type:blob" json:"video"`
	Name     string    `gorm:"column:name;type:varchar(11)" json:"name"`
	Path     string    `gorm:"column:path;type:varchar(11)" json:"path"`
	CutFirst time.Time `gorm:"column:CutFirst;type:datetime(6)" json:"CutFirst"`
	CutLate  time.Time `gorm:"column:CutLate;type:datetime(6)" json:"CutLate"`
}

// TableName Video's table name
func (talbe Video) TableName() string {
	return TableNameVideo
}

func GetVideosList() {
	data := make([]*Video, 0)
	DB.Find(&data)
	for _, v := range data {
		fmt.Printf("Problem ==> %v \n", v)
	}
}
