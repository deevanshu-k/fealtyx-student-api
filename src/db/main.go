package db

import (
	"errors"

	"github.com/deevanshu-k/fealtyx-student-api/src/structs"
	"github.com/deevanshu-k/fealtyx-student-api/src/utils"
)

var db struct {
	students map[string]structs.Student
}

func init() {
	db.students = make(map[string]structs.Student)
}

func GetAllStudents() ([]structs.Student, error) {
	var students []structs.Student
	for _, student := range db.students {
		students = append(students, student)
	}
	return students, nil
}

func GetStudent(id string) (structs.Student, error) {
	if student, ok := db.students[id]; ok {
		return student, nil
	}
	return structs.Student{}, errors.New("student not found")
}

func CreateStudent(student structs.Student) error {
	nanoId, err := utils.GenerateNanoId(10)
	if err != nil {
		return err
	}
	student.ID = nanoId

	if _, ok := db.students[student.ID]; ok {
		return errors.New("student already exists")
	}
	db.students[student.ID] = student
	return nil
}

func UpdateStudent(id string, student structs.Student) error {
	if _, ok := db.students[id]; !ok {
		return errors.New("student not found")
	}
	db.students[id] = student
	return nil
}

func DeleteStudent(id string) error {
	if _, ok := db.students[id]; !ok {
		return errors.New("student not found")
	}
	delete(db.students, id)
	return nil
}
