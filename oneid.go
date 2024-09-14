package oneid_go_sdk

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"os"
	"time"
)

// bind 用于用户绑定两个账号
func bind(ctx context.Context) OneIdResponse {
	//1.按照规则进行信息生成
	var bindMessage string
	//2.发送信息给pulsar
	client, clientErr := pulsar.NewClient(pulsar.ClientOptions{
		URL:               os.Getenv("oneId_pulsar"), // Pulsar服务地址
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if clientErr != nil {
		return OneIdResponse{Code: 500, Message: clientErr.Error()}
	}
	producer, producerErr := client.CreateProducer(pulsar.ProducerOptions{
		Topic: os.Getenv("oneId_pulsar_topic"),
	})
	if producerErr != nil {
		return OneIdResponse{Code: 500, Message: producerErr.Error()}
	}
	defer producer.Close()

	msg := pulsar.ProducerMessage{
		Payload: []byte(bindMessage),
	}
	if _, sendErr := producer.Send(ctx, &msg); sendErr != nil {
		return OneIdResponse{Code: 500, Message: sendErr.Error()}
	}
	return OneIdResponse{Code: 200, Message: "success"}
}

// unbind 用于用户解绑两个账号
func unbind(ctx context.Context) OneIdResponse {
	//1.判断账号是否合理
	//2.按照规则进行信息生成
	//3.发送信息给pulsar
	client, clientErr := pulsar.NewClient(pulsar.ClientOptions{
		URL:               os.Getenv("oneId_pulsar"), // Pulsar服务地址
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if clientErr != nil {
		return OneIdResponse{Code: 500, Message: clientErr.Error()}
	}
	producer, producerErr := client.CreateProducer(pulsar.ProducerOptions{
		Topic: os.Getenv("oneId_pulsar_topic"),
	})
	if producerErr != nil {
		return OneIdResponse{Code: 500, Message: producerErr.Error()}
	}
	defer producer.Close()

	msg := pulsar.ProducerMessage{
		Payload: []byte("Hello, Pulsar!"),
	}
	if _, sendErr := producer.Send(ctx, &msg); sendErr != nil {
		return OneIdResponse{Code: 500, Message: sendErr.Error()}
	}
	return OneIdResponse{Code: 200, Message: "success"}
}
