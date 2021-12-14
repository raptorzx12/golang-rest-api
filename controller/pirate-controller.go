package controller

import (
	"fmt"
	"net/http"
	"programming/golang/rest-api/dto"
	"programming/golang/rest-api/entity"
	"programming/golang/rest-api/helper"
	"programming/golang/rest-api/service"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PirateController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type pirateController struct {
	pirateService service.PirateService
	jwtService    service.JWTService
}

//NewPirateController create a new instances of PirateController
func NewPirateController(pirateServ service.PirateService, jwtServ service.JWTService) PirateController {
	return &pirateController{
		pirateService: pirateServ,
		jwtService:    jwtServ,
	}
}

func (c *pirateController) All(context *gin.Context) {
	var pirates []entity.Pirate = c.pirateService.All()
	res := helper.BuildResponse(true, "OK!", pirates)
	context.JSON(http.StatusOK, res)
}

func (c *pirateController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var pirate entity.Pirate = c.pirateService.FindByID(id)
	if (pirate == entity.Pirate{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK!", pirate)
		context.JSON(http.StatusOK, res)
	}
}

func (c *pirateController) Insert(context *gin.Context) {
	var pirateCreateDTO dto.PirateCreateDTO
	errDTO := context.ShouldBind(&pirateCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		// userID := c.getUserIDByToken(authHeader)
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			pirateCreateDTO.UserID = convertedUserID
		}
		result := c.pirateService.Insert(pirateCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *pirateController) Update(context *gin.Context) {
	var pirateUpdateDTO dto.PirateUpdateDTO
	errDTO := context.ShouldBind(&pirateUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.pirateService.IsAllowedToEdit(userID, pirateUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			pirateUpdateDTO.UserID = id
		}
		result := c.pirateService.Update(pirateUpdateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *pirateController) Delete(context *gin.Context) {
	var pirate entity.Pirate
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id was found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	pirate.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.pirateService.IsAllowedToEdit(userID, pirate.ID) {
		c.pirateService.Delete(pirate)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *pirateController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
