package controller

import (
	"finalassignment/controllers/response"
	"finalassignment/database"
	"finalassignment/helpers"
	"finalassignment/models"
	"finalassignment/variable"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var APPJSON = "application/json"

// RegisterUser godoc
// @Summary      Register User
// @Description  Register User to acces the app
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   body request.RegisReq  true  "Register Request"
// @Success      201  {object}  response.RegisResp
// @Failure      400  {object}  response.ErrorMessage
// @Router       /users/register [post]
func UserRegis(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	user := models.Users{}

	if contentType == APPJSON {
		c.ShouldBindJSON(&user)

	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Register User",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.RegisResp{
		Age:      user.Age,
		Email:    user.Email,
		Id:       variable.Id{Id: user.ID},
		Username: user.Username,
	})
}

// LoginUser godoc
// @Summary      Login User
// @Description  Login User to acces the app
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        Request   body request.LoginReq  true  "Login Request"
// @Success      200  {object}  response.LoginResp
// @Failure      400  {object}  response.ErrorMessage
// @Router       /users/login [post]
func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	user := models.Users{}
	pass := ""
	if contentType == APPJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}
	//Password nya disimpan ke variabel dulu
	pass = user.Password
	// Lalu user.password di overwrite dengan pass dari database
	err := db.Debug().Where("email=?", user.Email).First(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request",
			Message: "Invalid email/password",
		})
	}
	//Dibandingkan password yang dari inputan dan dari database
	comparePass := helpers.ComparePass([]byte(user.Password), []byte(pass))

	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request",
			Message: "Invalid email/password",
		})
		return

	}

	token := helpers.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, response.LoginResp{
		Token: token,
	})

}

// UpdateUser godoc
// @Summary      Update User
// @Description  Update User data
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userId   path int  true  "userId"
// @Param        Request   body request.UpdateUserReq  true  "Update Request"
// @Success      200  {object}  response.UpdateUserResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /users/{userId} [put]
func UpdateUser(ctx *gin.Context) {
	db := database.GetDB()
	var ID = ctx.Param("userId")
	userID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request",
			Message: err.Error(),
		})
		return
	}

	user := models.Users{}

	err = db.Debug().Where("id=?", userID).First(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check User Exist",
			Message: fmt.Sprintf("Data with id = %d not found", userID),
		})
		return
	}

	contentType := helpers.GetContentType(ctx)
	if contentType == APPJSON {
		ctx.ShouldBindJSON(&user)

	} else {
		ctx.ShouldBind(&user)
	}

	err = db.Model(&user).Where("id=?", userID).Updates(models.Users{
		Email:    user.Email,
		Password: user.Password}).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Update User",
			Message: err.Error(),
		})
		return
	}
	// tidak ada yg ke overwite
	ctx.JSON(http.StatusOK, response.UpdateUserResp{
		Id:       variable.Id{Id: uint(userID)},
		Email:    user.Email,
		Username: user.Username,
		Age:      user.Age,
		Updated:  variable.Updated{UpdatedAt: user.UpdatedAt},
	})
}

// DeleteUser godoc
// @Summary      Delete User
// @Description  Delete User data
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userId   path int  true  "userId"
// @Success      200  {object}  response.DelUserResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Failure      500  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /users/{userId} [delete]
func DeleteUser(ctx *gin.Context) {
	var ID = ctx.Param("userId")
	db := database.GetDB()
	userID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request",
			Message: err.Error(),
		})
		return
	}
	user := models.Users{}

	err = db.Debug().Where("id=?", userID).First(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check User Exist",
			Message: fmt.Sprintf("Data with id = %d not found", userID),
		})
		return
	}

	err = db.Where("id=?", userID).Delete(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Error:   "Internal Server Error | Delete User",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.DelUserResp{
		Message: "Your account has been succesfully deleted",
	})

}
