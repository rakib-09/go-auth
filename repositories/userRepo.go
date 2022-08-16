package repositories

import (
	"go-auth/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) UpsertUser(user *models.User) (*models.User, error) {
	if err := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: false,
		DoUpdates: clause.AssignmentColumns([]string{
			"id",
			"name",
			"email",
			"password",
		}),
	}, clause.Returning{}).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserById(userId int) (*models.User, error) {
	var user models.User
	if err := repo.db.Model(&models.User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(userEmail string) (*models.User, error) {
	var user models.User
	if err := repo.db.Model(&models.User{}).Where("email = ?", userEmail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
