import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

transaction(recipient: Address, withdrawID: UInt64) {

    let tempNFT: @NonFungibleToken.NFT

    prepare(signer: AuthAccount) {

        let collectionRef = signer.borrow<&RoxContract.Collection>(from: RoxContract.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the owner's collection")

        self.tempNFT <- collectionRef.withdraw(withdrawID: withdrawID)
    }
    execute {

        let recipient = getAccount(recipient)

        let depositRef = recipient.getCapability(RoxContract.CollectionPublicPath)!.borrow<&{NonFungibleToken.CollectionPublic}>()!

        depositRef.deposit(token: <- self.tempNFT)

        log("NFT Transferred")
  }
}
