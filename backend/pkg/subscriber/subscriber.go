package subscriber

import (
	"fmt"
	pahomqtt "github.com/eclipse/paho.mqtt.golang"
	denonavr "pimview.thelabshack.com/pkg/denonavr"
	mqtt "pimview.thelabshack.com/pkg/mqtt"
	"time"
)

func Run() {
	//s
	h := denonavr.New()
	client := mqtt.GetClient("pimview-denonavr")

	for {
		//Sub to mqtt topic (clean up later)
		Subscribe(client, h.ProcessMessages)
	}
}

func Subscribe(cc pahomqtt.Client, handler pahomqtt.MessageHandler) {
	topic := "denonavr/volume"
	token := cc.Subscribe(topic, 1, handler)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
	time.Sleep(time.Second * 120)
}
