package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = NewRepositoryMock(mock.Mock{})
var userService = NewService(userRepository)

func TestGetByIdFail(t *testing.T) {

	//program mock
	userRepository.Mock.On("FindById", 1).Return(nil)

	_, err := userService.GetById(1)
	assert.Nil(t, nil)
	assert.NotNil(t, err)
}

func TestGetByIdSuccess(t *testing.T) {
	user := User{
		Id:       2,
		Name:     "fahmi",
		Email:    "fahmi@test.com",
		Password: "password",
	}

	userRepository.Mock.On("FindById", 2).Return(user)

	result, err := userService.GetById(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.Name, result.Name)
}

func TestGetAllFail(t *testing.T) {
	userRepository.Mock.On("FindAll").Return(nil)

	user, err := userService.GetAll()
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestGetAllSuccess(t *testing.T) {
	var users = []User{
		{2, "test1", "test1@gmail.com", "pass1"},
		{3, "test2", "test2@gmail.com", "pass2"},
	}

	userRepository.Mock.On("FindAll").Return(users)

	result, err := userService.GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, users[0].Id, result[0].Id)
	assert.Equal(t, users[0].Name, result[0].Name)
}

func TestRegisterUserFail(t *testing.T) {
	//var errors error
	input := UserRegisterInput{}
	user := User{}
	userRepository.Mock.On("Save", user).Return(nil)

	_, err := userService.RegisterUser(input)
	assert.Nil(t, nil)
	assert.NotNil(t, err)
}

func TestRegisterUserSuccess(t *testing.T) {
	var userInput = UserRegisterInput{Name: "test1", Email: "test1@gmail.com", Password: "pass1"}
	var user = User{Name: userInput.Name, Email: userInput.Email, Password: userInput.Password}

	userRepository.Mock.On("Save", user).Return(user)

	result, err := userService.RegisterUser(userInput)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.Name, result.Name)
}

func TestUpdateUserSuccess(t *testing.T) {
	userA := User{
		Id:       0,
		Name:     "fahmi",
		Email:    "fahmi@test.com",
		Password: "password",
	}
	var userInput = UpdateInput{Name: "test1", Email: "test1@gmail.com", Password: "pass1"}
	var user = User{Name: userInput.Name, Email: userInput.Email, Password: userInput.Password}
	userPass := User{
		Id:       0,
		Name:     "test1",
		Email:    "test1@gmail.com",
		Password: "pass1",
	}

	userRepository.Mock.On("FindById", 0).Return(userA)
	userRepository.Mock.On("UpdateUser", user).Return(userPass)

	result, err := userService.UpdateUser(0, userInput)

	assert.Nil(t, err)
	assert.NotNil(t, userPass)
	assert.Equal(t, userPass.Id, result.Id)
	assert.Equal(t, userPass.Name, result.Name)
}

func TestUpdateUserFail(t *testing.T) {

	input := UpdateInput{}
	user := User{}

	userRepository.Mock.On("FindById", 0).Return(nil)
	userRepository.Mock.On("UpdateUser", user).Return(nil)

	_, err := userService.UpdateUser(0, input)

	assert.Nil(t, nil)
	assert.NotNil(t, err)
}

func TestDeleteUserSuccess(t *testing.T) {
	userA := User{
		Id:       1,
		Name:     "fahmi",
		Email:    "fahmi@test.com",
		Password: "password",
	}

	userRepository.Mock.On("FindById", 1).Return(userA)
	userRepository.Mock.On("DeleteUser", userA).Return(nil)

	result, _ := userService.DeleteUser(1)
	assert.Nil(t, nil)
	assert.NotNil(t, result)

}

func TestDeleteUserFail(t *testing.T) {
	userA := User{}
	var errors error

	userRepository.Mock.On("FindById", 1).Return(nil)
	userRepository.Mock.On("DeleteUser", userA).Return(errors)

	_, err := userService.DeleteUser(1)
	assert.Nil(t, errors)
	assert.NotNil(t, err)

}
