package models

import (
	"github.com/google/uuid"
)

type StudentProfile struct {
	ID          string    `gorm:"type:char(32);primaryKey" json:"id_student"`
	TeamID      uuid.UUID `gorm:"type:char(36);default:null" json:"team_id"`
	TagID       int32     `gorm:"type:int;default:null" json:"tag_id"`
	StudentName string    `gorm:"not null; type:varchar(50)" json:"student_name"`
	DateOfBirth string    `gorm:"not null;type:datetime" json:"date_of_birth"`
	Role        string    `gorm:"not null;type:varchar(10)" json:"role"`
	IsLeader    bool      `gorm:"default:false" json:"is_leader"`
	Major       string    `gorm:"not null;type:varchar(50)" json:"major"`
	University  string    `gorm:"not null;type:varchar(50)" json:"university"`
	ClassOf     string    `gorm:"not null;type:varchar(50)" json:"class_of"`
	Team        Team      `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Tags        Tags      `gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}
