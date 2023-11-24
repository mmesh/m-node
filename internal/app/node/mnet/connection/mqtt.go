package connection

import (
	"crypto/tls"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const mqttPort int = 1883

var mqttConnectionWatcherCh = make(chan struct{}, 4)

var ribDataMsgRxQueue = make(chan []byte, 256)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	xlog.Debugf("[mqtt] Received routing message: %s from topic: %s", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	xlog.Info("[mqtt] Routing session connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	xlog.Warnf("[mqtt] Routing session lost: %v", err)
	mqttConnectionWatcherCh <- struct{}{}
}

func (c *connection) newRoutingClient(controllerHost string) error {
	if c.mqttClient != nil {
		c.mqttClient.Disconnect(250)
		c.mqttClient = nil
	}

	mqttClientID := fmt.Sprintf("%s:%s:%s:%s:%s",
		c.node.AccountID, c.node.TenantID, c.node.Cfg.NetID, c.node.Cfg.SubnetID, c.node.NodeID)

	opts := mqtt.NewClientOptions()
	// opts.AddBroker(fmt.Sprintf("tcp://%s:%d", controllerHost, mqttPort))
	opts.AddBroker(fmt.Sprintf("ssl://%s:%d", controllerHost, mqttPort))
	opts.SetTLSConfig(&tls.Config{
		// MinVersion: tls.VersionTLS13,
		MinVersion: tls.VersionTLS12,
	})
	opts.SetClientID(mqttClientID)
	opts.SetUsername(c.node.NodeID)
	opts.SetPassword(c.authKey.Key)
	opts.SetOrderMatters(false)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(true)

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetOnConnectHandler(connectHandler)
	opts.SetConnectionLostHandler(connectLostHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err := token.Error()
		return errors.Wrapf(err, "[%v] function client.Connect()", errors.Trace())
	}

	c.mqttClient = client

	return nil
}

func (c *connection) newRoutingSubscription(topic string) error {
	if c.mqttClient == nil {
		return fmt.Errorf("[mqtt] unable to create a new routing subscription: nil client")
	}

	if token := c.mqttClient.Subscribe(topic, 0, c.ribDataMsgHandler); token.Wait() && token.Error() != nil {
		err := token.Error()
		return errors.Wrapf(err, "[%v] function c.mqttClient.Subscribe()", errors.Trace())
	}

	xlog.Infof("[mqtt] Routing session subscribed to %s", topic)

	return nil
}

func (c *connection) ribDataMsgHandler(client mqtt.Client, msg mqtt.Message) {
	xlog.Debugf("[mqtt] Received ribData message from topic: %s", msg.Topic())

	ribDataMsgRxQueue <- msg.Payload()
}
