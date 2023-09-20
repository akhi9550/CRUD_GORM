package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// record-Student
var id int
var name, domain string

type student struct {
	Id     int
	Name   string
	Domain string
}

func main() {
	var choice int
	db := connectPostgresDB()
	db.AutoMigrate(&student{})
	for {
		fmt.Println("\nChoose\n1.Insert\n2.Read\n3.Update\n4.Delete\n5.Exit")
		fmt.Println("Enter Your Choice:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Insert(db)
		case 2:
			Read(db)
		case 3:
			Update(db)
		case 4:
			Delete(db)
		case 5:
			os.Exit(0)
		}
	}
}
func connectPostgresDB() *gorm.DB {
	connectTo := "host=localhost user=postgres password=******* dbname=databasename port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func Insert(db *gorm.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	fmt.Println("Enter the name:")
	fmt.Scan(&name)
	fmt.Println("Enter the domain:")
	fmt.Scan(&domain)
	data := student{Id: id, Name: name, Domain: domain}
	db.Create(&data)
	fmt.Println("Value Inserted!!!")
}
func Read(db *gorm.DB) {

	var Student []student
	db.Find(&Student)
	fmt.Println("id   name   domain")
	for _, Student := range Student {
		fmt.Printf("%d - %s - %s \n", Student.Id, Student.Name, Student.Domain)
	}
}

func Update(db *gorm.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	fmt.Println("Enter new name:")
	fmt.Scan(&name)
	fmt.Println("Enter new domain:")
	fmt.Scan(&domain)
	var Student student
	db.First(&Student, id)
	Student.Name = name
	Student.Domain = domain
	db.Save(&Student)
	fmt.Println("Updated!!!")
}
func Delete(db *gorm.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	var Student student
	db.First(&Student, id)
	db.Delete(&Student)
	fmt.Println("Deleted!!!")
}
