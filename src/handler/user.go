package handler

import (
	"fmt"
	"gin/sdk/message"
	"gin/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserRegister(ctx *gin.Context) {
	userInput := model.UserRegister{}

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err.Error())
		return
	}

	user, err := h.uc.User.Register(ctx.Request.Context(), userInput)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "User registered", &user)
}

func (h *Handler) UserLogin(ctx *gin.Context) {
	userInput := model.UserLogin{}

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err.Error())
		return
	}

	userResponse, err := h.uc.User.Login(ctx.Request.Context(), userInput)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to login", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "User Succesfully Login", userResponse)
}

func (h *Handler) UserUpdate(ctx *gin.Context) {
	userInput := model.UserUpdate{}
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err.Error())
		return
	}

	userUpdate, err := h.uc.User.Update(ctx.Request.Context(), userInput, userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Update Successfully", userUpdate)
}

func (h *Handler) UploadPhotoProfile(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	photoUser, err := ctx.FormFile("photoProfile")
	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to get user image", err.Error())
		return
	}

	userUpdated, err := h.uc.User.UploadPhotoProfile(ctx, userID, photoUser)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to upload photo profile, please try again!", err.Error())
	}

	message.SuccessResponse(ctx, http.StatusOK, "Photo successfully uploaded", userUpdated)
}

func (h *Handler) UploadComment(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)
	fmt.Println(userID)
	ecoIDStr := ctx.Param(":id")
	ecoID, err := strconv.Atoi(ecoIDStr)
	fmt.Println(ecoID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}

}
