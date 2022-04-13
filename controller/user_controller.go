package controller

import (
	"api_unit_test/helper"
	"api_unit_test/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//Register
//GetById
//GetAll
//UpdateUser
//DeleteUser

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	var input user.UserRegisterInput
	err := c.Bind(&input)
	if err != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	newUser, errUser := h.userService.RegisterUser(input)
	if errUser != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

		return c.JSON(http.StatusBadRequest, res)
	}

	// token, errToken := h.authService.GenerateTokenJWT(newUser.Id, newUser.Name)
	// if errToken != nil {
	// 	res := helper.ApiResponse("New User Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

	// 	return c.JSON(http.StatusBadRequest, res)
	// }

	formatter := user.FormatUser(newUser, "token")

	res := helper.ApiResponse("New User Data Has Been Created", http.StatusCreated, "success", formatter)

	return c.JSON(http.StatusCreated, res)
}

func (h *userHandler) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userById, err := h.userService.GetById(id)
	if err != nil {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	if userById.Id == 0 {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	res := helper.ApiResponse("Successfuly Login", http.StatusOK, "success", userById)

	return c.JSON(http.StatusCreated, res)

}

func (h *userHandler) GetAll(c echo.Context) error {
	users, err := h.userService.GetAll()
	if err != nil {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	res := helper.ApiResponse("Successfuly Login", http.StatusOK, "success", users)

	return c.JSON(http.StatusCreated, res)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var input user.UpdateInput
	err := c.Bind(&input)
	if err != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	userUpdate, errUpdate := h.userService.UpdateUser(id, input)
	if errUpdate != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusBadRequest, "failed", errUpdate)

		return c.JSON(http.StatusBadRequest, res)
	}

	res := helper.ApiResponse("Successfuly Update", http.StatusOK, "success", userUpdate)

	return c.JSON(http.StatusCreated, res)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userById, err := h.userService.GetById(id)
	if err != nil {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	if userById.Id == 0 {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	userDelete, errDel := h.userService.DeleteUser(userById.Id)
	if errDel != nil {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", errDel)

		return c.JSON(http.StatusBadRequest, res)
	}

	cekUser, errCek := h.userService.GetById(id)
	if errCek != nil {
		res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", errCek)

		return c.JSON(http.StatusBadRequest, res)
	}

	if cekUser.Id == 0 {
		res := helper.ApiResponse("Successfuly Delete User", http.StatusOK, "success", nil)

		return c.JSON(http.StatusOK, res)
	}

	res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", userDelete)

	return c.JSON(http.StatusCreated, res)
}
