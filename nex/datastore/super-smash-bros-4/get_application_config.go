package nex_datastore_super_smash_bros_4

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	"github.com/PretendoNetwork/super-smash-bros-3ds/globals"
)

func GetApplicationConfig(err error, packet nex.PacketInterface, callID uint32, applicationID types.UInt32) (*nex.RMCMessage, *nex.Error) {
	//fmt.Printf("Application ID: %d\n", applicationID.Value)

	config := types.List[types.String]{
		types.NewString("1000000"),
		types.NewString("100"),
		types.NewString("0"),
		types.NewString("0"),
		types.NewString("1"),
		types.NewString("1"),
		types.NewString("0"),
		types.NewString("30"),
		types.NewString("50"),
		types.NewString("45"),
		types.NewString("30"),
		types.NewString("5"),
		types.NewString("8"),
		types.NewString("50"),
		types.NewString("87"),
		types.NewString("50"),
		types.NewString("85"),
		types.NewString("60"),
		types.NewString("60"),
		types.NewString("120"),
		types.NewString("600"),
		types.NewString("5"),
		types.NewString("600"),
		types.NewString("0"),
		types.NewString("45"),
		types.NewString("40"),
		types.NewString("100"),
		types.NewString("100"),
		types.NewString("82"),
		types.NewString("100"),
		types.NewString("95"),
		types.NewString("79"),
		types.NewString("10"),
		types.NewString("4"),
		types.NewString("2"),
		types.NewString("2"),
		types.NewString("2"),
		types.NewString("1"),
		types.NewString("5"),
		types.NewString("1"),
		types.NewString("2"),
		types.NewString("30"),
		types.NewString("9"),
		types.NewString("15"),
		types.NewString("0"),
		types.NewString("10"),
		types.NewString("10"),
		types.NewString("0"),
		types.NewString("0"),
		types.NewString("0"),
		types.NewString("7"),
		types.NewString("20"),
		types.NewString("0"),
		types.NewString("0"),
	}

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	config.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_smash_bros_4_protocol.ProtocolID
	rmcResponse.MethodID = datastore_super_smash_bros_4_protocol.MethodGetApplicationConfig
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
