package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

func calc_hash(stringToHash string) string {

	fmt.Printf("String recieved to hash : %s\n ", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))

}

type Student struct {
	rollnumber int
	name       string
	address    string
	subjects   []string
}

func NewStudent(rollno int, name string, address string, subjects []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subjects = subjects
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, subjects []string) *Student {
	st := NewStudent(rollno, name, address, subjects)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) Print() {
	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student Rollno: %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student Name: %s\n", ls.list[i].name)
		fmt.Printf("Student Address: %s\n\n", ls.list[i].address)
		fmt.Printf("Student Subjects: %s\n\n", ls.list[i].subjects)
		strroll := strconv.Itoa(ls.list[i].rollnumber)
		var str string
		str += strroll + ls.list[i].name + ls.list[i].address // + ls.list[i].subjects
		fmt.Printf("Hash of student block: %s\n\n", calc_hash(str))

	}
}

func main() {
	var subjects1 = []string{"Math1", "Urdu1", "English1"}
	var subjects2 = []string{"Math2", "Urdu2", "English2"}
	var subjects3 = []string{"Math3", "Urdu3", "English3"}
	student := new(StudentList)
	student.CreateStudent(2000, "Noman", "NewYork", subjects1)

	student.CreateStudent(1965, "Ilyas", "ISB", subjects2)

	student.CreateStudent(7853, "Xmen", "Arctic", subjects3)
	student.Print()
}
