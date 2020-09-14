package models

type User struct {
	Id       int64
	Name     string `json:"name"`
	Password string `json:"pasword"`
	Email    string `json:"email"`
}

const userShema string = `CREATE TABLE users (
         id int(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
         name VARCHAR(30) NOT NULL,
         password VARCHAR(60) NOT NULL,
         email VARCHAR (40),
         created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`

var users = make(map[int]User)

type Users []User

//NewUser name y password
func NewUser(username, password, email string) *User {
	user := &User{Name: username, Password: password, Email: email}
	return user
}

func (this *User) Save() error {
	if this.Id == 0 {
		return this.insert()
	} else {
		return this.Update()
	}
}
func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "select id ,name, password, email from users where id=?"
	rows, _ := Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	}
	return user
}
func GetUsers() Users {
	users := Users{}
	sql := "select id ,name, password, email from users"
	rows, _ := Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}
func (this *User) insert() error {
	sql := "insert users set name=?,password=?,email=?"
	result, err := Exec(sql, this.Name, this.Password, this.Email)
	this.Id, _ = result.LastInsertId()
	return err
}
func (this *User) Update() error {
	sql := "update users set name=?,password=?,email=?"
	_, err := Exec(sql, this.Name, this.Password, this.Name)
	return err
}
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func (this *User) Delete() {
	sql := "delete from users where id=?"
	Exec(sql, this.Id)
}
