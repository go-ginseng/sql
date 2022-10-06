package trace

import (
	"time"

	"github.com/go-ginseng/sql"
	"gorm.io/gorm"
)

type Model struct {
	sql.Model
	Version uint `gorm:"version" json:"version"`
}

type IModel interface {
	GetID() uint
	InitVersion()
	IncrementVersion()
}

func (m *Model) GetID() uint {
	return m.ID
}

func (m *Model) InitVersion() {
	m.Version = 1
}

func (m *Model) IncrementVersion() {
	m.Version++
}

type Trace struct {
	Model
	RecordID    uint      `gorm:"not null" json:"record_id"` // equal to the primary key of the record
	TraceTime   time.Time `gorm:"not null" json:"trace_time"`
	TraceAction string    `gorm:"not null" json:"trace_action"`
}

type ITrace interface {
	SetRecordID(id uint)
	SetTraceInfo(action string)
}

func (m *Trace) SetRecordID(id uint) {
	m.RecordID = id
}

func (m *Trace) SetTraceInfo(action string) {
	m.ID = 0
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	m.DeletedAt = gorm.DeletedAt{}
	m.TraceTime = time.Now()
	m.TraceAction = action
}
