import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

transaction(recipient: Address, collectibleId: String, tier: String, mintNumber : UInt64) {
    
    let minter: &RoxContract.NFTMinter

    prepare(adminAcc: AuthAccount) {

        self.minter = adminAcc.borrow<&RoxContract.NFTMinter>(from: RoxContract.MinterStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        let recipient = getAccount(recipient)

        let receiver = recipient
            .getCapability(RoxContract.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        self.minter.mintNFT(recipient: receiver, collectibleId: collectibleId, tier: tier, mintNumber: mintNumber)
    }
}