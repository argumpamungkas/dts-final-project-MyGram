package controllers

import (
	"DTS/Chapter-3/final-project-myGram/helpers"
	"DTS/Chapter-3/final-project-myGram/models"
	"DTS/Chapter-3/final-project-myGram/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJson = "application/json"

func RegisterUser(ctx *gin.Context) {
	var user models.User

	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"age":      user.Age,
		"email":    user.Email,
		"username": user.Username,
	})
}

// LoginUser godoc
// @Summary Login User
// @Description Login user needed for crud of the photo, social media and comment because if you login you have token for that
// @Tags Login user
// @Accept json
// @Produce json
// @Param models.User body models.User true "login user"
// @Success 200 {object} models.User
// @Router /user/login [post]
func LoginUser(ctx *gin.Context) {
	var user models.User

	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	password := user.Password

	err := db.Debug().Where("username = ?", user.Username).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username",
		})
		return
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
