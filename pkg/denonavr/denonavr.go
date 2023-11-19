package denonavr

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	telnet "github.com/reiver/go-telnet"
	"pimview.thelabshack.com/pkg/config"
	"pimview.thelabshack.com/pkg/log"
	"strings"
)

type DenonAVR struct {
	Host   string
	Client telnet.Caller
}

const (
	volume = "volume"
	up     = "up"
	down   = "down"
	mute   = "mute"

	DenonMaster       = "MV"
	DenonVolumeUP     = "UP"
	DenonVolumeDown   = "DOWN"
	DenonVolumeMute   = "MUON"
	DenonVolumeUNMute = "MUOFF"

	DenonMuteState = "MU?"
)

var (
	logger = log.NewLogger()
	cfg    = config.GetDeviceConfig()
)

func New() *DenonAVR {
	t := telnet.StandardCaller
	device := fmt.Sprintf("%s:%s", cfg.Address, cfg.Port)

	return &DenonAVR{
		Host:   device,
		Client: t,
	}
}

// ProcessMessages process mqtt message queue and dispatch to handlers
// client mqtt broker client
// message mqtt message including topic and payload
func (h *DenonAVR) ProcessMessages(client mqtt.Client, message mqtt.Message) {

	logger.Info("processing messages ...")
	logger.Info(message.Topic())

	switch {
	case strings.Contains(message.Topic(), volume):
		h.Volume(message.Payload())
	}
}

// Volume volume handler
// direction mqtt message payload
func (h *DenonAVR) Volume(direction []byte) {
	d := string(direction)

	logger.Info(d)
	switch d {
	case up:
		logger.Infof("Volume UP")
		err := h.VolumeUp()
		if err != nil {
			logger.Error(err)
		}
	case down:
		logger.Infof("Volume DOWN")
		err := h.VolumeDown()

		if err != nil {
			logger.Error(err)
		}
	case mute:
		logger.Infof("Toggle Volume Mute")
		err := h.ToggleMute()

		if err != nil {
			logger.Error(err)
		}
	}
}

func (h *DenonAVR) VolumeUp() error {

	conn, _ := telnet.DialTo(h.Host)
	defer conn.Close()

	cmd := fmt.Sprintf("%s%s", DenonMaster, DenonVolumeUP)

	_, err := conn.Write([]byte(cmd))

	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (h *DenonAVR) VolumeDown() error {
	conn, _ := telnet.DialTo(h.Host)
	defer conn.Close()

	cmd := fmt.Sprintf("%s%s", DenonMaster, DenonVolumeDown)

	_, err := conn.Write([]byte(cmd))

	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (h *DenonAVR) SetVolume(v int) error {

	return nil
}

func (h *DenonAVR) ToggleMute() error {

	var newMuteState string

	conn, _ := telnet.DialTo(h.Host)
	defer conn.Close()

	conn.Write([]byte(DenonMuteState))

	commandResponse := make([]byte, 5)
	_, err := conn.Read(commandResponse)
	if err != nil {
		logger.Error(err)
		return err
	}

	currentState := strings.TrimSpace(string(commandResponse))

	switch currentState {
	case DenonVolumeMute:
		newMuteState = DenonVolumeUNMute
		logger.Info("unmuting receiver")
	case DenonVolumeUNMute:
		newMuteState = DenonVolumeMute
		logger.Info("muting receiver")
	default:
		err := fmt.Errorf("unable to query current mute state for %s : response %s", h.Host, currentState)
		logger.Error(err)
		return err
	}

	_, err = conn.Write([]byte(newMuteState))

	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
