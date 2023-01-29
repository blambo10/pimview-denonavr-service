package config

import (
	"github.com/gobuffalo/envy"
	l "github.com/sirupsen/logrus"
	"strconv"
)

type Config struct {
	logging Logging
	mqtt    MQTT
}

type MQTT struct {
	Address  string
	Port     int
	ClientID string
}

type Logging struct {
	Level int
}

func GetLogEnvs() Logging {
	level, err := strconv.Atoi(envy.Get("LOGGING_LEVEL", "4"))

	if err != nil {
		level = 4
	}

	if level > len(l.AllLevels)-1 || level < 0 {
		level = 4
	}

	return Logging{
		Level: level,
	}
}

func NewConfig() *Config {
	envy.Load()

	port, err := strconv.Atoi(envy.Get("MQTT_BROKER_PORT", ""))

	if err != nil {
		port = 1883
	}

	return &Config{
		logging: GetLogEnvs(),
		mqtt: MQTT{
			Address:  envy.Get("MQTT_BROKER_ADDRESS", "localhost"),
			Port:     port,
			ClientID: envy.Get("MQTT_BROKER_CLIENTID", "pimview-test"),
		},
	}
}

func GetLogger() Logging {
	return NewConfig().logging
}

func GetMQTT() MQTT {
	return NewConfig().mqtt
}
