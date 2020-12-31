package sqlite

import (
	"log"
	"tracker/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type proImpl struct {
}

//NewProjectRepo creates a new Project repo
func NewProjectRepo() domain.ProjectRepo {
	return proImpl{}
}

func (p proImpl) AddProject(project domain.Project) (domain.Project, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return project, err
	}

	err = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&project).Error

	return project, err
}

func (p proImpl) UpdateProject(project domain.Project) (domain.Project, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return project, err
	}

	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&project)

	return project, err
}

func (p proImpl) GetProjectByName(name string) (*domain.Project, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var pr domain.Project
	var count int64
	db.Where("name = ?", name).First(&pr).Count(&count)

	if count > 0 {
		return &pr, nil
	}

	return nil, nil
}

func (p proImpl) GetProject(id int) (domain.Project, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		return domain.Project{}, err
	}

	var b domain.Project

	err = db.Preload("Tasks").First(&b, id).Error

	return b, err
}

func (p proImpl) Delete(id int) error {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	pr, _ := p.GetProject(id)
	trepo := NewTaskRepo()

	for _, t := range pr.Tasks {
		trepo.DeleteTask(int(t.ID))
	}

	return db.Where("id = ?", id).Delete(&domain.Project{}).Error
}

func (p proImpl) GetAll() ([]domain.Project, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	var all []domain.Project

	if err != nil {
		log.Println(err.Error())
		return all, err
	}

	err = db.Order("created_at desc").Preload("Tasks").Find(&all).Error

	return all, err
}
