package wire

import (
	"testing"

	"github.com/ruk1ng001/quic-go-upgrade/internal/protocol"
	"github.com/stretchr/testify/require"
)

func TestWriteHandshakeDoneSampleFrame(t *testing.T) {
	frame := HandshakeDoneFrame{}
	b, err := frame.Append(nil, protocol.Version1)
	require.NoError(t, err)
	require.Equal(t, []byte{handshakeDoneFrameType}, b)
	require.Equal(t, protocol.ByteCount(1), frame.Length(protocol.Version1))
}
