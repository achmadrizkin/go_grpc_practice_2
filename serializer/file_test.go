package serializer_test

import (
	pb "go_grpc_client_server/proto"
	"go_grpc_client_server/sample"
	"go_grpc_client_server/serializer"

	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestTextFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtoBufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	// check if same val
	require.True(t, proto.Equal(laptop1, laptop2))

	// proto to json
	err = serializer.WriteProtoBuffToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
