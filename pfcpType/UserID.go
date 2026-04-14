package pfcpType

import "fmt"

// UserID represents the User ID IE (type 141) as defined in TS 29.244 clause 8.2.101.
//
// Octet 5 is a flags byte; each set flag causes a length-prefixed field to follow
// in order: IMSI, IMEI, MSISDN, NAI, SUPI, GPSI, PEI.
type UserID struct {
	IMSI   []byte // present when IMSIF=1  (bit 1)
	IMEI   []byte // present when IMEIF=1  (bit 2)
	MSISDN []byte // present when MSISDNF=1 (bit 3)
	NAI    []byte // present when NAIF=1   (bit 4)
	SUPI   []byte // present when SUPIF=1  (bit 5)
	GPSI   []byte // present when GPSIF=1  (bit 6)
	PEI    []byte // present when PEIF=1   (bit 7)
}

func (u *UserID) MarshalBinary() ([]byte, error) {
	flags := btou(len(u.IMSI) > 0) |
		btou(len(u.IMEI) > 0)<<1 |
		btou(len(u.MSISDN) > 0)<<2 |
		btou(len(u.NAI) > 0)<<3 |
		btou(len(u.SUPI) > 0)<<4 |
		btou(len(u.GPSI) > 0)<<5 |
		btou(len(u.PEI) > 0)<<6

	data := []byte{flags}
	for _, field := range [][]byte{u.IMSI, u.IMEI, u.MSISDN, u.NAI, u.SUPI, u.GPSI, u.PEI} {
		if len(field) > 0 {
			data = append(data, byte(len(field)))
			data = append(data, field...)
		}
	}
	return data, nil
}

func (u *UserID) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("UserID IE too short: %d bytes", len(data))
	}
	flags := data[0]
	idx := 1

	fields := []*[]byte{&u.IMSI, &u.IMEI, &u.MSISDN, &u.NAI, &u.SUPI, &u.GPSI, &u.PEI}
	names := []string{"IMSI", "IMEI", "MSISDN", "NAI", "SUPI", "GPSI", "PEI"}
	for i, fieldPtr := range fields {
		if flags&(1<<i) == 0 {
			continue
		}
		if idx >= len(data) {
			return fmt.Errorf("UserID IE truncated at %s length byte", names[i])
		}
		fieldLen := int(data[idx])
		idx++
		if idx+fieldLen > len(data) {
			return fmt.Errorf("UserID IE truncated in %s data", names[i])
		}
		*fieldPtr = data[idx : idx+fieldLen]
		idx += fieldLen
	}
	return nil
}
