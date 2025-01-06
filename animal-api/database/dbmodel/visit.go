package dbmodel

import (
	"context"
	"time"

	"feldrise.com/animal-api/pkg/model"
	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model

	Date  time.Time `gorm:"not null"`
	CatID uint      `gorm:"not null"`

	// Foreign object
	Cat Cat `gorm:"foreignKey:CatID"`
}

func (visit *Visit) ToModel() *model.Visit {
	var cat *model.Cat

	if visit.CatID != 0 {
		cat = visit.Cat.ToModel()
	}

	return &model.Visit{
		Date: visit.Date,
		Cat:  cat,
	}
}

type VisitsFieldsToInclude struct {
	Cat bool
}

type VisitsFilter struct {
	CatID uint
}

type VisitsRepository interface {
	FindByID(id uint, fields *VisitsFieldsToInclude) (*Visit, error)
	FindAll(filter *VisitsFilter, fields *VisitsFieldsToInclude) ([]*Visit, error)
	Create(visit *Visit) (*Visit, error)
	Update(visit *Visit) (*Visit, error)
	Delete(visit *Visit) error
}

type visitsRepository struct {
	db *gorm.DB
}

func NewVisitsRepository(db *gorm.DB) VisitsRepository {
	return &visitsRepository{
		db: db,
	}
}

func (r *visitsRepository) FindByID(id uint, fields *VisitsFieldsToInclude) (*Visit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var visit Visit
	tx := r.db.WithContext(ctx).Model(&visit)

	if fields != nil && fields.Cat {
		tx = tx.Preload("Cat")
	}

	err := tx.Where("id = ?", id).First(&visit).Error

	if err != nil {
		return nil, err
	}

	return &visit, nil
}

func (r *visitsRepository) FindAll(filter *VisitsFilter, fields *VisitsFieldsToInclude) ([]*Visit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var visits []*Visit
	tx := r.db.WithContext(ctx).Model(&visits)

	if fields != nil && fields.Cat {
		tx = tx.Preload("Cat")
	}

	if filter != nil {
		tx = tx.Where("cat_id = ?", filter.CatID)
	}

	err := tx.Find(&visits).Error

	if err != nil {
		return nil, err
	}

	return visits, nil
}

func (r *visitsRepository) Create(visit *Visit) (*Visit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Create(visit).Error

	if err != nil {
		return nil, err
	}

	return visit, nil
}

func (r *visitsRepository) Update(visit *Visit) (*Visit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Save(visit).Error

	if err != nil {
		return nil, err
	}

	return visit, nil
}

func (r *visitsRepository) Delete(visit *Visit) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Delete(visit).Error

	if err != nil {
		return err
	}

	return nil
}
