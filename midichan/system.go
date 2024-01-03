package midichan

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
)

type System struct {
	connected bool
	device    string
	inPort    drivers.In
	stopFunc  func()

	OnMessage func(msg midi.Message)
}

func NewSystem() *System {
	return &System{}
}

func (sys *System) Connect(device string) error {
	sys.connected = false

	inPort, err := midi.FindInPort(device)
	if err != nil {
		return err
	}

	stopFunc, err := midi.ListenTo(inPort, sys.onMessage, midi.UseSysEx())
	if err != nil {
		return err
	}

	sys.device = device
	sys.stopFunc = stopFunc
	sys.inPort = inPort
	sys.setConnected(true)

	return nil
}

func (sys *System) setConnected(connected bool) {
	if connected {
		fmt.Printf("%q connected\n", sys.device)
	} else {
		fmt.Printf("%q disconnected\n", sys.device)
	}
	sys.connected = connected
}

func (sys *System) onMessage(msg midi.Message, timestampm int32) {
	if !sys.connected {
		return
	}
	sys.OnMessage(msg)
}

func (sys *System) Update() {
	if sys.connected && sys.inPort != nil {
		if !sys.inPort.IsOpen() {
			sys.setConnected(false)
			sys.inPort = nil
		}
	}
}
