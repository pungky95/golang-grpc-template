package repositories

import (
	"github.com/pungky95/golang-grpc-template/entities"
	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	FindOne(whereClause interface{}, whereNotClause interface{}, relations []string) (entities.Role, error)
	Find(whereClause interface{}, whereNotClause interface{}, relations []string) ([]entities.Role, error)
	Save(*entities.Role, *gorm.DB) error
	Delete(entities.Role, *gorm.DB) error
}

func NewRoleRepository(db *gorm.DB) RoleRepositoryInterface {
	return &RoleRepository{db: db}
}

type RoleRepository struct {
	db *gorm.DB
}

func (r RoleRepository) FindOne(whereClause interface{}, whereNotClause interface{}, relations []string) (role entities.Role, err error) {
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
	err = queryBuilder.First(&role).Error
	return role, err
}

func (r RoleRepository) Find(whereClause interface{}, whereNotClause interface{}, relations []string) (roles []entities.Role, err error) {
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
	err = queryBuilder.Find(&roles).Error
	return roles, err
}

func (r RoleRepository) Save(role *entities.Role, tx *gorm.DB) (err error) {
	if tx != nil {
		return tx.Save(&role).Error
	}
	return r.db.Save(&role).Error
}

func (r RoleRepository) Delete(role entities.Role, tx *gorm.DB) error {
	if tx != nil {
		return tx.Delete(&role).Error
	}
	return r.db.Delete(&role).Error
}
