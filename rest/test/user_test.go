package test

import (
	"testing"

	"../orm"
)

const (
	Name     = "luciano"
	Password = "123456"
	Email    = "oliluciano_11@gmail.com"
)

func TestUser(t *testing.T) {
	user := orm.NewUser(Name, Password, Email)
	if user.Name != Name || user.Password != Password || user.Email != Email {
		t.Error("no es posible crear onjeto", nil)
	}

}

func TestSave(t *testing.T) {
	user := orm.NewUser(Name, Password, Email)
	if err := user.Save(); err != nil {
		t.Error("no se puede guardar ", nil)
	}
}

