package templates

//go:generate go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions

import (
	"fmt"
	"strings"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	defaultNFTAddress        = "NFTADDRESS"
	defaultContractAddress   = "NFTCONTRACTADDRESS"
	defaultNFTName           = "RoxItems"
	defaultNFTStorage        = "RoxItemsCollection"
	createCollectionFilename = "SetupUser.cdc"
)

// GenerateCreateCollectionScript Creates a script that instantiates a new
// NFT collection instance, stores the collection in memory, then stores a
// reference to the collection.
func GenerateCreateCollectionScript(nftAddr, tokenAddr, tokenContractName string, storageName string) []byte {
	code := assets.MustAssetString(createCollectionFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultNFTAddress,
		"0x"+nftAddr,
	)
	code = strings.ReplaceAll(
		code,
		"0x"+defaultContractAddress,
		"0x"+tokenAddr,
	)
	code = strings.ReplaceAll(
		code,
		"0x"+defaultNFTName,
		"0x"+tokenContractName,
	)
	code = strings.ReplaceAll(
		code,
		"0x"+defaultNFTStorage,
		"0x"+storageName,
	)

	return []byte(code)
}

// GenerateMintNFTScript Creates a script that uses the admin resource
// to mint a new NFT and deposit it into a user's collection
func GenerateMintNFTTransaction(nftAddr, tokenAddr) []byte {
	template := `
	import NonFungibleToken from 0x%s
	import RoxItems from 0x%s

	transaction(recipient: Address, typeID: UInt64) {
    
		let minter: &RoxItems.NFTMinter
	
		prepare(adminAcc: AuthAccount) {
	
			self.minter = adminAcc.borrow<&RoxItems.NFTMinter>(from: RoxItems.MinterStoragePath)
				?? panic("Could not borrow a reference to the NFT minter")
		}
	
		execute {
			let recipient = getAccount(recipient)
	
			let receiver = recipient
				.getCapability(RoxItems.CollectionPublicPath)
				.borrow<&{NonFungibleToken.CollectionPublic}>()
				?? panic("Could not get receiver reference to the NFT Collection")
	
			self.minter.mintNFT(recipient: receiver, typeID: typeID)
		}
	}
	`

	return []byte(fmt.Sprintf(template, nftAddr, tokenAddr, receiverAddr))
}

// GenerateTransferScript creates a script that withdraws an NFT token
// from a collection and deposits it to another collection
func GenerateTransferScript(nftAddr, tokenAddr flow.Address, tokenContractName, storageLocation string, receiverAddr flow.Address, transferNFTID int) []byte {
	template := `
		import NonFungibleToken from 0x%s
		import %s from 0x%s

		transaction {
		  prepare(acct: AuthAccount) {
			let recipient = getAccount(0x%s)

			let collectionRef = acct.borrow<&%s.Collection>(from: /storage/%s)!
			let depositRef = recipient.getCapability(/public/%s)!.borrow<&{NonFungibleToken.CollectionPublic}>()!

			let nft <- collectionRef.withdraw(withdrawID: %d)

			depositRef.deposit(token: <-nft)
		  }
		}
	`

	return []byte(fmt.Sprintf(template, nftAddr, tokenContractName, tokenAddr, receiverAddr.String(), tokenContractName, storageLocation, storageLocation, transferNFTID))
}

// GenerateDestroyScript creates a script that withdraws an NFT token
// from a collection and destroys it
func GenerateDestroyScript(nftAddr, tokenAddr flow.Address, tokenContractName, storageLocation string, destroyNFTID int) []byte {
	template := `
		import NonFungibleToken from 0x%s
		import %s from 0x%s

		transaction {
		  prepare(acct: AuthAccount) {

			let collection <- acct.load<@%s.Collection>(from:/storage/%s)!

			let nft <- collection.withdraw(withdrawID: %d)

			destroy nft
			
			acct.save(<-collection, to: /storage/%s)
		  }
		}
	`

	return []byte(fmt.Sprintf(template, nftAddr, tokenContractName, tokenAddr.String(), tokenContractName, storageLocation, destroyNFTID, storageLocation))
}
