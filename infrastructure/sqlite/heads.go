package sqlite

import (
	"log"
	"tracker/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

const KEY_PROJECT = "CURRENT_PROJECT"
const KEY_TASK = "CURRENT_TASK"

type cuImpl struct{}

//NewHeadsRepo creates a new currents repo
func NewHeadsRepo() domain.HeadsRepo {
	return cuImpl{}
}

func (c cuImpl) SaveCurrentProject(project domain.Project) error {
	return saveCurrent(KEY_PROJECT, project.ID)
}

func (c cuImpl) SaveCurrentTask(task domain.Task) error {
	return saveCurrent(KEY_TASK, task.ID)
}

func saveCurrent(key string, val uint) error {

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	cc := domain.Heads{
		Key:   key,
		Value: val,
	}

	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}}, // key columne
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(&cc).Error

	return err
}

/**
func (c cuImpl) GetCurrentProject() (*domain.Project, error) {

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var h domain.Heads

	if err := db.Where("key = ?", KEY_PROJECT).Find(&h).Error; err != nil {
		return nil, err
	}

	prepo := NewProjectRepo()

	var pr domain.Project

	if pr, err = prepo.GetProject(int(h.Value)); err != nil {
		return nil, err
	}

	return &pr, nil
}**/

func (c cuImpl) GetCurrentTask() (*domain.Task, *domain.Project, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	var ht domain.Heads

	if err := db.Where("key = ?", KEY_TASK).Find(&ht).Error; err != nil {
		return nil, nil, err
	}

	trepo := NewTaskRepo()

	t, err := trepo.GetTaskByID(ht.Value)

	if err != nil {
		return nil, nil, err
	}

	prepo := NewProjectRepo()

	var pr domain.Project

	if pr, err = prepo.GetProject(int(t.ProjectID)); err != nil {
		return nil, nil, err
	}

	return t, &pr, nil
}
