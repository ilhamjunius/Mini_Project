package project

import (
	"Mini_Project/delivery/formatter"
	request "Mini_Project/delivery/request"
	"Mini_Project/delivery/response"
	"Mini_Project/entities"
	"Mini_Project/repository/project"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectController struct {
	repo project.Project
}

func NewController(repo project.Project) *ProjectController {
	return &ProjectController{repo}
}

func (c *ProjectController) GetAllProject(e echo.Context) error {
	res := response.Response{}

	projects, _ := c.repo.GetAllProject()

	if len(projects) == 0 {
		res.Status = http.StatusNotFound
		res.Message = "no data found projects"
		res.Data = projects

		return e.JSON(http.StatusNotFound, res)
	}

	formatData := formatter.FormatProjects(projects)

	res.Status = http.StatusOK
	res.Message = "success get all project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) GetProjectByID(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	project_id, _ := strconv.Atoi(id)

	project, _ := c.repo.GetProjectByID(project_id)

	if project.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "project not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	formatData := formatter.FormatProject(project)

	res.Status = http.StatusOK
	res.Message = "success get project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) CreateProject(e echo.Context) error {
	res := response.Response{}

	request := request.CreateProjectInput{}

	e.Bind(&request)

	e.Validate(&request)

	project := entities.Project{}

	project.Name = request.Name

	newProject, _ := c.repo.CreateProject(project)

	formatData := formatter.FormatProject(newProject)

	res.Status = http.StatusOK
	res.Message = "success create project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) UpdateProject(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	project := entities.Project{}

	project_id, _ := strconv.Atoi(id)

	project, _ = c.repo.GetProjectByID(project_id)

	if project.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "project not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	request := request.UpdateProjectInput{}

	e.Bind(&request)

	e.Validate(&request)

	project.Name = request.Name

	updatedProject, _ := c.repo.UpdateProject(project)

	formatData := formatter.FormatProject(updatedProject)

	res.Status = http.StatusOK
	res.Message = "success update project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) DeleteProject(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	project := entities.Project{}

	project_id, _ := strconv.Atoi(id)

	project, _ = c.repo.GetProjectByID(project_id)

	if project.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "project not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	oldProject := project

	_, err := c.repo.DeleteProject(project)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatProject(oldProject)

	res.Status = http.StatusOK
	res.Message = "success delete project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}
