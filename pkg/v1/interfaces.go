package v1

import "github.com/leodahal4/grpc-clean/internal/models"


type RepoInterface interface{
  // creates a user with the data supplied
  Create(models.User) (models.User, error)

  // get retrieves the user instance
  Get(id string) (models.User, error)

  // update method updates the user and returns if any error occurred
  Update(models.User) (error)
  
  // Deletes the user whose ID is supplied
  Delete(id string) (error)

  // this method returns user instance by email
  GetByEmail(email string) (models.User, error)
}

type UseCaseInterface interface{
  // creates a user with the data supplied
  Create(models.User) (models.User, error)

  // get retrieves the user instance
  Get(id string) (models.User, error)

  // update method updates the user and returns if any error occurred
  Update(models.User) (error)
  
  // Deletes the user whose ID is supplied
  Delete(id string) (error)
}
