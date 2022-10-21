package sql

import "gorm.io/gorm"

// Create creates a record
func Create[T any](tx *gorm.DB, model *T) (*T, error) {
	err := tx.Create(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// Update updates a record
func Update[T any](tx *gorm.DB, model *T) (*T, error) {
	err := tx.Save(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// Delete deletes a record
func Delete[T any](tx *gorm.DB, model *T) error {
	return tx.Delete(model).Error
}

// DeleteBy deletes a record by statement
func DeleteBy[T any](tx *gorm.DB, stm *Statement) error {
	tx = consumeStatement(tx, stm)
	return tx.Delete(new(T)).Error
}
