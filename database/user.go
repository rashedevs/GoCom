package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LirstName   string `json:"lirst_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

func Find(email, pass string) *User {
	for _, user := range users {
		if user.Email == email && user.Password == pass {
			return &user
		}
	}
	return nil
}

// old
func (u User) List() []User {
	return users
}

func (u User) GetById(uId int) *User {
	for _, User := range users {
		if User.ID == uId {
			return &User
		}
	}
	return nil
}

func (u User) Update(usr User) {
	for idx, User := range users {
		if User.ID == u.ID {
			users[idx] = usr
		}
	}
}

func (u User) Delete(uId int) {
	var tempList []User
	for _, User := range users {
		if User.ID != uId {
			tempList = append(tempList, User)
			// tempList[idx] = User

		}
	}
	users = tempList
	// users = append(users, tempList...)
}

// func init() {

// 	prd1 := User{
// 		ID:          1,
// 		Title:       "Orange",
// 		Description: "This is orange.",
// 		Price:       150.00,
// 		ImgURL:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRwJf4xPOftZLLgkWjr2eMMumu9XuRdKiGc1eZXFku9OmA4lSymIPBm_vU0bFIM_BjpYOsI_pX7O9c9sRUxfkXxE0N1x_bWSERw7SXHP3A",
// 	}
// 	prd2 := User{
// 		ID:          2,
// 		Title:       "Apple",
// 		Description: "This is apple.",
// 		Price:       120.00,
// 		ImgURL:      "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcSbWqLu3A4mBXzZNi1POVqPrGV4UAuPw3Bl9RwTDjJqjCQrMiid5VgbDGHbAIA0rcFkAWdpjba5sqFdLdIvRXZ27_H_dPxXDpebJuqHiHQ",
// 	}
// 	prd3 := User{
// 		ID:          3,
// 		Title:       "Banana",
// 		Description: "This is banana.",
// 		Price:       40.00,
// 		ImgURL:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
// 	}
// 	prd4 := User{
// 		ID:          4,
// 		Title:       "Grapes",
// 		Description: "This is grapes.",
// 		Price:       140.00,
// 		ImgURL:      "https://d2t8nl1y0ie1km.cloudfront.net/images/6706bbb12b01cd80b29a0488_Grapes-Red-N-500gm_1.webp",
// 	}
// 	prd5 := User{
// 		ID:          5,
// 		Title:       "Mango",
// 		Description: "This is mango.",
// 		Price:       80.00,
// 		ImgURL:      "https://belescooverseas.com/wp-content/uploads/2024/04/Mango.jpeg",
// 	}

// 	users = append(users, prd1)
// 	users = append(users, prd2)
// 	users = append(users, prd3)
// 	users = append(users, prd4)
// 	users = append(users, prd5)
// }
