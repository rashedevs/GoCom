package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
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
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) Create(usr User) (*User, error) {
	query := `
        INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
        VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
        RETURNING id
    `
	var userID int
	rows, err := r.db.NamedQuery(query, usr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userID)
	}

	usr.ID = userID
	return &usr, nil
}

func (r userRepo) Find(email, pass string) (*User, error) {
	var user User
	query := `
        SELECT id, first_name, last_name, email, password, is_shop_owner
        FROM users
        WHERE email=$1 AND password=$2
        LIMIT 1
    `

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		fmt.Println("haha", err)
		// sqlx.Get returns sql.ErrNoRows if nothing is found
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r userRepo) Get(uId int) (*User, error) {
	// for _, User := range r.users {
	// 	if User.ID == uId {
	// 		return &User, nil
	// 	}
	// }
	return nil, nil
}

// func (u *userRepo) List() ([]*User, error) {
// 	return &u.users, nil
// }

func (r userRepo) Update(usr User) (*User, error) {
	// for idx, User := range u.users {
	// 	if User.ID == usr.ID {
	// 		u.users[idx] = usr
	// 		return nil, nil
	// 	}
	// }
	return nil, nil
}

func (r userRepo) Delete(uId int) error {
	// var tempList []User
	// for _, user := range u.users {
	// 	if user.ID != uId {
	// 		tempList = append(tempList, user)
	// 		// tempList[idx] = User

	// 	}
	// }
	// u.users = tempList
	// users = append(users, tempList...)
	return nil
}
