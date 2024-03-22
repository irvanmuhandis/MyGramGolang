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

// CreateSocialMedia godoc
// @Summary      Create SocialMedia
// @Description  Create SocialMedia
// @Tags         socialmedias
// @Accept       json
// @Produce      json
// @Param        Request   body request.PostSocialReq  true  "Create socialmedia Request"
// @Success      201  {object}  response.PostSocialResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /socialmedias [post]
func CreateSocialMedia(ctx *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialmedia := models.SocialMedias{}

	if contentType == APPJSON {
		ctx.ShouldBindJSON(&socialmedia)
	} else {
		ctx.ShouldBind(&socialmedia)
	}
	socialmedia.UserId = userID

	err := db.Debug().Create(&socialmedia).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Create SocialMedia",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.PostSocialResp{
		Id:        variable.Id{Id: socialmedia.ID},
		Name:      socialmedia.Name,
		SocialUrl: socialmedia.SocialMediaUrl,
		UserId:    variable.UserId{UserId: socialmedia.UserId},
		Created:   variable.Created{CreatedAt: socialmedia.CreatedAt},
	})
}

// Show SocialMedia godoc
// @Summary      Show SocialMedia
// @Description  Show SocialMedia
// @Tags         socialmedias
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.GetSocialResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /socialmedias [get]
func ShowSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialmedias := []models.SocialMedias{}

	err := db.Model(&socialmedias).Where("user_id=?", userID).Preload("User").Find(&socialmedias).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Get SocialMedia",
			Message: err.Error(),
		})
		return
	}
	socialmediaResponse := response.GetSocialResp{}
	var socialItem = []response.SocialMediaItem{}

	for _, p := range socialmedias {
		socialItem = append(socialItem,
			response.SocialMediaItem{
				Id:        variable.Id{Id: p.ID},
				Name:      p.Name,
				SocialUrl: p.SocialMediaUrl,
				UserId:    variable.UserId{UserId: p.UserId},
				Updated:   variable.Updated{UpdatedAt: p.UpdatedAt},
				Created:   variable.Created{CreatedAt: p.CreatedAt},
				User: response.GetUserSocial{
					Id:       variable.Id{Id: uint(p.UserId)},
					Username: p.User.Username,
					Profile:  p.SocialMediaUrl,
				},
			})
	}
	socialmediaResponse.SocialMedia = socialItem
	ctx.JSON(http.StatusOK, socialmediaResponse)
}

// UpdateSocialMedia godoc
// @Summary      Update SocialMedia
// @Description  Update SocialMedia data
// @Tags         socialmedias
// @Accept       json
// @Produce      json
// @Param        socialMediaId   path int  true  "socialMediaId"
// @Param        Request   body request.UpdateSocialReq  true  "Update Request"
// @Success      200  {object}  response.UpdateSocialResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /socialmedias/{socialMediaId} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	var ID = ctx.Param("socialMediaId")
	socialMediaId, err := strconv.Atoi(ID)

	socialmedia := models.SocialMedias{}
	socialmediaExist := models.SocialMedias{}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check SocialMedia Exist",
			Message: err.Error(),
		})
		return
	}

	//Cek data exist saja
	err = db.Debug().Where("id=?", socialMediaId).First(&socialmediaExist).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check SocialMedia Exist",
			Message: err.Error(),
		})
		return
	}

	//Assign param ke variable
	contentType := helpers.GetContentType(ctx)
	if contentType == APPJSON {
		ctx.ShouldBindJSON(&socialmedia)
	} else {
		ctx.ShouldBind(&socialmedia)
	}

	//Update name dan social url
	err = db.Model(&socialmedia).Where("id=?", socialMediaId).Updates(models.SocialMedias{
		Name:           socialmedia.Name,
		SocialMediaUrl: socialmedia.SocialMediaUrl,
	}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Update SocialMedia",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdateSocialResp{
		Id:        variable.Id{Id: uint(socialMediaId)},
		Name:      socialmedia.Name,
		SocialUrl: socialmedia.SocialMediaUrl,
		UserId:    variable.UserId{UserId: socialmediaExist.UserId},
		Updated:   variable.Updated{UpdatedAt: socialmedia.UpdatedAt},
	})
}

// DeleteSocialMedia godoc
// @Summary      Delete SocialMedia
// @Description  Delete SocialMedia data
// @Tags         socialmedias
// @Accept       json
// @Produce      json
// @Param        socialMediaId   path int  true  "socialMediaId"
// @Success      200  {object}  response.DelSocialResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Failure      500  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /socialmedias/{socialMediaId} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	var ID = ctx.Param("socialMediaId")
	db := database.GetDB()
	socialMediaId, err := strconv.Atoi(ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Check SocialMedia Exist",
			Message: err.Error(),
		})
		return
	}
	socialmedia := models.SocialMedias{}
	err = db.Debug().Where("id=?", socialMediaId).First(&socialmedia).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check SocialMedia Exist",
			Message: err.Error(),
		})
		return
	}
	err = db.Where("id=?", socialMediaId).Delete(&socialmedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Error:   "Internal Server Error | Delete SocialMedia",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.DelSocialResp{
		Message: "Your socialmedia has been succesfully deleted",
	})

}
