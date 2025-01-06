package dbmodel

import (
	"context"
	"time"

	"feldrise.com/animal-api/pkg/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"not null;unique"`
	EmailToken     *string
	EmailVerifedAt *time.Time

	PasswordHash       string `gorm:"not null"`
	PasswordResetToken *string

	RoleID uint `gorm:"not null;DEFAULT:2;"`
	Role   Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserProfile *UserProfile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserFilter struct {
	SearchString string
}

func (user *User) ToModel(includeEmail bool, includePhone bool) *model.User {
	var userProfile *model.UserProfile

	if user.UserProfile != nil {
		userProfile = user.UserProfile.ToModel()
	}

	email := ""
	if includeEmail {
		email = user.Email
	}

	return &model.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		Email:     email,
		Profile:   userProfile,
	}
}

type UserRepository interface {
	FindByID(id uint, includeProfile bool) (*User, error)
	FindByEmail(email string, includeProfile bool) (*User, error)
	FindAll(filter *UserFilter) ([]*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	// Delete(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByID(id uint, includeProfile bool) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	tx := r.db.WithContext(ctx).Model(&user)

	if includeProfile {
		tx = tx.Preload("UserProfile")
	}

	err := tx.Where("id = ?", id).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string, includeProfile bool) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	tx := r.db.WithContext(ctx).Model(&user)

	if includeProfile {
		tx = tx.Joins("UserProfile")
	}

	err := tx.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByPhone(phone string, includeProfile bool) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	tx := r.db.WithContext(ctx).Model(&user)

	if includeProfile {
		tx = tx.Joins("UserProfile")
	}

	err := tx.Where("phone = ?", phone).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll(filter *UserFilter) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []*User
	tx := r.db.WithContext(ctx).Model(&User{})
	tx = tx.Preload("UserProfile")

	if filter != nil && filter.SearchString != "" {
		searchString := "%" + filter.SearchString + "%"
		tx = tx.Joins("LEFT JOIN user_profiles ON user_profiles.user_id = users.id").
			Where("users.email ILIKE ? OR user_profiles.first_name ILIKE ? OR user_profiles.last_name ILIKE ?", searchString, searchString, searchString)
	}

	err := tx.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Create(user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Save(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
