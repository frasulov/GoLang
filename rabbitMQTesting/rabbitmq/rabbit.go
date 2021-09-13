package rabbitmq

import (
	"github.com/streadway/amqp"
)



func ConfigureRabbit(url string) (*amqp.Connection, *amqp.Channel){
	conn , err := amqp.Dial(url)

	if err != nil{
		panic(err.Error())
		return nil, nil
	}

	ch, err := conn.Channel()

	if err != nil{
		panic(err.Error())
		return nil,nil
	}
	return conn, ch
}
