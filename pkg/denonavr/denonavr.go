package denonavr

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	telnet "github.com/reiver/go-telnet"
	"strings"
)

type DenonAVR struct {
	Host   string
	Client telnet.Caller

	//ClientKey string
	//Client heosapi.Heos
}

const (
	volume = "volume"
	up     = "up"
	down   = "down"
	mute   = "mute"

	denonmaster = "MV"
	denonup     = "UP"
	denondown   = "DOWN"
	denonmute   = "MUON"
	denonunmute = "MUOFF"
)

func New() *DenonAVR {
	//h := heosapi.NewHeos("192.168.1.206:1255")

	t := telnet.StandardCaller

	return &DenonAVR{
		Host:   "192.168.1.206:23",
		Client: t,
	}
}

// ProcessMessages process mqtt message queue and dispatch to handlers
// client mqtt broker client
// message mqtt message including topic and payload
func (h *DenonAVR) ProcessMessages(client mqtt.Client, message mqtt.Message) {

	fmt.Println("processing messages ...")
	fmt.Println(message.Topic())

	switch {
	case strings.Contains(message.Topic(), volume):
		h.Volume(message.Payload())
	}
}

// Volume volume handler
// direction mqtt message payload
func (h *DenonAVR) Volume(direction []byte) {
	d := string(direction)

	fmt.Println("asdfasd")
	fmt.Println(d)
	switch d {
	case up:
		fmt.Println("Volume UP")
		err := h.VolumeUp()
		if err != nil {
			fmt.Println(err)
		}
	case down:
		fmt.Println("Volume DOWN")
		err := h.VolumeDown()

		if err != nil {
			fmt.Println(err)
		}
	case mute:
		fmt.Println("Volume Mute")
		err := h.Mute()

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *DenonAVR) VolumeUp() error {

	conn, _ := telnet.DialTo(h.Host)
	defer conn.Close()

	cmd := fmt.Sprintf("%s%s", denonmaster, denonup)

	_, err := conn.Write([]byte(cmd))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (h *DenonAVR) VolumeDown() error {
	conn, _ := telnet.DialTo(h.Host)
	defer conn.Close()

	cmd := fmt.Sprintf("%s%s", denonmaster, denondown)

	_, err := conn.Write([]byte(cmd))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (h *DenonAVR) SetVolume(v int) error {

	return nil
}

func (h *DenonAVR) Mute() error {

	conn, _ := telnet.DialTo(h.Host)
	defer conn.Close()

	_, err := conn.Write([]byte(denonmute))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//
//func (h *DenonAVR) VolumeUp() error {
//	h.Client.Connect()
//	defer h.Client.Disconnect()
//
//	cmd := heosapi.Command{
//		Group:   "player",
//		Command: "volume_up",
//	}
//
//	resp, err := h.Client.Send(cmd, nil)
//	if err != nil {
//		//todo: change to logger
//		fmt.Println(resp)
//	}
//
//	return nil
//}
//
//func (h *DenonAVR) VolumeDown() error {
//	err := h.Client.Connect()
//
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//
//	defer func() {
//		err := h.Client.Disconnect()
//
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(2)
//		}
//	}()
//
//	cmd := heosapi.Command{
//		Group:   "player",
//		Command: "volume_down&step=1",
//	}
//
//	resp, err := h.Client.Send(cmd, map[string]string{})
//	if err != nil {
//		//todo: change to logger
//		fmt.Println(resp)
//	}
//
//	fmt.Println(resp)
//
//	return nil
//}
//
//func (h *DenonAVR) SetVolume(v int) error {
//	h.Client.Connect()
//	defer h.Client.Disconnect()
//
//	_ = v
//
//	return nil
//}
//
//func (h *DenonAVR) Mute() error {
//	h.Client.Connect()
//	defer h.Client.Disconnect()
//
//	cmd := heosapi.Command{
//		Group: "",
//		//Command: "playertoggle_mute",
//		Command: "get_volume?pid=1",
//		//Command: "get_groups",
//	}
//
//	//h.Client.
//	resp, err := h.Client.Send(cmd, map[string]string{})
//	if err != nil {
//		//todo: change to logger
//		fmt.Println(resp)
//	}
//
//	//fmt.Println("%+v\n", resp)
//
//	fmt.Println(resp.Payload)
//	fmt.Println(resp.Heos.Result)
//
//	return nil
//}
