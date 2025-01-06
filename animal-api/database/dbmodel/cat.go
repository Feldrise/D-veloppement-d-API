package dbmodel

import (
	"context"
	"time"

	"feldrise.com/animal-api/pkg/model"
	"gorm.io/gorm"
)

type Cat struct {
	gorm.Model

	Name string `gorm:"not null"`

	Visists []*Visit `gorm:"foreignKey:CatID"`
}

func (cat *Cat) ToModel() *model.Cat {
	return &model.Cat{
		Name: cat.Name,
	}
}

type CatsFieldsToInclude struct {
	Visists bool
}

type CatsRepository interface {
	FindByID(id uint, fields *CatsFieldsToInclude) (*Cat, error)
	FindAll(fields *CatsFieldsToInclude) ([]*Cat, error)
	Create(cat *Cat) (*Cat, error)
	Update(cat *Cat) (*Cat, error)
	Delete(cat *Cat) error
}

type catsRepository struct {
	db *gorm.DB
}

func NewCatsRepository(db *gorm.DB) CatsRepository {
	return &catsRepository{
		db: db,
	}
}

func (r *catsRepository) FindByID(id uint, fields *CatsFieldsToInclude) (*Cat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cat Cat
	tx := r.db.WithContext(ctx).Model(&cat)

	if fields != nil && fields.Visists {
		tx = tx.Preload("Visits")
	}

	err := tx.Where("id = ?", id).First(&cat).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &cat, nil
}

func (r *catsRepository) FindAll(fields *CatsFieldsToInclude) ([]*Cat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cats []*Cat
	tx := r.db.WithContext(ctx).Model(&Cat{})

	if fields != nil && fields.Visists {
		tx = tx.Preload("Visits")
	}

	err := tx.Find(&cats).Error

	if err != nil {
		return nil, err
	}

	return cats, nil
}

func (r *catsRepository) Create(cat *Cat) (*Cat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Create(cat).Error

	if err != nil {
		return nil, err
	}

	return cat, nil
}

func (r *catsRepository) Update(cat *Cat) (*Cat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Model(cat).Updates(cat).Error

	if err != nil {
		return nil, err
	}

	return cat, nil
}

func (r *catsRepository) Delete(cat *Cat) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Delete(cat).Error

	if err != nil {
		return err
	}

	return nil
}
