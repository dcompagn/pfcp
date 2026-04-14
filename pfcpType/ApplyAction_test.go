package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalApplyAction(t *testing.T) {
	testData := ApplyAction{
		Dupl: true,
		Nocp: false,
		Buff: true,
		Forw: false,
		Drop: true,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{21}, buf)
}

func TestUnmarshalApplyAction(t *testing.T) {
	buf := []byte{21}
	var testData ApplyAction
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := ApplyAction{
		Dupl: true,
		Nocp: false,
		Buff: true,
		Forw: false,
		Drop: true,
	}
	assert.Equal(t, expectData, testData)
}

// TestApplyActionRel16 verifies that a 2-octet Apply Action (Rel-16) is decoded
// without error and that the Rel-16 bits are correctly extracted.
func TestApplyActionRel16(t *testing.T) {
	// Octet 5: FORW (bit 2) = 0x02
	// Octet 6: DDPN (bit 3) = 0x04, BDPN (bit 2) = 0x02  →  0x06
	buf := []byte{0x02, 0x06}

	var got ApplyAction
	err := got.UnmarshalBinary(buf)
	assert.Nil(t, err)
	assert.True(t, got.Forw)
	assert.True(t, got.Ddpn)
	assert.True(t, got.Bdpn)
	// All other fields must be false
	assert.False(t, got.Drop)
	assert.False(t, got.Buff)
	assert.False(t, got.Nocp)
	assert.False(t, got.Dupl)
	assert.False(t, got.Edrt)
}

// TestApplyActionRel16RoundTrip verifies marshal→unmarshal identity for a Rel-16 value.
func TestApplyActionRel16RoundTrip(t *testing.T) {
	orig := ApplyAction{Forw: true, Ddpn: true, Bdpn: true}
	buf, err := orig.MarshalBinary()
	assert.Nil(t, err)
	assert.Len(t, buf, 2)

	var got ApplyAction
	err = got.UnmarshalBinary(buf)
	assert.Nil(t, err)
	assert.Equal(t, orig, got)
}
