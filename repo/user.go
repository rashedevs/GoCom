package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LirstName   string `json:"lirst_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(usr User) (*User, error)
	Find(email, pass string) (*User, error)
	// List() ([]*User, error)
	Get(uId int) (*User, error)
	Update(usr User) (*User, error)
	Delete(uId int) error
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u userRepo) Create(usr User) (*User, error) {
	if usr.ID != 0 {
		return &usr, nil
	}
	usr.ID = len(u.users) + 1
	u.users = append(u.users, usr)
	return &usr, nil
}

func (u userRepo) Find(email, pass string) (*User, error) {
	for _, user := range u.users {
		if user.Email == email && user.Password == pass {
			return &user, nil
		}
	}
	return nil, nil
}

func (u userRepo) Get(uId int) (*User, error) {
	for _, User := range u.users {
		if User.ID == uId {
			return &User, nil
		}
	}
	return nil, nil
}

// func (u *userRepo) List() ([]*User, error) {
// 	return &u.users, nil
// }

func (u userRepo) Update(usr User) (*User, error) {
	for idx, User := range u.users {
		if User.ID == usr.ID {
			u.users[idx] = usr
			return nil, nil
		}
	}
	return nil, nil
}

func (u userRepo) Delete(uId int) error {
	var tempList []User
	for _, user := range u.users {
		if user.ID != uId {
			tempList = append(tempList, user)
			// tempList[idx] = User

		}
	}
	u.users = tempList
	// users = append(users, tempList...)
	return nil
}
