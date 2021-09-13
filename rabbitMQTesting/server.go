package main

import (
	"github.com/gin-gonic/gin"
	"rabbitMQTesting/CONSTANTS"
	"rabbitMQTesting/rabbitmq"
)


func main() {

	server := gin.Default()

	// sending message
	message := rabbitmq.RabbitMessage{
		Queue: CONSTANTS.QUEUE_TEST,
		ExchangeKey: CONSTANTS.EXCHANGE_KEY_TEST,
		ContentType: "text/plain",
		Message: []byte("It is for tessaddsting"),
	}
	message.Send()

	// listen message
	// Go Listen function to add you logic
	rabbitmq.Listen(CONSTANTS.QUEUE_TEST,  CONSTANTS.EXCHANGE_KEY_TEST)

	server.Run(":8001")

}

