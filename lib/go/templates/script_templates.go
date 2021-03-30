package templates

import (
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	roxItemsTotalSupplyFileName = "scripts/rox_items_total_supply.cdc"
	nftCollectionLengthFileName = "scripts/nft_collection_length.cdc"
	borrowNftFileName           = "scripts/borrow_nft.cdc"
)

func GenerateBorrowNftScript(env Environment) []byte {
	code := assets.MustAssetString(borrowNftFileName)

	return []byte(replaceAddresses(code, env))
}

func GenerateRoxItemsCollectionLengthScript(env Environment) []byte {
	code := assets.MustAssetString(nftCollectionLengthFileName)

	return []byte(replaceAddresses(code, env))
}

func GenerateRoxItemsTotalSupplyScript(env Environment) []byte {
	code := assets.MustAssetString(roxItemsTotalSupplyFileName)

	return []byte(replaceAddresses(code, env))
}
