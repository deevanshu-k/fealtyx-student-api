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
	var students = make([]structs.Student, 0)
	for _, student := range db.students {
		students = append(students, student)
	}
	return students, nil
}

func GetStudent(id string) (structs.Student, error) {
	if student, ok := db.students[id]; ok {
		return student, nil
	}
	return structs.Student{}, errors.New("STUDENT_NOT_FOUND")
}

func CreateStudent(student structs.Student) (structs.Student, error) {
	nanoId, err := utils.GenerateNanoId(10)
	if err != nil {
		return structs.Student{}, err
	}
	student.ID = nanoId

	if _, ok := db.students[student.ID]; ok {
		return structs.Student{}, errors.New("STUDENT_ALREADY_EXISTS")
	}
	db.students[student.ID] = student
	return student, nil
}

func UpdateStudent(student structs.Student) (structs.Student, error) {
	if _, ok := db.students[student.ID]; !ok {
		return structs.Student{}, errors.New("STUDENT_NOT_FOUND")
	}
	db.students[student.ID] = student
	return student, nil
}

func DeleteStudent(id string) error {
	if _, ok := db.students[id]; !ok {
		return errors.New("STUDENT_NOT_FOUND")
	}
	delete(db.students, id)
	return nil
}
