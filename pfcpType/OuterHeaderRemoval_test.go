package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalOuterHeaderRemoval(t *testing.T) {
	testData := OuterHeaderRemoval{
		OuterHeaderRemovalDescription: OuterHeaderRemovalGtpUUdpIpv4,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{OuterHeaderRemovalGtpUUdpIpv4}, buf)
}

func TestMarshalOuterHeaderRemovalWithExtHeaderDeletion(t *testing.T) {
	extHeaderDeletion := GtpuExtHeaderDeletionPduSessionContainer
	testData := OuterHeaderRemoval{
		OuterHeaderRemovalDescription: OuterHeaderRemovalGtpUUdpIpv4,
		GtpuExtensionHeaderDeletion:   &extHeaderDeletion,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{OuterHeaderRemovalGtpUUdpIpv4, GtpuExtHeaderDeletionPduSessionContainer}, buf)
}

func TestUnmarshalOuterHeaderRemoval(t *testing.T) {
	buf := []byte{OuterHeaderRemovalGtpUUdpIpv4}
	var testData OuterHeaderRemoval
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := OuterHeaderRemoval{
		OuterHeaderRemovalDescription: OuterHeaderRemovalGtpUUdpIpv4,
	}
	assert.Equal(t, expectData, testData)
}

func TestUnmarshalOuterHeaderRemovalWithExtHeaderDeletion(t *testing.T) {
	buf := []byte{OuterHeaderRemovalGtpUUdpIpv6, GtpuExtHeaderDeletionPduSessionContainer}
	var testData OuterHeaderRemoval
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	extHeaderDeletion := GtpuExtHeaderDeletionPduSessionContainer
	expectData := OuterHeaderRemoval{
		OuterHeaderRemovalDescription: OuterHeaderRemovalGtpUUdpIpv6,
		GtpuExtensionHeaderDeletion:   &extHeaderDeletion,
	}
	assert.Equal(t, expectData, testData)
}

func TestOuterHeaderRemovalDescriptionValues(t *testing.T) {
	// Verify constants match specification Table 8.2.64-1
	assert.Equal(t, uint8(0), OuterHeaderRemovalGtpUUdpIpv4)
	assert.Equal(t, uint8(1), OuterHeaderRemovalGtpUUdpIpv6)
	assert.Equal(t, uint8(2), OuterHeaderRemovalUdpIpv4)
	assert.Equal(t, uint8(3), OuterHeaderRemovalUdpIpv6)
	assert.Equal(t, uint8(4), OuterHeaderRemovalIpv4)
	assert.Equal(t, uint8(5), OuterHeaderRemovalIpv6)
	assert.Equal(t, uint8(6), OuterHeaderRemovalGtpUUdpIp)
	assert.Equal(t, uint8(7), OuterHeaderRemovalVlanStagPop)
	assert.Equal(t, uint8(8), OuterHeaderRemovalVlanStagCtag)
}
