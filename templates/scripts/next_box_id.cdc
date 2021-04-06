import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

pub fun main(): UInt32 {
    return RoxContract.nextBoxId
}
