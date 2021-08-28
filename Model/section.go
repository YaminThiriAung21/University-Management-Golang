package Model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Section struct {
	Id       uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name     string     `gorm:"size:255;not null" json:"name"`
	Students []Student  `gorm:"ForeignKey:SectionID"`
	Teachers []*Teacher `gorm:"many2many:teacher_section"`
	ClassID  uint       `gorm:"column:class_id"`
	Class    Class
}

func (sec *Section) GetAllSection(db *gorm.DB) (*[]Section, error) {

	sections := []Section{}

	err := db.Debug().Model(&Section{}).Preload("Students").Preload("Teachers").Preload("Class").Limit(100).Find(&sections).Error
	if err != nil {
		return &[]Section{}, err
	}
	return &sections, err
}
func (sec *Section) GetSectionByID(db *gorm.DB, secid uint) (*Section, error) {

	err := db.Debug().Model(Section{}).Preload("Class").Preload("Students").Preload("Teachers").Where("id=?", secid).Take(&sec).Error
	if err != nil {
		return &Section{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Section{}, errors.New("section not found")
	}
	return sec, err
}
func (sec *Section) CreateSection(db *gorm.DB) (*Section, error) {

	err := db.Omit("Class").Preload("Class").Create(&sec).Error
	if err != nil {
		return &Section{}, err
	}
	return sec, nil
}
func (sec *Section) UpdateASection(db *gorm.DB, secid uint) (*Section, error) {
	db = db.Debug().Model(&Section{}).Preload("Class").Where("id = ?", secid).Take(&Section{}).UpdateColumns(
		map[string]interface{}{
			"name":     sec.Name,
			"class_id": sec.ClassID,
		},
	)
	if db.Error != nil {
		return &Section{}, db.Error
	}
	err := db.Debug().Model(&Section{}).Preload("Class").Where("id = ?", secid).Take(&sec).Error
	if err != nil {
		return &Section{}, err
	}
	return sec, nil
}

func (sec *Section) DeleteASection(db *gorm.DB, secid uint) (int64, error) {
	db = db.Debug().Model(&Section{}).Preload("Class").Where("id = ?", secid).Take(&Section{}).Delete(&Section{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
func (sec *Section) GetSectionsByClassID(db *gorm.DB, cid uint) (*[]Section, error) {
	sections := []Section{}

	class := Class{}
	db.First(&class, "id = ?", cid)
	db.Model(&class).Association("Sections").Find(&sections)
	//err := db.Debug().Model(&Class{}).Preload("Teachers").Where("teacher_id = ?", tid).Take(&Class{}).Limit(100).Find(&classes).Error
	if db.Error != nil {
		return &[]Section{}, db.Error
	}
	return &sections, nil
}
func (sec *Section) GetSectionsByTeacherID(db *gorm.DB, tid uint) (*[]Section, error) {
	sections := []Section{}

	teacher := Teacher{}
	db.First(&teacher, "id = ?", tid)
	db.Model(&teacher).Association("Sections").Find(&sections)
	//err := db.Debug().Model(&Class{}).Preload("Teachers").Where("teacher_id = ?", tid).Take(&Class{}).Limit(100).Find(&classes).Error
	if db.Error != nil {
		return &[]Section{}, db.Error
	}
	return &sections, nil
}
func (s *Section) ToString() string {

	return fmt.Sprintf("SectionId: %d  SectionName: %s", s.Id, s.Name)
}
func (s *Section) ToStringSec() string {
	return fmt.Sprintf("SectionId: %d  \nSectionName: %s  \nClass: %s", s.Id, s.Name, s.Class.Name)
}
