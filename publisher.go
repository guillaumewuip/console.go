package console

//import (
//    "encoding/json"
//    "github.com/streadway/amqp"
//    "log"
//)

////Stores filename and file of a log
//type Location struct {
//    Filename string `json:"filename"`
//    Line     int    `json:"line"`
//}

////A Log
////will be converted to json
//type Log struct {
//    Type     string   `json:"type"`
//    Tags     []string `json:"tags"`
//    Location `json:"location"`
//    Time     int    `json:"time"`
//    Message  string `json:"message"`
//}

//func NewPublisher(amqpURI, queue, corrId string, message *Message) error {

//    var err error

//    connection, err := amqp.Dial("amqp://localhost")
//    if err != nil {
//        log.Printf("con err %s", err)
//    }

//    channel, err := connection.Channel()
//    if err != nil {
//        log.Printf("channel err %s", err)
//    }

//    //convert json
//    m, _ := json.Marshal(message)

//    log.Printf("publishing to %s", queue)
//    channel.Publish(
//        "",
//        queue,
//        false,
//        false,
//        amqp.Publishing{
//            Headers:         amqp.Table{},
//            ContentType:     "application/json",
//            ContentEncoding: "utf8",
//            Body:            []byte(string(m)),
//            DeliveryMode:    amqp.Transient,
//            Priority:        0,
//            CorrelationId:   corrId,
//        },
//    )

//    return connection.Close()
//}
