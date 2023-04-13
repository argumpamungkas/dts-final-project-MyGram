package controllers

import (
	"DTS/Chapter-3/final-project-myGram/helpers"
	"DTS/Chapter-3/final-project-myGram/models"
	"DTS/Chapter-3/final-project-myGram/repo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateComment godoc
// @Summary Post Comment
// @Description Post a new Comment, NOTE : id auto increment, so in body id is deleted. and add query parameter photo_id for comment
// @Tags Create Comment
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "create comment"
// @Success 201 {object} models.Comment
// @Router /comment/post [post]
func CreateComment(ctx *gin.Context) {
	var comment models.Comment
	var photo models.Photo

	db := repo.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	photoID, err := strconv.Atoi(ctx.Query("photo_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Input photo_id with number",
		})
		return
	}

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&comment); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&comment); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	err = db.Debug().Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Photo not found",
		})
		return
	}

	comment.UserID = userID
	comment.PhotoID = uint(photoID)

	err = db.Debug().Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)

}

// GetAllComment godoc
// @Summary Get details of All comment
// @Description Get details of all comment or add query parameter photo_id for all comment from photo_id
// @Tags Get All Comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comment/getAll [get]
func GetAllComent(ctx *gin.Context) {
	var comment []models.Comment
	var photo models.Photo

	db := repo.GetDB()

	if _, ok := ctx.GetQuery("photo_id"); ok {
		photoID, err := strconv.Atoi(ctx.Query("photo_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Input photo_id with number",
			})
			return
		}

		err = db.Debug().Where("id = ?", photoID).First(&photo).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Photo not found",
			})
			return
		}

		err = db.Debug().Order("id").Where("photo_id = ?", photoID).Find(&comment).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		if len(comment) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "There are no comments for this photo",
			})
			return
		}

	} else {

		err := db.Debug().Order("id").Find(&comment).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": comment,
	})
}

// GetOneComment godoc
// @Summary Get details for a given commentID
// @Description Get details of comment corresponding to the input commentID
// @Tags Get Comment by ID
// @Accept json
// @Produce json
// @Param commentID path integer true "ID of the photo"
// @Success 200 {object} models.Comment
// @Failed 404 if id comment not found
// @Router /comment/getOne/{commentID} [get]
func GetOneComment(ctx *gin.Context) {
	var comment models.Comment

	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter",
		})
		return
	}

	db := repo.GetDB()

	err = db.Debug().Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Comment not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// UpdateComment godoc
// @Summary Updated data comment with commentID
// @Description Update data comment by id
// @Tags Update Comment
// @Accept json
// @Produce json
// @Param commentID path integer true "commentID of the data comment to be updated"
// @Param models.Comment body models.Comment true "updated comment"
// @Success 200 {object} models.Comment
// @Failed 400 {object} if bad request
// @Failed 404 if id comment not found
// @Router /comment/update/{commentID} [put]
func UpdateComment(ctx *gin.Context) {
	var comment, findComment models.Comment

	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Input parameter with id",
		})
		return
	}

	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	err = db.Debug().Where("id = ?", commentID).First(&findComment).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Comment with id %d not found", commentID),
		})
		return
	}

	comment = models.Comment{
		Message: comment.Message,
	}

	comment.ID = uint(commentID)
	comment.CreatedAt = findComment.CreatedAt
	comment.PhotoID = findComment.PhotoID

	err = db.Debug().Model(&comment).Where("id = ?", commentID).Updates(comment).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary Delete data comment with commentID
// @Description Delete data comment by id
// @Tags Delete Comment
// @Accept json
// @Produce json
// @Param commentID path integer true "commentID of the data comment to be deleted"
// @Success 200 {object} models.Comment
// @Failed 404 if id comment not found
// @Router /comment/delete/{commentID} [delete]
func DeleteComent(ctx *gin.Context) {
	var comment models.Comment

	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Input parameter with id",
		})
		return
	}

	db := repo.GetDB()

	err = db.Debug().Where("id = ?", commentID).First(&comment).Delete(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Comment with id %d not found", commentID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Comment message '%s' successfully deleted", comment.Message),
	})
}
