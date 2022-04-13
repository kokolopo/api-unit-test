package user

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock interface {
	Save(user User) (User, error)
	FindById(id int) (User, error)
	FindAll() ([]User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(user User) (User, error)
}

type repositoryMock struct {
	Mock mock.Mock
}

func NewRepositoryMock(mock mock.Mock) *repositoryMock {
	return &repositoryMock{mock}
}

func (r *repositoryMock) Save(user User) (User, error) {

	argument := r.Mock.Called(user)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(User)
		return newUser, nil
	}
}

func (r *repositoryMock) FindById(id int) (User, error) {
	var user User
	argument := r.Mock.Called(id)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		user := argument.Get(0).(User)
		return user, nil
	}
}

func (r *repositoryMock) FindAll() ([]User, error) {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("ada yang salah")
	} else {
		user := arguments.Get(0).([]User)
		return user, nil
	}
}

func (r *repositoryMock) UpdateUser(user User) (User, error) {
	argument := r.Mock.Called(user)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(User)
		return newUser, nil
	}
}

func (r *repositoryMock) DeleteUser(user User) (User, error) {
	argument := r.Mock.Called(user)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(User)
		return newUser, nil
	}

}
