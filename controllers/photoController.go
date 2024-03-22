package controller

import (
	"finalassignment/controllers/response"
	"finalassignment/database"
	"finalassignment/helpers"
	"finalassignment/models"
	"finalassignment/variable"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
// @Summary      Create Photo
// @Description  Create Photo
// @Tags         photos
// @Accept       json
// @Produce      json
// @Param        Request   body request.PostPhotoReq  true  "Create Photo Request"
// @Success      201  {object}  response.PostCommentResp
// @Failure      500  {object}  response.ErrorMessage
// @Failure      400  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /photos [post]
func CreatePhoto(ctx *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photo := models.Photos{}

	if contentType == APPJSON {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}
	photo.UserID = userID

	err := db.Debug().Create(&photo).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Create Photo",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.PostPhotoResp{
		Id:       variable.Id{Id: photo.ID},
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId:   variable.UserId{UserId: photo.UserID},
		Created:  variable.Created{CreatedAt: photo.CreatedAt},
	})
}

// Show Photo godoc
// @Summary      Show Photo
// @Description  Show Photo
// @Tags         photos
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.GetPhotoResp
// @Failure      404  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /photos [get]
func ShowPhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photos := []models.Photos{}

	err := db.Model(&photos).Where("user_id=?", userID).Preload("User").Find(&photos).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Get Photo",
			Message: err.Error(),
		})
		return
	}
	photoResponse := []response.GetPhotoResp{}

	for _, p := range photos {
		photoResponse = append(photoResponse, response.GetPhotoResp{
			Caption:  p.Caption,
			Created:  variable.Created{CreatedAt: p.CreatedAt}, // Convert time.Time to string
			Id:       variable.Id{Id: p.ID},
			PhotoUrl: p.PhotoUrl,
			Title:    p.Title,
			Updated:  variable.Updated{UpdatedAt: p.UpdatedAt}, // Convert time.Time to string
			UserId:   variable.UserId{UserId: p.UserID},
			User: response.GetUserItemPhoto{
				Email:    p.User.Email,
				Username: p.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, photoResponse)
}

// UpdatePhoto godoc
// @Summary      Update photo
// @Description  Update photo data
// @Tags         photos
// @Accept       json
// @Produce      json
// @Param        photoId   path int  true  "photoId"
// @Param        Request   body request.UpdatePhotoReq  true  "Update Request"
// @Success      200  {object}  response.UpdatePhotoResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /photos/{photoId} [put]
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	var ID = ctx.Param("photoId")
	photoId, err := strconv.Atoi(ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request",
			Message: err.Error(),
		})
		return
	}

	photo := models.Photos{}

	err = db.Debug().Where("id=?", photoId).First(&photo).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check Photo Exist",
			Message: err.Error(),
		})
		return
	}

	contentType := helpers.GetContentType(ctx)
	if contentType == APPJSON {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	err = db.Model(&photo).Where("id=?", photoId).Updates(models.Photos{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Update Photo",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdatePhotoResp{
		Id:       variable.Id{Id: uint(photoId)},
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId:   variable.UserId{UserId: photo.UserID},
		Updated:  variable.Updated{UpdatedAt: photo.UpdatedAt},
	})
}

// DeletePhoto godoc
// @Summary      Delete Photo
// @Description  Delete Photo data
// @Tags         photos
// @Accept       json
// @Produce      json
// @Param        photoId   path int  true  "photoId"
// @Success      200  {object}  response.DelPhotoResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /photos/{photoId} [delete]
func DeletePhoto(ctx *gin.Context) {
	var ID = ctx.Param("photoId")
	db := database.GetDB()
	photoId, err := strconv.Atoi(ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request",
			Message: err.Error(),
		})
		return
	}
	photo := models.Photos{}

	err = db.Where("id=?", photoId).Delete(&photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Delete Photo",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.DelPhotoResp{
		Message: "Your photo has been succesfully deleted",
	})

}
