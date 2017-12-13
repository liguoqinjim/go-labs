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

	//struct reflect--------------------------------------------------------------------------------
	//FieldByName()返回的是StructField
	log.Println("struct reflect----------------------------------------------------")
	typClass := reflect.TypeOf(c)
	log.Println("typ1=", typClass.String())
	typClassKind := typClass.Kind()
	log.Println("typ1.Kind()=", typClassKind)
	//查看字段的类型
	typClassField1, found := typClass.FieldByName("Cno")
	if found {
		log.Printf("typClassField1.Name=%s,type=%s\n", typClassField1.Name, typClassField1.Type.Name())
	}
	typClassField2, found := typClass.FieldByName("School")
	if !found {
		_ = typClassField2
		log.Println("typClassField School not found")
	}
	typClassField3, found := typClass.FieldByName("Name")
	if found {
		log.Printf("typClassField3.Name=%s,type=%s\n", typClassField3.Name, typClassField3.Type.Name())
	}

	//查看字段的值--------------------------------------------------------------------------------
	valClass := reflect.ValueOf(c)
	log.Println("the class'cno is", valClass.FieldByName("Cno").Int())
	log.Println("the class'name is", valClass.FieldByName("Name").String())

	//pointer reflect
	log.Println("pointer reflect----------------------------------------------------")
	typ2 := reflect.TypeOf(s)
	log.Println("typ2=", typ2.String())
	typ2Kind := typ2.Kind()
	log.Println("typ2.Kind()=", typ2Kind)
	val2 := reflect.ValueOf(s)
	//FieldByName 只能用在struct上，在ptr上用的话会报错()
	//log.Println("the student'name is", val2.FieldByName("Name"))

	//ptr的时候，还要再调用一次Elem()
	val2 = val2.Elem()
	typ2 = typ2.Elem()
	log.Println("ptr.NumField()=", val2.NumField())
	for i := 0; i < val2.NumField(); i++ {
		fv := val2.Field(i)
		ft := typ2.Field(i)

		switch fv.Kind() {
		case reflect.Int:
			log.Printf("student's %dth field name=[%s],type[%s],value[%d]\n", i, ft.Name, ft.Type.Name(), fv.Int())
		case reflect.String:
			log.Printf("student's %dth field name=[%s],type[%s],value[%s]\n", i, ft.Name, ft.Type.Name(), fv.String())
		case reflect.Ptr:
			log.Printf("student's %dth field name=[%s],type[%s],value[%v]\n", i, ft.Name, fv.Kind(), fv.Pointer())
		}
	}
}
