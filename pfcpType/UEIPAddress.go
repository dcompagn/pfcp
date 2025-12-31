package pfcpType

import (
	"fmt"
	"net"
)

type UEIPAddress struct {
	Ip6pl                    bool   // Bit 7 - IPv6 Prefix Length present
	Chv6                     bool   // Bit 6 - Choose IPv6
	Chv4                     bool   // Bit 5 - Choose IPv4
	Ipv6d                    bool   // Bit 4 - IPv6 Prefix Delegation
	Sd                       bool   // Bit 3 - Source/Destination
	V4                       bool   // Bit 2 - IPv4 address present
	V6                       bool   // Bit 1 - IPv6 address present
	Ipv4Address              net.IP // Octets m to (m+3)
	Ipv6Address              net.IP // Octets p to (p+15)
	Ipv6PrefixDelegationBits uint8  // Octet r
	Ipv6PrefixLength         uint8  // Octet s
}

func (u *UEIPAddress) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(u.Ip6pl)<<6 |
		btou(u.Chv6)<<5 |
		btou(u.Chv4)<<4 |
		btou(u.Ipv6d)<<3 |
		btou(u.Sd)<<2 |
		btou(u.V4)<<1 |
		btou(u.V6)
	data = append([]byte(""), tmpUint8)

	// Octet m to (m+3)
	if u.V4 {
		if u.Ipv4Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("IPv4 address shall be present if V4 is set")
		}
		data = append(data, u.Ipv4Address.To4()...)
	}

	// Octet p to (p+15)
	if u.V6 {
		if u.Ipv6Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("IPv6 address shall be present if V6 is set")
		}
		data = append(data, u.Ipv6Address.To16()...)
	}

	// Octet r - IPv6 Prefix Delegation Bits
	if (u.V6 || u.Chv6) && u.Ipv6d {
		data = append(data, u.Ipv6PrefixDelegationBits)
	}

	// Octet s - IPv6 Prefix Length
	if (u.V6 || u.Chv6) && !u.Ipv6d && u.Ip6pl {
		data = append(data, u.Ipv6PrefixLength)
	}

	return data, nil
}

func (u *UEIPAddress) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("inadequate TLV length: %d", length)
	}
	u.Ip6pl = utob(data[idx] & BitMask7)
	u.Chv6 = utob(data[idx] & BitMask6)
	u.Chv4 = utob(data[idx] & BitMask5)
	u.Ipv6d = utob(data[idx] & BitMask4)
	u.Sd = utob(data[idx] & BitMask3)
	u.V4 = utob(data[idx] & BitMask2)
	u.V6 = utob(data[idx] & BitMask1)
	idx = idx + 1

	// Octet m to (m+3)
	if u.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("inadequate TLV length: %d", length)
		}
		u.Ipv4Address = net.IP(data[idx : idx+net.IPv4len])
		idx = idx + net.IPv4len
	}

	// Octet p to (p+15)
	if u.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("inadequate TLV length: %d", length)
		}
		u.Ipv6Address = net.IP(data[idx : idx+net.IPv6len])
		idx = idx + net.IPv6len
	}

	// Octet r - IPv6 Prefix Delegation Bits
	if (u.V6 || u.Chv6) && u.Ipv6d {
		if length < idx+1 {
			return fmt.Errorf("inadequate TLV length: %d", length)
		}
		u.Ipv6PrefixDelegationBits = data[idx]
		idx = idx + 1
	}

	// Octet s - IPv6 Prefix Length
	if (u.V6 || u.Chv6) && !u.Ipv6d && u.Ip6pl {
		if length < idx+1 {
			return fmt.Errorf("inadequate TLV length: %d", length)
		}
		u.Ipv6PrefixLength = data[idx]
		idx = idx + 1
	}

	if length != idx {
		return fmt.Errorf("inadequate TLV length: %d", length)
	}

	return nil
}
