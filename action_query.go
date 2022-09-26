package sql

import "gorm.io/gorm"

type Pagination struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

type Sort struct {
	By  string `json:"by" form:"by"`
	Asc bool   `json:"asc" form:"asc"`
}

// FindOne finds one record
func FindOne[T any](tx *gorm.DB, cls *Clause) (*T, error) {
	tx = _buildClause(tx, cls)
	result := new(T)
	err := tx.First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll finds all records
func FindAll[T any](tx *gorm.DB, cls *Clause, p *Pagination, s *Sort) ([]T, error) {
	tx = _buildClause(tx, cls)
	tx = _buildFindAllOption(tx, p, s)
	results := make([]T, 0)
	err := tx.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Count counts records
func Count(tx *gorm.DB, cls *Clause) (int64, error) {
	tx = _buildClause(tx, cls)
	var count int64
	err := tx.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func _buildClause(tx *gorm.DB, clause *Clause) *gorm.DB {
	if clause == nil {
		return tx
	}
	s, v := clause.Build()
	return tx.Where(s, v...)
}

func _buildFindAllOption(tx *gorm.DB, p *Pagination, s *Sort) *gorm.DB {
	if p != nil {
		if p.Page > 0 {
			tx = tx.Offset(p.Page - 1*p.Size)
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
