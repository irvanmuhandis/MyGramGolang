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

// CreateComment godoc
// @Summary      Create Comment
// @Description  Create Comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        Request   body request.PostCommentReq  true  "Create comment Request"
// @Success      201  {object}  response.PostCommentResp
// @Failure      500  {object}  response.ErrorMessage
// @Failure      400  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /comments [post]
func CreateComment(ctx *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	comment := models.Comments{}

	if contentType == APPJSON {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}
	comment.UserID = userID
	photo := models.Photos{}
	photo.ID = comment.PhotoID
	err := db.Debug().Where("id=?", comment.PhotoID).First(&photo).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check Photo Exist",
			Message: err.Error(),
		})
		return
	}

	err = db.Debug().Create(&comment).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Error:   "Internal Server Error | Create Comment",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.PostCommentResp{
		Id:      variable.Id{Id: comment.ID},
		Message: comment.Message,
		PhotoId: comment.PhotoID,
		UserId:  variable.UserId{UserId: comment.UserID},
		Created: variable.Created{CreatedAt: comment.CreatedAt},
	})
}

// Show Comment godoc
// @Summary      Show Comment
// @Description  Show Comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.GetCommentResp
// @Failure      404  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /comments [get]
func ShowComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	comments := []models.Comments{}

	err := db.Model(&comments).Where("user_id=?", userID).Preload("User").Preload("Photo").Find(&comments).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Get Comment",
			Message: err.Error(),
		})
		return
	}
	commentResponse := []response.GetCommentResp{}

	for _, p := range comments {
		commentResponse = append(commentResponse, response.GetCommentResp{

			Id:      variable.Id{Id: p.ID},
			Message: p.Message,
			PhotoId: p.PhotoID,
			UserId:  variable.UserId{UserId: p.UserID},
			Updated: variable.Updated{UpdatedAt: p.UpdatedAt},
			Created: variable.Created{CreatedAt: p.CreatedAt},
			Photo: response.GetPhotoItemComment{
				Id:       variable.Id{Id: uint(p.PhotoID)},
				Title:    p.Photo.Title,
				Caption:  p.Photo.Caption,
				PhotoUrl: p.Photo.PhotoUrl,
				UserId:   variable.UserId{UserId: p.UserID},
			},
			User: response.GetUserItemComment{
				Id:       variable.Id{Id: p.UserID},
				Email:    p.User.Email,
				Username: p.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, commentResponse)
}

// UpdateComment godoc
// @Summary      Update Comment
// @Description  Update Comment data
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        commentId   path int  true  "commentId"
// @Param        Request   body request.UpdateCommentReq  true  "Update Request"
// @Success      200  {object}  response.UpdateCommentResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Failure      500  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /comments/{commentId} [put]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	var ID = ctx.Param("commentId")
	commentId, err := strconv.Atoi(ID)

	comment := models.Comments{}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Check Comment Exist",
			Message: err.Error(),
		})
		return
	}

	//Cek data exist saja
	err = db.Debug().Where("id=?", commentId).Preload("Photo").First(&comment).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check Comment Exist",
			Message: err.Error(),
		})
		return
	}

	//Assign param ke variable
	contentType := helpers.GetContentType(ctx)
	if contentType == APPJSON {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	//Update message
	err = db.Model(&comment).Where("id=?", commentId).Updates(models.Comments{
		Message: comment.Message}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Error:   "Internal Server Error | Update Comment",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdateCommentResp{
		Id:       variable.Id{Id: uint(commentId)},
		Title:    comment.Photo.Title,
		Caption:  comment.Photo.Caption,
		PhotoUrl: comment.Photo.PhotoUrl,
		UserId:   variable.UserId{UserId: comment.UserID},
		Updated:  variable.Updated{UpdatedAt: comment.UpdatedAt},
	})
}

// DeleteComment godoc
// @Summary      Delete Comment
// @Description  Delete Comment data
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        commentId   path int  true  "commentId"
// @Success      200  {object}  response.DelCommentResp
// @Failure      400  {object}  response.ErrorMessage
// @Failure      401  {object}  response.ErrorMessage
// @Failure      404  {object}  response.ErrorMessage
// @Failure      500  {object}  response.ErrorMessage
// @Security ApiKeyAuth
// @Router       /comments/{commentId} [delete]
func DeleteComment(ctx *gin.Context) {
	var ID = ctx.Param("commentId")
	db := database.GetDB()
	commentId, err := strconv.Atoi(ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Error:   "Bad Request | Check Comment Exist",
			Message: err.Error(),
		})
		return
	}
	comment := models.Comments{}
	err = db.Debug().Where("id=?", commentId).First(&comment).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
			Error:   "Not Found | Check Comment Exist",
			Message: err.Error(),
		})
		return
	}
	err = db.Where("id=?", commentId).Delete(&comment).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Error:   "Internal Server Error | Delete Comment",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.DelCommentResp{
		Message: "Your comment has been succesfully deleted",
	})

}
