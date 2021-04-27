import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

transaction(recipient: Address, boxId: UInt32, roxId: String, tier: String, metadata: {String: String}) {
    
    let collectionRef: &RoxContract.Collection

    prepare(adminAcc: AuthAccount) {

        self.adminRef = adminAcc.borrow<&RoxContract.Collection>(from: RoxContract.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the Rox Collection")

    }

    execute {
        let boxRef = self.collectionRef.borrowBox(boxId: boxId)

        boxRef.ownedNFTs

        let recipient = getAccount(recipient)

        let receiver = recipient
            .getCapability(RoxContract.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        boxRef.mintRox(recipient: receiver, roxId: roxId, tier: tier, metadata: metadata)
    }
}