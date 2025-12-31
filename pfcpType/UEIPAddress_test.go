package pfcpType

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUEIPAddress(t *testing.T) {
	testData := UEIPAddress{
		Ip6pl:                    false,
		Chv6:                     false,
		Chv4:                     false,
		Ipv6d:                    true,
		Sd:                       true,
		V4:                       true,
		V6:                       true,
		Ipv4Address:              net.ParseIP("12.34.56.78").To4(),
		Ipv6Address:              net.ParseIP("2001:db8::68").To16(),
		Ipv6PrefixDelegationBits: 4,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{15, 12, 34, 56, 78, 32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 104, 4}, buf)
}

func TestUnmarshalUEIPAddress(t *testing.T) {
	buf := []byte{15, 12, 34, 56, 78, 32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 104, 4}
	var testData UEIPAddress
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := UEIPAddress{
		Ip6pl:                    false,
		Chv6:                     false,
		Chv4:                     false,
		Ipv6d:                    true,
		Sd:                       true,
		V4:                       true,
		V6:                       true,
		Ipv4Address:              net.ParseIP("12.34.56.78").To4(),
		Ipv6Address:              net.ParseIP("2001:db8::68").To16(),
		Ipv6PrefixDelegationBits: 4,
	}
	assert.Equal(t, expectData, testData)
}

func TestMarshalUEIPAddressWithCHV6(t *testing.T) {
	testData := UEIPAddress{
		Ip6pl:                    false,
		Chv6:                     true,
		Chv4:                     false,
		Ipv6d:                    true,
		Sd:                       false,
		V4:                       false,
		V6:                       false,
		Ipv6PrefixDelegationBits: 8,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	// Octet 5: CHV6=1(bit6), Ipv6d=1(bit4) => 0b00101000 = 40
	assert.Equal(t, []byte{40, 8}, buf)
}

func TestUnmarshalUEIPAddressWithCHV6(t *testing.T) {
	buf := []byte{40, 8}
	var testData UEIPAddress
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := UEIPAddress{
		Ip6pl:                    false,
		Chv6:                     true,
		Chv4:                     false,
		Ipv6d:                    true,
		Sd:                       false,
		V4:                       false,
		V6:                       false,
		Ipv6PrefixDelegationBits: 8,
	}
	assert.Equal(t, expectData, testData)
}

func TestMarshalUEIPAddressWithIP6PL(t *testing.T) {
	testData := UEIPAddress{
		Ip6pl:            true,
		Chv6:             false,
		Chv4:             false,
		Ipv6d:            false,
		Sd:               false,
		V4:               false,
		V6:               true,
		Ipv6Address:      net.ParseIP("2001:db8::1").To16(),
		Ipv6PrefixLength: 56,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	// Octet 5: IP6PL=1(bit7), V6=1(bit1) => 0b01000001 = 65
	assert.Equal(t, []byte{65, 32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 56}, buf)
}

func TestUnmarshalUEIPAddressWithIP6PL(t *testing.T) {
	buf := []byte{65, 32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 56}
	var testData UEIPAddress
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := UEIPAddress{
		Ip6pl:            true,
		Chv6:             false,
		Chv4:             false,
		Ipv6d:            false,
		Sd:               false,
		V4:               false,
		V6:               true,
		Ipv6Address:      net.ParseIP("2001:db8::1").To16(),
		Ipv6PrefixLength: 56,
	}
	assert.Equal(t, expectData, testData)
}

func TestMarshalUEIPAddressWithCHV4(t *testing.T) {
	testData := UEIPAddress{
		Ip6pl: false,
		Chv6:  false,
		Chv4:  true,
		Ipv6d: false,
		Sd:    false,
		V4:    false,
		V6:    false,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	// Octet 5: CHV4=1(bit5) => 0b00010000 = 16
	assert.Equal(t, []byte{16}, buf)
}

func TestUnmarshalUEIPAddressWithCHV4(t *testing.T) {
	buf := []byte{16}
	var testData UEIPAddress
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := UEIPAddress{
		Ip6pl: false,
		Chv6:  false,
		Chv4:  true,
		Ipv6d: false,
		Sd:    false,
		V4:    false,
		V6:    false,
	}
	assert.Equal(t, expectData, testData)
}
