package test

import (
	"fmt"
	"os"
	"testing"

	"../orm"
)

func Testmain(m *testing.M) {
	BeforeTest()
	result := m.Run()
	AfterTest()
	os.Exit(result)
}

func BeforeTest() {
	fmt.Println("Antes de las pruebas ")
	orm.CreatConnection()
	orm.CreateTables()
}
func AfterTest() {
	fmt.Println("Despues de las pruebas ")
}
