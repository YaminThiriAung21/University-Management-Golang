package Controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/YaminThiriAung21/UniversityGolang/Database"
	"github.com/YaminThiriAung21/UniversityGolang/Model"
	"github.com/YaminThiriAung21/UniversityGolang/Response"
	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	student := Model.Student{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = json.Unmarshal(body, &student)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	studentCreated, err := student.CreateStudent(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, studentCreated.ToString())
	fmt.Fprintln(w, "========================")
}

func GetStudents(w http.ResponseWriter, r *http.Request) {

	students := Model.Student{}
	//getStudents := Database.Database.Model(&students).Association("Class").Find(&students)
	getStudents, err := students.GetAllStudents(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprintln(w, "Total Students: ", len(*getStudents))
	fmt.Fprintln(w, "---------------------------------------")

	if len(*getStudents) > 0 {
		for _, students := range *getStudents {
			fmt.Fprintln(w, students.ToString())
			fmt.Fprintln(w, "========================")
		}
	}

}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	student := Model.Student{}
	getStudent, err := student.GetStudentByID(Database.Database, uint(sid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, getStudent.ToString())
	fmt.Fprintln(w, "========================")
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student := Model.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedStudent, err := student.UpdateAStudent(Database.Database, uint(sid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, updatedStudent.ToString())
	fmt.Fprintln(w, "========================")
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student := Model.Student{}
	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = student.DeleteAStudent(Database.Database, uint(sid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, "Deleted")
	fmt.Fprintln(w, "========================")
}

func GetStudentsByClassID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	students := Model.Student{}
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getStudents, err := students.GetStudentsByClassID(Database.Database, uint(cid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	classes := Model.Class{}
	class, _ := classes.GetClassByID(Database.Database, uint(cid))
	fmt.Fprintln(w, class.ToString())
	fmt.Fprintln(w, "Total Students: ", len(*getStudents))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getStudents) > 0 {
		for _, students := range *getStudents {
			fmt.Fprintln(w, "\n", students.ToString())
			fmt.Fprintln(w, "========================")
		}
	}

}
func GetStudentsByClassSectionID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	students := Model.Student{}
	cid, err := strconv.ParseUint(vars["cid"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	sid, err := strconv.ParseUint(vars["sid"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getStudents, err := students.GetStudentsByClassSectionID(Database.Database, uint(cid), uint(sid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	classes := Model.Class{}
	class, _ := classes.GetClassByID(Database.Database, uint(cid))
	sections := Model.Section{}
	section, _ := sections.GetSectionByID(Database.Database, uint(sid))
	fmt.Fprintln(w, "Class\n", class.ToString())
	fmt.Fprintln(w, "Section\n", section.ToString())
	fmt.Fprintln(w, "Total Students: ", len(*getStudents))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getStudents) > 0 {
		for _, students := range *getStudents {
			fmt.Fprintln(w, students.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
