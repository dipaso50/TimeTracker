package domain

import "github.com/jinzhu/gorm"

//Heads table
type Heads struct {
	gorm.Model
	Key   string `gorm:"primaryKey;column:Key;unique_index"`
	Value uint
}

//HeadsRepo operations
type HeadsRepo interface {
	SaveCurrentProject(project Project) error
	SaveCurrentTask(task Task) error
	//GetCurrentProject() (*Project, error)
	GetCurrentTask() (*Task, *Project, error)
}
