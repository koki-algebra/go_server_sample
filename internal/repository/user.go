package repository

type UserRepository interface{}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

type userRepositoryImpl struct{}
