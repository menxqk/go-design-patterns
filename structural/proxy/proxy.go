package structural

import (
	"fmt"
)

// The Proxy design pattern is used to provide an object that acts as a substitute for
// a real service object used by a client.
type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	// Search for the object in the cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}

	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}
	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.LastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

type UserList []User

func (l *UserList) addUser(newUser User) {
	*l = append(*l, newUser)
}

func (l *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*l); i++ {
		if (*l)[i].ID == id {
			return (*l)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found", id)
}

type User struct {
	ID int32
}
