package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
	"strings"
	"time"
)

var MQTT *autopaho.ConnectionManager

type MqttPayload struct {
	Action  string      `json:"action,omitempty"`
	ID      string      `json:"id,omitempty"`
	Name    string      `json:"name,omitempty"`
	Source  string      `json:"src,omitempty"`
	Error   error       `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  string      `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func InitMQTT() {
	log.Info("MQTT: Init")

	serverURL, err := url.Parse(viper.GetString("mqtt.url"))
	if err != nil {
		log.Errorf("MQTT: Init error: %s", err)
	}

	cliCfg := autopaho.ClientConfig{
		BrokerUrls:        []*url.URL{serverURL},
		KeepAlive:         10,
		ConnectRetryDelay: 3 * time.Second,
		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
			log.Info("MQTT: Connection up")
			if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
				Subscriptions: map[string]paho.SubscribeOptions{
					viper.GetString("mqtt.topic"): {QoS: byte(1)},
				},
			}); err != nil {
				log.Errorf("MQTT: Subscribe error: %s", err)
				return
			}
			log.Info("MQTT: Subscription made")
		},
		OnConnectError: func(err error) {
			log.Errorf("MQTT: Attempting connection error: %s", err)
		},
		ClientConfig: paho.ClientConfig{
			ClientID: viper.GetString("mqtt.client_id"),
			//Router: paho.RegisterHandler(common.WorkflowExec, m.execMessage),
			Router: paho.NewStandardRouter(),
			OnClientError: func(err error) {
				log.Errorf("MQTT: Client error: %s", err)
			},
			OnServerDisconnect: func(d *paho.Disconnect) {
				if d.Properties != nil {
					log.Errorf("MQTT: Server requested disconnect: %s", d.Properties.ReasonString)
				} else {
					log.Errorf("MQTT: Server requested disconnect: %s", d.ReasonCode)
				}
			},
		},
	}

	cliCfg.SetUsernamePassword(viper.GetString("mqtt.user"), []byte(viper.GetString("mqtt.password")))

	if viper.GetString("mqtt.debug") == "true" {
		debugLog := NewPahoLogAdapter(1)
		cliCfg.Debug = debugLog
		cliCfg.PahoDebug = debugLog
	}

	MQTT, err = autopaho.NewConnection(context.Background(), cliCfg)
	if err != nil {
		log.Errorf("MQTT: Fail to connect: %s", err)
	}

	cliCfg.Router.RegisterHandler(viper.GetString("mqtt.topic"), execMessage)
}

func execMessage(m *paho.Publish) {
	log.Debugf("MQTT: Received message: %s from topic: %s\n", string(m.Payload), m.Topic)
	id := "false"
	s := strings.Split(m.Topic, "/")
	p := string(m.Payload)

	if s[0] == "kli" && len(s) == 5 {
		id = s[4]
	} else if s[0] == "exec" && len(s) == 4 {
		id = s[3]
	}

	if id == "false" {
		switch p {
		case "start":
			//	go a.startExecMqtt(p)
			//case "stop":
			//	go a.stopExecMqtt(p)
			//case "status":
			//	go a.execStatusMqtt(p)
		}
	}

	if id != "false" {
		switch p {
		case "start":
			//	go a.startExecMqttByID(p, id)
			//case "stop":
			//	go a.stopExecMqttByID(p, id)
			//case "status":
			//	go a.execStatusMqttByID(p, id)
			//case "cmdstat":
			//	go a.cmdStatMqtt(p, id)
			//case "progress":
			//	go a.getProgressMqtt(p, id)
			//case "report":
			//	go a.getReportMqtt(p, id)
			//case "alive":
			//	go a.isAliveMqtt(p, id)
		}
	}
}

func SendRespond(id string, m *MqttPayload) {
	var topic string

	if id == "false" {
		topic = viper.GetString("mqtt.srv_topic")
	} else {
		topic = viper.GetString("mqtt.srv_topic") + "/" + id
	}
	message, err := json.Marshal(m)
	if err != nil {
		log.Errorf("MQTT: Message parsing error: %s", err)
	}

	pa, err := MQTT.Publish(context.Background(), &paho.Publish{
		QoS:     byte(1),
		Retain:  false,
		Topic:   topic,
		Payload: message,
	})
	if err != nil {
		log.Errorf("MQTT: Publish error: %s, reason: %s", err, pa.Properties.ReasonString)
	}
}

func SendMessage(id string) {
	var topic string
	var m interface{}
	//date := time.Now().Format("2006-01-02")

	//if id == "ingest" {
	//	topic = common.MonitorIngestTopic
	//	m, _ = models.FindIngest(DB, "date", date)
	//}
	//
	//if id == "trimmer" {
	//	topic = common.MonitorTrimmerTopic
	//	m, _ = models.FindTrimmer(a.DB, "date", date)
	//}
	//
	//if id == "archive" {
	//	topic = common.MonitorArchiveTopic
	//	m, _ = models.FindKmFiles(a.DB, "date", date)
	//}
	//
	//if id == "trim" {
	//	topic = common.StateTrimmerTopic
	//	m, _ = models.GetFilesToTrim(a.DB)
	//}
	//
	//if id == "drim" {
	//	topic = common.StateDgimaTopic
	//	m, _ = models.GetFilesToDgima(a.DB)
	//}
	//
	//if id == "bdika" {
	//	topic = common.StateArichaTopic
	//	m, _ = models.GetBdika(a.DB)
	//}
	//
	//if id == "jobs" {
	//	topic = common.StateJobsTopic
	//	m, _ = models.GetActiveJobs(a.DB)
	//}
	//
	//if id == "langcheck" {
	//	topic = common.StateLangcheckTopic
	//	var s models.State
	//	s.StateID = "langcheck"
	//	_ = s.GetState(a.DB)
	//	m = s.Data
	//}

	message, err := json.Marshal(m)
	if err != nil {
		log.Errorf("MQTT: Message parsing error: %s", err)
	}

	pa, err := MQTT.Publish(context.Background(), &paho.Publish{
		QoS:     byte(1),
		Retain:  true,
		Topic:   topic,
		Payload: message,
	})
	if err != nil {
		log.Errorf("MQTT: Publish error: %s, reason: %s", err, pa.Properties.ReasonString)
	}

	log.Debugf("MQTT: Send message: %s to topic: %s\n", string(message), topic)
}

type PahoLogAdapter struct {
	level log.Level
}

func NewPahoLogAdapter(level log.Level) *PahoLogAdapter {
	return &PahoLogAdapter{level: level}
}

func (a *PahoLogAdapter) Println(v ...interface{}) {
	log.Infof("MQTT: %s", fmt.Sprint(v...))
}

func (a *PahoLogAdapter) Printf(format string, v ...interface{}) {
	log.Infof("MQTT: %s", fmt.Sprintf(format, v...))
}
