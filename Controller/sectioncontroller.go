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

func CreateSection(w http.ResponseWriter, r *http.Request) {
	section := Model.Section{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = json.Unmarshal(body, &section)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	secCreated, err := section.CreateSection(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, secCreated.ToStringSec())
	fmt.Fprintln(w, "========================")
}

func GetSections(w http.ResponseWriter, r *http.Request) {

	section := Model.Section{}
	//getStudents := Database.Database.Model(&students).Association("Class").Find(&students)
	getsections, err := section.GetAllSection(Database.Database)
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprintln(w, "Total Sections: ", len(*getsections))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getsections) > 0 {
		for _, sec := range *getsections {
			fmt.Fprint(w, sec.ToStringSec())

			fmt.Fprintln(w, "\n========================")
		}
	}

}

func GetSectionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	section := Model.Section{}
	getsection, err := section.GetSectionByID(Database.Database, uint(secid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprint(w, getsection.ToStringSec())
	fmt.Fprint(w, "\nStudents: ")
	for _, v := range section.Students {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}
	fmt.Fprint(w, "\nTeachers: ")
	for _, v := range section.Teachers {
		fmt.Fprint(w, v.Name)
		fmt.Fprint(w, ",")
	}

	fmt.Fprintln(w, "\n========================")
}
func UpdateSection(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	secid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	section := Model.Section{}
	err = json.Unmarshal(body, &section)
	if err != nil {
		Response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedSection, err := section.UpdateASection(Database.Database, uint(secid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, updatedSection.ToStringSec())
	fmt.Fprintln(w, "========================")
}
func DeleteSection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	section := Model.Section{}
	secid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = section.DeleteASection(Database.Database, uint(secid))
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Fprintln(w, "Deleted")
	fmt.Fprintln(w, "========================")
}

func GetSectionsByClassID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	section := Model.Section{}
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getSections, err := section.GetSectionsByClassID(Database.Database, uint(cid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	classes := Model.Class{}
	class, _ := classes.GetClassByID(Database.Database, uint(cid))
	fmt.Fprintln(w, class.ToString())
	fmt.Fprintln(w, "Total Sections: ", len(*getSections))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getSections) > 0 {
		for _, sec := range *getSections {
			fmt.Fprintln(w, "\n", sec.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
func GetSectionsByTeacherID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	section := Model.Section{}
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	getSections, err := section.GetSectionsByTeacherID(Database.Database, uint(tid))
	if err != nil {
		Response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	teachers := Model.Teacher{}
	teacher, _ := teachers.GetTeacherByID(Database.Database, uint(tid))
	fmt.Fprintln(w, teacher.ToString())
	fmt.Fprintln(w, "Total Sections: ", len(*getSections))
	fmt.Fprintln(w, "---------------------------------------")
	if len(*getSections) > 0 {
		for _, sec := range *getSections {
			fmt.Fprintln(w, "\n", sec.ToString())
			fmt.Fprintln(w, "========================")
		}
	}
}
