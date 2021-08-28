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

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	teacher := Model.Teacher{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teachercreated, err := teacher.CreateTeacher(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, teachercreated.ToString())
	fmt.Fprintln(w, "========================")
}

func GetTeachers(w http.ResponseWriter, r *http.Request) {

	teacher := Model.Teacher{}
	//getStudents := Database.Database.Model(&students).Association("Class").Find(&students)
	getteachers, err := teacher.GetAllTeachers(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	//json.NewEncoder(w).Encode(getteachers)
	fmt.Fprintln(w, "Total Teachers: ", len(*getteachers))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getteachers) > 0 {
		for _, teacher := range *getteachers {
			fmt.Fprint(w, teacher.ToStringtr())
			fmt.Fprint(w, "\nClasses: ")
			for _, v := range teacher.Classes {
				fmt.Fprint(w, v.Name)
				fmt.Fprint(w, ",")
			}
			fmt.Fprint(w, "\nSections: ")
			for _, v := range teacher.Sections {
				fmt.Fprint(w, v.Name)
				fmt.Fprint(w, ",")
			}
			fmt.Fprint(w, "\nSubjects: ")
			for _, v := range teacher.Subjects {
				fmt.Fprint(w, v.Name)
				fmt.Fprint(w, ",")
			}
			fmt.Fprintln(w, "\n========================")

		}
	}

}

func GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	teacher := Model.Teacher{}
	getteacher, err := teacher.GetTeacherByID(Database.Database, uint(tid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(w, getteacher.ToStringtr())
	fmt.Fprint(w, "\nClasses: ")
	for _, v := range teacher.Classes {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprint(w, "\nSections: ")
	for _, v := range teacher.Sections {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprint(w, "\nSubjects: ")
	for _, v := range teacher.Subjects {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprintln(w, "\n========================")
}
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacher := Model.Teacher{}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedTeacher, err := teacher.UpdateATeacher(Database.Database, uint(tid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, updatedTeacher.ToString())
	fmt.Fprintln(w, "========================")
}
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacher := Model.Teacher{}
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = teacher.DeleteATeacher(Database.Database, uint(tid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, "Deleted")
	fmt.Fprintln(w, "========================")
}

func GetTeachersByClassID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teacher := Model.Teacher{}
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getTeacher, err := teacher.GetTeacherByClassID(Database.Database, uint(cid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	classes := Model.Class{}
	class, _ := classes.GetClassByID(Database.Database, uint(cid))
	fmt.Fprintln(w, class.ToString())
	fmt.Fprintln(w, "Total Teachers: ", len(*getTeacher))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getTeacher) > 0 {
		for _, teacher := range *getTeacher {
			fmt.Fprintln(w, "\n", teacher.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
func GetTeachersBySectionID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacher := Model.Teacher{}
	secid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getTeacher, err := teacher.GetTeacherBySectionID(Database.Database, uint(secid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	sections := Model.Section{}
	sec, _ := sections.GetSectionByID(Database.Database, uint(secid))
	fmt.Fprintln(w, sec.ToString())
	fmt.Fprintln(w, "Total Teachers: ", len(*getTeacher))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getTeacher) > 0 {
		for _, teacher := range *getTeacher {
			fmt.Fprintln(w, "\n", teacher.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
func GetTeachersBySubjectID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacher := Model.Teacher{}
	subid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getTeacher, err := teacher.GetTeacherBySubjectID(Database.Database, uint(subid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	subject := Model.Subject{}
	sec, _ := subject.GetSubjectByID(Database.Database, uint(subid))
	fmt.Fprintln(w, sec.ToString())
	fmt.Fprintln(w, "Total Teachers: ", len(*getTeacher))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getTeacher) > 0 {
		for _, teacher := range *getTeacher {
			fmt.Fprintln(w, "\n", teacher.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
