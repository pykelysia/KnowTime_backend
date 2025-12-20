package route

import (
	"knowtime/database"
	"knowtime/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func userLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type (
			UserLoginReq struct {
				Name     string `json:"name"`
				Password string `json:"password"`
			}
			LoginData struct {
				Token string `json:"token"`
				UId   uint   `json:"u_id"`
			}
		)
		var (
			req UserLoginReq
		)
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"net_message": BaseMsg{
					Code:    400,
					Message: "Invalid request body",
				},
			})
			return
		}

		// TODO: add hash and move to internal/
		userEngine := database.NewUser()
		userFromDB, err := userEngine.GetByName(req.Name)
		if err != nil || req.Password != userFromDB.Password {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"net_message": BaseMsg{
					Code:    400,
					Message: "Invalid username or password error",
				},
			})
			return
		}

		// 生成JWT token
		token, err := middleware.GenerateJWT(userFromDB.UId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"net_message": BaseMsg{
					Code:    500,
					Message: "Failed to generate token",
				},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"net_message": BaseMsg{
				Code:    200,
				Message: "Login in successful",
			},
			"data": LoginData{
				Token: token,
				UId:   userFromDB.UId,
			},
		})
	}
}

func userLogupHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type (
			UserLogupReq struct {
				Name     string `json:"name"`
				Password string `json:"password"`
			}
			LogupData struct {
				UId uint `json:"u_id"`
			}
		)
		var (
			userFromReq UserLogupReq
		)
		if err := ctx.ShouldBindBodyWithJSON(&userFromReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"net_message": BaseMsg{
					Code:    400,
					Message: "Invalid request body",
				},
			})
			return
		}

		// TODO: add hash and move to internal/
		userEngine := database.NewUser()
		userToDB := database.User{
			Name:     userFromReq.Name,
			Password: userFromReq.Password,
		}
		uid, err := userEngine.Create(&userToDB)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"net_message": BaseMsg{
					Code:    500,
					Message: "Could not log up user",
				},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"net_message": BaseMsg{
				Code:    200,
				Message: "Log up successful",
			},
			"data": LogupData{
				UId: uid,
			},
		})
	}
}

func userInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type (
			UserInfoResp struct {
				Name string `json:"name"`
				UId  uint   `json:"u_id"`
			}
		)
		userIDFromJWT, exists := ctx.Get("user_id")
		userIDFromParma, err := strconv.Atoi(ctx.Param("u_id"))
		if !exists || userIDFromJWT.(uint) != uint(userIDFromParma) || err != nil {
			ctx.JSON(http.StatusUnauthorized, BaseMsg{
				Code:    401,
				Message: "User not authenticated",
			})
			return
		}

		userEngine := database.NewUser()
		user, err := userEngine.Get(userIDFromJWT.(uint))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"net_message": BaseMsg{
					Code:    404,
					Message: "User not found",
				},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"net_message": BaseMsg{
				Code:    200,
				Message: "User founded",
			},
			"data": UserInfoResp{
				Name: user.Name,
				UId:  user.UId,
			},
		})
	}
}
