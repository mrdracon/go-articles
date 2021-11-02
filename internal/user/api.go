package user

import (
	"github.com/gin-gonic/gin"
	"go-articles/model"
	"go-articles/pkg/log"
	"go-articles/store"
	"net/http"
)

type UserCreateRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password"`
}

func HandleCreateUser(c *gin.Context, st *store.Store, logger log.Logger) {
	var req UserCreateRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{
		Email: req.Email,
		EncryptedPassword:  req.Password,
	}
	logger.Infof("Created user: %s", req.Email)

	uid, err := st.User().Create(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error creating User": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"userId": uid,
	})
}

func HandleGetUser(c *gin.Context, st *store.Store, logger log.Logger) {
	uid := c.Param("userId")

	u, err := st.User().FindById(uid)
	if err != nil {
		if err == store.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"User does not exist": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{
		"User": u,
	})
	//logger.Infof("initial url is: %s", initialUrl)
	//
	//c.Redirect(302, initialUrl)

}