package repo

import (
	"github.com/leodahal4/grpc-clean/internal/models"
	"gorm.io/gorm"
	interfaces "github.com/leodahal4/grpc-clean/pkg/v1"
)

type Repo struct {
  db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
  return &Repo{db}
}

// Create
//
// This function creates a new User which was supplied as the argument
func (repo *Repo) Create(user models.User) (models.User, error){
  err := repo.db.Create(&user).Error

  return user, err
}

// Get
//
// This function returns the user instance which is 
// saved on the DB and returns to the usecase
func (repo *Repo) Get(id string) (models.User, error){
  var user models.User
  err := repo.db.Where("id = ?", id).First(&user).Error

  return user, err
}

// Update
// 
// This function updates the user name to the one specified,
// as we have email, id and name, so name only gets changed
func (repo *Repo) Update(user models.User) error{
  var dbUser models.User
  if err := repo.db.Where("id = ?", user.ID).First(&dbUser).Error; err != nil {
    return err
  }
  dbUser.Name = user.Name
  err := repo.db.Save(dbUser).Error

  return err
}

// Delete
//
// This function creates a the User whose ID was supplied as the argument
func (repo *Repo) Delete(id string) error{
  err := repo.db.Where("id = ?", id).Delete(&models.User{}).Error

  return err
}

// GetByEmail
//
// This function returns the user instance which is 
// saved on the DB and returns to the usecase
func (repo *Repo) GetByEmail(email string) (models.User, error){
  var user models.User
  err := repo.db.Where("email = ?", email).First(&user).Error

  return user, err
}
