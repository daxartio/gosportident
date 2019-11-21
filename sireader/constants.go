package sireader

//CardType is enum for type of cards
type CardType int

//Variables of CardType
const (
	SI5 = iota + 1
	SI6
	SI8
	SI9
	SI10
	SItCard
	SIpCard
)

//Weekday encoding (only for reference, currently unused)
const (
	DSunday    = 0x000
	DMonday    = 0x001
	DTuesday   = 0x002
	DWednesday = 0x003
	DThursday  = 0x004
	DFriday    = 0x005
	DSaturday  = 0x006
	DUnknown   = 0x007
)

//RecLen Only in extended protocol, otherwise 6!
const RecLen = 8

//EE General card data structure value
const EE = 0xEE

//SI Card data structures
const (
	CardSi5Cn2 = 6
	CardSi5Cn1 = 4
	CardSi5Cn0 = 5
	CardSi5Std = 0
	CardSi5St  = 19
	CardSi5Ftd = 0
	CardSi5Ft  = 21
	CardSi5Ctd = 21
	CardSi5Ct  = 25
	CardSi5Ltd = 0
	CardSi5Lt  = 0
	CardSi5Rc  = 23
	CardSi5P1  = 32
	CardSi5Pl  = 3
	CardSi5Pm  = 30
	CardSi5Cn  = 0
	CardSi5Ptd = 0
	CardSi5Pth = 1
	CardSi5Ptl = 2
)

//SI Card data structures
const (
	CardSi6Cn2 = 11
	CardSi6Cn1 = 12
	CardSi6Cn0 = 13
	CardSi6Std = 24
	CardSi6St  = 26
	CardSi6Ftd = 20
	CardSi6Ft  = 22
	CardSi6Ctd = 28
	CardSi6Ct  = 30
	CardSi6Ltd = 32
	CardSi6Lt  = 34
	CardSi6Rc  = 18
	CardSi6P1  = 128
	CardSi6Pl  = 4
	CardSi6Pm  = 64
	CardSi6Ptd = 0
	CardSi6Cn  = 1
	CardSi6Pth = 2
	CardSi6Ptl = 3
)

//SI Card data structures
const (
	CardSi8Cn2 = 25
	CardSi8Cn1 = 26
	CardSi8Cn0 = 27
	CardSi8Std = 12
	CardSi8St  = 14
	CardSi8Ftd = 16
	CardSi8Ft  = 18
	CardSi8Ctd = 8
	CardSi8Ct  = 10
	CardSi8Ltd = 0
	CardSi8Lt  = 0
	CardSi8Rc  = 22
	CardSi8P1  = 136
	CardSi8Pl  = 4
	CardSi8Pm  = 50
	CardSi8Ptd = 0
	CardSi8Cn  = 1
	CardSi8Pth = 2
	CardSi8Ptl = 3
	CardSi8Bc  = 2
)

//SI Card data structures
const (
	CardSi10Cn2 = 25
	CardSi10Cn1 = 26
	CardSi10Cn0 = 27
	CardSi10Std = 12
	CardSi10St  = 14
	CardSi10Ftd = 16
	CardSi10Ft  = 18
	CardSi10Ctd = 8
	CardSi10Ct  = 10
	CardSi10Ltd = 0
	CardSi10Lt  = 0
	CardSi10Rc  = 22
	CardSi10P1  = 128
	CardSi10Pl  = 4
	CardSi10Pm  = 64
	CardSi10Ptd = 0
	CardSi10Cn  = 1
	CardSi10Pth = 2
	CardSi10Ptl = 3
	CardSi10Bc  = 8
)

//Punch trigger in control mode data structure
const (
	TOffset = 8
	TCn     = 0
	TTime   = 5
)

//Backup memory in control mode
const (
	BcCn   = 3
	BcTime = 8
)

//Crc
const (
	CrcPolynom = 0x8005
	CrcBitf    = 0x8000
)

//Protocol characters
const (
	STX    byte = 0x02
	ETX    byte = 0x03
	ACK    byte = 0x06
	NAK    byte = 0x15
	DLE    byte = 0x10
	WAKEUP byte = 0xFF
)

//Basic protocol commands, currently unused
const (
	BcSetCardNo    byte = 0x30
	BcGetSi5       byte = 0x31
	BcTransRec     byte = 0x33
	BcSi5Write     byte = 0x43
	BcSi5Det       byte = 0x46
	BcTransRec2    byte = 0x53
	BcTransTime    byte = 0x54
	BcGetSi6       byte = 0x61
	BcSi6WritePage byte = 0x62
	BcSi6ReadWord  byte = 0x63
	BcSi6WriteWord byte = 0x64
	BcSi6Det       byte = 0x66
	BcSetMs        byte = 0x70
	BcGetMs        byte = 0x71
	BcSetSysVal    byte = 0x72
	BcGetSysVal    byte = 0x73
	BcGetBackup    byte = 0x74
	BcEraseBackup  byte = 0x75
	BcSetTime      byte = 0x76
	BcGetTime      byte = 0x77
	BcOff          byte = 0x78
	BcReset        byte = 0x79
	BcGetBackup2   byte = 0x7A
	BcSetBaud      byte = 0x7E
)

//Extended protocol commands
const (
	CGetBackup byte = 0x81
	CSetSysVal byte = 0x82
	CGetSysVal byte = 0x83
	CSrrWrite  byte = 0xA2
	CSrrRead   byte = 0xA3
	CSrrQuery  byte = 0xA6
	CSrrPing   byte = 0xA7
	CSrrAdhoc  byte = 0xA8
	CGetSi5    byte = 0xB1
	CSi5Write  byte = 0xC3
	CTransRec  byte = 0xD3
	CClearCard byte = 0xE0
	CGetSi6    byte = 0xE1
	CSi5Det    byte = 0xE5
	CSi6Det    byte = 0xE6
	//SI-card removed
	CSiRem byte = 0xE7
	//SI-card 8/9/10/11/p/t inserted
	CSi9Det      byte = 0xE8
	CSi9Write    byte = 0xEA
	CGetSi9      byte = 0xEF
	CSetMs       byte = 0xF0
	CGetMs       byte = 0xF1
	CEraseBackup byte = 0xF5
	CSetTime     byte = 0xF6
	CGetTime     byte = 0xF7
	COff         byte = 0xF8
	CBeep        byte = 0xF9
	CSetBaud     byte = 0xFE
)

//Protocol Parameters
const (
	PMsDirect   byte = 0x4D
	PMsIndirect byte = 0x53
	PSi6Cb      byte = 0x08
)

//Offsets in system data
const (
	OOldSerial byte = 0x00
	OOldCPUID  byte = 0x02
	OSerialNo  byte = 0x00

	OSrrCfg byte = 0x04

	OFirmware  byte = 0x05
	OBuildDate byte = 0x08
	OModelID   byte = 0x0B
)

//Service
const (
	OMemSize   byte = 0x0D
	OBatDate   byte = 0x15
	OBatCap    byte = 0x19
	OBackupPtr byte = 0x1C
	OSi6Cb     byte = 0x33

	OSrrChannel  byte = 0x34
	OMemOverflow byte = 0x3D
	OProgram     byte = 0x70
	OMode        byte = 0x71
	OStationCode byte = 0x72
	OFeedback    byte = 0x73

	OProto byte = 0x74

	OWakeupDate byte = 0x75
	OWakeupTime byte = 0x78
	OSleepTime  byte = 0x7B
)

//SI station modes
const (
	MSIACSpecial = 0x01
	MControl     = 0x02
	MStart       = 0x03
	MFinish      = 0x04
	MReadout     = 0x05
	MClearOld    = 0x06
	MClear       = 0x07
	MCheck       = 0x0A
	MPrintout    = 0x0B
	MStartTrig   = 0x0C
	MFinishTrig  = 0x0D
	MBcControl   = 0x12
	MBcStart     = 0x13
	MBcFinish    = 0x14
	MBcReadout   = 0x15
)

var (
	//TimeReset General card data structure values
	TimeReset = []byte{EE, EE}
	//BeepTwice = []byte{STX, CBeep, 0x01, 0x02, 0x14, 0x0A, ETX}
)
