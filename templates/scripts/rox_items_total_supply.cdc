import NonFungibleToken from 0xNFTADDRESS
import RoxContract from 0xNFTCONTRACTADDRESS

pub fun main(): UInt64 {
    return RoxContract.totalSupply
}
