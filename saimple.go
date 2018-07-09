package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/streadway/amqp"
)

var (
	amqpUri          string           = "amqp://guest:guest@192.168.56.101:5672/"
	rabbitCloseError chan *amqp.Error = make(chan *amqp.Error)
	started          chan bool        = make(chan bool)
	forever          chan bool        = make(chan bool)
)

func connectToRabbitMQ(uri string) *amqp.Connection {
	for {
		conn, err := amqp.Dial(uri)
		if err == nil {
			return conn
		}

		log.Println(err)
		log.Printf("Trying to reconnect to RabbitMQ at %s\n", uri)
		time.Sleep(500 * time.Millisecond)
	}
}

func rabbitConnector() {

	for {
		fmt.Println("go!")

		rabbitErr := <-rabbitCloseError

		if rabbitErr != nil {
			log.Printf("Connecting to %s\n..", amqpUri)

			r_connection := connectToRabbitMQ(amqpUri)

			r_channel, err := r_connection.Channel()
			fmt.Println(reflect.TypeOf(err))

			failOnError(err, "Failed to open a channel")

			q, err := r_channel.QueueDeclare(
				"hello", // name
				false,   // durable
				false,   // delete when usused
				false,   // exclusive
				false,   // no-wait
				nil,     // arguments
			)
			failOnError(err, "Failed to declare a queue")
			r_queue := &q

			r_messages, err := r_channel.Consume(
				r_queue.Name, // queue
				"",           // consumer
				true,         // auto-ack
				false,        // exclusive
				false,        // no-local
				false,        // no-wait
				nil,          // args
			)
			//				r_messages = r_msg
			failOnError(err, "Failed to channel consume")

			log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

			r_connection.NotifyClose(rabbitCloseError)
			fmt.Println("Connected.")
			for {
				select {
				case msg := <-r_messages:
					{
						log.Printf("Received a message: %s", msg.Body)
					}
				case rabbitCloseError <- amqp.ErrClosed:
					{
						break
					}

				}
			}
		}

	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	go rabbitConnector()
	rabbitCloseError <- amqp.ErrClosed
	<-forever

}
