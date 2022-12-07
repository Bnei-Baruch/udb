package models

import "github.com/jackc/pgtype"
import "github.com/lib/pq"

func (Trimmer) TableName() string {
	return "trimmer"
}

type Trimmer struct {
	ID        int             `gorm:"autoIncrement" json:"id"`
	TrimID    string          `gorm:"primaryKey" json:"trim_id"`
	Date      string          `json:"date"`
	FileName  string          `json:"file_name"`
	Inpoints  pq.Float32Array `gorm:"type:numeric" json:"inpoints"`
	Outpoints pq.Float32Array `gorm:"type:numeric" json:"outpoints"`
	Parent    pgtype.JSONB    `json:"parent"`
	Line      pgtype.JSONB    `gorm:"type:jsonb" json:"line"`
	Original  pgtype.JSONB    `gorm:"type:jsonb" json:"original"`
	Proxy     pgtype.JSONB    `gorm:"type:jsonb" json:"proxy"`
	Wfstatus  pgtype.JSONB    `gorm:"type:jsonb" json:"wfstatus"`
}
