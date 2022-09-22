package mq

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

// 创建 topic
func TestCreateTopic(t *testing.T) {
	// to create topics when auto.create.topics.enable='false'
	//topic := "article"
	//
	//conn, err := kafka.Dial("tcp", "localhost:9092")
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer conn.Close()
	//
	//controller, err := conn.Controller()
	//if err != nil {
	//	panic(err.Error())
	//}
	//var controllerConn *kafka.Conn
	//controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer controllerConn.Close()
	//
	//
	//topicConfigs := []kafka.TopicConfig{
	//	{
	//		Topic:             topic,
	//		NumPartitions:     1,
	//		ReplicationFactor: 1,
	//	},
	//}
	//
	//err = controllerConn.CreateTopics(topicConfigs...)
	//if err != nil {
	//	panic(err.Error())
	//}
}


// 列出 Topic
func TestListTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

