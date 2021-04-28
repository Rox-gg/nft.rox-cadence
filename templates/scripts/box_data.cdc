import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

pub fun main(boxId: UInt32): RoxContract.BoxData {
    return RoxContract.BoxData(boxId: boxId)
}