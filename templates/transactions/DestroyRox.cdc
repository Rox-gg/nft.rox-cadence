import NonFungibleToken from 0xNFTADDRESS
import RoxItems from 0xNFTCONTRACTADDRESS

transaction(id: UInt64) {
	prepare(acct: AuthAccount) {

		let collection <- acct.load<@RoxItems.CollectionPrivate>(from: RoxItems.CollectionStoragePath)!

		let nft <- collection.withdraw(withdrawID: id)

		destroy nft

		acct.save(<-collection, to: RoxItems.CollectionStoragePath)
	}
}