package main

import "fmt"

type Db interface {
	Connect() string
}

type DbAdapter struct {
	*SpecificDb
}

func NewDbAdapter(db *SpecificDb) Db {
	return &DbAdapter{db}
}

func (adapter *DbAdapter) Connect() string {
	return adapter.SpecificConnect()
}

type SpecificDb struct {
	name string
}

func (db *SpecificDb) SpecificConnect() string {
	return fmt.Sprintf("Connected %s", db.name)
}

func main() {
	db := NewDbAdapter(&SpecificDb{"mysql"})
	fmt.Println(db.Connect())
}
