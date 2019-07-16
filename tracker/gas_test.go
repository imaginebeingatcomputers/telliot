package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestGas(t *testing.T) {
	client := rpc.NewMockClientWithValues(big.NewInt(7000000000), 1, big.NewInt(7000000000))
	DB, err := db.Open(filepath.Join(os.TempDir(), "test_gas"))
	if err != nil {
		t.Fatal(err)
	}
	tracker := &GasTracker{}
	ctx := context.WithValue(context.Background(), ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.GasKey)
	if err != nil {
		t.Fatal(err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Gas Price stored: %v\n", string(v))
	if b.Cmp(big.NewInt(7000000000)) != 0 {
		t.Fatalf("Balance from client did not match what should have been stored in DB. %s != %s", b, "Should be 1")
	}
}
