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

	ecoIDStr := ctx.Param("id")
	ecoID, err := strconv.Atoi(ecoIDStr)
	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["thumbnail"]
	rating := ctx.PostForm("rating")
	body := ctx.PostForm("body")

	photoComment := []*multipart.FileHeader{}
	data := make([]string, 2)
	data[0] = rating
	data[1] = body

	photoComment = append(photoComment, files...)

	ecotourism, err := h.uc.User.Comment(ctx.Request.Context(), uint(ecoID), userID, photoComment, data)
	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to upload comment", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Comment successfully uploaded", ecotourism)
}

func (h *Handler) GetProfile(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	user, err := h.uc.User.Profile(ctx.Request.Context(), userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user profile", nil)
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "User Found!", user)
}

func (h *Handler) AddCart(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	ecoIDStr := ctx.Param("id")
	ecoID, err := strconv.Atoi(ecoIDStr)
	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}

	newProductCart := model.CartProduct{}

	if err := ctx.ShouldBindJSON(&newProductCart); err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err.Error())
		return
	}

	cart, err := h.uc.User.AddCart(ctx.Request.Context(), userID, uint(ecoID), newProductCart)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to add into cart", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Added to cart", cart)
}

func (h *Handler) GetCart(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	cart, err := h.uc.User.GetCart(ctx.Request.Context(), userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Cart found", cart)
}

func (h *Handler) DeleteCartContent(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	ecoIDStr := ctx.Param("id")
	ecoID, err := strconv.Atoi(ecoIDStr)
	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}
	err = h.uc.User.DeleteCartContent(ctx.Request.Context(), userID, uint(ecoID))

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Content deleted", nil)
}

func (h *Handler) DeleteAccount(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
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
