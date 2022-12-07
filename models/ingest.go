package models

import (
	"github.com/jackc/pgtype"
)

type Tabler interface {
	TableName() string
}

func (Ingest) TableName() string {
	return "ingest"
}

type Ingest struct {
	ID        int          `gorm:"autoIncrement" json:"id"`
	CaptureID string       `gorm:"primaryKey" json:"capture_id"`
	Capsrc    string       `json:"capture_src"`
	Date      string       `json:"date"`
	StartName string       `json:"start_name"`
	StopName  string       `json:"stop_name"`
	Sha1      string       `json:"sha1"`
	Line      pgtype.JSONB `gorm:"type:jsonb" json:"line"`
	Original  pgtype.JSONB `gorm:"type:jsonb" json:"original"`
	Proxy     pgtype.JSONB `gorm:"type:jsonb" json:"proxy"`
	Wfstatus  pgtype.JSONB `gorm:"type:jsonb" json:"wfstatus"`
}
