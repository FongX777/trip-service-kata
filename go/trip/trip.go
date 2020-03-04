package trip

import (
	"errors"
)

type Trip struct{}

type Service struct{}

var getLoggedUser = func() (*User, error) {
	return NewSession().GetLoggedUser()
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTripsByUser(user User) ([]Trip, error) {
	loggedUser, err := getLoggedUser()
	if err != nil {
		return nil, err
	}

	isFriend := false
	if loggedUser != nil {
		for _, friend := range user.GetFriends() {
			if friend.GetName() == loggedUser.GetName() {
				isFriend = true
				break
			}
		}
	} else {
		return nil, errors.New("loggedUser is nil")
	}

	trips := []Trip{}
	if isFriend {
		trips, err = FindTripsByUser(user)
		if err != nil {
			return nil, err
		}
	}

	return trips, nil
}
