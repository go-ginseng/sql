package sql

import "gorm.io/gorm"

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Size  int   `json:"size" form:"size"`
	Total int64 `json:"total"`
}

type Sort struct {
	By  string `json:"by" form:"by"`
	Asc bool   `json:"asc" form:"asc"`
}

// FindOne finds one record
func FindOne[T any](tx *gorm.DB, stm *Statement) (*T, error) {
	tx = consumeStatement(tx, stm)
	result := new(T)
	err := tx.First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnscopedFindOne finds one record including soft deleted records
func UnscopedFindOne[T any](tx *gorm.DB, stm *Statement) (*T, error) {
	tx = consumeStatement(tx, stm)
	result := new(T)
	err := tx.Unscoped().First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll finds all records
func FindAll[T any](tx *gorm.DB, stm *Statement, p *Pagination, s *Sort) ([]T, error) {
	tx = consumeStatement(tx, stm)
	tx = consumePaginationAndSort(tx, p, s)
	results := make([]T, 0)
	err := tx.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// UnscopedFindAll finds all records including soft deleted records
func UnscopedFindAll[T any](tx *gorm.DB, stm *Statement, p *Pagination, s *Sort) ([]T, error) {
	tx = consumeStatement(tx, stm)
	tx = consumePaginationAndSort(tx, p, s)
	results := make([]T, 0)
	err := tx.Unscoped().Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Count counts records
func Count[T any](tx *gorm.DB, stm *Statement) (int64, error) {
	tx = consumeStatement(tx, stm)
	var count int64
	err := tx.Model(new(T)).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// UnscopedCount counts records including soft deleted records
func UnscopedCount[T any](tx *gorm.DB, stm *Statement) (int64, error) {
	tx = consumeStatement(tx, stm)
	var count int64
	err := tx.Unscoped().Model(new(T)).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func consumeStatement(tx *gorm.DB, stm *Statement) *gorm.DB {
	if stm == nil {
		return tx
	}
	s, v := stm.Build()
	return tx.Where(s, v...)
}

func consumePaginationAndSort(tx *gorm.DB, p *Pagination, s *Sort) *gorm.DB {
	if p != nil {
		if p.Page > 0 {
			tx = tx.Offset((p.Page - 1) * p.Size)
		}
		if p.Size > 0 {
			tx = tx.Limit(p.Size)
		}
	}

	if s != nil {
		if s.By != "" {
			var order string
			if s.Asc {
				order = "ASC"
			} else {
				order = "DESC"
			}
			tx = tx.Order(s.By + " " + order)
		}
	}

	return tx
}
