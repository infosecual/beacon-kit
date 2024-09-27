package deposit

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_BeaconDepositContractCaller_DepositAuth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCaller
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.DepositAuth(opts, arg0)
	})
}

func Fuzz_Nosy_BeaconDepositContractCaller_DepositCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCaller
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.DepositCount(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCaller
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.Owner(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractCaller_OwnershipHandoverExpiresAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCaller
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pendingOwner common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.OwnershipHandoverExpiresAt(opts, pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractCaller_SupportsInterface__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCaller
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var interfaceId [4]byte
		fill_err = tp.Fill(&interfaceId)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.SupportsInterface(opts, interfaceId)
	})
}

// skipping Fuzz_Nosy_BeaconDepositContractCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_BeaconDepositContractCallerSession_DepositAuth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCallerSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.DepositAuth(arg0)
	})
}

func Fuzz_Nosy_BeaconDepositContractCallerSession_DepositCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCallerSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.DepositCount()
	})
}

func Fuzz_Nosy_BeaconDepositContractCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCallerSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.Owner()
	})
}

func Fuzz_Nosy_BeaconDepositContractCallerSession_OwnershipHandoverExpiresAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCallerSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var pendingOwner common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.OwnershipHandoverExpiresAt(pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractCallerSession_SupportsInterface__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractCallerSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var interfaceId [4]byte
		fill_err = tp.Fill(&interfaceId)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.SupportsInterface(interfaceId)
	})
}

func Fuzz_Nosy_BeaconDepositContractDepositIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_BeaconDepositContractDepositIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_BeaconDepositContractDepositIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_FilterDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.FilterDeposit(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_FilterOwnershipHandoverCanceled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pendingOwner []common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.FilterOwnershipHandoverCanceled(opts, pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_FilterOwnershipHandoverRequested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pendingOwner []common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.FilterOwnershipHandoverRequested(opts, pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var oldOwner []common.Address
		fill_err = tp.Fill(&oldOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.FilterOwnershipTransferred(opts, oldOwner, newOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_ParseDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.ParseDeposit(log)
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_ParseOwnershipHandoverCanceled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.ParseOwnershipHandoverCanceled(log)
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_ParseOwnershipHandoverRequested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.ParseOwnershipHandoverRequested(log)
	})
}

func Fuzz_Nosy_BeaconDepositContractFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractFilterer
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.ParseOwnershipTransferred(log)
	})
}

// skipping Fuzz_Nosy_BeaconDepositContractFilterer_WatchDeposit__ because parameters include func, chan, or unsupported interface: chan<- *github.com/berachain/beacon-kit/mod/geth-primitives/pkg/deposit.BeaconDepositContractDeposit

// skipping Fuzz_Nosy_BeaconDepositContractFilterer_WatchOwnershipHandoverCanceled__ because parameters include func, chan, or unsupported interface: chan<- *github.com/berachain/beacon-kit/mod/geth-primitives/pkg/deposit.BeaconDepositContractOwnershipHandoverCanceled

// skipping Fuzz_Nosy_BeaconDepositContractFilterer_WatchOwnershipHandoverRequested__ because parameters include func, chan, or unsupported interface: chan<- *github.com/berachain/beacon-kit/mod/geth-primitives/pkg/deposit.BeaconDepositContractOwnershipHandoverRequested

// skipping Fuzz_Nosy_BeaconDepositContractFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/berachain/beacon-kit/mod/geth-primitives/pkg/deposit.BeaconDepositContractOwnershipTransferred

func Fuzz_Nosy_BeaconDepositContractOwnershipHandoverCanceledIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipHandoverCanceledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipHandoverCanceledIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipHandoverCanceledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipHandoverCanceledIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipHandoverCanceledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipHandoverRequestedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipHandoverRequestedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipHandoverRequestedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipHandoverRequestedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipHandoverRequestedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipHandoverRequestedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_BeaconDepositContractOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *BeaconDepositContractOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_BeaconDepositContractRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_BeaconDepositContractRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_BeaconDepositContractRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractRaw
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.Transfer(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_AllowDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var depositor common.Address
		fill_err = tp.Fill(&depositor)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.AllowDeposit(depositor, number)
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_CancelOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.CancelOwnershipHandover()
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_CompleteOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var pendingOwner common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.CompleteOwnershipHandover(pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawal_credentials []byte
		fill_err = tp.Fill(&withdrawal_credentials)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.Deposit(pubkey, withdrawal_credentials, amount, signature)
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_DepositAuth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.DepositAuth(arg0)
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_DepositCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.DepositCount()
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.Owner()
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_OwnershipHandoverExpiresAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var pendingOwner common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.OwnershipHandoverExpiresAt(pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.RenounceOwnership()
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_RequestOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.RequestOwnershipHandover()
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_SupportsInterface__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var interfaceId [4]byte
		fill_err = tp.Fill(&interfaceId)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.SupportsInterface(interfaceId)
	})
}

func Fuzz_Nosy_BeaconDepositContractSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactor_AllowDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactor
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var depositor common.Address
		fill_err = tp.Fill(&depositor)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.AllowDeposit(opts, depositor, number)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactor_CancelOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactor
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.CancelOwnershipHandover(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactor_CompleteOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactor
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pendingOwner common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.CompleteOwnershipHandover(opts, pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactor_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactor
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawal_credentials []byte
		fill_err = tp.Fill(&withdrawal_credentials)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.Deposit(opts, pubkey, withdrawal_credentials, amount, signature)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactor
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactor_RequestOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactor
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.RequestOwnershipHandover(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactor
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.TransferOwnership(opts, newOwner)
	})
}

// skipping Fuzz_Nosy_BeaconDepositContractTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_BeaconDepositContractTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorRaw
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil || opts == nil {
			return
		}

		_BeaconDepositContract.Transfer(opts)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactorSession_AllowDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var depositor common.Address
		fill_err = tp.Fill(&depositor)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.AllowDeposit(depositor, number)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactorSession_CancelOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.CancelOwnershipHandover()
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactorSession_CompleteOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var pendingOwner common.Address
		fill_err = tp.Fill(&pendingOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.CompleteOwnershipHandover(pendingOwner)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactorSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawal_credentials []byte
		fill_err = tp.Fill(&withdrawal_credentials)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.Deposit(pubkey, withdrawal_credentials, amount, signature)
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.RenounceOwnership()
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactorSession_RequestOwnershipHandover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.RequestOwnershipHandover()
	})
}

func Fuzz_Nosy_BeaconDepositContractTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _BeaconDepositContract *BeaconDepositContractTransactorSession
		fill_err = tp.Fill(&_BeaconDepositContract)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _BeaconDepositContract == nil {
			return
		}

		_BeaconDepositContract.TransferOwnership(newOwner)
	})
}

// skipping Fuzz_Nosy_DeployBeaconDepositContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
