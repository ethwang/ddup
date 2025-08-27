package demo

import (
	"context"
)

type KafkaConfig struct {
	Brokers   string `json:"brokers"`
	Consumers struct {
		Topic string `json:"topic"`
		Group string `json:"group"`
		Init  bool   `json:"init"`
	} `json:"consumers"`

	ConsumerThirdPartyReport struct {
		Topic string `json:"topic"`
		Group string `json:"group"`
		Init  bool   `json:"init"`
	} `json:"consumerThirdPartyReport"`
}

/*
	一个 interface （打埋点方法） Event
	两个 struct （实现 interface） LogEvent,KaflkaEvent
	三个 成员变量 (struct有三个成员变量)  Topic,GetHeader,HandleFailed
*/

func NewLogEvent(kafkaCfg *KafkaConfig, topic string, getHeader func(uid int64) (map[string]string, error), handleFailed func(err error)) *LogEvent {

	// 初始化 kaflka
	//kafkaProducer, err := kafka.NewAsyncProducer(kafkaCfg.Brokers,
	//	kafka.ConfigProducerMiddleware(),
	//)
	//if err != nil {
	//	panic(err)
	//}

	//// 回调业务处理 err
	//go func() {
	//	for {
	//		select {
	//		case err := <-kafkaProducer.Errors():
	//			handleFailed(err)
	//		}
	//	}
	//}()
	//// 返回实体
	//return &LogEvent{
	//	Topic:        topic,
	//	GetHeader:    getHeader,
	//	HandleFailed: handleFailed,
	//}
	return nil
}

func NewKalkaEvent(kafkaCfg *KafkaConfig, topic string, getHeader func(uid int64) (map[string]string, error), handleFailed func(err error)) *KaflkaEvent {

	// 初始化 kaflka
	//kafkaProducer, err := kafka.NewAsyncProducer(kafkaCfg.Brokers,
	//	kafka.ConfigProducerMiddleware(),
	//)
	//if err != nil {
	//	panic(err)
	//}

	//// 回调业务处理 err
	//go func() {
	//	for {
	//		select {
	//		case err := <-kafkaProducer.Errors():
	//			handleFailed(err)
	//		}
	//	}
	//}()
	//// 返回实体
	//return &KaflkaEvent{
	//	Topic:        topic,
	//	GetHeader:    getHeader,
	//	HandleFailed: handleFailed,
	//}
	return nil
}

type Event interface {
	AddEvent(ctx context.Context, eventName string, uid int64, extra map[string]string) error
}

type LogEvent struct {
	Topic        string
	GetHeader    func(uid int64) (map[string]string, error) // 获取 header
	HandleFailed func(err error)                            // 异步发送失败回调方法
}

func (le *LogEvent) AddEvent(ctx context.Context, eventName string, uid int64, extra map[string]string) error {
	/*
		1. 获取头信息：le.GetHeader, 报错返回error;
		2. 写入 log 报错返回error;
		3. 异步写入 kaflka(le.Topic);
	*/
	return nil
}

type KaflkaEvent struct {
	Topic        string
	GetHeader    func(uid int64) (map[string]string, error) // 获取 header
	HandleFailed func(err error)                            // 异步发送失败回调方法
}

func (ke *KaflkaEvent) AddEvent(ctx context.Context, eventName string, uid int64, extra map[string]string) error {
	/*
		1. 消费kalka; 兼容旧版本
		1. 获取头信息：ke.GetHeader, 报错返回error;
		3. 异步写入 kaflka(ke.Topic);
	*/
	return nil
}
