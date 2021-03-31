import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

transaction(id: UInt64) {
	prepare(acct: AuthAccount) {

		let collection <- acct.load<@RoxContract.Collection>(from: RoxContract.CollectionStoragePath)!

		let nft <- collection.withdraw(withdrawID: id)

		destroy nft

		acct.save(<-collection, to: RoxContract.CollectionStoragePath)
	}
}