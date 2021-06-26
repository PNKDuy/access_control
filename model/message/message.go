package message

import (
	"access_control/model/role"
	"bytes"
	"context"
	"encoding/json"
	fmt "fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
	"time"
)

type Message struct {
	Action	string	`json:"action"`
	Type	string	`json:"type"`
	Value 	string	`json:"value"`
}

func ConvertFromByteArrayToMessageModel(value []byte) (msg Message, err error) {
	err = json.Unmarshal(value, &msg); if err != nil {
		return msg, err
	}
	return msg, nil
}

func ProduceMassge(msg Message) error {
	topic := "access-control-topic"
	partition := 0
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(msg)
	if err != nil {
		return err
	}

	conn , err := kafka.DialLeader(context.Background(), "tcp", "45.32.117.131:9092", topic, partition)
	if err != nil {
		return err
	}
	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}
	_, err = conn.WriteMessages(
		kafka.Message{Value: reqBodyBytes.Bytes()},
	)
	if err != nil {
		return err
	}
	if err = conn.Close(); err != nil {
		return err
	}
	return nil
}

func ConsumeMessageAndSyncDatabase(ctx context.Context) {
	conn := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"45.32.117.131:9092"},
		Topic:       "account-service-topic",
		StartOffset: kafka.LastOffset,
		GroupID:     "account-service-topic",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := conn.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		msgModel, err := ConvertFromByteArrayToMessageModel(msg.Value)
		if err != nil {
			log.Fatal(err.Error())
		}
		switch msgModel.Type {

		case "role":
			{
				var roleModel role.Role
				if strings.Contains(msgModel.Action, "create") {
					if err = json.Unmarshal([]byte(msgModel.Value), &roleModel); err != nil {
						log.Fatal(err)
					}
					if _, err = roleModel.Create(); err != nil {
						log.Fatal(err)
					}
				}
				if strings.Contains(msgModel.Action, "delete") {
					if err = role.Delete(msgModel.Value); err != nil {
						log.Fatal(err)
					}
				}
				if strings.Contains(msgModel.Action, "update") {
					if err = json.Unmarshal([]byte(msgModel.Value), &roleModel); err != nil {
						log.Fatal(err)
					}
					if _, err = roleModel.Update(); err != nil {
						log.Fatal(err)
					}
				}
			}
		default:
			log.Fatal("type is not supported")
		}
	}
}