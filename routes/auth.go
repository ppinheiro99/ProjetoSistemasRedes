package routes

import (
	"github.com/gestaoFrota/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Realizar registro
// @Description Regista um utilizador
// @Accept  json
// @Produce  json
// @Router /auth/register [post]
// @Param evaluation body model.Users true "Do register"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func RegisterUser(c *gin.Context) {
	controllers.RegisterHandler(c)
}

// @Summary Realizar autenticação
// @Description Autentica o utilizador e gera o token para os próximos acessos
// @Accept  json
// @Produce  json
// @Router /auth/login [post]
// @Param evaluation body model.Users true "Do login"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func GenerateToken(c *gin.Context) {
	controllers.LoginHandler(c)
}

// @Summary Realizar autenticação
// @Description Autentica o utilizador e gera o token para os próximos acessos
// @Accept  json
// @Produce  json
// @Router /auth/login [post]
// @Param evaluation body model.Users true "Do login"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func MobileGenerateToken(c *gin.Context) {
	controllers.MobileLoginHandler(c)
}

// @Summary Atualiza token de autenticação
// @Description Atualiza o token de autenticação do usuário
// @Accept  json
// @Produce  json
// @Router /auth/refresh_token [put]
// @param Authorization header string true "Token"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func RefreshToken(c *gin.Context) {
	controllers.RefreshHandler(c)
}

// @Summary Realizar a verificaçao de um email
// @Description Autentica o utilizador e gera o token para os próximos acessos
// @Accept  json
// @Produce  json
// @Router /auth/login [post]
// @Param evaluation body model.Users true "Do login"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func VerififyEmail(c *gin.Context) {
	controllers.VerifyEmailHandler(c)
}

func CheckToken(c *gin.Context) {
	controllers.CheckTokenHandler(c)
}

func SetPassword(c *gin.Context) {
	controllers.SetPasswordHandler(c)
}

func ChangePassword(c *gin.Context) {
	controllers.ChangePasswordHandler(c)
}

func GetTravelMap(c *gin.Context){
	controllers.GetTravelMap(c)
}