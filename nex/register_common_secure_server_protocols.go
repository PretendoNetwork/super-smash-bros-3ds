package nex

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	"github.com/PretendoNetwork/super-smash-bros-3ds/database"
	"github.com/PretendoNetwork/super-smash-bros-3ds/globals"
	local_globals "github.com/PretendoNetwork/super-smash-bros-3ds/globals"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	commonSecureProtocol := common_secure.NewCommonProtocol(secureProtocol)
	commonSecureProtocol.EnableInsecureRegister()
	commonSecureProtocol.CreateReportDBRecord = func(pid types.PID, reportID types.UInt32, reportData types.QBuffer) error {
		return nil
	}

	matchmakingManager := common_globals.NewMatchmakingManager(local_globals.SecureEndpoint, database.Postgres)
	matchmakingManager.GetUserFriendPIDs = globals.GetUserFriendPids

	natTraversalProtocol := nat_traversal.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := match_making.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	commonMatchMakingProtocol := common_match_making.NewCommonProtocol(matchMakingProtocol)
	commonMatchMakingProtocol.SetManager(matchmakingManager)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol := common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol.SetManager(matchmakingManager)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol.SetManager(matchmakingManager)

	//commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = func(matchmakeSession *mm_types.MatchmakeSession) {
	//	globals.Logger.Infof("%v", matchmakeSession.FormatToString(1))
	//}
	//
	//commonMatchmakeExtensionProtocol.OnAfterAutoMatchmakeWithSearchCriteriaPostpone = func(packet nex.PacketInterface, lstSearchCriteria types.List[mm_types.MatchmakeSessionSearchCriteria], anyGathering types.AnyObjectHolder[mm_types.GatheringInterface], strMessage types.String) {
	//	globals.Logger.Infof("=== OnAfter for pid %v ===", packet.Sender().PID())
	//	for _, criteria := range lstSearchCriteria {
	//		globals.Logger.Infof("%v", criteria.FormatToString(1))
	//	}
	//	globals.Logger.Infof("%v", anyGathering.FormatToString(1))
	//	globals.Logger.Infof(" %v", strMessage)
	//}
}
