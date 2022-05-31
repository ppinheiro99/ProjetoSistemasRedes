package controllers

import (
	"net/http"
	"sort"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context){
	var messages []model.Messages
	services.OpenDatabase()
	services.Db.Find(&messages)

	if len(messages) <= 0 {
		defer services.Db.Close()
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": messages})
}

func GetMessagesByUser(c *gin.Context){
	var messages_sender []model.Messages
	var messages_receiver []model.Messages
	var user_messages []model.Messages
	var user model.Users
	
	// este var id é apenas um teste
	id := c.Param("id")
	
	services.OpenDatabase()
	services.Db.Find(&user, " id = ? ", id) // vamos buscar o user à bd
	if user.Email == ""{
		defer services.Db.Close()
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}
	services.OpenDatabase()
	services.Db.Find(&messages_sender, " sender_id = ? ", user.ID)  // vamos buscar as mensagens que esse user enviou
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&messages_receiver, " receiver_id = ? ", user.ID) // vamos buscar as mensagens que esse user recebeu 
	defer services.Db.Close()

	// Como temos um array de mensagens recebidas e outro de mensagens enviadas, temos de os juntar 
	user_messages = append(messages_receiver,messages_sender...)
	// Ao juntar o array fica "baralhado", para isso precisamos de ordenar pelo ID
	sort.Slice(user_messages[:], func(i, j int) bool {
		return user_messages[i].ID < user_messages[j].ID
	  })

	
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "messagesArray": user_messages})
	
}

func Messages(c *gin.Context) {
	var messages model.Messages
	var users_sender model.Users
	var users_receiver model.Users

	if err := c.ShouldBindJSON(&messages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if messages.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	// Ver se os ids de receiver e sender estão na bd de users
	services.OpenDatabase()
	services.Db.Find(&users_sender, "id = ? ", messages.SenderID)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&users_receiver, "id = ? ", messages.ReceiverID)

	if users_sender.Email == "" || users_receiver.Email == "" {
		defer services.Db.Close()
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	// Adiciona à BD de messages
	services.Db.Save(&messages)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
