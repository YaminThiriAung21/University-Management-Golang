package Model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	Id          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	ClassID     uint   `gorm:"column:class_id"`
	Class       Class
	SectionID   uint `gorm:"column:section_id"`
	Section     Section
	Age         uint   `json:"age"`
	Gender      string `json:"gender"`
	Address     string `gorm:"size:255;" json:"address"`
	ParentsName string `gorm:"size:255;" json:"parentsname"`
}

func (s *Student) GetAllStudents(db *gorm.DB) (*[]Student, error) {

	students := []Student{}

	err := db.Debug().Model(&Student{}).Limit(100).Preload("Class").Preload("Section").Find(&students).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &students, err
}
func (s *Student) GetStudentByID(db *gorm.DB, sid uint) (*Student, error) {

	err := db.Debug().Model(Student{}).Preload("Class").Preload("Section").Where("id=?", sid).Take(&s).Error
	if err != nil {
		return &Student{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Student{}, errors.New("student not found")
	}
	return s, err
}
func (s *Student) CreateStudent(db *gorm.DB) (*Student, error) {

	err := db.Omit("Section").Preload("Class").Preload("Section").Create(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil
}
func (s *Student) UpdateAStudent(db *gorm.DB, sid uint) (*Student, error) {
	db = db.Debug().Model(&Student{}).Where("id = ?", sid).Take(&Student{}).UpdateColumns(
		map[string]interface{}{
			"name":         s.Name,
			"class_id":     s.ClassID,
			"section_id":   s.SectionID,
			"age":          s.Age,
			"gender":       s.Gender,
			"address":      s.Address,
			"parents_name": s.ParentsName,
		},
	)
	if db.Error != nil {
		return &Student{}, db.Error
	}
	err := db.Debug().Model(&Student{}).Preload("Class").Preload("Section").Where("id = ?", sid).Take(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil
}

func (s *Student) DeleteAStudent(db *gorm.DB, sid uint) (int64, error) {
	db = db.Debug().Model(&Student{}).Where("id = ?", sid).Take(&Student{}).Delete(&Student{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *Student) GetStudentsByClassID(db *gorm.DB, cid uint) (*[]Student, error) {
	students := []Student{}

	err := db.Debug().Model(&Student{}).Preload("Class").Preload("Section").Where("class_id=?", cid).Take(&Student{}).Limit(100).Find(&students).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &students, err
}
func (s *Student) GetStudentsByClassSectionID(db *gorm.DB, cid uint, secid uint) (*[]Student, error) {
	students := []Student{}

	err := db.Debug().Model(&Student{}).Preload("Class").Preload("Section").Where("class_id = ? AND section_id = ?", cid, secid).Take(&Student{}).Limit(100).Find(&students).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &students, err
}
func (s *Student) ToString() string {

	return fmt.Sprintf("StudentId: %d\nName: %s\nAge: %d\nAddress: %s\nGender: %s\nParents Name: %s\nClass: %s\nSession: %s", s.Id, s.Name, s.Age, s.Address, s.Gender, s.ParentsName, s.Class.Name, s.Section.Name)
}
