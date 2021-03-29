import NonFungibleToken from 0xNFTADDRESS
import RoxItems from 0xNFTCONTRACTADDRESS

pub fun main(accAddress: Address): Int {

    let acct = getAccount(accAddress)
			let collectionRef = acct.getCapability(/public/RoxItemsCollection)!.borrow<&{NonFungibleToken.CollectionPublic}>()
				?? panic("Could not borrow capability from public collection")

    return collectionRef.getIDs().length
}
