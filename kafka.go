package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"sync"
)

//全局注释 kafka 虽然性能比 rabbitmq要快 但是他丢失数据库的可能性更大，而且还会存在重复接受消息的情况

var Topic = "266"
var partition = int32(0)

func main() {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_2_0_0
	cfg.Producer.Return.Errors = true
	cfg.Net.SASL.Enable = false
	cfg.Producer.Return.Successes = true //这个是关键，否则读取不到消息
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewManualPartitioner //允许指定分组
	//cfg.Consumer.Return.Errors
	//cfg.Group.Return.Notifications = true
	cfg.ClientID = "service-exchange-api"
	var kafka = KafkaConfig{
		Addrs:  []string{"127.0.0.1:9092"},
		Config: cfg,
	}
	_, _, err := NewKafkaClient(kafka)
	fmt.Println("err:", err)
}

type log struct{}

func (log) Print(v ...interface{}) {
	//fmt.Println("implement me")
}

func (log) Printf(format string, v ...interface{}) {
	//fmt.Println("implement me")
}

func (log) Println(v ...interface{}) {
	fmt.Println("implement me")
}

//发送消息 此为异步发送消息
func NewAsyncProducer(client sarama.Client, i int) (sarama.AsyncProducer, func(), error) {
	c, err := sarama.NewAsyncProducerFromClient(client)
	//sarama.NewSyncProducerFromClient() 此为同步
	if err != nil {
		return nil, nil, err
	}
	//Topic 为主题，Partition为区域 Partition如果不给默认为0 记得设置cfg.Producer.Partitioner = sarama.NewManualPartitioner 这里为允许设置指定的分区
	//分区是从0开始，记得在启动配置文件时修改Partition的分区
	//不同的主题包括不同的分区都是有着不同的offset
	c.Input() <- &sarama.ProducerMessage{Topic: Topic, Partition: partition, Key: sarama.StringEncoder(fmt.Sprintf("/topic/market/order-trade")), Value: sarama.StringEncoder("消息发送成功拉ssssssss！！！！" + strconv.Itoa(i))}
	select {
	//case msg := <-producer.Successes():
	//    log.Printf("Produced message successes: [%s]\n",msg.Value)
	case err := <-c.Errors():
		fmt.Println("Produced message failure: ", err)
	default:
		//fmt.Println("Produced message success",err)
	}
	return c, func() {
		err := c.Close()
		if err != nil {
			fmt.Println(err)
		}
	}, nil
}

//客户端接收消息
func NewKafkaClient(cfg KafkaConfig) (sarama.Client, func(), error) {
	sarama.Logger = log{}

	//创建链接 创建客户机
	c, err := sarama.NewClient(cfg.Addrs, cfg.Config)
	if err != nil {
		return nil, nil, err
	}
	//获取该主题的所有区
	/*_, err =  c.Partitions(Topic)
	if err != nil {
		fmt.Println(err)
	}*/
	/*consumer, err := sarama.NewConsumerFromClient(c)
	if err != nil {
		fmt.Println(err)
	}*/
	//获取该主题该分区哪个时间段之后的offset
	/*offset,err :=c.GetOffset(Topic,0,1564416000)
	if err != nil {
		fmt.Println(err)
	}*/
	//必须为空，否则会报没有创建自己发挥想象去解决吧
	offsetManager, err := sarama.NewOffsetManagerFromClient("", c)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		consumer, err := sarama.NewConsumerFromClient(c)
		if err != nil {
			fmt.Println(err)
		}
		//sarama.NewConsumerGroupFromClient("group-1",c)
		loopConsumer(consumer, Topic, partition, "b", offsetManager)
		consumer.Close()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			NewAsyncProducer(c, i)
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	return c, func() {
		err := c.Close()
		if err != nil {
			fmt.Print(err)
		}
	}, nil
}

func loopConsumer(consumer sarama.Consumer, topic string, partition int32, item string, om sarama.OffsetManager) {

	pom, err := om.ManagePartition(topic, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pom.Close()

	offset, _ := pom.NextOffset()
	if offset == -1 {
		offset = sarama.OffsetOldest
	}

	//接受的主题，分区，和从多少偏移值 offset 为 sarama.OffsetNewest会一直刷新最新的
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		pom.MarkOffset(msg.Offset+1, "备注")
		fmt.Printf("[%s] : Consumed message: [%s], offset: [%d]\n", item, string(msg.Value), msg.Offset)
	}
}

type KafkaConfig struct {
	Addrs  []string
	Config *sarama.Config
}
