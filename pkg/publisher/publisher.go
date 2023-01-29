package publisher

import (
	"fmt"
	mqtt "pimview.thelabshack.com/pkg/mqtt"
	"time"
)

func Publish() {
	client := mqtt.GetClient("pimviewpub1")

	text := fmt.Sprintf("on")
	token := client.Publish("denonavr/power", 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}
