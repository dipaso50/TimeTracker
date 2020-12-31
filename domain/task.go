package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Task table
type Task struct {
	gorm.Model
	Name        string `gorm:"column:name;unique_index"`
	Description string
	ProjectID   uint
	Times       []TaskTime `gorm:"foreignkey:TaskID;"`
}

//TaskTime table
type TaskTime struct {
	gorm.Model
	TaskID          uint
	Start           time.Time
	End             time.Time
	DurationMinutes int
}

//TableName a name for this table
func (ttime *TaskTime) TableName() string {
	return "taskTime"
}

//TaskTimeRepo all repo operations
type TaskTimeRepo interface {
	RegisterTrack(task Task) error
	GetTaskByID(idTask uint) (*Task, error)
	GetTaskByName(taskName string) (*Task, error)
	GetAggregates(taskID int) (today int, week int, month int, total int, err error)
	DeleteTask(taskID int)
}
