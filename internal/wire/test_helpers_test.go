package wire

import (
	"bytes"
	"encoding/binary"
	"log"
	"testing"

	"github.com/ruk1ng001/quic-go-upgrade/internal/protocol"
	"github.com/ruk1ng001/quic-go-upgrade/internal/utils"
	"github.com/ruk1ng001/quic-go-upgrade/quicvarint"
)

func encodeVarInt(i uint64) []byte {
	return quicvarint.Append(nil, i)
}

func appendVersion(data []byte, v protocol.Version) []byte {
	offset := len(data)
	data = append(data, []byte{0, 0, 0, 0}...)
	binary.BigEndian.PutUint32(data[offset:], uint32(v))
	return data
}

func setupLogTest(t *testing.T, buf *bytes.Buffer) utils.Logger {
	logger := utils.DefaultLogger
	logger.SetLogLevel(utils.LogLevelDebug)
	originalOutput := log.Writer()
	log.SetOutput(buf)
	t.Cleanup(func() { log.SetOutput(originalOutput) })
	return logger
}
