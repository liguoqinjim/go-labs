package main

import (
	sq "github.com/Masterminds/squirrel"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	query1()
}

func example() {
	users := sq.Select("*").From("users").Join("emails USING (email_id)")

	active := users.Where(sq.Eq{"deleted_at": nil})

	sql, args, err := active.ToSql()
	if err != nil {
		log.Fatalf("active.ToSql error:%v", err)
	}
	log.Println(args)
	log.Println(sql)
}

func query1() {
	username := "li"
	companyId := 1
	log.Println(username, companyId)

	s := sq.Select("*").From("t_user")

	s = s.Offset(10).Limit(10)

	//where
	if companyId != 0 {
		s = s.Where(sq.Eq{"company_id": companyId})
	}
	if username != "" {
		s = s.Where(sq.Eq{"username": username})
	}

	sql, args, err := s.ToSql()
	if err != nil {
		log.Fatalf("s.ToSql error:%v", err)
	}
	_ = args
	log.Println(sql)
}
