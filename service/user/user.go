package user_service

type UserService interface{}

type UserServiceImpl struct {
}

func New() *UserServiceImpl {
	return &UserServiceImpl{}
}
