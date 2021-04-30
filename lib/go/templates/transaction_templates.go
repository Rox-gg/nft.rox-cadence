package templates

import (
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	setupUserFilename         = "transactions/setupUser.cdc"
	adminMintBoxFilename      = "transactions/adminMintBox.cdc"
	adminLockBoxFilename      = "transactions/adminLockBox.cdc"
	adminMintRoxFilename      = "transactions/adminMintRox.cdc"
	adminBatchMintRoxFilename = "transactions/adminBatchMintRox.cdc"
	transferRoxFilename       = "transactions/transferRox.cdc"
	destroyRoxFilename        = "transactions/destroyRox.cdc"
)

// GenerateCreateCollectionScript Creates a script that instantiates a new
// NFT collection instance, stores the collection in memory, then stores a
// reference to the collection.
func GenerateCreateCollectionScript(env Environment) []byte {
	code := assets.MustAssetString(setupUserFilename)

	return []byte(replaceAddresses(code, env))
}

func GenerateMintBoxTransaction(env Environment) []byte {
	code := assets.MustAssetString(adminMintBoxFilename)

	return []byte(replaceAddresses(code, env))
}

func GenerateLockBoxTransaction(env Environment) []byte {
	code := assets.MustAssetString(adminLockBoxFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateMintNFTScript Creates a script that uses the admin resource
// to mint a new NFT and deposit it into a user's collection
func GenerateMintNFTTransaction(env Environment) []byte {
	code := assets.MustAssetString(adminMintRoxFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateMintNFTScript Creates a script that uses the admin resource
// to mint a new NFT and deposit it into a user's collection
func GenerateBatchMintNFTTransaction(env Environment) []byte {
	code := assets.MustAssetString(adminBatchMintRoxFilename)

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
func GenerateDestroyScript(env Environment) []byte {
	code := assets.MustAssetString(destroyRoxFilename)

	return []byte(replaceAddresses(code, env))
}
