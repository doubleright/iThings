package stores

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Time struct {
	CreatedTime time.Time    `gorm:"column:created_time;index;sort:desc;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedTime time.Time    `gorm:"column:updated_time;autoUpdateTime;default:CURRENT_TIMESTAMP;NOT NULL"`
	DeletedTime sql.NullTime `gorm:"column:deleted_time"`
}

type NoDelTime struct {
	CreatedTime time.Time `gorm:"column:created_time;index;sort:desc;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;autoUpdateTime;default:CURRENT_TIMESTAMP;NOT NULL"`
}

type DeletedTime = gorm.DeletedAt

type SoftTime struct {
	CreatedTime time.Time   `gorm:"column:created_time;index;sort:desc;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedTime time.Time   `gorm:"column:updated_time;autoUpdateTime;default:CURRENT_TIMESTAMP;NOT NULL"`
	DeletedTime DeletedTime `gorm:"column:deleted_time;index"`
}
