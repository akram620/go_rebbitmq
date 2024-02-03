package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	log.Println("Publisher is running...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("successfully connected to RebbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test_queue_name", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		panic(err)
	}

	log.Println(q)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World 1"),
		})

	if err != nil {
		panic(err)
	}

	log.Println("successfully published message to queue")
}
