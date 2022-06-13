package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

//UserService is a contract.....
type UserService interface {
	Update(user schema.UserUpdateDTO) schema.User
	Profile(userID string) schema.User
}

type userService struct {
	userRepository database.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo database.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user schema.UserUpdateDTO) schema.User {
	userToUpdate := schema.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) schema.User {
	return service.userRepository.ProfileUser(userID)
}
