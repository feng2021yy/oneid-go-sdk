package oneid_go_sdk

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"os"
	"time"
)

const ()

var channelMessage map[string]string //存储变量信息
var isInit bool = false              //是否进行初始化

// channelInit 将两个账号的渠道信息初始化到变量中，避免之后重复输入
func channelInit(channelCode1, channelType1, channelCode2, channelType2 string) {
	channelMessage = make(map[string]string, 4)
	channelMessage["channelCode1"] = channelCode1
	channelMessage["channelType1"] = channelType1
	channelMessage["channelCode2"] = channelCode2
	channelMessage["channelType2"] = channelType2
	isInit = true
}

// bind 用于用户绑定两个账号，生成字符串传输给pulsar消息队列
func bind(ctx context.Context, sourceID, destID string) OneIdResponse {
	//检查是否成功进行初始化
	if isInit == false {
		return OneIdResponse{
			Code:    401,
			Message: "Please Init",
		}
	}
	//1.按照规则进行信息生成
	var bindMessage string
	bindMessage = "bind|" + channelMessage["channelCode1"] + "|" + channelMessage["channelType1"] + "|" + sourceID + "|" + channelMessage["channelCode2"] + "|" + channelMessage["channelType2"] + "|" + destID
	//2.发送信息给pulsar
	client, clientErr := pulsar.NewClient(pulsar.ClientOptions{
		URL:               os.Getenv("oneId_pulsar"), // Pulsar服务地址
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if clientErr != nil {
		return OneIdResponse{Code: 400, Message: clientErr.Error()}
	}
	producer, producerErr := client.CreateProducer(pulsar.ProducerOptions{
		Topic: os.Getenv("oneId_pulsar_topic"),
	})
	if producerErr != nil {
		return OneIdResponse{Code: 400, Message: producerErr.Error()}
	}
	defer producer.Close()

	msg := pulsar.ProducerMessage{
		Payload: []byte(bindMessage),
	}
	if _, sendErr := producer.Send(ctx, &msg); sendErr != nil {
		return OneIdResponse{Code: 400, Message: sendErr.Error()}
	}
	return OneIdResponse{Code: 0, Message: "success"}
}

// unbind 用于用户解绑两个账号，生成字符串传输给pulsar消息队列
func unbind(ctx context.Context, sourceID, sourceType, destID, destType string) OneIdResponse {
	//检查是否成功进行初始化
	if isInit == false {
		return OneIdResponse{
			Code:    401,
			Message: "Please Init",
		}
	}
	//1.按照规则进行信息生成
	var bindMessage string
	bindMessage = "bind|" + channelMessage["channelCode1"] + "|" + channelMessage["channelType1"] + "|" + sourceID + "|" + channelMessage["channelCode2"] + "|" + channelMessage["channelType2"] + "|" + destID
	//2.发送信息给pulsar
	client, clientErr := pulsar.NewClient(pulsar.ClientOptions{
		URL:               os.Getenv("oneId_pulsar"), // Pulsar服务地址
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if clientErr != nil {
		return OneIdResponse{Code: 400, Message: clientErr.Error()}
	}
	producer, producerErr := client.CreateProducer(pulsar.ProducerOptions{
		Topic: os.Getenv("oneId_pulsar_topic"),
	})
	if producerErr != nil {
		return OneIdResponse{Code: 400, Message: producerErr.Error()}
	}
	defer producer.Close()

	msg := pulsar.ProducerMessage{
		Payload: []byte(bindMessage),
	}
	if _, sendErr := producer.Send(ctx, &msg); sendErr != nil {
		return OneIdResponse{Code: 400, Message: sendErr.Error()}
	}
	return OneIdResponse{Code: 0, Message: "success"}
}
