package xbee_test

import (
	"testing"

	"github.com/cswank/xbee"
	"github.com/stretchr/testify/assert"
)

func TestXbee(t *testing.T) {
	d := []byte{0x92, 0x00, 0x13, 0xA2, 0x00, 0x40, 0x4C, 0x0E, 0xBE, 0x61, 0x59, 0x01, 0x01, 0x00, 0x18, 0x03, 0x00, 0x10, 0x02, 0x2F, 0x01, 0xFE, 0x49}
	x, err := xbee.NewMessage(d)
	assert.NoError(t, err)

	adc, err := x.GetAnalog()
	assert.NoError(t, err)
	assert.InDelta(t, 655.7184, adc["adc0"], 0.01)
	assert.InDelta(t, 598.2405, adc["adc1"], 0.01)

	dio, err := x.GetDigital()
	assert.NoError(t, err)
	v, ok := dio["dio3"]
	assert.True(t, ok)
	assert.False(t, v)
	assert.True(t, dio["dio4"])
}
