package cometbft

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/config"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_Commit__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_CreateQueryContext__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_Info__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_InitChain__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_PrepareProposal__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_ProcessProposal__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_getContextForProposal__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_initChainer__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_internalFinalizeBlock__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_setState__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_workingHash__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_AppVersion__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_CheckTx__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_Close__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_CommitMultiStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_DefaultGenesis__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_FinalizeBlock__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_GetBlockRetentionHeight__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_LastBlockHeight__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_MountStore__ because parameters include func, chan, or unsupported interface: cosmossdk.io/store/types.StoreKey

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_SetName__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_SetVersion__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_Start__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_ValidateGenesis__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_setInterBlockCache__ because parameters include func, chan, or unsupported interface: cosmossdk.io/store/types.MultiStorePersistentCache

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_setMinRetainBlocks__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_validateFinalizeBlockHeight__ as it appears to be an interface:

func Fuzz_Nosy_state_CacheMultiStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *state
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		st.CacheMultiStore()
	})
}

func Fuzz_Nosy_state_Context__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *state
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		st.Context()
	})
}

func Fuzz_Nosy_state_SetContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *state
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		var ctx types.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		st.SetContext(ctx)
	})
}

// skipping Fuzz_Nosy_AttestationData[AttestationDataT any]_GetIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.AttestationData[AttestationDataT any]

// skipping Fuzz_Nosy_AttestationData[AttestationDataT any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.AttestationData[AttestationDataT any]

// skipping Fuzz_Nosy_BeaconState_HashTreeRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.BeaconState

// skipping Fuzz_Nosy_BeaconState_ValidatorIndexByCometBFTAddress__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.BeaconState

// skipping Fuzz_Nosy_MiddlewareI_FinalizeBlock__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.MiddlewareI

// skipping Fuzz_Nosy_MiddlewareI_InitGenesis__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.MiddlewareI

// skipping Fuzz_Nosy_MiddlewareI_PrepareProposal__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.MiddlewareI

// skipping Fuzz_Nosy_MiddlewareI_ProcessProposal__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.MiddlewareI

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_ApplySnapshotChunk__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_ExtendVote__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_ListSnapshots__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_LoadSnapshotChunk__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_OfferSnapshot__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_Query__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[LoggerT log.AdvancedLogger[LoggerT]]_VerifyVoteExtension__ as it appears to be an interface:

// skipping Fuzz_Nosy_SlashingInfo[SlashingInfoT any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.SlashingInfo[SlashingInfoT any]

// skipping Fuzz_Nosy_SlotData[AttestationDataT, SlashingInfoT, SlotDataT any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.SlotData[AttestationDataT, SlashingInfoT, SlotDataT any]

// skipping Fuzz_Nosy_StorageBackend[BeaconStateT BeaconState]_StateFromContext__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.StorageBackend[BeaconStateT github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service.BeaconState]

func Fuzz_Nosy_GetGenDocProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		GetGenDocProvider(cfg)
	})
}

func Fuzz_Nosy_SetChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID string) {
		SetChainID(chainID)
	})
}

func Fuzz_Nosy_SetIAVLCacheSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
		SetIAVLCacheSize(size)
	})
}

func Fuzz_Nosy_SetIAVLDisableFastNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, disable bool) {
		SetIAVLDisableFastNode(disable)
	})
}

// skipping Fuzz_Nosy_SetInterBlockCache__ because parameters include func, chan, or unsupported interface: cosmossdk.io/store/types.MultiStorePersistentCache

func Fuzz_Nosy_SetMinRetainBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, minRetainBlocks uint64) {
		SetMinRetainBlocks(minRetainBlocks)
	})
}

func Fuzz_Nosy_SetPruning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var opts types.PruningOptions
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}

		SetPruning(opts)
	})
}

func Fuzz_Nosy_convertValidatorUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u **transition.ValidatorUpdate
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		convertValidatorUpdate(u)
	})
}
