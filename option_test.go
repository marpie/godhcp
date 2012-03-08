package dhcp

import (
	"testing"
)

var (
	testDataOption01 = []byte{0x35, 0x01, 0x07}
	testDataOption02 = []byte{0x36, 0x04, 0xc0, 0xa8, 0x00, 0x50}
	testDataOption03 = []byte{0x3c, 0x33, 0x64, 0x68, 0x63, 0x70, 0x63, 0x64, 0x2d, 0x35, 0x2e, 0x35, 0x2e, 0x34, 0x3a, 0x4c, 0x69, 0x6e, 0x75, 0x78, 0x2d, 0x33, 0x2e, 0x32, 0x2e, 0x38, 0x2d, 0x31, 0x2d, 0x41, 0x52, 0x43, 0x48, 0x3a, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x3a, 0x47, 0x65, 0x6e, 0x75, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c}
)

func TestReadOption01(t *testing.T) {
	opt, err := ReadOption(testDataOption01)
	if err != nil {
		t.Fatal(err)
	}

	if opt.Code != OptionCodeDHCPMessageType {
		t.Fatalf("Should be DHCP Message Type (53) but got %q", opt.Code)
	}

	if opt.Length != 1 {
		t.Fatalf("Length should be 1 but got %d", opt.Length)
	}

	if opt.Value[0] != 0x07 {
		t.Fatalf("Value should be 0x07 but got %q", opt.Value[0])
	}
}

func TestReadOption02(t *testing.T) {
	opt, err := ReadOption(testDataOption02)
	if err != nil {
		t.Fatal(err)
	}

	if opt.Code != OptionCodeDHCPServerIdentifier {
		t.Fatalf("Should be DHCP Server Identifier but got %q", opt.Code)
	}

	if opt.Length != 4 {
		t.Fatalf("Length should be 1 but got %d", opt.Length)
	}

	if opt.Value[0] != 0xC0 {
		t.Fatalf("First value should be 0xC0 but got %q", opt.Value[0])
	}

	if opt.Value[3] != 0x50 {
		t.Fatalf("Last value should be 0x50 but got %q", opt.Value[3])
	}
}

func TestReadOption03(t *testing.T) {
	opt, err := ReadOption(testDataOption03)
	if err != nil {
		t.Fatal(err)
	}

	if opt.Code != OptionCodeDHCPVendorclassidentifier {
		t.Fatalf("Should be DHCP Vendor Class Identifier but got %q", opt.Code)
	}

	if opt.Length != 51 {
		t.Fatalf("Length should be 51 but got %d", opt.Length)
	}

	if string(opt.Value) != "dhcpcd-5.5.4:Linux-3.2.8-1-ARCH:x86_64:GenuineIntel" {
		t.Fatalf("Value should be 'dhcpcd-5.5.4:Linux-3.2.8-1-ARCH:x86_64:GenuineIntel' but got %q", string(opt.Value))
	}
}
