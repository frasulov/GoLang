package rabbitmq

import (
	"github.com/streadway/amqp"
	"rabbitMQTesting/CONSTANTS"
)

type RabbitMessage struct {
	Queue 		string
	ExchangeKey string
	Message 	[]byte
	ContentType	string
}

func (rm *RabbitMessage) generateQueue(ch * amqp.Channel) {
	_, err := ch.QueueDeclare(rm.Queue, false, false, false,false, nil)
	if err != nil {
		panic(err.Error())
		return
	}
}

func (rm* RabbitMessage) Send(){
	connection, channel := ConfigureRabbit(CONSTANTS.RABBITMQ_URL)

	rm.generateQueue(channel)

	err := channel.Publish(rm.ExchangeKey, rm.Queue, false, false, amqp.Publishing{
		ContentType: rm.ContentType,
		Body: rm.Message,
	})

	if err != nil{
		panic(err.Error())
		return
	}
	defer connection.Close()
	defer channel.Close()
}