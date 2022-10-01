package common

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

type SQLModel struct {
	Id       uuid.UUID  `json:"id" gorm:"column:id; default:uuid_generate_v3();"`
	Status   int32      `json:"status" gorm:"column:status; default:1;"`
	CreateAt *time.Time `json:"create_at,omitempty" gorm:"create_at; default:now()"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"update_at; default:now()"`
}

func (m *SQLModel) GenUID(dbType int) {
}
