package sqlrepository

import (
	"github.com/ad0791/todoServices/api/v1/database"
	sqlmodel "github.com/ad0791/todoServices/api/v1/models/sql_model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func CreateUser(user *sqlmodel.User) error {
	if err := database.DB.Create(&user).Error; err != nil {
		log.Errorf("Failed to create user: %v", err)
		return err
	}
	log.Infof("User created with the ID %d", user.ID)
	return nil
}

func GetAllUSers() ([]*sqlmodel.User, error) {
	var users []*sqlmodel.User
	if err := database.DB.Find(&users).Error; err != nil {
		log.Errorf("Failed to fetch users: %v", err)
		return nil, err
	}
	log.Infof("Fetched %d users", len(users))
	return users, nil
}

func GetUserByID(id uint) (*sqlmodel.User, error) {
	var user sqlmodel.User
	if err := database.DB.First(&user, id).Error; err != nil {
		log.Errorf("User with ID %d not found: %v", id, err)
		return nil, err
	}
	log.Infof("User with ID %d fetched successfully", id)
	return &user, nil
}

func UpdateUser(user *sqlmodel.User) error {
	if err := database.DB.Save(&user).Error; err != nil {
		log.Errorf("Failed to update user with ID %d: %v", &user.ID, err)
		return err
	}
	log.Infof("User with ID %d updated successfully", &user.ID)
	return nil
}

func DeleteUser(id uint) (*sqlmodel.User, error) {
	var user sqlmodel.User
	res := database.DB.Where("id = ?", id).Delete(&user)

	if res.Error != nil {
		log.Errorf("Failed to soft-delete user with ID %d: %v", id, res.Error)
		return nil, res.Error
	}

	// Check if a record was deleted
	if res.RowsAffected == 0 {
		log.Warnf("No user found to delete with ID %d", id)
		return nil, gorm.ErrRecordNotFound
	}

	if err := database.DB.Unscoped().First(&user, id).Error; err != nil {
		log.Errorf("Failed to retrieve soft-deleted todo with ID %d: %v", id, err)
		return nil, err
	}

	return &user, nil

}
