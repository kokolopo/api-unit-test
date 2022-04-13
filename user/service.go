package user

type Service interface {
	RegisterUser(input UserRegisterInput) (User, error)
	GetById(id int) (User, error)
	GetAll() ([]User, error)
	UpdateUser(id int, input UpdateInput) (User, error)
	DeleteUser(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input UserRegisterInput) (User, error) {

	user := User{}

	// enkripsi data password
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	// if err != nil {
	// 	return user, err
	// }

	// memindahkan input ke struct User
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	newUser, errNewuser := s.repository.Save(user)
	if errNewuser != nil {
		return newUser, errNewuser
	}

	return newUser, nil

}

func (s *service) GetById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(id int, input UpdateInput) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	// passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	// if errHash != nil {
	// 	return user, errHash
	// }

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	userUpdated, errUpdate := s.repository.UpdateUser(user)
	if errUpdate != nil {
		return userUpdated, errUpdate
	}

	return userUpdated, nil

}

func (s *service) DeleteUser(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	deleteUser, errDel := s.repository.DeleteUser(user)
	if errDel != nil {
		return deleteUser, errDel
	}

	return deleteUser, nil
}
