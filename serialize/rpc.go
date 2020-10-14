package serialize

import (
	"JFFun/data/command"
	"encoding/json"
	"errors"

	"github.com/golang/protobuf/proto"
)

var ErrNoCommandDecoder = errors.New("No Command Decoder")
var ErrNoCommandEncoder = errors.New("No Command Encoder")
var ErrEncode = errors.New("Encode Error")

func DecodeReq(cmd command.Command, mode SerializeMode, data []byte) (interface{}, error) {
	switch cmd {
	case command.Command_ping:
		fallthrough
	case command.Command_getOnlinePlayerCount:
		return nil, nil
	}
	return nil, ErrNoCommandDecoder
}

var emptyBytes []byte = []byte{}

func EncodeResp(cmd command.Command, mode SerializeMode, data interface{}) ([]byte, error) {
	switch cmd {
	case command.Command_ping:
		return emptyBytes, nil
	case command.Command_getOnlinePlayerCount:
		return encode(mode, data)
	}
	return nil, ErrNoCommandEncoder
}

func encode(mode SerializeMode, data interface{}) ([]byte, error) {
	switch mode {
	case JSON:
		return json.Marshal(data)
	case Protobuf:
		if pd, ok := data.(proto.Message); ok {
			return proto.Marshal(pd)
		}
	}
	return nil, ErrEncode
}