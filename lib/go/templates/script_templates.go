package templates

import (
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	roxItemsTotalSupplyFileName = "scripts/rox_nfts_total_supply.cdc"
	nextBoxIdFileName           = "scripts/next_box_id.cdc"
	nftCollectionLengthFileName = "scripts/nft_collection_length.cdc"
	borrowNftFileName           = "scripts/borrow_nft.cdc"
	boxDataFileName             = "scripts/box_data.cdc"
	borrowRoxNftFileName        = "scripts/borrow_rox_nft.cdc"
)

func GenerateBorrowNftScript(env Environment) []byte {
	code := assets.MustAssetString(borrowNftFileName)

	return []byte(replaceAddresses(code, env))
}

func GenerateCollectionLengthScript(env Environment) []byte {
	code := assets.MustAssetString(nftCollectionLengthFileName)

	return []byte(replaceAddresses(code, env))
}

func GenerateTotalSupplyScript(env Environment) []byte {
	code := assets.MustAssetString(roxItemsTotalSupplyFileName)

	return []byte(replaceAddresses(code, env))
}

func GenerateNextBoxIdScript(env Environment) []byte {
	code := assets.MustAssetString(nextBoxIdFileName)

	return []byte(replaceAddresses(code, env))
}

func GenerateBoxDataScript(env Environment) []byte {
	code := assets.MustAssetString(boxDataFileName)

	return []byte(replaceAddresses(code, env))
}

func GenerateBorrowRoxNftScript(env Environment) []byte {
	code := assets.MustAssetString(borrowRoxNftFileName)

	return []byte(replaceAddresses(code, env))
}
