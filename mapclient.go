package ethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/mapprotocol/atlas/consensus/istanbul/backend"
	"github.com/mapprotocol/atlas/core/types"
)

// MAPHeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (ec *Client) MAPHeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var head *types.Header
	err := ec.c.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}
	return head, err
}

func (ec *Client) GetSnapshot(ctx context.Context, number *big.Int) (*backend.Snapshot, error) {
	var snap *backend.Snapshot
	err := ec.c.CallContext(ctx, &snap, "istanbul_getSnapshot", toBlockNumArg(number))
	if err != nil {
		return nil, err
	}
	return snap, err
}
