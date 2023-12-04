package repositories

import (
	"github.com/pungky95/golang-grpc-template/entities"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindOne(whereClause interface{}, whereNotClause interface{}, relations []string) (entities.User, error)
	Find(whereClause interface{}, whereNotClause interface{}, relations []string) ([]entities.User, error)
	Save(*entities.User, *gorm.DB) error
	Delete(entities.User, *gorm.DB) error
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

type UserRepository struct {
	db *gorm.DB
}

func (r UserRepository) FindOne(whereClause interface{}, whereNotClause interface{}, relations []string) (user entities.User, err error) {
	queryBuilder := r.db
	if whereClause != nil {
		queryBuilder = queryBuilder.Where(whereClause)
	}
	if whereNotClause != nil {
		queryBuilder = queryBuilder.Not(whereNotClause)
	}
	for _, relation := range relations {
		queryBuilder = queryBuilder.Preload(relation)
	}
	err = queryBuilder.First(&user).Error
	return user, err
}

func (r UserRepository) Find(whereClause interface{}, whereNotClause interface{}, relations []string) (users []entities.User, err error) {
	queryBuilder := r.db
	if whereClause != nil {
		queryBuilder = queryBuilder.Where(whereClause)
	}
	if whereNotClause != nil {
		queryBuilder = queryBuilder.Not(whereNotClause)
	}
	for _, relation := range relations {
		queryBuilder = queryBuilder.Preload(relation)
	}
	err = queryBuilder.Find(&users).Error
	return users, err
}

func (r UserRepository) Save(user *entities.User, tx *gorm.DB) (err error) {
	if tx != nil {
		return tx.Save(&user).Error
	}
	return r.db.Save(&user).Error
}

func (r UserRepository) Delete(user entities.User, tx *gorm.DB) error {
	if tx != nil {
		return tx.Delete(&user).Error
	}
	return r.db.Delete(&user).Error
}
