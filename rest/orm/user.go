package orm

type User struct {
	Id       int
	Name     string `json:"name"`
	Password string `json:"pasword"`
	Email    string `json:"email"`
}
type Users []User

func NewUser(username, password, email string) *User {
	user := &User{Name: username, Password: password, Email: email}
	return user
}
func GetUsers() Users {
	users := Users{}
	db.Find(&users)
	return users
}

func GetUser(id int) *User {
	user := &User{}
	db.Where("id=?", id).First(user)
	return user
}

//Save asdasd
func (this *User) Save() error {
	if this.Id == 0 {
		db.Create(this)
	} else {
		this.update()
	}
	return db.Error
}

func (this *User) update() error {
	user := User{Name: this.Name, Password: this.Password, Email: this.Email}
	db.Model(&this).UpdateColumns(user)
	return db.Error
}

func (this *User) Delete() {
	db.Delete(&this)
}
