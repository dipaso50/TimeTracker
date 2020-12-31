package domain

import (
	"github.com/jinzhu/gorm"
)

//Project table
type Project struct {
	gorm.Model
	Name        string `gorm:"primaryKey;column:name;unique_index"`
	Description string
	Tasks       []Task `gorm:"foreignkey:ProjectID;"`
}

//ProjectRepo all repo operations
type ProjectRepo interface {
	AddProject(project Project) (Project, error)
	GetProject(id int) (Project, error)
	GetProjectByName(name string) (*Project, error)
	UpdateProject(project Project) (Project, error)
	Delete(id int) error
	GetAll() ([]Project, error)
}
