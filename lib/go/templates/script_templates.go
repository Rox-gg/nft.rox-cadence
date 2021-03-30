package templates

import (
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	roxItemsTotalSupplyFileName = "scripts/rox_items_total_supply.cdc"
	nftCollectionLengthFileName = "scripts/nft_collection_length.cdc"
	borrowNftFileName           = "scripts/read_nft_id.cdc"
)

// GenerateInspectCollectionScript creates a script that retrieves an NFT collection
// from storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateInspectCollectionScript(env Environment) []byte {
	code := assets.MustAssetString(borrowNftFileName)

	return []byte(replaceAddresses(code, env))
}

// GenerateInspectCollectionLenScript creates a script that retrieves an NFT collection
// from storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateInspectCollectionLenScript(env Environment) []byte {
	code := assets.MustAssetString(nftCollectionLengthFileName)

	return []byte(replaceAddresses(code, env))
}

// GenerateInspectNFTSupplyScript creates a script that reads
// the total supply of tokens in existence
// and makes assertions about the number
func GenerateInspectNFTSupplyScript(env Environment) []byte {
	code := assets.MustAssetString(roxItemsTotalSupplyFileName)

	return []byte(replaceAddresses(code, env))
}
