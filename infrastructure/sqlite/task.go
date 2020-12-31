package sqlite

import (
	"log"
	"tracker/domain"

	"github.com/jinzhu/now"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type taskImpl struct{}

var formatFF = "2006-01-02 15:04"

//NewTaskRepo new taskTrack repo
func NewTaskRepo() domain.TaskTimeRepo {
	return taskImpl{}
}

func (t taskImpl) RegisterTrack(task domain.Task) error {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&task).Error
}

func (t taskImpl) GetTaskByID(idTask uint) (*domain.Task, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var b domain.Task

	db.Preload("Times").First(&b, idTask)

	return &b, nil
}

func (t taskImpl) GetTaskByName(taskName string) (*domain.Task, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var b domain.Task

	err = db.Where("name = ?", taskName).Preload("Times").First(&b).Error

	return &b, err
}

func (t taskImpl) GetAggregates(taskID int) (today int, week int, month int, total int, err error) {
	return aggregates(taskID)
}

func aggregates(taskID int) (today int, week int, month int, total int, err error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return 0, 0, 0, 0, err
	}

	now.BeginningOfMonth()
	now.EndOfMonth()

	mondayt := now.BeginningOfWeek().Format(formatFF)
	sundayt := now.EndOfWeek().Format(formatFF)

	todayinit := now.BeginningOfDay().Format(formatFF)
	todayend := now.EndOfDay().Format(formatFF)

	monthInit := now.BeginningOfMonth().Format(formatFF)
	monthEnd := now.EndOfMonth().Format(formatFF)

	db.Table("taskTime").Select("sum(duration_minutes)").Where("task_id = ? AND created_at between ? and ? ", taskID, todayinit, todayend).Row().Scan(&today)

	db.Table("taskTime").Select("sum(duration_minutes)").Where("task_id = ? ", taskID).Row().Scan(&total)

	db.Table("taskTime").Select("sum(duration_minutes)").Where("task_id = ? AND created_at between ? and ? ", taskID, mondayt, sundayt).Row().Scan(&week)

	db.Table("taskTime").Select("sum(duration_minutes)").Where("task_id = ? AND created_at between ? and ? ", taskID, monthInit, monthEnd).Row().Scan(&month)

	return today, week, month, total, nil
}

func (t taskImpl) DeleteTask(taskID int) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err.Error())
		return
	}

	db.Delete(&domain.Task{}, taskID)

	//gorm or sqlite does not support ondelete cascade constraint, i have to do this manually
	db.Where("task_id = ?", taskID).Delete(&domain.TaskTime{})
}
