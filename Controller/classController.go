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

func CreateClass(w http.ResponseWriter, r *http.Request) {
	class := Model.Class{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = json.Unmarshal(body, &class)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	classCreated, err := class.CreateClass(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, classCreated.ToString())
	fmt.Fprintln(w, "========================")
}

func GetClasses(w http.ResponseWriter, r *http.Request) {

	class := Model.Class{}
	//getStudents := Database.Database.Model(&students).Association("Class").Find(&students)
	getclassess, err := class.GetAllClasses(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprintln(w, "Total Classes: ", len(*getclassess))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getclassess) > 0 {
		for _, classes := range *getclassess {
			fmt.Fprintln(w, classes.ToString())
			fmt.Fprint(w, "\nStudents: ")
			for _, v := range classes.Students {
				fmt.Fprint(w, v.Name)
				fmt.Fprint(w, ",")
			}
			fmt.Fprint(w, "\nTeachers: ")
			for _, v := range classes.Teachers {
				fmt.Fprint(w, v.Name)
				fmt.Fprint(w, ",")
			}
			fmt.Fprint(w, "\nSections: ")
			for _, v := range classes.Sections {
				fmt.Fprint(w, v.Name)
				fmt.Fprint(w, ",")
			}
			fmt.Fprint(w, "\nSubjects: ")
			for _, v := range classes.Subjects {
				fmt.Fprint(w, v.Name)
				fmt.Fprint(w, ",")
			}
			fmt.Fprintln(w, "\n========================")
		}
	}

}

func GetClassByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	classes := Model.Class{}
	getclass, err := classes.GetClassByID(Database.Database, uint(cid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, getclass.ToString())
	fmt.Fprint(w, "\nStudents: ")
	for _, v := range classes.Students {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprint(w, "\nTeachers: ")
	for _, v := range classes.Teachers {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprint(w, "\nSections: ")
	for _, v := range classes.Sections {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprint(w, "\nSubjects: ")
	for _, v := range classes.Subjects {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprintln(w, "\n========================")
}
func UpdateClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	class := Model.Class{}
	err = json.Unmarshal(body, &class)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedClass, err := class.UpdateAClass(Database.Database, uint(cid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, updatedClass.ToString())
	fmt.Fprintln(w, "========================")
}
func DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	class := Model.Class{}
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = class.DeleteAClass(Database.Database, uint(cid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, "Deleted")
	fmt.Fprintln(w, "========================")
}
func GetClassByTeacherID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	class := Model.Class{}
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getclassess, err := class.GetClassByTeacherID(Database.Database, uint(tid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	teachers := Model.Teacher{}
	teacher, _ := teachers.GetTeacherByID(Database.Database, uint(tid))
	fmt.Fprintln(w, teacher.ToString())
	fmt.Fprintln(w, "Total Classes: ", len(*getclassess))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getclassess) > 0 {
		for _, sec := range *getclassess {
			fmt.Fprintln(w, "\n", sec.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
func GetClassBySubjectID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	class := Model.Class{}
	subid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getclassess, err := class.GetClassBySubjectID(Database.Database, uint(subid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	subjects := Model.Subject{}
	sub, _ := subjects.GetSubjectByID(Database.Database, uint(subid))
	fmt.Fprintln(w, sub.ToString())
	fmt.Fprintln(w, "Total Classes: ", len(*getclassess))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getclassess) > 0 {
		for _, sec := range *getclassess {
			fmt.Fprintln(w, "\n", sec.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
