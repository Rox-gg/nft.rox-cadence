import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

// This script reads metadata about an NFT in a user's collection
pub fun main(account: Address, id: UInt64): &RoxContract.NFT? {
    
    // Borrow a reference to a specific NFT in the collection
    let nft = RoxContract.fetch(account, itemID: id)

    return nft
}