import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

// This script reads metadata about an NFT in a user's collection
pub fun main(account: Address, id: UInt64): UInt64 {

    // Get the public collection of the owner of the token
    let collectionRef = getAccount(account)
        .getCapability(RoxContract.CollectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")

    // Borrow a reference to a specific NFT in the collection
    let nft = collectionRef.borrowNFT(id: id)

    return nft.id
}