package templates

//go:generate go-bindata -prefix ../../../templates/... -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../templates/...

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	transactionsPath         = "../../../templates/"
	createCollectionFilename = "transactions/SetupUser.cdc"
	adminAssignRoxFilename   = "transactions/AdminAssignRox.cdc"
	transferRoxFilename      = "transactions/TransferRox.cdc"
)

// GenerateCreateCollectionScript Creates a script that instantiates a new
// NFT collection instance, stores the collection in memory, then stores a
// reference to the collection.
func GenerateCreateCollectionScript(env Environment) []byte {
	code := assets.MustAssetString(createCollectionFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateMintNFTScript Creates a script that uses the admin resource
// to mint a new NFT and deposit it into a user's collection
func GenerateMintNFTTransaction(env Environment) []byte {
	code := assets.MustAssetString(adminAssignRoxFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateTransferScript creates a script that withdraws an NFT token
// from a collection and deposits it to another collection
func GenerateTransferScript(env Environment) []byte {
	code := assets.MustAssetString(transferRoxFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateDestroyScript creates a script that withdraws an NFT token
// from a collection and destroys it
func GenerateDestroyScript(nftAddr, tokenAddr flow.Address, tokenContractName, storageLocation string, destroyNFTID int) []byte {
	template := `
		import NonFungibleToken from 0x%s
		import %s from 0x%s

		transaction {
		  prepare(acct: AuthAccount) {

			let collection <- acct.load<@%s.CollectionPrivate>(from:/storage/%s)!

			let nft <- collection.withdraw(withdrawID: %d)

			destroy nft

			acct.save(<-collection, to: /storage/%s)
		  }
		}
	`
	return []byte(fmt.Sprintf(template, nftAddr, tokenContractName, tokenAddr.String(), tokenContractName, storageLocation, destroyNFTID, storageLocation))
}
