import RoxContract from 0xNFTCONTRACTADDRESS

transaction(boxId: UInt32) {
    
    let adminRef: &RoxContract.Admin

    prepare(adminAcc: AuthAccount) {

        self.adminRef = adminAcc.borrow<&RoxContract.Admin>(from: RoxContract.AdminStoragePath)
            ?? panic("Could not borrow a reference to the NFT admin")
    }

    execute {
        let boxRef = self.adminRef.borrowBox(boxId: boxId)

        boxRef.lock()
    }
}