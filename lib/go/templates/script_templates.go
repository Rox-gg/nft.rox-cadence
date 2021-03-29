package templates

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	roxItemsTotalSupplyFileName = "scripts/rox_items_total_supply.cdc"
	nftCollectionLengthFileName = "scripts/nft_collection_length.cdc"
)

// GenerateInspectCollectionScript creates a script that retrieves an NFT collection
// from storage and tries to borrow a reference for an NFT that it owns.
// If it owns it, it will not fail.
func GenerateInspectCollectionScript(nftAddr, tokenAddr, userAddr flow.Address, tokenContractName, storageLocation string, nftID int) []byte {
	template := `
		import NonFungibleToken from 0x%s
		import %s from 0x%s

		pub fun main() {
			let acct = getAccount(0x%s)
			let collectionRef = acct.getCapability(/public/%s)!.borrow<&{NonFungibleToken.CollectionPublic}>()
				?? panic("Could not borrow capability from public collection")
			
			let tokenRef = collectionRef.borrowNFT(id: UInt64(%d))
		}
	`

	return []byte(fmt.Sprintf(template, nftAddr, tokenContractName, tokenAddr, userAddr, storageLocation, nftID))
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
