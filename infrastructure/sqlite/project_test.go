package sqlite

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"tracker/domain"
)

func TestAddProject(t *testing.T) {

	repo := NewProjectRepo()

	pr := domain.Project{
		Name:        "tests2",
		Description: "desdss",
		Tasks:       []domain.Task{{Name: "one", Description: "dd"}},
	}

	var pp, pp2 domain.Project
	var err error

	defer func() {
		repo.Delete(int(pp.ID))
		repo.Delete(int(pp2.ID))
	}()

	pp, err = repo.AddProject(pr)

	if err != nil {
		t.Fatalf("Error %v", err)
	}

	pp2, err = repo.AddProject(pr)

	if err != nil {
		t.Fatalf("Error %v", err)
	}
}

func TestMain(m *testing.M) {

	log.SetOutput(ioutil.Discard)
	databaseForTest := "/home/diego/tracker.db"
	InitialMigration(databaseForTest)

	exitVal := m.Run()

	os.Exit(exitVal)
}
