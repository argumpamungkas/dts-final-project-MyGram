package controllers

import (
	"DTS/Chapter-3/final-project-myGram/models"
	"DTS/Chapter-3/final-project-myGram/repo"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
// @Summary Post Photo
// @Description Post a new Photo, NOTE : id auto increment, so in body id is deleted
// @Tags Create Photo
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "create photo"
// @Success 201 {object} models.Photo
// @Router /photo/post [post]
func CreatePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := repo.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	username := userData["username"].(string)

	_ = ctx.ShouldBind(&photo)
	file, err := ctx.FormFile("photo_url")
	if err == nil {

		charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
		stringRandom := make([]byte, rand.Intn(100))
		for i := range stringRandom {
			stringRandom[i] = charset[rand.Intn(len(charset))]
		}

		ext := strings.Split(file.Filename, ".")[1]

		log.Println("file ext ->", ext)

		if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "webp" {
			dst := "./assets/" + username + "-" + string(stringRandom) + "." + "jpg"
			ctx.SaveUploadedFile(file, dst)

			urlPhoto := "http://" + ctx.Request.Host + "/img/" + username + "-" + string(stringRandom) + "." + "jpg"
			photo.PhotoUrl = urlPhoto
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "File is not image/photo",
			})
			return
		}

	}

	photo.UserID = userID

	err = db.Debug().Create(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// GetAllPhoto godoc
// @Summary Get details of All photo
// @Description Get details of all photo or add query parameter user_id for all photo from user_id
// @Tags Get All Photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photo/getAll [get]
func GetAllPhoto(ctx *gin.Context) {
	var photos []models.Photo

	db := repo.GetDB()

	if _, ok := ctx.GetQuery("user_id"); ok {
		user_id, err := strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Input user_id with number",
			})
			return
		}

		err = db.Debug().Preload("Comments").Order("id").Where("user_id = ?", user_id).Find(&photos).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

	} else {

		err := db.Debug().Preload("Comments").Order("id").Find(&photos).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": photos,
	})
}

// GetOnePhoto godoc
// @Summary Get details for a given photoID
// @Description Get details of photo corresponding to the input photoID
// @Tags Get Photo by ID
// @Accept json
// @Produce json
// @Param photoID path integer true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photo/getOne/{photoID} [get]
func GetOnePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := repo.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Preload("Comments").Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// UpdatePhoto godoc
// @Summary Updated data photo with socialMediaID
// @Description Update data social media by id, NOTE: photo is not updated, just title and caption can be updated, so in the body photo_url doesn't use
// @Tags Update Photo
// @Accept json
// @Produce json
// @Param photoID path integer true "photoID of the data photo to be updated"
// @Param models.Photo body models.Photo true "updated photo"
// @Success 200 {object} models.Photo
// @Failed 400 {object} if bad request
// @Failed 404 if id photo not found
// @Router /photo/update/{photoID} [put]
func UpdatePhoto(ctx *gin.Context) {
	var photo, findPhoto models.Photo

	db := repo.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Where("id = ?", photoID).First(&findPhoto).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Photo not found",
		})
		return
	}

	_ = ctx.ShouldBind(&photo)

	photo = models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: findPhoto.PhotoUrl,
	}

	photo.ID = uint(photoID)
	photo.CreatedAt = findPhoto.CreatedAt

	err = db.Debug().Model(&photo).Where("id = ?", photoID).Updates(photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Preload("Comments").Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// DeletePhoto godoc
// @Summary Delete data photo with photoID
// @Description Delete data photo by id
// @Tags Delete Photo
// @Accept json
// @Produce json
// @Param photoID path integer true "photoID of the data photo to be deleted"
// @Success 200 {object} models.Photo
// @Failed 404 if id photo not found
// @Router /photo/delete/{photoID} [delete]
func DeletePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := repo.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Where("id = ?", photoID).First(&photo).Delete(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Photo with title '%s' successfully deleted", photo.Title),
	})
}
