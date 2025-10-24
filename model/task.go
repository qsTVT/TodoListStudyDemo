package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User   User   `gorm:"Foreignkey:Uid"`
	Uid    uint   `gorm:"not null"`
	Title  string `gorm:"not null"`
	Status int    `gorm:"default:0"` //0未完成，1已完成
	//	View      int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64  //备忘录开始时间
	EndTime   int64  //备忘录结束时间
}
