import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

transaction(name: String, metadata: {String: String}) {
    
    let adminRef: &RoxContract.Admin

    prepare(adminAcc: AuthAccount) {

        self.adminRef = adminAcc.borrow<&RoxContract.Admin>(from: RoxContract.AdminStoragePath)
            ?? panic("Could not borrow a reference to the RoxAdmin")
    }

    execute {
        self.adminRef.mintBox(name: name, metadata: metadata)

        log("Created box")
        log(name)
    }
}