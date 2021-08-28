package Model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Teacher struct {
	Id       uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name     string     `gorm:"size:255;not null" json:"name"`
	Subjects []*Subject `gorm:"many2many:teacher_subject"`
	Classes  []*Class   `gorm:"many2many:teacher_class"`
	Sections []*Section `gorm:"many2many:teacher_section"`
}

func (t *Teacher) GetAllTeachers(db *gorm.DB) (*[]Teacher, error) {

	teachers := []Teacher{}
	db.Model(&teachers).Preload("Subjects").Preload("Classes").Preload("Sections").Find(&teachers)

	//err := db.Debug().Model(&Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Limit(100).Find(&teachers).Error
	if db.Error != nil {
		return &[]Teacher{}, db.Error
	}

	return &teachers, db.Error
}
func (t *Teacher) GetTeacherByID(db *gorm.DB, tid uint) (*Teacher, error) {

	err := db.Debug().Model(Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Where("id=?", tid).Take(&t).Error
	if err != nil {
		return &Teacher{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Teacher{}, errors.New("teacher not found")
	}
	return t, err
}
func (t *Teacher) CreateTeacher(db *gorm.DB) (*Teacher, error) {
	teacher := Teacher{
		Id:   1,
		Name: "tr1",
		Subjects: []*Subject{
			{Id: 1},
			{Id: 2},
		},
		Classes: []*Class{
			{Id: 1},
		},
		Sections: []*Section{
			{Id: 1},
		},
	}
	db.Create(&teacher)
	if db.Error != nil {
		return &Teacher{}, db.Error
	}
	return &teacher, nil
}
func (t *Teacher) UpdateATeacher(db *gorm.DB, tid uint) (*Teacher, error) {
	db = db.Debug().Model(&Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Where("id = ?", tid).Take(&Teacher{}).UpdateColumns(
		map[string]interface{}{
			"name": t.Name,
		},
	)
	if db.Error != nil {
		return &Teacher{}, db.Error
	}
	err := db.Debug().Model(&Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Where("id = ?", tid).Take(&t).Error
	if err != nil {
		return &Teacher{}, err
	}
	return t, nil
}

func (t *Teacher) DeleteATeacher(db *gorm.DB, tid uint) (int64, error) {
	db = db.Debug().Model(&Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Where("id = ?", tid).Take(&Teacher{}).Delete(&Teacher{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (t *Teacher) GetTeacherByClassID(db *gorm.DB, cid uint) (*[]Teacher, error) {
	teachers := []Teacher{}

	class := Class{}

	db.First(&class, "id = ?", cid)

	//err := db.Model(&class).Related(&teachers, "Teachers")
	db.Model(&class).Association("Teachers").Find(&teachers)

	//db.Debug().Model(&Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Where("class_id = ?", cid).Take(&Teacher{}).Limit(100).Find(&teachers).Error
	if db.Error != nil {
		return &[]Teacher{}, db.Error
	}
	return &teachers, db.Error
}
func (t *Teacher) GetTeacherBySectionID(db *gorm.DB, secid uint) (*[]Teacher, error) {
	teachers := []Teacher{}
	section := Section{}
	db.First(&section, "id = ?", secid)
	db.Debug().Model(&section).Association("Teachers").Find(&teachers)
	//err := db.Debug().Model(&Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Where("section_id = ?", secid).Take(&Teacher{}).Limit(100).Find(&teachers).Error
	if db.Error != nil {
		return &[]Teacher{}, db.Error
	}
	return &teachers, db.Error
}
func (t *Teacher) GetTeacherBySubjectID(db *gorm.DB, sub uint) (*[]Teacher, error) {
	teachers := []Teacher{}
	subject := Subject{}
	db.First(&subject, "id = ?", sub)
	db.Model(&subject).Association("Teachers").Find(&teachers)
	//err := db.Debug().Model(&Teacher{}).Preload("Subjects").Preload("Classes").Preload("Sections").Where("subject_id = ?", sub).Take(&Teacher{}).Limit(100).Find(&teachers).Error
	if db.Error != nil {
		return &[]Teacher{}, db.Error
	}
	return &teachers, db.Error
}
func (t *Teacher) ToString() string {

	return fmt.Sprintf("TeacherId: %d  TeacherName: %s", t.Id, t.Name)

}
func (t *Teacher) ToStringtr() string {

	return fmt.Sprintf("TeacherId: %d  \nTeacherName: %s", t.Id, t.Name)

}
