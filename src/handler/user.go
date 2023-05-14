package handler

import (
	"gin/sdk/message"
	"gin/src/model"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserRegister(ctx *gin.Context) {
	userInput := model.UserRegister{}

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		message.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Failed to bind JSON", err.Error())
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
		message.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Failed to bind JSON", err.Error())
		return
	}

	userResponse, err := h.uc.User.Login(ctx.Request.Context(), userInput)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to login!", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "User success to login!", &userResponse)
}

func (h *Handler) UserUpdate(ctx *gin.Context) {
	userInput := model.UserUpdate{}

	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		message.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Failed to bind JSON", err.Error())
		return
	}

	userUpdate, err := h.uc.User.Update(ctx.Request.Context(), userInput, userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Update Successfully", &userUpdate)
}

func (h *Handler) UploadPhotoProfile(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	var photo *multipart.FileHeader

	userID := user.(uint)

	photoUser, _ := ctx.FormFile("photoProfile")

	if photoUser != nil {
		photo = photoUser
	} else {
		photo = nil
	}

	// if err != nil {
	// 	message.ErrorResponse(ctx, http.StatusUnsupportedMediaType, "Failed to get file!", err.Error())
	// 	return
	// }

	// if photoUser == nil {
	// 	message.ErrorResponse(ctx, http.StatusUnsupportedMediaType, "Please select a photo!", nil)
	// 	return
	// }

	userUpdate, err := h.uc.User.UploadPhotoProfile(ctx, userID, photo)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to upload photo profile!", err.Error())
	}

	message.SuccessResponse(ctx, http.StatusOK, "Photo profile success uploaded!", &userUpdate)
}

func (h *Handler) UploadComment(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	ecoIDStr := ctx.Param("id")

	ecoID, err := strconv.Atoi(ecoIDStr)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		message.ErrorResponse(ctx, http.StatusUnsupportedMediaType, "Failed to get file", err.Error())
		return
	}

	files := form.File["thumbnail"]
	rating := ctx.PostForm("rating")
	body := ctx.PostForm("body")

	photoComment := []*multipart.FileHeader{}
	data := []string{rating, body}

	photoComment = append(photoComment, files...)

	if len(photoComment) == 0 && len(data) < 2 {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Please fill some column", nil)
		return
	}

	ecotourism, err := h.uc.User.Comment(ctx.Request.Context(), uint(ecoID), userID, photoComment, data)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to upload comment", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Comment success uploaded", &ecotourism)
}

func (h *Handler) GetProfile(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	user, err := h.uc.User.Profile(ctx.Request.Context(), userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user profile!", nil)
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "User found!", &user)
}

func (h *Handler) DeleteAccount(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	err := h.uc.User.Delete(ctx.Request.Context(), userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete account", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Account deleted", nil)
}
