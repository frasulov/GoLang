package rabbitmq

import (
	"fmt"
	"rabbitMQTesting/CONSTANTS"
)

func Listen(queue, exchangeKey string) {
	connection, channel := ConfigureRabbit(CONSTANTS.RABBITMQ_URL)

	messages, err := channel.Consume(queue, exchangeKey, true, false, false, false, nil)

	if err != nil{
		panic(err.Error())
	}

	forever := make(chan bool)
	go func() {
		for d :=range messages{
			// ad your logic here
			fmt.Println("Message: ", d.Body)
		}
	}()

	<- forever
	defer connection.Close()
	defer channel.Close()
}