package nex

import (
	datastore_super_smash_bros_4_protocol "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4"
	matchmake_referee "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-referee"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	"github.com/PretendoNetwork/super-smash-bros-3ds/globals"
	datastore_super_smash_bros_4 "github.com/PretendoNetwork/super-smash-bros-3ds/nex/datastore/super-smash-bros-4"
)

func registerSecureServerNEXProtocols() {
	datastoreProtocol := datastore_super_smash_bros_4_protocol.NewProtocol(globals.SecureEndpoint)
	datastoreProtocol.GetApplicationConfig = datastore_super_smash_bros_4.GetApplicationConfig
	datastoreProtocol.GetProfiles = datastore_super_smash_bros_4.GetProfiles
	datastoreProtocol.PostProfile = datastore_super_smash_bros_4.PostProfile
	datastoreProtocol.GetFightingPowerChartAll = datastore_super_smash_bros_4.GetFightingPowerChartAll
	datastoreProtocol.GetWorldPlayReport = datastore_super_smash_bros_4.GetWorldPlayReport
	//datastoreProtocol.GetNextReplay = datastore_super_smash_bros_4.GetNextReplay
	//datastoreProtocol.PreparePostReplay = datastore_super_smash_bros_4.PreparePostReplay
	//datastoreProtocol.CheckPostReplay = datastore_super_smash_bros_4.CheckPostReplay
	globals.SecureEndpoint.RegisterServiceProtocol(datastoreProtocol)

	matchmakeRefereeProtocol := matchmake_referee.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeRefereeProtocol)

	// TODO - Add legacy ranking protocol!
	rankingProtocol := ranking.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(rankingProtocol)
}
