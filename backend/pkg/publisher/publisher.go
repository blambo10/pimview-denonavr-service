package publisher

import (
	"pimview.thelabshack.com/pkg/denonavr"
	"pimview.thelabshack.com/pkg/log"
	mqtt "pimview.thelabshack.com/pkg/mqtt"
	"time"
)

const (
	topic = "denonavr/volume"
)

var (
	logger = log.NewLogger()
)

func Run(client mqtt.Client) {
	avr := denonavr.New()

	currentVolume, err := avr.GetVolume()
	if err != nil {
		logger.Info(err)
	}

	Publish(client, currentVolume)
}

func Publish(client mqtt.Client, msg string) {
	topic := "denonavr/currentvolume"

	logger.Info(topic + " " + msg)
	//text := fmt.Sprintf("on")

	token := client.Publish(topic, 0, false, msg)
	token.Wait()
	time.Sleep(time.Second)
}
