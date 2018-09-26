package user

import (
	"context"
	"fmt"
)

// interface for connect to Resource
type Resource interface {
	GetUser(int64) (User, error)
	Register(*User) error
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

func (s *Service) IsUserActive(ctx context.Context, userId int64) (bool, error) {

	fmt.Println("isUserActive <<<====")

	return true, nil
}

func (s *Service) Register(ctx context.Context, model *User) error {
	if err := s.resource.Register(model); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context) User {
	u, err := s.resource.GetUser(1)
	if err != nil {
		fmt.Println("error getuser ", err)
	}
	return u
}
