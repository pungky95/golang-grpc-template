package repositories

import (
	"github.com/pungky95/golang-grpc-template/entities"
	"gorm.io/gorm"
)

type PermissionRepositoryInterface interface {
	FindOne(whereClause interface{}, whereNotClause interface{}, relations []string) (entities.Permission, error)
	Find(whereClause interface{}, whereNotClause interface{}, relations []string) ([]entities.Permission, error)
	Save(*entities.Permission, *gorm.DB) error
	Delete(entities.Permission, *gorm.DB) error
}

func NewPermissionRepository(db *gorm.DB) PermissionRepositoryInterface {
	return &PermissionRepository{db: db}
}

type PermissionRepository struct {
	db *gorm.DB
}

func (r PermissionRepository) FindOne(whereClause interface{}, whereNotClause interface{}, relations []string) (permission entities.Permission, err error) {
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
	err = queryBuilder.First(&permission).Error
	return permission, err
}

func (r PermissionRepository) Find(whereClause interface{}, whereNotClause interface{}, relations []string) (permissions []entities.Permission, err error) {
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
	err = queryBuilder.Find(&permissions).Error
	return permissions, err
}

func (r PermissionRepository) Save(permission *entities.Permission, tx *gorm.DB) (err error) {
	if tx != nil {
		return tx.Save(&permission).Error
	}
	return r.db.Save(&permission).Error
}

func (r PermissionRepository) Delete(permission entities.Permission, tx *gorm.DB) error {
	if tx != nil {
		return tx.Delete(&permission).Error
	}
	return r.db.Delete(&permission).Error
}
