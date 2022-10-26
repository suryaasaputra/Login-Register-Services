package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	RegisterUser(models.User) (models.CreateUserResponse, error)
	EditUser(string, models.User) models.EditUserResponse
	GetUserByEmail(string) models.User
	GetUserByUserName(string) models.GetUserResponse
}

type UserDb struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (u UserDb) RegisterUser(user models.User) (models.CreateUserResponse, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	return models.CreateUserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FullName:  user.FullName,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u UserDb) GetUserByEmail(email string) models.User {
	User := models.User{}

	u.db.Where("email =? ", email).First(&User)

	return User
}

func (u UserDb) GetUserByUserName(userName string) models.GetUserResponse {
	User := models.User{}

	u.db.Where("user_name =? ", userName).First(&User)
	response := models.GetUserResponse{
		ID:         User.ID,
		UserName:   User.UserName,
		Email:      User.Email,
		FullName:   User.FullName,
		Age:        User.Age,
		Created_at: User.CreatedAt,
		Updated_at: User.UpdatedAt,
	}
	return response
}

func (u UserDb) EditUser(username string, newUser models.User) models.EditUserResponse {
	User := models.User{}
	u.db.Model(&User).Clauses(clause.Returning{}).Where("user_name=?", username).Updates(models.User{
		Email:    newUser.Email,
		UserName: newUser.UserName,
		FullName: newUser.FullName,
		Age:      newUser.Age,
	})

	response := models.EditUserResponse{
		ID:         User.ID,
		UserName:   User.UserName,
		Email:      User.Email,
		FullName:   User.FullName,
		Age:        User.Age,
		Updated_at: User.UpdatedAt,
	}

	return response
}
