// http://rafalgolarz.com/blog/2018/02/20/rabbitmq_essentials_with_go_examples/
// Publisher.go
package main

import (
    "log"
    "fmt"
    "github.com/streadway/amqp"
)

// ***********************************************************************
// Main
// ***********************************************************************
func main() {


      // Запись в разные каналы
      go WriteToQue("hey",1000)    
      go WriteToQue("Wrk",100)    
      WriteToQue("Integrat",1234)    
      
      // Readque("hey")
      // Readque("hey")
      // Readque("hello")
}

// ***********************************************************************
// Запись in loop
// ***********************************************************************
func WriteToQue(Q string, Qnt int){
         
     for i:=1;i<=Qnt;i++  {
           l:=`{"Id":"`+InttoStr(i)+`","Strs":"Testing"}` 
           SentQueu(Q, l)    
     }
}

// ***********************************************************************
// Sent in giues
// ***********************************************************************
func SentQueu(Chanel, Bodys string){
   // Make a connection
    conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
    defer conn.Close()

    // Ccreate a channel
    ch, _ := conn.Channel()
    defer ch.Close()

    // Declare a queue
    // Параметры - описание
    // name of the queue
    // should the message be persistent? also queue will survive if the cluster gets reset
    // autodelete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
    // exclusive means I should get an error if any other consumer subsribes to this queue
    // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
    // arguments for more advanced configuration
    q, err := ch.QueueDeclare(Chanel, false, false, false, false, nil)

    if err!=nil{
       log.Println(err)
     }

    // Publish a message
    // exchange, routing key, mandatory, immediate
    err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType:"text/plain", Body:[]byte(Bodys)})
    log.Printf("%s -> %s",Chanel, Bodys)
}


// *********************************************************************
// Чтение из канала
// *********************************************************************
func Readque(Chanel string){
      // Make a connection
    conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
    defer conn.Close()

    //Ccreate a channel
    ch, _ := conn.Channel()
    defer ch.Close()
    
    // Запрос к каналу - Hey
    q, err := ch.QueueDeclare(Chanel,false,false,false,false,nil)

    if err!=nil{
       log.Println(err)
     }

    // queue, consumer,auto-ack,exclusive, no-local,no-wait, args
    msgs, err := ch.Consume(q.Name,"", true, false, false, false, nil)

    // Чтение сообщений с канала:
    forever := make(chan bool)

    go func() {
       for d := range msgs {
           log.Printf("Получение c канала : %s %s", Chanel, d.Body)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

    <-forever
}

/******************************************************************
 * Конвертация Int to Str
 ******************************************************************/
func InttoStr(Ints int) string {
    //str := strconv.FormatInt(Intt64, 10)      // Выдает конвертацию 2000-wqut
    //str := strconv.Itoa64(Int64)              // use base 10 for sanity purpose
    str := fmt.Sprintf("%d", Ints)
    return str
}

/******************************************************************
 * Конвертация Int64 to Str
 ******************************************************************/
func Int64toStr(Int64 int64) string {
    //str := strconv.FormatInt(Intt64, 10)      // Выдает конвертацию 2000-wqut
    //str := strconv.Itoa64(Int64)              // use base 10 for sanity purpose
    str := fmt.Sprintf("%d", Int64)
    return str
}
