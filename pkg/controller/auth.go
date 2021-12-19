package controller

import (
	"github.com/amper-pw/auth/pkg/forms"
	"github.com/amper-pw/auth/pkg/responses"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Controller) testAction(ctx *gin.Context) {
	var form forms.SignIn

	if err := ctx.BindJSON(&form); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := s.services.AuthService.GenerateToken(form.Username, form.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, structs.Map(&responses.SignIn{
		Token: token,
	}))

}

func (s *Controller) signIn(ctx *gin.Context) {
	var form forms.SignIn

	if err := ctx.BindJSON(&form); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := s.services.AuthService.GenerateToken(form.Username, form.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, structs.Map(&responses.SignIn{
		Token: token,
	}))

}

func (c *Controller) signUp(ctx *gin.Context) {
	var form forms.SignUp
	translator := ru.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("ru")
	if !found {
		logrus.Fatal("translator not found")
	}

	validate := validator.New()

	if err := ru_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		logrus.Fatal("translator not found")
	}

	if err := ctx.ShouldBindJSON(&form); err != nil {
		logrus.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, structs.Map(&responses.ErrorResponse{
			Message: "error",
		}))
		return
	}

	if err := validate.Struct(&form); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, structs.Map(&responses.ErrorResponse{
				Message: e.Translate(trans),
			}))
		}
		return
	}

	user, err := c.services.AuthService.RegisterUser(form.Username, form.Password)
	if err != nil {
		logrus.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, structs.Map(&responses.ErrorResponse{
			Message: "internal server error",
		}))
		return
	}

	ctx.JSON(http.StatusOK, structs.Map(&responses.SignUp{
		Id: user.Id.String(),
	}))

}
