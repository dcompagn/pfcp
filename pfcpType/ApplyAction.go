package pfcpType

import "fmt"

// ApplyAction represents the Apply Action IE (TS 29.244 §8.2.26).
// Octet 5 is mandatory (Rel-15 base bits + Rel-16 IPMA/IPMD/DFRT).
// Octet 6 is optional and was added in Rel-16 (EDRT/BDPN/DDPN/FSSM/MBSU).
// Additional octets are silently ignored for forward compatibility.
type ApplyAction struct {
	// Octet 5
	Drop bool // Bit 1 — Drop packets
	Forw bool // Bit 2 — Forward packets
	Buff bool // Bit 3 — Buffer packets
	Nocp bool // Bit 4 — Notify CP function about first buffered DL packet
	Dupl bool // Bit 5 — Duplicate packets
	Ipma bool // Bit 6 — IP Multicast Accept (IPTV)
	Ipmd bool // Bit 7 — IP Multicast Deny (IPTV)
	Dfrt bool // Bit 8 — Duplicate for Redundant Transmission (URLLC)
	// Octet 6 (Rel-16, present only when explicitly specified)
	Edrt bool // Bit 1 — Eliminate Duplicate Packets for Redundant Transmission
	Bdpn bool // Bit 2 — Buffered Downlink Packet Notification (DDDS)
	Ddpn bool // Bit 3 — Discarded Downlink Packet Notification (DDDS)
	Fssm bool // Bit 4 — Forward packets to lower layer SSM (5MBS)
	Mbsu bool // Bit 5 — Forward and replicate MBS data using Unicast transport
}

func (a *ApplyAction) MarshalBinary() (data []byte, err error) {
	// Octet 5
	octet5 := btou(a.Dfrt)<<7 |
		btou(a.Ipmd)<<6 |
		btou(a.Ipma)<<5 |
		btou(a.Dupl)<<4 |
		btou(a.Nocp)<<3 |
		btou(a.Buff)<<2 |
		btou(a.Forw)<<1 |
		btou(a.Drop)
	data = []byte{octet5}

	// Octet 6 — only emitted when at least one Rel-16 bit is set
	if a.Mbsu || a.Fssm || a.Ddpn || a.Bdpn || a.Edrt {
		octet6 := btou(a.Mbsu)<<4 |
			btou(a.Fssm)<<3 |
			btou(a.Ddpn)<<2 |
			btou(a.Bdpn)<<1 |
			btou(a.Edrt)
		data = append(data, octet6)
	}

	return data, nil
}

func (a *ApplyAction) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("inadequate TLV length: %d", len(data))
	}

	// Octet 5
	a.Dfrt = utob(data[0] & BitMask8)
	a.Ipmd = utob(data[0] & BitMask7)
	a.Ipma = utob(data[0] & BitMask6)
	a.Dupl = utob(data[0] & BitMask5)
	a.Nocp = utob(data[0] & BitMask4)
	a.Buff = utob(data[0] & BitMask3)
	a.Forw = utob(data[0] & BitMask2)
	a.Drop = utob(data[0] & BitMask1)

	// Octet 6 (optional Rel-16 extension)
	if len(data) >= 2 {
		a.Mbsu = utob(data[1] & BitMask5)
		a.Fssm = utob(data[1] & BitMask4)
		a.Ddpn = utob(data[1] & BitMask3)
		a.Bdpn = utob(data[1] & BitMask2)
		a.Edrt = utob(data[1] & BitMask1)
	}

	return nil
}
