package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"net/smtp"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	var creds model.Users
	var usr model.Users

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	services.OpenDatabase()
	services.Db.Find(&usr, "email = ?", creds.Email)
	defer services.Db.Close()
	if usr.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	if !CheckPasswordHash(creds.Password, usr.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	token := services.GenerateTokenJWT(creds)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Access denied!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!","ID":usr.ID ,"Role": usr.RoleId, "Name": usr.FirstName,"LastName": usr.LastName,"Address": usr.Address ,"Email": usr.Email, "token": token})
}

func MobileLoginHandler(c *gin.Context) {
	var creds model.Users
	var usr model.Users

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	services.OpenDatabase()
	services.Db.Find(&usr, "email = ?", creds.Email)
	defer services.Db.Close()
	if usr.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	if !CheckPasswordHash(creds.Password, usr.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	if usr.RoleId != 4{
		defer services.Db.Close()
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	token := services.GenerateTokenJWT(creds)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Access denied!"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!","ID":usr.ID ,"Role": usr.RoleId, "Name": usr.FirstName,"LastName": usr.LastName,"Address": usr.Address ,"Email": usr.Email, "token": token})
}


func RegisterHandler(c *gin.Context) {
	var creds model.Users

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if len(creds.Email) < 8 || len(creds.Email) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid email!"})
		return
	}

	if len(creds.FirstName) < 3 || len(creds.FirstName) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid nome!"})
		return
	}

	if len(creds.Country) < 2 || len(creds.Country) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid pais!"})
		return
	}

	if len(creds.LastName) < 3 || len(creds.LastName) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Apelido!"})
		return
	}

	if len(creds.Password) < 8 || len(creds.Password) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid password!"})
		return
	}

	if creds.RoleId <= 1 || creds.RoleId >= 5{
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid cargo!"})
		return
	}

	if services.IsDuplicateEmail(creds.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Email already taken!"})
		return
	}

	creds.Password, _ = HashPassword(creds.Password)
	services.OpenDatabase()
	services.Db.Save(&creds)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "User ID": creds.ID, "user": creds})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RefreshHandler(c *gin.Context) {

	user := model.Users{
		Email: c.GetHeader("email"),
	}

	token := services.GenerateTokenJWT(user)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Acesso não autorizado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusNoContent, "message": "Token atualizado com sucesso!", "token": token})
}

func VerifyEmailHandler(c *gin.Context) {
	var creds model.Users
	var usr model.Users
	var passrecover model.PassRecover
	var passrecoverAux model.PassRecover

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	services.OpenDatabase()
	services.Db.Find(&usr, "email = ? ", creds.Email)
	defer services.Db.Close()
	if usr.Email == "" {
		
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Email!"})
		return
	}

	//Criamos um token para incluir no link de recuperação
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)

	// String token
	passrecover.Email = creds.Email
	passrecover.Token = hex.EncodeToString(hasher.Sum(nil))
	services.OpenDatabase()

	// Verificar se email já existe
	// Caso exista faz update
	services.Db.Find(&passrecoverAux, "email = ? ", passrecover.Email)
	if passrecoverAux.Email == "" {
		// Caso nao exista faz save
		services.Db.Save(&passrecover)
	} else { // Caso exista faz update
		services.Db.Exec("update pass_recovers set token = ? where email = ?", passrecover.Token, passrecover.Email)
	}
	defer services.Db.Close()

	//Enviamos email ao utilizador com o token para recuperar a pass
	// Choose auth method and set it up

	auth := smtp.PlainAuth("", "36fd747e0aa65a", "74d9fe1468aea1", "smtp.mailtrap.io")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"bill@gates.com"}
	msg := []byte("To: bill@gates.com\r\n" +
		"Subject: Recuperacao de password\r\n" +
		"\r\n" +
		"Aqui está o link para a recuperacao da password : https://localhost:4200/passrecover/" + passrecover.Token)
	err = smtp.SendMail("smtp.mailtrap.io:25", auth, "botaquetem@teste.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}

	
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "token": passrecover.Token})

}

func CheckTokenHandler(c *gin.Context) {
	var passrecover model.PassRecover
	var creds model.PassRecover

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	services.OpenDatabase()
	services.Db.Find(&passrecover, "token = ? ", creds.Token)
	defer services.Db.Close()
	if passrecover.Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token!"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "recover": true, "message": "Valid token!"})
	}

}

func SetPasswordHandler(c *gin.Context) {
	var passrecover model.PassRecover
	var usr model.Users
	var creds struct {
		Password        string `json:"password"`
		Token           string `json:"token"`
		ConfirmPassword string `json:"confirm_password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if len(creds.Password) < 8 || len(creds.Password) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid password!"})
		return
	}

	if creds.Password != creds.ConfirmPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Passwords sao diferentes!"})
		return
	}

	services.OpenDatabase()
	//fazemos o find pelo token
	services.Db.Find(&passrecover, "token = ? ", creds.Token)
	defer services.Db.Close()
	if passrecover.Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token!"})
		return
	}
	//procuramos o user com o email respetivo do token que recebemos
	services.OpenDatabase()
	services.Db.Find(&usr, "email = ? ", passrecover.Email)
	if usr.Email == "" {
		defer services.Db.Close()
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token!"})
		return
	}
	creds.Password, _ = HashPassword(creds.Password)
	services.Db.Exec("update users set password = ? where email = ?", creds.Password, usr.Email)
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Password Alterada com Sucesso!"})
}

func ChangePasswordHandler(c *gin.Context) {
	var usr model.Users
	var creds struct {
		Oldpassword   string `json:"oldpassword"`
		Newpassword   string `json:"newpassword"`
		NewpasswordC  string `json:"newpasswordC"`
		Email 		  string `json:"Email"`
	}


	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	
	if len(creds.Newpassword) < 8 || len(creds.Newpassword) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid password!"})
		return
	}

	if creds.Newpassword != creds.NewpasswordC {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Passwords sao diferentes!"})
		return
	}

	services.OpenDatabase()
	//fazemos o find pelo email
	services.Db.Find(&usr, "email = ? ", creds.Email)
	if usr.Email == "" {
		defer services.Db.Close()
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token!"})
		return
	}
	//Comparar a password antiga com a da base de dados
	if !CheckPasswordHash(creds.Oldpassword, usr.Password) {
		defer services.Db.Close()
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}
	//Se chegarmos aqui , significa que a pass antiga deu certo , entao podemos alterar
	
	creds.Newpassword, _ = HashPassword(creds.Newpassword)
	services.Db.Exec("update users set password = ? where email = ?", creds.Newpassword, usr.Email)
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Password Alterada com Sucesso!"})
}
