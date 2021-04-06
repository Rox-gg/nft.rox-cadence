import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

transaction {
    prepare(signer: AuthAccount) {
      // if the account doesn't already have a collection
      if signer.borrow<&RoxContract.Collection>(from: RoxContract.CollectionStoragePath) != nil {
        log("Signer has completed set up before")
        return
      }
      // create a new empty collection
      let collection <- RoxContract.createEmptyCollection()
          
      // save it to the account
      signer.save(<-collection, to: RoxContract.CollectionStoragePath)

      // create a public capability for the collection
      signer.link<&RoxContract.Collection{NonFungibleToken.CollectionPublic, RoxContract.RoxCollectionPublic}>(RoxContract.CollectionPublicPath, target: RoxContract.CollectionStoragePath)
      
      log("Completed setup for account:")
      log(signer.address)
    }
}
