package Route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YaminThiriAung21/UniversityGolang/Controller"
	"github.com/gorilla/mux"
)

func LoadRoute() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)

	router.HandleFunc("/student", Controller.GetStudents).Methods("GET")
	router.HandleFunc("/student", Controller.CreateStudent).Methods("POST")
	router.HandleFunc("/student/{id}", Controller.GetStudentByID).Methods("GET")
	router.HandleFunc("/student/{id}", Controller.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{id}", Controller.DeleteStudent).Methods("DELETE")

	router.HandleFunc("/studentbyclass/{id}", Controller.GetStudentsByClassID).Methods("GET")
	router.HandleFunc("/studentbyclasssection/{cid}{sid}", Controller.GetStudentsByClassSectionID).Methods("GET")
	//-------------------

	router.HandleFunc("/teacher", Controller.GetTeachers).Methods("GET")
	router.HandleFunc("/teacher", Controller.CreateTeacher).Methods("POST")
	router.HandleFunc("/teacher/{id}", Controller.GetTeacherByID).Methods("GET")
	router.HandleFunc("/teacher/{id}", Controller.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/teacher/{id}", Controller.DeleteTeacher).Methods("DELETE")

	router.HandleFunc("/teacherbyclass/{id}", Controller.GetTeachersByClassID).Methods("GET")
	router.HandleFunc("/teacherbysection/{id}", Controller.GetTeachersBySectionID).Methods("GET")
	router.HandleFunc("/teacherbysubject/{id}", Controller.GetTeachersBySubjectID).Methods("GET")
	//---------------

	router.HandleFunc("/class", Controller.GetClasses).Methods("GET")
	router.HandleFunc("/class", Controller.CreateClass).Methods("POST")
	router.HandleFunc("/class/{id}", Controller.GetClassByID).Methods("GET")
	router.HandleFunc("/class/{id}", Controller.UpdateClass).Methods("PUT")
	router.HandleFunc("/class/{id}", Controller.DeleteClass).Methods("DELETE")

	router.HandleFunc("/classbyteacher/{id}", Controller.GetClassByTeacherID).Methods("GET")
	router.HandleFunc("/classbysubject/{id}", Controller.GetClassBySubjectID).Methods("GET")
	//--------------

	router.HandleFunc("/section", Controller.GetSections).Methods("GET")
	router.HandleFunc("/section", Controller.CreateSection).Methods("POST")
	router.HandleFunc("/section/{id}", Controller.GetSectionByID).Methods("GET")
	router.HandleFunc("/section/{id}", Controller.UpdateSection).Methods("PUT")
	router.HandleFunc("/section/{id}", Controller.DeleteSection).Methods("DELETE")

	router.HandleFunc("/sectionbyclass/{id}", Controller.GetSectionsByClassID).Methods("GET")
	router.HandleFunc("/sectionbyteacher/{id}", Controller.GetSectionsByTeacherID).Methods("GET")
	//--------------

	router.HandleFunc("/subject", Controller.GetSubjects).Methods("GET")
	router.HandleFunc("/subject", Controller.CreateSubject).Methods("POST")
	router.HandleFunc("/subject/{id}", Controller.GetSubjectByID).Methods("GET")
	router.HandleFunc("/subject/{id}", Controller.UpdateSubject).Methods("PUT")
	router.HandleFunc("/subject/{id}", Controller.DeleteSubject).Methods("DELETE")

	router.HandleFunc("/subjectbyclass/{id}", Controller.GetSubjectByClassID).Methods("GET")
	router.HandleFunc("/subjectbyteacher/{id}", Controller.GetSubjectByTeacherID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome From School Management Project GoLang")

}
