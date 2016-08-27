package midievent

type Control byte

const (
	BankSelectMSB      Control = 0x00
	ModulationWheelMSB Control = 0x01
	BreathControlMSB   Control = 0x02

	UndefinedMSB Control = 0x03

	FootControllerMSB Control = 0x04
	PortamentoTimeMSB Control = 0x05
	DataEntryMSB      Control = 0x06
	ChannelVolumeMSB  Control = 0x07
	BalanceMSB        Control = 0x08

	UndefinedMSB2 Control = 0x09

	PanMSB                  Control = 0x0A
	ExpressionControllerMSB Control = 0x0B
	EffectControl1MSB       Control = 0x0C
	EffectControl2MSB       Control = 0x0D

	UndefinedMSB3 Control = 0x0E
	UndefinedMSB4 Control = 0x0F

	GeneralPurposeController1MSB Control = 0x10
	GeneralPurposeController2MSB Control = 0x11
	GeneralPurposeController3MSB Control = 0x12
	GeneralPurposeController4MSB Control = 0x13

	UndefinedMSB5  Control = 0x14
	UndefinedMSB6  Control = 0x15
	UndefinedMSB7  Control = 0x16
	UndefinedMSB8  Control = 0x17
	UndefinedMSB9  Control = 0x18
	UndefinedMSB10 Control = 0x19
	UndefinedMSB11 Control = 0x1A
	UndefinedMSB12 Control = 0x1B
	UndefinedMSB13 Control = 0x1C
	UndefinedMSB14 Control = 0x1D
	UndefinedMSB15 Control = 0x1E
	UndefinedMSB16 Control = 0x1F

	BankSelectLSB      Control = 0x20
	ModulationWheelLSB Control = 0x21
	BreathControlLSB   Control = 0x22

	UndefinedLSB Control = 0x23

	FootControllerLSB Control = 0x24
	PortamentoTimeLSB Control = 0x25
	DataEntryLSB      Control = 0x26
	ChannelVolumeLSB  Control = 0x27
	BalanceLSB        Control = 0x28

	UndefinedLSB2 Control = 0x29

	PanLSB                  Control = 0x2A
	ExpressionControllerLSB Control = 0x2B
	EffectControl1LSB       Control = 0x2C
	EffectControl2LSB       Control = 0x2D

	UndefinedLSB3 Control = 0x2E
	UndefinedLSB4 Control = 0x2F

	GeneralPurposeController1LSB Control = 0x30
	GeneralPurposeController2LSB Control = 0x31
	GeneralPurposeController3LSB Control = 0x32
	GeneralPurposeController4LSB Control = 0x33

	UndefinedLSB5  Control = 0x34
	UndefinedLSB6  Control = 0x35
	UndefinedLSB7  Control = 0x36
	UndefinedLSB8  Control = 0x37
	UndefinedLSB9  Control = 0x38
	UndefinedLSB10 Control = 0x39
	UndefinedLSB11 Control = 0x3A
	UndefinedLSB12 Control = 0x3B
	UndefinedLSB13 Control = 0x3C
	UndefinedLSB14 Control = 0x3D
	UndefinedLSB15 Control = 0x3E
	UndefinedLSB16 Control = 0x3F

	DamperPedalOnOff Control = 0x40 // sustain ≤63 off, ≥64 on
	PortamentoOnOff  Control = 0x41 // ≤63 off, ≥64 on
	SostenutoOnOff   Control = 0x42 // ≤63 off, ≥64 on
	SoftPedalOnOff   Control = 0x43 // ≤63 off, ≥64 on
	LegatoFootswitch Control = 0x44 // ≤63 Normal, ≥64 Legato
	Hold2            Control = 0x45 // ≤63 off, ≥64 on

	SoundController1LSB  Control = 0x46 // default: Sound Variation
	SoundController2LSB  Control = 0x47 // default: Timbre/Harmonic Intens.
	SoundController3LSB  Control = 0x48 // default: Release Time
	SoundController4LSB  Control = 0x49 // default: Attack Time
	SoundController5LSB  Control = 0x4A // default: Brightness
	SoundController6LSB  Control = 0x4B // default: Decay Time - see MMA RP-021
	SoundController7LSB  Control = 0x4C // default: Vibrato Rate - see MMA RP-021
	SoundController8LSB  Control = 0x4D // default: Vibrato Depth - see MMA RP-021
	SoundController9LSB  Control = 0x4E // default: Vibrato Delay - see MMA RP-021
	SoundController10LSB Control = 0x4F // default: undefined - see MMA RP-021

	GeneralPurposeController5LSB Control = 0x50
	GeneralPurposeController6LSB Control = 0x51
	GeneralPurposeController7LSB Control = 0x52
	GeneralPurposeController8LSB Control = 0x53

	PortamentoControlLSB Control = 0x54

	Undefined  Control = 0x55
	Undefined2 Control = 0x56
	Undefined3 Control = 0x57

	HighResolutionVelocityPrefixLSB Control = 0x58

	Undefined4 Control = 0x59
	Undefined5 Control = 0x5A

	Effects1Depth Control = 0x5B // default: Reverb Send Level - see MMA RP-023
	Effects2Depth Control = 0x5C // default: Tremolo Depth
	Effects3Depth Control = 0x5D // default: Chorus Send Level - see MMA RP-023
	Effects4Depth Control = 0x5E // default: Celeste [Detune] Depth
	Effects5Depth Control = 0x5F // default: Phaser Depth

	DataIncrement Control = 0x60 // Data Entry +1, see MMA RP-018
	DataDecrement Control = 0x61 // Data Entry -1, see MMA RP-018

	NonRegisteredParameterNumberLSB Control = 0x62
	NonRegisteredParameterNumberMSB Control = 0x63
	RegisteredParameterNumberLSB    Control = 0x64
	RegisteredParameterNumberMSB    Control = 0x65

	Undefined6  Control = 0x66
	Undefined7  Control = 0x67
	Undefined8  Control = 0x68
	Undefined9  Control = 0x69
	Undefined10 Control = 0x6A
	Undefined11 Control = 0x6B
	Undefined12 Control = 0x6C
	Undefined13 Control = 0x6D
	Undefined14 Control = 0x6E
	Undefined15 Control = 0x6F
	Undefined16 Control = 0x70
	Undefined17 Control = 0x71
	Undefined18 Control = 0x72
	Undefined19 Control = 0x73
	Undefined20 Control = 0x74
	Undefined21 Control = 0x75
	Undefined22 Control = 0x76
	Undefined23 Control = 0x77

	// Controller numbers 120-127 are reserved for Channel Mode Messages,
	// which rather than controlling sound parameters,
	// affect the channel's operating mode. (See also Table 1.)

	AllSoundOff         Control = 0x78
	ResetAllControllers Control = 0x79 // see MMA RP-015
	LocalControlOnOff   Control = 0x7A
	AllNotesOff         Control = 0x7B
	OmniModeOff         Control = 0x7C // + all notes off
	OmniModeOn          Control = 0x7D // + all notes off
	MonoModeOn          Control = 0x7E // + poly off, + all notes off
	PolyModeOn          Control = 0x7F // + mono off, + all notes off
)
