package project

import (
	"Mini_Project/entities"

	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *projectRepository {
	return &projectRepository{db}
}

func (br *projectRepository) GetAllProject() ([]entities.Project, error) {
	projects := []entities.Project{}

	br.db.Find(&projects)

	return projects, nil
}

func (br *projectRepository) GetProjectByID(id int) (entities.Project, error) {
	project := entities.Project{}

	br.db.Where("ID = ?", id).Find(&project)

	return project, nil
}

func (br *projectRepository) CreateProject(project entities.Project) (entities.Project, error) {
	br.db.Save(&project)

	return project, nil
}

func (br *projectRepository) FindProjectByID(id int) (entities.Project, error) {
	project := entities.Project{}

	br.db.Where("id = ?", id).Find(&project)

	return project, nil
}

func (br *projectRepository) UpdateProject(project entities.Project) (entities.Project, error) {
	br.db.Save(&project)

	return project, nil
}

func (br *projectRepository) DeleteProject(project entities.Project) (entities.Project, error) {
	br.db.Delete(&project)

	return project, nil
}
