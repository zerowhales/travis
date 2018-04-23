package stake

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/go-crypto"
)

/**** code to parse accounts from genesis docs ***/

// GenesisValidator - genesis validator parameters
type genesisValidator struct {
	Address   common.Address `json:"address"`
	PubKey    crypto.PubKey  `json:"pub_key"`
	Power     int64          `json:"power"`
	Name      string         `json:"name"`
	MaxAmount string         `json:"max_amount"`
	Cut       float64        `json:"cut"`
}
