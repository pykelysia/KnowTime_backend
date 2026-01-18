package route

import (
	"knowtime/database"
	"knowtime/internal"
	"knowtime/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserLogin godoc
//
//	@Summary		用户登录
//	@Description	用户登录获取JWT token
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		UserLoginReq	true	"用户登录信息"
//	@Success		200				{object}	map[string]interface{}
//	@Failure		400				{object}	map[string]interface{}
//	@Failure		500				{object}	map[string]interface{}
//	@Router			/user/login [post]
func userLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		uid, b, err := internal.UserLoginInternal(req.Name, req.Password)
		if err != nil {
			ctx.JSON(b.Code, gin.H{
				"net_message": b,
			})
			return
		}

		// 生成JWT token
		token, err := middleware.GenerateJWT(uid)
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
				UId:   uid,
			},
		})
	}
}

// UserRegister godoc
//
//	@Summary		用户注册
//	@Description	新用户注册
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		UserLogupReq	true	"用户注册信息"
//	@Success		200				{object}	map[string]interface{}
//	@Failure		400				{object}	map[string]interface{}
//	@Failure		500				{object}	map[string]interface{}
//	@Router			/user/logup [post]
func userLogupHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		uid, b, err := internal.UserLogupInternal(userFromReq.Name, userFromReq.Password)
		if err != nil {
			ctx.JSON(b.Code, gin.H{
				"net_message": b,
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

// UserInfo godoc
//
//	@Summary		获取用户信息
//	@Description	获取指定用户的信息
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			u_id	path		int	true	"用户ID"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		401		{object}	map[string]interface{}
//	@Failure		404		{object}	map[string]interface{}
//	@Security		BearerAuth
//	@Router			/user/info/{u_id} [get]
func userInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
