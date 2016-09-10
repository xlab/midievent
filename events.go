package midievent

type Event byte

// Note Off Events
// 2nd byte: Note Number (0-127)
// 3rd byte: Note Velocity (0-127)
const (
	Chan1NoteOffEvent  Event = 0x80
	Chan2NoteOffEvent  Event = 0x81
	Chan3NoteOffEvent  Event = 0x82
	Chan4NoteOffEvent  Event = 0x83
	Chan5NoteOffEvent  Event = 0x84
	Chan6NoteOffEvent  Event = 0x85
	Chan7NoteOffEvent  Event = 0x86
	Chan8NoteOffEvent  Event = 0x87
	Chan9NoteOffEvent  Event = 0x88
	Chan10NoteOffEvent Event = 0x89
	Chan11NoteOffEvent Event = 0x8A
	Chan12NoteOffEvent Event = 0x8B
	Chan13NoteOffEvent Event = 0x8C
	Chan14NoteOffEvent Event = 0x8D
	Chan15NoteOffEvent Event = 0x8E
	Chan16NoteOffEvent Event = 0x8F
)

// Note On Events
// 2nd byte: Note Number (0-127)
// 3rd byte: Note Velocity (0-127)
const (
	Chan1NoteOnEvent  Event = 0x90
	Chan2NoteOnEvent  Event = 0x91
	Chan3NoteOnEvent  Event = 0x92
	Chan4NoteOnEvent  Event = 0x93
	Chan5NoteOnEvent  Event = 0x94
	Chan6NoteOnEvent  Event = 0x95
	Chan7NoteOnEvent  Event = 0x96
	Chan8NoteOnEvent  Event = 0x97
	Chan9NoteOnEvent  Event = 0x98
	Chan10NoteOnEvent Event = 0x99
	Chan11NoteOnEvent Event = 0x9A
	Chan12NoteOnEvent Event = 0x9B
	Chan13NoteOnEvent Event = 0x9C
	Chan14NoteOnEvent Event = 0x9D
	Chan15NoteOnEvent Event = 0x9E
	Chan16NoteOnEvent Event = 0x9F
)

// Polyphonic Aftertouch Events
// 2nd byte: Note Number (0-127)
// 3rd byte: Note Velocity (0-127)
const (
	Chan1PolyphonicAftertouchEvent  Event = 0xA0
	Chan2PolyphonicAftertouchEvent  Event = 0xA1
	Chan3PolyphonicAftertouchEvent  Event = 0xA2
	Chan4PolyphonicAftertouchEvent  Event = 0xA3
	Chan5PolyphonicAftertouchEvent  Event = 0xA4
	Chan6PolyphonicAftertouchEvent  Event = 0xA5
	Chan7PolyphonicAftertouchEvent  Event = 0xA6
	Chan8PolyphonicAftertouchEvent  Event = 0xA7
	Chan9PolyphonicAftertouchEvent  Event = 0xA8
	Chan10PolyphonicAftertouchEvent Event = 0xA9
	Chan11PolyphonicAftertouchEvent Event = 0xAA
	Chan12PolyphonicAftertouchEvent Event = 0xAB
	Chan13PolyphonicAftertouchEvent Event = 0xAC
	Chan14PolyphonicAftertouchEvent Event = 0xAD
	Chan15PolyphonicAftertouchEvent Event = 0xAE
	Chan16PolyphonicAftertouchEvent Event = 0xAF
)

// Control Mode Change Events, Chan1PolyphonicAftertouchEvent
const (
	Chan1ControlModeChangeEvent  Event = 0xB0
	Chan2ControlModeChangeEvent  Event = 0xB1
	Chan3ControlModeChangeEvent  Event = 0xB2
	Chan4ControlModeChangeEvent  Event = 0xB3
	Chan5ControlModeChangeEvent  Event = 0xB4
	Chan6ControlModeChangeEvent  Event = 0xB5
	Chan7ControlModeChangeEvent  Event = 0xB6
	Chan8ControlModeChangeEvent  Event = 0xB7
	Chan9ControlModeChangeEvent  Event = 0xB8
	Chan10ControlModeChangeEvent Event = 0xB9
	Chan11ControlModeChangeEvent Event = 0xBA
	Chan12ControlModeChangeEvent Event = 0xBB
	Chan13ControlModeChangeEvent Event = 0xBC
	Chan14ControlModeChangeEvent Event = 0xBD
	Chan15ControlModeChangeEvent Event = 0xBE
	Chan16ControlModeChangeEvent Event = 0xBF
)

// Program Change Event
// 2nd byte: Program # (0-127)
// 3rd byte: -
const (
	Chan1ProgramChangeEvent  Event = 0xC0
	Chan2ProgramChangeEvent  Event = 0xC1
	Chan3ProgramChangeEvent  Event = 0xC2
	Chan4ProgramChangeEvent  Event = 0xC3
	Chan5ProgramChangeEvent  Event = 0xC4
	Chan6ProgramChangeEvent  Event = 0xC5
	Chan7ProgramChangeEvent  Event = 0xC6
	Chan8ProgramChangeEvent  Event = 0xC7
	Chan9ProgramChangeEvent  Event = 0xC8
	Chan10ProgramChangeEvent Event = 0xC9
	Chan11ProgramChangeEvent Event = 0xCA
	Chan12ProgramChangeEvent Event = 0xCB
	Chan13ProgramChangeEvent Event = 0xCC
	Chan14ProgramChangeEvent Event = 0xCD
	Chan15ProgramChangeEvent Event = 0xCE
	Chan16ProgramChangeEvent Event = 0xCF
)

// Channel Aftertouch Events
// 2nd byte: Aftertouch pressure (0-127)
// 3rd byte: -
const (
	Chan1ChannelAftertouchEvent  Event = 0xD0
	Chan2ChannelAftertouchEvent  Event = 0xD1
	Chan3ChannelAftertouchEvent  Event = 0xD2
	Chan4ChannelAftertouchEvent  Event = 0xD3
	Chan5ChannelAftertouchEvent  Event = 0xD4
	Chan6ChannelAftertouchEvent  Event = 0xD5
	Chan7ChannelAftertouchEvent  Event = 0xD6
	Chan8ChannelAftertouchEvent  Event = 0xD7
	Chan9ChannelAftertouchEvent  Event = 0xD8
	Chan10ChannelAftertouchEvent Event = 0xD9
	Chan11ChannelAftertouchEvent Event = 0xDA
	Chan12ChannelAftertouchEvent Event = 0xDB
	Chan13ChannelAftertouchEvent Event = 0xDC
	Chan14ChannelAftertouchEvent Event = 0xDD
	Chan15ChannelAftertouchEvent Event = 0xDE
	Chan16ChannelAftertouchEvent Event = 0xDF
)

// Pitch Wheel Range Events
// 2nd byte: Pitch wheel LSB (0-127)
// 3rd byte: Pitch wheel MSB (0-127)
const (
	Chan1PitchWheelRangeEvent  Event = 0xE0
	Chan2PitchWheelRangeEvent  Event = 0xE1
	Chan3PitchWheelRangeEvent  Event = 0xE2
	Chan4PitchWheelRangeEvent  Event = 0xE3
	Chan5PitchWheelRangeEvent  Event = 0xE4
	Chan6PitchWheelRangeEvent  Event = 0xE5
	Chan7PitchWheelRangeEvent  Event = 0xE6
	Chan8PitchWheelRangeEvent  Event = 0xE7
	Chan9PitchWheelRangeEvent  Event = 0xE8
	Chan10PitchWheelRangeEvent Event = 0xE9
	Chan11PitchWheelRangeEvent Event = 0xEA
	Chan12PitchWheelRangeEvent Event = 0xEB
	Chan13PitchWheelRangeEvent Event = 0xEC
	Chan14PitchWheelRangeEvent Event = 0xED
	Chan15PitchWheelRangeEvent Event = 0xEE
	Chan16PitchWheelRangeEvent Event = 0xEF
)

// Common Events
const (
	// System Exclusive (data dump) 2nd byte= Vendor ID followed by more data bytes and ending with EOX (0x7F).
	SysExclusiveEvent        Event = 0xF0
	SysUndefinedEvent        Event = 0xF1
	SongPositionPointerEvent Event = 0xF2 // LSB	MSB
	SongSelectEvent          Event = 0xF3 // (0-127)	-
	SysUndefinedEvent2       Event = 0xF4
	SysUndefinedEvent3       Event = 0xF5
	SysTuneRequestEvent      Event = 0xF6
	SysEndOfSysExEvent       Event = 0xF7
	SysTimingClockEvent      Event = 0xF8
	SysUndefinedEvent4       Event = 0xF9
	SysStartEvent            Event = 0xFA
	SysContinueEvent         Event = 0xFB
	SysStopEvent             Event = 0xFC
	SysUndefinedEvent5       Event = 0xFD
	SysActiveSensingEvent    Event = 0xFE
	SysResetEvent            Event = 0xFF
)

func ChanOf(ev Event) (ch int, ok bool) {
	ok = true
	switch ev {
	case Chan1NoteOffEvent, Chan1NoteOnEvent, Chan1PolyphonicAftertouchEvent,
		Chan1ControlModeChangeEvent, Chan1ProgramChangeEvent,
		Chan1ChannelAftertouchEvent, Chan1PitchWheelRangeEvent:
		ch = 1
	case Chan2NoteOffEvent, Chan2NoteOnEvent, Chan2PolyphonicAftertouchEvent,
		Chan2ControlModeChangeEvent, Chan2ProgramChangeEvent,
		Chan2ChannelAftertouchEvent, Chan2PitchWheelRangeEvent:
		ch = 2
	case Chan3NoteOffEvent, Chan3NoteOnEvent, Chan3PolyphonicAftertouchEvent,
		Chan3ControlModeChangeEvent, Chan3ProgramChangeEvent,
		Chan3ChannelAftertouchEvent, Chan3PitchWheelRangeEvent:
		ch = 3
	case Chan4NoteOffEvent, Chan4NoteOnEvent, Chan4PolyphonicAftertouchEvent,
		Chan4ControlModeChangeEvent, Chan4ProgramChangeEvent,
		Chan4ChannelAftertouchEvent, Chan4PitchWheelRangeEvent:
		ch = 4
	case Chan5NoteOffEvent, Chan5NoteOnEvent, Chan5PolyphonicAftertouchEvent,
		Chan5ControlModeChangeEvent, Chan5ProgramChangeEvent,
		Chan5ChannelAftertouchEvent, Chan5PitchWheelRangeEvent:
		ch = 5
	case Chan6NoteOffEvent, Chan6NoteOnEvent, Chan6PolyphonicAftertouchEvent,
		Chan6ControlModeChangeEvent, Chan6ProgramChangeEvent,
		Chan6ChannelAftertouchEvent, Chan6PitchWheelRangeEvent:
		ch = 6
	case Chan7NoteOffEvent, Chan7NoteOnEvent, Chan7PolyphonicAftertouchEvent,
		Chan7ControlModeChangeEvent, Chan7ProgramChangeEvent,
		Chan7ChannelAftertouchEvent, Chan7PitchWheelRangeEvent:
		ch = 7
	case Chan8NoteOffEvent, Chan8NoteOnEvent, Chan8PolyphonicAftertouchEvent,
		Chan8ControlModeChangeEvent, Chan8ProgramChangeEvent,
		Chan8ChannelAftertouchEvent, Chan8PitchWheelRangeEvent:
		ch = 8
	case Chan9NoteOffEvent, Chan9NoteOnEvent, Chan9PolyphonicAftertouchEvent,
		Chan9ControlModeChangeEvent, Chan9ProgramChangeEvent,
		Chan9ChannelAftertouchEvent, Chan9PitchWheelRangeEvent:
		ch = 9
	case Chan10NoteOffEvent, Chan10NoteOnEvent, Chan10PolyphonicAftertouchEvent,
		Chan10ControlModeChangeEvent, Chan10ProgramChangeEvent,
		Chan10ChannelAftertouchEvent, Chan10PitchWheelRangeEvent:
		ch = 10
	case Chan11NoteOffEvent, Chan11NoteOnEvent, Chan11PolyphonicAftertouchEvent,
		Chan11ControlModeChangeEvent, Chan11ProgramChangeEvent,
		Chan11ChannelAftertouchEvent, Chan11PitchWheelRangeEvent:
		ch = 11
	case Chan12NoteOffEvent, Chan12NoteOnEvent, Chan12PolyphonicAftertouchEvent,
		Chan12ControlModeChangeEvent, Chan12ProgramChangeEvent,
		Chan12ChannelAftertouchEvent, Chan12PitchWheelRangeEvent:
		ch = 12
	case Chan13NoteOffEvent, Chan13NoteOnEvent, Chan13PolyphonicAftertouchEvent,
		Chan13ControlModeChangeEvent, Chan13ProgramChangeEvent,
		Chan13ChannelAftertouchEvent, Chan13PitchWheelRangeEvent:
		ch = 13
	case Chan14NoteOffEvent, Chan14NoteOnEvent, Chan14PolyphonicAftertouchEvent,
		Chan14ControlModeChangeEvent, Chan14ProgramChangeEvent,
		Chan14ChannelAftertouchEvent, Chan14PitchWheelRangeEvent:
		ch = 14
	case Chan15NoteOffEvent, Chan15NoteOnEvent, Chan15PolyphonicAftertouchEvent,
		Chan15ControlModeChangeEvent, Chan15ProgramChangeEvent,
		Chan15ChannelAftertouchEvent, Chan15PitchWheelRangeEvent:
		ch = 15
	case Chan16NoteOffEvent, Chan16NoteOnEvent, Chan16PolyphonicAftertouchEvent,
		Chan16ControlModeChangeEvent, Chan16ProgramChangeEvent,
		Chan16ChannelAftertouchEvent, Chan16PitchWheelRangeEvent:
		ch = 16
	default: // unknown
		return -1, false
	}
	return
}

func IsNoteOn(ev Event) bool {
	switch ev {
	case Chan1NoteOnEvent,
		Chan2NoteOnEvent,
		Chan3NoteOnEvent,
		Chan4NoteOnEvent,
		Chan5NoteOnEvent,
		Chan6NoteOnEvent,
		Chan7NoteOnEvent,
		Chan8NoteOnEvent,
		Chan9NoteOnEvent,
		Chan10NoteOnEvent,
		Chan11NoteOnEvent,
		Chan12NoteOnEvent,
		Chan13NoteOnEvent,
		Chan14NoteOnEvent,
		Chan15NoteOnEvent,
		Chan16NoteOnEvent:
		return true
	}
	return false
}

func IsNoteOff(ev Event) bool {
	switch ev {
	case Chan1NoteOffEvent,
		Chan2NoteOffEvent,
		Chan3NoteOffEvent,
		Chan4NoteOffEvent,
		Chan5NoteOffEvent,
		Chan6NoteOffEvent,
		Chan7NoteOffEvent,
		Chan8NoteOffEvent,
		Chan9NoteOffEvent,
		Chan10NoteOffEvent,
		Chan11NoteOffEvent,
		Chan12NoteOffEvent,
		Chan13NoteOffEvent,
		Chan14NoteOffEvent,
		Chan15NoteOffEvent,
		Chan16NoteOffEvent:
		return true
	}
	return false
}
