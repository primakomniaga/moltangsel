package user

import (
	"context"
	"fmt"
)

// interface for connect to Resource
type Resource interface {
	GetUser(int) (map[string]interface{}, error)
	Register(*User) error
	Login(*Login) (*User, error)
	ValidEmailPhone(*User) bool
}

// Service of user
type Service struct {
	resource Resource
}

// New user service
func New(userResource Resource) *Service {
	s := Service{
		resource: userResource,
	}
	return &s
}

func (s *Service) IsUserActive(ctx context.Context, userId int) (bool, error) {

	fmt.Println("isUserActive <<<====")

	return true, nil
}

func (s *Service) Register(ctx context.Context, model *User) error {
	if err := s.resource.Register(model); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUser(userId int) (map[string]interface{}, error) {
	u, err := s.resource.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) Login(m *Login) (*User, error) {
	user, err := s.resource.Login(m)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) ValidEmailPhone(m *User) bool {
	if s.resource.ValidEmailPhone(m) {
		return true
	}
	return false
}
