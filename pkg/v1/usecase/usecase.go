package usecase

import (
	"errors"

	"github.com/leodahal4/grpc-clean/internal/models"
	interfaces "github.com/leodahal4/grpc-clean/pkg/v1"
	"gorm.io/gorm"
)

type UseCase struct {
  repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
  return &UseCase{repo}
}

// Create
//
// This function creates a new User which was supplied as the argument
func (uc *UseCase) Create(user models.User) (models.User, error) {
  // check if valid email is supplied
  if _, err := uc.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound){
    return models.User{}, errors.New("the email is already associated with another user")
  }

  // email doesnot exist so, now proceed
  return uc.repo.Create(user)
}

// Get
//
// This function returns the user instance which is 
// saved on the DB and returns to the usecase
func (uc *UseCase) Get(id string) (models.User, error){
  var user models.User
  var err error

  if user, err = uc.repo.Get(id); err != nil{
    if errors.Is(err, gorm.ErrRecordNotFound){
      return models.User{}, errors.New("no such user with the id supplied")
    }
    // handle the error properly as the error was something else then record not found,
    // and return to the caller of this function
    return models.User{}, err
  }

  return user, nil
}

// Update
// 
// This function updates the user name to the one specified,
// as we have email, id and name, so name only gets changed
func (uc *UseCase) Update(updateUser models.User) error{
  var user models.User
  var err error
  // check if user exists
  if user, err = uc.Get(string(updateUser.ID)); err != nil {
    return err
  }

  // check if only name is going to change,
  // as the email cannot be changed
  if user.Email != updateUser.Email {
    return errors.New("email cannot be changed")
  }
  
  err = uc.repo.Update(updateUser)
  if err != nil {
    // handle the error properly as the error might be something worth to debug
  }

  return nil
}

// Delete
//
// This function creates a the User whose ID was supplied as the argument
func (uc *UseCase) Delete(id string) error{
  var err error
  // check if user exists
  if _, err = uc.Get(id); err != nil {
    return err
  }

  err = uc.repo.Delete(id)
  if err != nil {
    // handle the error as it might be something worth to debug
    return err
  }

  return nil
}
