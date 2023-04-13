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

// CreateSocialMedia godoc
// @Summary Post Social media
// @Description Post create a new social media, NOTE : id auto increment, so in body id is deleted
// @Tags Social Media
// @Accept json
// @Produce json
// @Param models.SocialMedia body models.SocialMedia true "create social media"
// @Success 201 {object} models.SocialMedia
// @Router /social-media/create [post]
func CreateSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := repo.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMedia); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	socialMedia.UserID = userID

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, socialMedia)

}

// GetAllSocialMedia godoc
// @Summary Get details of All social media
// @Description Get details of all social media or add query parameter user_id for all social media from it
// @Tags Social Media
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /social-media/getAll [get]
func GetAllSocialMedia(ctx *gin.Context) {
	var socialMedia []models.SocialMedia

	db := repo.GetDB()

	if _, ok := ctx.GetQuery("user_id"); ok {
		user_id, err := strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Input user_id with number",
			})
			return
		}

		err = db.Debug().Order("id").Where("user_id = ?", user_id).Find(&socialMedia).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if len(socialMedia) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("user_id %d doesn't have social media", user_id),
			})
			return
		}
	} else {
		err := db.Debug().Order("id").Find(&socialMedia).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": socialMedia,
	})
}

// GetOneSocialMedia godoc
// @Summary Get details for a given socialMediaID
// @Description Get details of social media corresponding to the input socialMediaID
// @Tags Social Media
// @Accept json
// @Produce json
// @Param socialMediaID path integer true "ID of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /social-media/getOne/{socialMediaID} [get]
func GetOneSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := repo.GetDB()

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter",
		})
		return
	}

	err = db.Debug().Where("id = ?", socialMediaID).First(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Social media not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)

}

// UpdateSocialMedia godoc
// @Summary Updated data social media with socialMediaID
// @Description Update data social media by id
// @Tags Social Media
// @Accept json
// @Produce json
// @Param socialMediaID path integer true "socialMediaID of the data social media to be updated"
// @Param models.SocialMedia body models.SocialMedia true "updated social media"
// @Success 200 {object} models.SocialMedia
// @Failed 400 {object} if bad request
// @Failed 404 if id social media not found
// @Router /social-media/update/{socialMediaID} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	var socialMedia, findSocialMedia models.SocialMedia

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMedia); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	err = db.Debug().Where("id = ?", socialMediaID).First(&findSocialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Social media not found",
		})
		return
	}

	socialMedia = models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	socialMedia.ID = uint(socialMediaID)
	socialMedia.CreatedAt = findSocialMedia.CreatedAt
	socialMedia.UserID = findSocialMedia.UserID

	err = db.Debug().Model(&socialMedia).Where("id = ?", socialMediaID).Updates(socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete data social media with socialMediaID
// @Description Delete data social media by id
// @Tags Social Media
// @Accept json
// @Produce json
// @Param socialMediaID path integer true "socialMediaID of the data social media to be deleted"
// @Success 200 {object} models.SocialMedia
// @Router /social-media/delete/{socialMediaID} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	db := repo.GetDB()

	err = db.Debug().Where("id = ?", socialMediaID).First(&socialMedia).Delete(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Social Media %s successfully deleted", socialMedia.Name),
	})
}
