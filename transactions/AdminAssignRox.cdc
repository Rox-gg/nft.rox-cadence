import NonFungibleToken from 0xNFTADDRESS
import RoxItems from 0xNFTCONTRACTADDRESS

transaction(recipient: Address, collectibleId: String, tier: String, mintNumber : UInt64) {
    
    let minter: &RoxItems.NFTMinter

    prepare(adminAcc: AuthAccount) {

        self.minter = adminAcc.borrow<&RoxItems.NFTMinter>(from: RoxItems.MinterStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        let recipient = getAccount(recipient)

        let receiver = recipient
            .getCapability(RoxItems.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        self.minter.mintNFT(recipient: receiver, collectibleId: collectibleId, tier: tier, mintNumber: mintNumber)
    }
}