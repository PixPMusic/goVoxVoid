package main

import (
	"fmt"
	"time"

	vzxapi "github.com/PixPMusic/goVoxVoid"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

const (

	// FX Selector Colors
	fxSelectorActive   = 53
	fxSelectorInactive = 55
)

func fxSelectorLightReset() {
	send(midi.NoteOn(0, 11, fxSelectorInactive))
	send(midi.NoteOn(0, 12, fxSelectorInactive))
	send(midi.NoteOn(0, 13, fxSelectorInactive))
	send(midi.NoteOn(0, 14, fxSelectorInactive))
	send(midi.NoteOn(0, 15, fxSelectorInactive))
	send(midi.NoteOn(0, 16, fxSelectorInactive))
	send(midi.NoteOn(0, 17, fxSelectorInactive))
	send(midi.NoteOn(0, 18, fxSelectorInactive))
	send(midi.NoteOn(0, 19, fxSelectorInactive))
}

func handle(key uint8) {
	switch key {
	case 11:
		vzxapi.PlayQueueCurrentItemFxLevelSet(0.0)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 11, fxSelectorActive))
	case 12:
		vzxapi.PlayQueueCurrentItemFxLevelSet(0.25)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 12, fxSelectorActive))
	case 13:
		vzxapi.PlayQueueCurrentItemFxLevelSet(0.50)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 13, fxSelectorActive))
	case 14:
		vzxapi.PlayQueueCurrentItemFxLevelSet(0.75)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 14, fxSelectorActive))
	case 15:
		vzxapi.PlayQueueCurrentItemFxLevelSet(1.0)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 15, fxSelectorActive))
	case 16:
		vzxapi.PlayQueueCurrentItemFxLevelSet(1.25)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 16, fxSelectorActive))
	case 17:
		vzxapi.PlayQueueCurrentItemFxLevelSet(1.50)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 17, fxSelectorActive))
	case 18:
		vzxapi.PlayQueueCurrentItemFxLevelSet(1.75)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 18, fxSelectorActive))
	case 19:
		vzxapi.PlayQueueCurrentItemFxLevelSet(2.0)
		fxSelectorLightReset()
		send(midi.NoteOn(0, 19, fxSelectorActive))
	}
}

var send func(midi.Message) error

func main() {
	defer midi.CloseDriver()

	in, err := midi.FindInPort("MIDIIN2 (LPMiniMK3 MIDI) 2")
	if err != nil {
		panic(err)
	}

	out, err := midi.FindOutPort("MIDIOUT2 (LPMiniMK3 MIDI) 3")
	if err != nil {
		panic(err)
	}

	send, err = midi.SendTo(out)

	send([]byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x00, 0x7F, 0xF7})

	fxSelectorLightReset()

	_, err = midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var ch, key, vel uint8
		if msg.GetNoteOn(&ch, &key, &vel) {
			handle(key)
		} else if msg.GetControlChange(&ch, &key, &vel) {
			if vel == 127 {
				handle(key)
			}
		}
	}, midi.UseSysEx())

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	for true {
		time.Sleep(1000)
	}
}
