package controller

import (
	"net/http"
	"project_pertama/model"
	"project_pertama/repository"
	"project_pertama/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type personController struct {
	personRepository repository.IPersonRepository
}

func NewPersonController(personRepository repository.IPersonRepository) *personController {
	return &personController{
		personRepository: personRepository,
	}
}

func (pc *personController) Create(ctx *gin.Context) {
	var newPerson model.Person

	err := ctx.ShouldBindJSON(&newPerson)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createdPerson, err := pc.personRepository.Create(newPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdPerson, ""))
}

func (pc *personController) GetAll(ctx *gin.Context) {

	persons, err := pc.personRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, persons, ""))
}

func (pc *personController) Update(ctx *gin.Context) {
	reqId := ctx.Param("id")
	id, err := strconv.Atoi(reqId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	var newPerson model.Person
	errP := ctx.ShouldBindJSON(&newPerson)
	if errP != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, errP.Error()))
	}

	updatedPerson, err := pc.personRepository.Update(id, newPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, updatedPerson, ""))
}

func (pc *personController) Delete(ctx *gin.Context) {
	reqId := ctx.Param("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	deletedPerson, err := pc.personRepository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, deletedPerson, ""))
}
