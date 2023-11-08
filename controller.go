package main

func GetAllUsers() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

func GetUserByID(id uint) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	return db.Create(&user).Error
}

func UpdateUser(user *User) error {
	return db.Save(&user).Error
}

func DeleteUser(id uint) error {
	return db.Delete(&User{}, id).Error
}
