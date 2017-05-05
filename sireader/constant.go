package sireader

const (
	D_SUNDAY    = 0x000
	D_MONDAY    = 0x001
	D_TUESDAY   = 0x002
	D_WEDNESDAY = 0x003
	D_THURSDAY  = 0x004
	D_FRIDAY    = 0x005
	D_SATURDAY  = 0x006
	D_UNKNOWN   = 0x007
)

const REC_LEN = 8

var TIME_RESET = []byte{0xEE, 0xEE}

const (
	CARD_SI5_CN2 = 6
	CARD_SI5_CN1 = 4
	CARD_SI5_CN0 = 5
	CARD_SI5_STD = 0
	CARD_SI5_ST  = 19
	CARD_SI5_FTD = 0
	CARD_SI5_FT  = 21
	CARD_SI5_CTD = 21
	CARD_SI5_CT  = 25
	CARD_SI5_LTD = 0
	CARD_SI5_LT  = 0
	CARD_SI5_RC  = 23
	CARD_SI5_P1  = 32
	CARD_SI5_PL  = 3
	CARD_SI5_PM  = 30
	CARD_SI5_CN  = 0
	CARD_SI5_PTD = 0
	CARD_SI5_PTH = 1
	CARD_SI5_PTL = 2
)

const (
	CARD_SI6_CN2 = 11
	CARD_SI6_CN1 = 12
	CARD_SI6_CN0 = 13
	CARD_SI6_STD = 24
	CARD_SI6_ST  = 26
	CARD_SI6_FTD = 20
	CARD_SI6_FT  = 22
	CARD_SI6_CTD = 28
	CARD_SI6_CT  = 30
	CARD_SI6_LTD = 32
	CARD_SI6_LT  = 34
	CARD_SI6_RC  = 18
	CARD_SI6_P1  = 128
	CARD_SI6_PL  = 4
	CARD_SI6_PM  = 64
	CARD_SI6_PTD = 0
	CARD_SI6_CN  = 1
	CARD_SI6_PTH = 2
	CARD_SI6_PTL = 3
)

const (
	CARD_SI8_CN2 = 25
	CARD_SI8_CN1 = 26
	CARD_SI8_CN0 = 27
	CARD_SI8_STD = 12
	CARD_SI8_ST  = 14
	CARD_SI8_FTD = 16
	CARD_SI8_FT  = 18
	CARD_SI8_CTD = 8
	CARD_SI8_CT  = 10
	CARD_SI8_LTD = 0
	CARD_SI8_LT  = 0
	CARD_SI8_RC  = 22
	CARD_SI8_P1  = 136
	CARD_SI8_PL  = 4
	CARD_SI8_PM  = 50
	CARD_SI8_PTD = 0
	CARD_SI8_CN  = 1
	CARD_SI8_PTH = 2
	CARD_SI8_PTL = 3
	CARD_SI8_BC  = 2
)

const (
	CARD_SI10_CN2 = 25
	CARD_SI10_CN1 = 26
	CARD_SI10_CN0 = 27
	CARD_SI10_STD = 12
	CARD_SI10_ST  = 14
	CARD_SI10_FTD = 16
	CARD_SI10_FT  = 18
	CARD_SI10_CTD = 8
	CARD_SI10_CT  = 10
	CARD_SI10_LTD = 0
	CARD_SI10_LT  = 0
	CARD_SI10_RC  = 22
	CARD_SI10_P1  = 128
	CARD_SI10_PL  = 4
	CARD_SI10_PM  = 64
	CARD_SI10_PTD = 0
	CARD_SI10_CN  = 1
	CARD_SI10_PTH = 2
	CARD_SI10_PTL = 3
	CARD_SI10_BC  = 8
)

const (
	T_OFFSET = 8
	T_CN     = 0
	T_TIME   = 5
)

const (
	BC_CN   = 3
	BC_TIME = 8
)

const (
	CRC_POLYNOM = 0x8005
	CRC_BITF    = 0x8000
)

var (
	STX    = []byte{0x02}
	ETX    = []byte{0x03}
	ACK    = []byte{0x06}
	NAK    = []byte{0x15}
	DLE    = []byte{0x10}
	WAKEUP = []byte{0xFF}
)

var (
	BC_SET_CARDNO    = []byte{0x30}
	BC_GET_SI5       = []byte{0x31}
	BC_TRANS_REC     = []byte{0x33}
	BC_SI5_WRITE     = []byte{0x43}
	BC_SI5_DET       = []byte{0x46}
	BC_TRANS_REC2    = []byte{0x53}
	BC_TRANS_TIME    = []byte{0x54}
	BC_GET_SI6       = []byte{0x61}
	BC_SI6_WRITEPAGE = []byte{0x62}
	BC_SI6_READWORD  = []byte{0x63}
	BC_SI6_WRITEWORD = []byte{0x64}
	BC_SI6_DET       = []byte{0x66}
	BC_SET_MS        = []byte{0x70}
	BC_GET_MS        = []byte{0x71}
	BC_SET_SYS_VAL   = []byte{0x72}
	BC_GET_SYS_VAL   = []byte{0x73}
	BC_GET_BACKUP    = []byte{0x74}
	BC_ERASE_BACKUP  = []byte{0x75}
	BC_SET_TIME      = []byte{0x76}
	BC_GET_TIME      = []byte{0x77}
	BC_OFF           = []byte{0x78}
	BC_RESET         = []byte{0x79}
	BC_GET_BACKUP2   = []byte{0x7A}
	BC_SET_BAUD      = []byte{0x7E}
)

var (
	C_GET_BACKUP   = []byte{0x81}
	C_SET_SYS_VAL  = []byte{0x82}
	C_GET_SYS_VAL  = []byte{0x83}
	C_SRR_WRITE    = []byte{0xA2}
	C_SRR_READ     = []byte{0xA3}
	C_SRR_QUERY    = []byte{0xA6}
	C_SRR_PING     = []byte{0xA7}
	C_SRR_ADHOC    = []byte{0xA8}
	C_GET_SI5      = []byte{0xB1}
	C_SI5_WRITE    = []byte{0xC3}
	C_TRANS_REC    = []byte{0xD3}
	C_CLEAR_CARD   = []byte{0xE0}
	C_GET_SI6      = []byte{0xE1}
	C_SI5_DET      = []byte{0xE5}
	C_SI6_DET      = []byte{0xE6}
	C_SI_REM       = []byte{0xE7}
	C_SI9_DET      = []byte{0xE8}
	C_SI9_WRITE    = []byte{0xEA}
	C_GET_SI9      = []byte{0xEF}
	C_SET_MS       = []byte{0xF0}
	C_GET_MS       = []byte{0xF1}
	C_ERASE_BACKUP = []byte{0xF5}
	C_SET_TIME     = []byte{0xF6}
	C_GET_TIME     = []byte{0xF7}
	C_OFF          = []byte{0xF8}
	C_BEEP         = []byte{0xF9}
	C_SET_BAUD     = []byte{0xFE}
)

var (
	P_MS_DIRECT   = []byte{0x4D}
	P_MS_INDIRECT = []byte{0x53}
	P_SI6_CB      = []byte{0x08}
)

var (
	O_OLD_SERIAL = []byte{0x00}
	O_OLD_CPU_ID = []byte{0x02}
	O_SERIAL_NO  = []byte{0x00}

	O_SRR_CFG = []byte{0x04}

	O_FIRMWARE   = []byte{0x05}
	O_BUILD_DATE = []byte{0x08}
	O_MODEL_ID   = []byte{0x0B}
)

var (
	O_MEM_SIZE   = []byte{0x0D}
	O_BAT_DATE   = []byte{0x15}
	O_BAT_CAP    = []byte{0x19}
	O_BACKUP_PTR = []byte{0x1C}
	O_SI6_CB     = []byte{0x33}

	O_SRR_CHANNEL  = []byte{0x34}
	O_MEM_OVERFLOW = []byte{0x3D}
	O_PROGRAM      = []byte{0x70}
	O_MODE         = []byte{0x71}
	O_STATION_CODE = []byte{0x72}
	O_FEEDBACK     = []byte{0x73}

	O_PROTO = []byte{0x74}

	O_WAKEUP_DATE = []byte{0x75}
	O_WAKEUP_TIME = []byte{0x78}
	O_SLEEP_TIME  = []byte{0x7B}
)

const (
	M_SIAC_SPECIAL = 0x01
	M_CONTROL      = 0x02
	M_START        = 0x03
	M_FINISH       = 0x04
	M_READOUT      = 0x05
	M_CLEAR_OLD    = 0x06
	M_CLEAR        = 0x07
	M_CHECK        = 0x0A
	M_PRINTOUT     = 0x0B
	M_START_TRIG   = 0x0C
	M_FINISH_TRIG  = 0x0D
	M_BC_CONTROL   = 0x12
	M_BC_START     = 0x13
	M_BC_FINISH    = 0x14
	M_BC_READOUT   = 0x15
)
