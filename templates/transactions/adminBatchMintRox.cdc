import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

transaction(recipient: Address, boxId: UInt32, quantity: UInt64, roxId: String, tier: String, metadata: {String: String}) {
    
    let adminRef: &RoxContract.Admin

    prepare(adminAcc: AuthAccount) {

        self.adminRef = adminAcc.borrow<&RoxContract.Admin>(from: RoxContract.AdminStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        let boxRef = self.adminRef.borrowBox(boxId: boxId)
        let recipient = getAccount(recipient)

        let receiver = recipient
            .getCapability(RoxContract.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        boxRef.batchMintRox(recipient: receiver, quantity: quantity, roxId: roxId, tier: tier, metadata: metadata)
    }
}