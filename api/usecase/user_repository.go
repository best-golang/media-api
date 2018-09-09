package usecase

import "api/domain"

type UserRepository interface {
	Create(*domain.User) (*interface{}, error)
	Index(*domain.Users) error
}
