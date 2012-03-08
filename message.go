package dhcp

import (
  "crypto/rand"
  "math/big"
	"net"
)

const (
	MagicCookie = 0x63825363

	MessageTypeRequest = 1
	MessageTypeReply   = 2

	MessageHardwareTypeEthernet               = 1
	MessageHardwareTypeExperimentalEthernet   = 2
	MessageHardwareTypeAmateurRadioAX25       = 3
	MessageHardwareTypeProteonProNETTokenRing = 4
	MessageHardwareTypeChaos                  = 5
	MessageHardwareTypeIEEE802Networks        = 6
	MessageHardwareTypeARCNET                 = 7
	MessageHardwareTypeHyperchannel           = 8
	MessageHardwareTypeLanstar                = 9
	MessageHardwareTypeAutonetShortAddress    = 10
	MessageHardwareTypeLocalTalk              = 11
	MessageHardwareTypeLocalNet               = 12
)

// Message implements the DHCP message structure.
type Message struct {
	Type                  uint8
	HardwareType          uint8
	HardwareAddressLength uint8
	Hops                  uint8
	TransactionID         uint32
	SecondsElapsed        uint16
	Flags                 uint16
	ClientIPAdress        net.IP
	YourIPAddress         net.IP
	NextServerIPAddress   net.IP
	RelayIPAddress        net.IP
	ClientMAC             net.HardwareAddr
	ServerHostName        string
	File                  string
	Options               map[uint8]*Option
}

func NewMessage() (msg *Message, err error) {
  msg = new(Message)

  r, err := rand.Int(rand.Reader, big.NewInt(2^32))
  if err != nil {
    return nil, err
  }
  msg.TransactionID = uint32(r.Int64())

  return
}

func ReadMessage(b []byte) (msg *Message, err error) {
	msg = new(Message)

	if len(b) < 240 || byteToUint32(b[236:240]) != MagicCookie {
		return nil, ErrInvalidFormat
	}

	msg.Type = uint8(b[0])
	msg.HardwareType = uint8(b[1])
	msg.HardwareAddressLength = uint8(b[2])
	msg.Hops = uint8(b[3])
	msg.TransactionID = byteToUint32(b[4:8])
	msg.SecondsElapsed = byteToUint16(b[8:10])
	msg.Flags = byteToUint16(b[10:12])
	msg.ClientIPAdress = b[12:16]
	msg.YourIPAddress = b[16:20]
	msg.NextServerIPAddress = b[20:24]
	msg.RelayIPAddress = b[24:28]
	msg.ClientMAC = b[28 : 28+int(msg.HardwareAddressLength)] // full length: b[28:44])
	if b[44] != 0x00 {
		msg.ServerHostName = string(b[44:108])
	}
	if b[108] != 0x00 {
		msg.File = string(b[108:236])
	}

	msg.Options = make(map[uint8]*Option)
	c := 240
	for c < len(b) {
		option, err := ReadOption(b[c:])

		if err != nil {
			return nil, ErrInvalidFormat
		}

		if option.Code == OptionCodeEnd {
			break
		}

		if _, ok := msg.Options[option.Code]; ok && option.Code != OptionCodePad {
			return nil, ErrDuplicateField
		} else {
			msg.Options[option.Code] = option
		}

		c += int(option.Length) + 2
	}

	return msg, nil
}
