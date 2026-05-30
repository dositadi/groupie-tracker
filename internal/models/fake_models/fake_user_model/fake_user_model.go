package fakeusermodel

import (
	"errors"

	"github.com/dositadi/groupie-tracker/internal/data"
)

type FakeUserModel struct {
	Users map[string]data.User
	Err   error
}

func NewFakeUserModel() FakeUserModel {
	return FakeUserModel{
		Users: make(map[string]data.User),
	}
}

/*
type UserModel interface {
	Delete(id string) error
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
	Insert(user data.User) error
	Update(id string, info data.UpdateUser) error
	EmailExists(email string) (bool, error)
	IDExists(id string) (bool, error)
}
*/

var (
	USER_NOT_FOUND = errors.New("User not found")
	USER_EXISTS    = errors.New("User exists")
	EMAIL_EXISTS   = errors.New("Email exists")
)

func (f *FakeUserModel) Delete(id string) error {
	user, ok := f.Users[id]
	if !ok {
		return USER_NOT_FOUND
	}

	delete(f.Users, user.Id)
	return nil
}

func (f *FakeUserModel) GetWithID(id string) (data.User, error) {
	user, ok := f.Users[id]
	if !ok {
		return data.User{}, USER_NOT_FOUND
	}
	return user, nil
}

func (f *FakeUserModel) GetWithEmail(email string) (data.User, error) {
	for _, user := range f.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return data.User{}, USER_NOT_FOUND
}

func (f *FakeUserModel) Insert(user data.User) error {
	_, ok := f.Users[user.Id]
	if ok {
		return USER_EXISTS
	}
	f.Users[user.Id] = user
	return nil
}

func (f *FakeUserModel) Update(id string, info data.UpdateUser) error {
	user, ok := f.Users[id]
	if !ok {
		return USER_NOT_FOUND
	}

	if info.HashedPassword != nil {
		user.HashedPassword = info.HashedPassword
	}

	if info.Username != nil {
		user.Username = *info.Username
	}

	if info.Email != nil {
		for _, user := range f.Users {
			if user.Email == *info.Email {
				return EMAIL_EXISTS
			}
		}
		user.Email = *info.Email
	}

	f.Users[user.Id] = user
	return nil
}

func (f *FakeUserModel) EmailExists(email string) (bool, error) {
	for _, user := range f.Users {
		if user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func (f *FakeUserModel) IDExists(id string) (bool, error) {
	_, ok := f.Users[id]
	if !ok {
		return false, USER_NOT_FOUND
	}
	return ok, nil
}
