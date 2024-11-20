#!/bin/bash -eu

export CXX="${CXX} -lresolv" # required by Go 1.20

compile_go_fuzzer github.com/ruk1ng001/quic-go-upgrade/fuzzing/frames Fuzz frame_fuzzer
compile_go_fuzzer github.com/ruk1ng001/quic-go-upgrade/fuzzing/header Fuzz header_fuzzer
compile_go_fuzzer github.com/ruk1ng001/quic-go-upgrade/fuzzing/transportparameters Fuzz transportparameter_fuzzer
compile_go_fuzzer github.com/ruk1ng001/quic-go-upgrade/fuzzing/tokens Fuzz token_fuzzer
compile_go_fuzzer github.com/ruk1ng001/quic-go-upgrade/fuzzing/handshake Fuzz handshake_fuzzer
