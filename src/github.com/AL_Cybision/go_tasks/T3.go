import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name string
	address string
}

func NewStudent(rollno int , name string  ,address string)
*Student {
	s:= new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}


type StudentList struct {
	list []*Student
}

func (1s *StudentList) CreateStudent (rollno int, name string, address string) *Student {
	st := NewStudent (rollno, name, address) 
	ls.list = append(ls.list, st) 
	return st
}

func main() {
		student := new (StudentList)
		student.CreateStudent (24, "Asim", "AAAAAA") 
		student.CreateStudent (25, "Naveed", "BBBBBB")
		fmt.println(student)
}