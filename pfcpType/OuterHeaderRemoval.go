package pfcpType

import (
	"fmt"
)

// Outer Header Removal Description values (Table 8.2.64-1)
const (
	OuterHeaderRemovalGtpUUdpIpv4  uint8 = 0 // GTP-U/UDP/IPv4
	OuterHeaderRemovalGtpUUdpIpv6  uint8 = 1 // GTP-U/UDP/IPv6
	OuterHeaderRemovalUdpIpv4      uint8 = 2 // UDP/IPv4
	OuterHeaderRemovalUdpIpv6      uint8 = 3 // UDP/IPv6
	OuterHeaderRemovalIpv4         uint8 = 4 // IPv4
	OuterHeaderRemovalIpv6         uint8 = 5 // IPv6
	OuterHeaderRemovalGtpUUdpIp    uint8 = 6 // GTP-U/UDP/IP (regardless IPv4 or IPv6)
	OuterHeaderRemovalVlanStagPop  uint8 = 7 // VLAN TAG POP (removal of one VLAN tag, S-TAG or C-TAG)
	OuterHeaderRemovalVlanStagCtag uint8 = 8 // VLAN TAGs POP-POP (removal of 2 VLAN tags, S-TAG and C-TAG)
)

// GTP-U Extension Header Deletion flags (Table 8.2.64-2)
// Bitmask indicating which GTP-U extension headers should be deleted
const (
	GtpuExtHeaderDeletionPduSessionContainer uint8 = 1 << 0 // Bit 1: PDU Session Container
)

type OuterHeaderRemoval struct {
	OuterHeaderRemovalDescription uint8
	// GTP-U Extension Header Deletion field (octet 6) - optional
	// Present only when GTP-U extension headers need to be deleted from incoming GTP-PDUs
	GtpuExtensionHeaderDeletion *uint8
}

func (o *OuterHeaderRemoval) MarshalBinary() (data []byte, err error) {
	// Octet 5: Outer Header Removal Description
	data = append([]byte(""), o.OuterHeaderRemovalDescription)

	// Octet 6: GTP-U Extension Header Deletion (optional)
	if o.GtpuExtensionHeaderDeletion != nil {
		data = append(data, *o.GtpuExtensionHeaderDeletion)
	}
	return data, nil
}

func (o *OuterHeaderRemoval) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	if length < 1 {
		return fmt.Errorf("inadequate TLV length: %d", length)
	}

	// Octet 5: Outer Header Removal Description
	o.OuterHeaderRemovalDescription = data[0]

	// Octet 6: GTP-U Extension Header Deletion (optional)
	// Present only if explicitly specified to delete GTP-U extension header(s)
	if length >= 2 {
		o.GtpuExtensionHeaderDeletion = new(uint8)
		*o.GtpuExtensionHeaderDeletion = data[1]
	}

	return nil
}
