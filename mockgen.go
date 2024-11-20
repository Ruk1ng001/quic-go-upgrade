//go:build gomock || generate

package quic

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_send_conn_test.go github.com/ruk1ng001/quic-go-upgrade SendConn"
type SendConn = sendConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_raw_conn_test.go github.com/ruk1ng001/quic-go-upgrade RawConn"
type RawConn = rawConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_sender_test.go github.com/ruk1ng001/quic-go-upgrade Sender"
type Sender = sender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_stream_internal_test.go github.com/ruk1ng001/quic-go-upgrade StreamI"
type StreamI = streamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_receive_stream_internal_test.go github.com/ruk1ng001/quic-go-upgrade ReceiveStreamI"
type ReceiveStreamI = receiveStreamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_send_stream_internal_test.go github.com/ruk1ng001/quic-go-upgrade SendStreamI"
type SendStreamI = sendStreamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_stream_sender_test.go github.com/ruk1ng001/quic-go-upgrade StreamSender"
type StreamSender = streamSender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_stream_control_frame_getter_test.go github.com/ruk1ng001/quic-go-upgrade StreamControlFrameGetter"
type StreamControlFrameGetter = streamControlFrameGetter

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_frame_source_test.go github.com/ruk1ng001/quic-go-upgrade FrameSource"
type FrameSource = frameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_ack_frame_source_test.go github.com/ruk1ng001/quic-go-upgrade AckFrameSource"
type AckFrameSource = ackFrameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_stream_manager_test.go github.com/ruk1ng001/quic-go-upgrade StreamManager"
type StreamManager = streamManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_sealing_manager_test.go github.com/ruk1ng001/quic-go-upgrade SealingManager"
type SealingManager = sealingManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_unpacker_test.go github.com/ruk1ng001/quic-go-upgrade Unpacker"
type Unpacker = unpacker

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_packer_test.go github.com/ruk1ng001/quic-go-upgrade Packer"
type Packer = packer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_mtu_discoverer_test.go github.com/ruk1ng001/quic-go-upgrade MTUDiscoverer"
type MTUDiscoverer = mtuDiscoverer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_conn_runner_test.go github.com/ruk1ng001/quic-go-upgrade ConnRunner"
type ConnRunner = connRunner

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_quic_conn_test.go github.com/ruk1ng001/quic-go-upgrade QUICConn"
type QUICConn = quicConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_packet_handler_test.go github.com/ruk1ng001/quic-go-upgrade PacketHandler"
type PacketHandler = packetHandler

//go:generate sh -c "go run go.uber.org/mock/mockgen -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_packet_handler_manager_test.go github.com/ruk1ng001/quic-go-upgrade PacketHandlerManager"

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_packet_handler_manager_test.go github.com/ruk1ng001/quic-go-upgrade PacketHandlerManager"
type PacketHandlerManager = packetHandlerManager

// Need to use source mode for the batchConn, since reflect mode follows type aliases.
// See https://github.com/golang/mock/issues/244 for details.
//
//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -source sys_conn_oob.go -destination mock_batch_conn_test.go -mock_names batchConn=MockBatchConn"

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package quic -self_package github.com/ruk1ng001/quic-go-upgrade -self_package github.com/ruk1ng001/quic-go-upgrade -destination mock_packetconn_test.go net PacketConn"
