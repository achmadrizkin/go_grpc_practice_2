package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, fileName string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto message to binary: %w", err)
	}

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("Cannot write binary data to file: %w", err)
	}

	return nil
}

func ReadProtoBufFromBinaryFile(fileName string, message proto.Message) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Cannot Read binary data to file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Cannot marshal binary to proto message: %w", err)
	}

	return nil
}

func WriteProtoBuffToJSONFile(message proto.Message, fileName string) error {
	data, err := ProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto message to JSON: %w", err)
	}

	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("Cannot write json data to file: %w", err)
	}

	return nil
}
