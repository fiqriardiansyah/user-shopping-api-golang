package helper

import (
	"fmt"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	*gorm.DB
}

func (r *Repository[T]) Create(tx *gorm.DB, data *T) error {
	return tx.Create(data).Error
}

func (r *Repository[T]) FindByColumName(tx *gorm.DB, entity *T, str string, col string) error {
	return tx.Where(fmt.Sprintf("%s = ?", col), str).First(entity).Error
}

func (r *Repository[T]) Update(tx *gorm.DB, data *T) error {
	return tx.Save(data).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}
