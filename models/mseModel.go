package models

import (
	"github.com/google/uuid"
)

type MseProfile struct {
	ID                string    `gorm:"type:char(32);primaryKey" json:"id_mse"`
	CollaborationWith uuid.UUID `gorm:"type:char(36);default:null" json:"team_id"`
	Role              string    `gorm:"not null;type:varchar(10)" json:"role"`
	OwnerName         string    `gorm:"not null;type:varchar(50)" json:"owner_name"`
	MseName           string    `gorm:"not null;type:varchar(50)" json:"mse_name"`
	MseType           string    `gorm:"not null;type:varchar(50)" json:"mse_type"`
	MseSince          string    `gorm:"not null;type:char(4)" json:"mse_since"`
	Team              Team      `gorm:"foreignKey:CollaborationWith;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

type Problem struct {
	ID             uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id_problem"`
	MseID          string     `gorm:"type:char(32);default:null" json:"mse_id"`
	TagID          int32      `gorm:"type:int;default:null" json:"tag_id"`
	Like           int64      `gorm:"type:int" json:"like"`
	CommentCount   int64      `gorm:"type:int" json:"comment_count"`
	Problem        string     `gorm:"not null;type:varchar(50)" json:"problem"`
	Description    string     `gorm:"not null;type:text" json:"description"`
	ProblemCreated string     `gorm:"not null;type:datetime" json:"problem_created"`
	Income         string     `gorm:"type:varchar(50)" json:"income"`
	Employees      int32      `gorm:"type:int" json:"employee"`
	LastSale       string     `gorm:"type:varchar(50)" json:"last_sale"`
	MediaSocial    string     `gorm:"type:varchar(50)" json:"media_social"`
	Goals          string     `gorm:"type:varchar(70)" json:"goals"`
	Address        string     `gorm:"type:varchar(50)" json:"address"`
	MseProfile     MseProfile `gorm:"foreignKey:MseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Tags           Tags       `gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}
