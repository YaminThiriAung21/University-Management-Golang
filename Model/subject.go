package Model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Subject struct {
	Id       uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name     string     `gorm:"size:255;not null" json:"name"`
	Teachers []*Teacher `gorm:"many2many:teacher_subject"`
	Classes  []*Class   `gorm:"many2many:class_subject"`
}

func (sbj *Subject) GetAllSubject(db *gorm.DB) (*[]Subject, error) {

	subjects := []Subject{}

	err := db.Debug().Model(&Subject{}).Preload("Teachers").Preload("Classes").Limit(100).Find(&subjects).Error
	if err != nil {
		return &[]Subject{}, err
	}
	return &subjects, err
}
func (sbj *Subject) GetSubjectByID(db *gorm.DB, sbid uint) (*Subject, error) {

	err := db.Debug().Model(Subject{}).Preload("Teachers").Preload("Classes").Where("id=?", sbid).Take(&sbj).Error
	if err != nil {
		return &Subject{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Subject{}, errors.New("subject not found")
	}
	return sbj, err
}
func (sbj *Subject) CreateSubject(db *gorm.DB) (*Subject, error) {

	err := db.Create(&sbj).Error
	if err != nil {
		return &Subject{}, err
	}
	return sbj, nil
}
func (sbj *Subject) UpdateASubject(db *gorm.DB, sbid uint) (*Subject, error) {
	db = db.Debug().Model(&Subject{}).Where("id = ?", sbid).Take(&Subject{}).UpdateColumns(
		map[string]interface{}{
			"name": sbj.Name,
		},
	)
	if db.Error != nil {
		return &Subject{}, db.Error
	}
	err := db.Debug().Model(&Subject{}).Where("id = ?", sbid).Take(&sbj).Error
	if err != nil {
		return &Subject{}, err
	}
	return sbj, nil
}

func (sbj *Subject) DeleteASubject(db *gorm.DB, sbid uint) (int64, error) {
	db = db.Debug().Model(&Subject{}).Where("id = ?", sbid).Take(&Subject{}).Delete(&Subject{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
func (s *Subject) GetSubjectByTeacher(db *gorm.DB, tid uint) (*[]Subject, error) {
	subject := []Subject{}
	teacher := Teacher{}
	db.First(&teacher, "id = ?", tid)
	db.Model(&teacher).Association("Subjects").Find(&subject)
	//err := db.Debug().Model(&Class{}).Preload("Teachers").Where("teacher_id = ?", tid).Take(&Class{}).Limit(100).Find(&classes).Error
	if db.Error != nil {
		return &[]Subject{}, db.Error
	}
	return &subject, nil
}
func (s *Subject) GetSubjectByClass(db *gorm.DB, cid uint) (*[]Subject, error) {
	classes := Class{}
	subject := []Subject{}
	db.First(&classes, "id = ?", cid)
	db.Model(&classes).Association("Subjects").Find(&subject)
	//err := db.Debug().Model(&Class{}).Preload("Teachers").Where("teacher_id = ?", tid).Take(&Class{}).Limit(100).Find(&classes).Error
	if db.Error != nil {
		return &[]Subject{}, db.Error
	}
	return &subject, nil
}

func (sub *Subject) ToString() string {

	return fmt.Sprintf("SubjectId: %d  SubjectName: %s", sub.Id, sub.Name)

}
