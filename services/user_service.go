package services

import (
	"github.com/student/ginblog/datamodels"
	"github.com/student/ginblog/repositories"
)

type IUserService interface {
	// GetUserByID(int64) (*datamodels.User, error)
	GetAllUser() ([]*datamodels.User, error)
	// GetAllUser()
	DeleteUserByID(int64) bool
	InsertUser(user *datamodels.User) int64
	UpdateUser(user *datamodels.User) error
	GetUserByName(*datamodels.User, string) int64
	GetOneUser(*datamodels.LoginUser, string) *datamodels.LoginUser
	GetOneUserById(*datamodels.User, int64) *datamodels.User
}

type UserService struct {
	userRepository repositories.IUser
}

func NewUserService(repository repositories.IUser) IUserService {
	return &UserService{repository}
}
func (p *UserService) InsertUser(user *datamodels.User) int64 {
	return p.userRepository.Insert(user)
}
func (p *UserService) GetUserByName(user *datamodels.User, name string) int64 {
	return p.userRepository.SelectByName(user, name)
}
func (p *UserService) GetOneUser(user *datamodels.LoginUser, name string) *datamodels.LoginUser {
	return p.userRepository.FindUser(user, name)
}
func (p *UserService) GetAllUser() ([]*datamodels.User, error) {
	return p.userRepository.SelectAll()
}
func (p *UserService) GetOneUserById(user *datamodels.User, id int64) *datamodels.User {
	return p.userRepository.FindUserById(user, id)
}
func (p *UserService) UpdateUser(user *datamodels.User) error {
	return p.userRepository.Update(user)
}
func (p *UserService) DeleteUserByID(id int64) bool {
	return p.userRepository.Delete(id)
}
