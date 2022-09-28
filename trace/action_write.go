package trace

import (
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Create creates a new record and a new trace record
func Create(tx *gorm.DB, m IModel, h ITrace) error {
	m.InitVersion()

	err := tx.Omit(clause.Associations).Create(m).Error
	if err != nil {
		return err
	}

	_updateModelFromModel(m, h)

	h.SetRecordID(m.GetID())
	h.SetTraceInfo("create")

	return tx.Omit(clause.Associations).Create(h).Error
}

// Update updates an existing record and creates a new trace record
func Update(tx *gorm.DB, m IModel, h ITrace) error {
	m.IncrementVersion()

	err := tx.Omit(clause.Associations).Save(m).Error
	if err != nil {
		return err
	}

	_updateModelFromModel(m, h)

	h.SetRecordID(m.GetID())
	h.SetTraceInfo("update")

	return tx.Omit(clause.Associations).Save(h).Error
}

// Delete delete an existing record and creates a new trace record
func Delete(tx *gorm.DB, m IModel, h ITrace) error {
	err := tx.Omit(clause.Associations).Delete(m).Error
	if err != nil {
		return err
	}

	_updateModelFromModel(m, h)

	h.SetRecordID(m.GetID())
	h.SetTraceInfo("delete")

	return tx.Omit(clause.Associations).Save(h).Error
}

func _updateModelFromModel(from interface{}, to interface{}) {
	j, err := json.Marshal(from)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(j, to)
	if err != nil {
		panic(err)
	}
}
