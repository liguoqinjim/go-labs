package main

import (
	"log"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Class *Class
}

type Class struct {
	Cno  int
	Name string
}

func main() {
	c := Class{Cno: 1, Name: "Class1st"}
	s := &Student{Name: "LiLei", Age: 18, Class: &c}

	//struct----------------------------------------------------------------------------------------------------
	log.Println("reflect struct-----------------------------------------------------")
	valClass := reflect.ValueOf(c)
	valField1 := valClass.FieldByName("Cno")
	if valField1.CanSet() {
		valField1.SetInt(2)
	} else {
		log.Println("valField1 can not set")
	}
	valField2 := valClass.FieldByName("Name")
	if valField2.CanSet() {
		valField2.SetString("Class2nd")
	} else {
		log.Println("valField2 can not set")
	}
	log.Printf("the class'cno[%d] name[%s]\n", c.Cno, c.Name)

	//pointer  (Elem())---------------------------------------------------------------------------------------------------
	log.Println("reflect pointer------------------------------------------------------")
	valStudent := reflect.ValueOf(s).Elem()
	typStudent := reflect.TypeOf(s).Elem()
	for i := 0; i < valStudent.NumField(); i++ {
		fv := valStudent.Field(i)
		ft := typStudent.Field(i)

		if fv.CanSet() {
			switch fv.Kind() {
			case reflect.Int:
				fv.SetInt(15)
			case reflect.String:
				fv.SetString("BaoMeimei")
			}
		} else {
			log.Printf("student's field[%s] can not set\n", ft.Name)
		}

	}
	log.Printf("student'name[%s] age[%d] \n", s.Name, s.Age)
}
