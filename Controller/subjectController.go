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

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	subject := Model.Subject{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = json.Unmarshal(body, &subject)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	subjectCreated, err := subject.CreateSubject(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, subjectCreated.ToString())
	fmt.Fprintln(w, "========================")
}

func GetSubjects(w http.ResponseWriter, r *http.Request) {

	subject := Model.Subject{}
	getsubjects, err := subject.GetAllSubject(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprintln(w, "Total Subjects: ", len(*getsubjects))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getsubjects) > 0 {
		for _, sub := range *getsubjects {
			fmt.Fprintln(w, sub.ToString())

			fmt.Fprintln(w, "\n========================")
		}
	}

}

func GetSubjectByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	subject := Model.Subject{}
	getsubject, err := subject.GetSubjectByID(Database.Database, uint(subid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, getsubject.ToString())
	fmt.Fprint(w, "\nTeachers: ")
	for _, v := range subject.Teachers {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprint(w, "\nClasses: ")
	for _, v := range subject.Classes {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprintln(w, "\n========================")
}
func UpdateSubject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	subid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	subject := Model.Subject{}
	err = json.Unmarshal(body, &subject)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedsubject, err := subject.UpdateASubject(Database.Database, uint(subid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, updatedsubject.ToString())
	fmt.Fprintln(w, "========================")
}
func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subject := Model.Subject{}
	subid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = subject.DeleteASubject(Database.Database, uint(subid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, "Deleted")
	fmt.Fprintln(w, "========================")
}

func GetSubjectByClassID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subject := Model.Subject{}
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getSub, err := subject.GetSubjectByClass(Database.Database, uint(cid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	classes := Model.Class{}
	class, _ := classes.GetClassByID(Database.Database, uint(cid))
	fmt.Fprintln(w, class.ToString())
	fmt.Fprintln(w, "Total Subjects: ", len(*getSub))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getSub) > 0 {
		for _, sub := range *getSub {
			fmt.Fprintln(w, "\n", sub.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
func GetSubjectByTeacherID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subject := Model.Subject{}
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getSub, err := subject.GetSubjectByTeacher(Database.Database, uint(tid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	teachers := Model.Teacher{}
	teacher, _ := teachers.GetTeacherByID(Database.Database, uint(tid))
	fmt.Fprintln(w, teacher.ToString())
	fmt.Fprintln(w, "Total Sections: ", len(*getSub))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getSub) > 0 {
		for _, sec := range *getSub {
			fmt.Fprintln(w, "\n", sec.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
