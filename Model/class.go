package Model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Class struct {
	Id       uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name     string     `gorm:"size:255;not null" json:"name"`
	Students []Student  `gorm:"ForeignKey:ClassID"`
	Teachers []*Teacher `gorm:"many2many:teacher_class"`
	Sections []Section  `gorm:"ForeignKey:ClassID"`
	Subjects []*Subject `gorm:"many2many:class_subject"`
}

func (c *Class) GetAllClasses(db *gorm.DB) (*[]Class, error) {

	classes := []Class{}

	err := db.Debug().Model(&Class{}).Preload("Students").Preload("Teachers").Preload("Sections").Preload("Subjects").Limit(100).Find(&classes).Error
	if err != nil {
		return &[]Class{}, err
	}
	return &classes, err
}
func (c *Class) GetClassByID(db *gorm.DB, cid uint) (*Class, error) {

	err := db.Debug().Model(Class{}).Preload("Students").Preload("Teachers").Preload("Sections").Preload("Subjects").Where("id=?", cid).Take(&c).Error
	if err != nil {
		return &Class{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Class{}, errors.New("class not found")
	}
	return c, err
}
func (c *Class) CreateClass(db *gorm.DB) (*Class, error) {

	err := db.Create(&c).Error
	if err != nil {
		return &Class{}, err
	}
	return c, nil
}
func (c *Class) UpdateAClass(db *gorm.DB, cid uint) (*Class, error) {
	db = db.Debug().Model(&Class{}).Where("id = ?", cid).Preload("Students").Preload("Teachers").Preload("Sections").Preload("Subjects").Take(&Class{}).UpdateColumns(
		map[string]interface{}{
			"name": c.Name,
		},
	)
	if db.Error != nil {
		return &Class{}, db.Error
	}
	err := db.Debug().Model(&Teacher{}).Where("id = ?", cid).Take(&c).Error
	if err != nil {
		return &Class{}, err
	}
	return c, nil
}

func (c *Class) DeleteAClass(db *gorm.DB, cid uint) (int64, error) {
	db = db.Debug().Model(&Class{}).Where("id = ?", cid).Take(&Class{}).Delete(&Class{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
func (c *Class) GetClassByTeacherID(db *gorm.DB, tid uint) (*[]Class, error) {
	classes := []Class{}
	teacher := Teacher{}
	db.First(&teacher, "id = ?", tid)
	db.Model(&teacher).Association("Classes").Find(&classes)
	//err := db.Debug().Model(&Class{}).Preload("Teachers").Where("teacher_id = ?", tid).Take(&Class{}).Limit(100).Find(&classes).Error
	if db.Error != nil {
		return &[]Class{}, db.Error
	}
	return &classes, nil
}

func (c *Class) GetClassBySubjectID(db *gorm.DB, subid uint) (*[]Class, error) {
	classes := []Class{}
	sub := Subject{}
	db.First(&sub, "id = ?", subid)
	db.Model(&sub).Association("Classes").Find(&classes)
	//err := db.Debug().Model(&Class{}).Preload("Teachers").Where("teacher_id = ?", tid).Take(&Class{}).Limit(100).Find(&classes).Error
	if db.Error != nil {
		return &[]Class{}, db.Error
	}
	return &classes, nil
}
func (c *Class) ToString() string {

	return fmt.Sprintf("ClassId: %d  ClassName: %s", c.Id, c.Name)
}
