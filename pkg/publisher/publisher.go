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

func Run() {
	avr := denonavr.New()

	currentVolume, err := avr.GetVolume()
	if err != nil {
		logger.Info(err)
	}

	Publish(currentVolume)
}

func Publish(msg string) {
	topic := "denonavr/currentvolume"
	client := mqtt.GetClient("pimviewpub1")

	logger.Info(topic + " " + msg)
	//text := fmt.Sprintf("on")

	token := client.Publish(topic, 0, false, msg)
	token.Wait()
	time.Sleep(time.Second)
}
