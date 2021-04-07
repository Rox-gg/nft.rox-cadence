import NonFungibleToken from "NonFungibleToken.cdc"

pub contract RoxContract: NonFungibleToken {

    // Events
    pub event ContractInitialized()

    pub event Withdraw(id: UInt64, from: Address?)
    pub event Deposit(id: UInt64, to: Address?)
    pub event Minted(id: UInt64, roxId: String)

    pub event BoxCreated(boxId: UInt32)
    // Emitted when a Box is locked, meaning Rox Nfts cannot be added
    pub event BoxLocked(boxId: UInt32)
    pub event RoxDestroyed(id: UInt64)

    // Named Paths
    pub let CollectionStoragePath: StoragePath
    pub let CollectionPublicPath: PublicPath
    pub let AdminStoragePath: StoragePath

    access(self) var boxes: @{UInt32: Box}
    access(self) var mintedNumberPerBox: {UInt32: UInt32}

    pub var nextBoxId: UInt32

    // The total number of Rox NFTs that have been minted
    pub var totalSupply: UInt64

    pub resource Box {

        // Unique ID for the box
        pub let boxId: UInt32
        pub let name: String
        pub let metadata: {String: String}
        pub var locked: Bool
        pub var mintedNumberPerRox: {String: UInt32}

        init(name: String, metadata: {String: String}) {
            pre {
                name.length > 0: "New Box name cannot be empty"
            }
            self.boxId = RoxContract.nextBoxId
            self.name = name 
            self.locked = false
            self.mintedNumberPerRox = {}
            self.metadata = metadata

            // Increment the boxID so that it isn't used again
            RoxContract.nextBoxId = RoxContract.nextBoxId + (1 as UInt32)
            RoxContract.mintedNumberPerBox[self.boxId] = 0

            emit BoxCreated(boxId: self.boxId)
        }

        // locks the Box so that no more Rox Nfts can be added to it
        pub fun lock() {
            if !self.locked {
                self.locked = true
                emit BoxLocked(boxId: self.boxId)
            }
        }

        pub fun mintRox(recipient: &{NonFungibleToken.CollectionPublic}, roxId: String, tier: String, metadata: {String: String}) {
            pre {
                !self.locked: "Cannot mint the rox: This box is locked"
            }

            if (self.mintedNumberPerRox[roxId] == nil){
                self.mintedNumberPerRox[roxId] = 0
            }
            self.mintedNumberPerRox[roxId] = self.mintedNumberPerRox[roxId]! + (1 as UInt32)
            RoxContract.mintedNumberPerBox[self.boxId] = RoxContract.mintedNumberPerBox[self.boxId]! + (1 as UInt32)

            // deposit it in the recipient's account using their reference
            recipient.deposit(token: <-create NFT(boxId: self.boxId,
                                                  roxId: roxId,
                                                  tier: tier,
                                                  mintNumber: self.mintedNumberPerRox[roxId]!,
                                                  metadata: metadata))

            emit Minted(id: RoxContract.totalSupply, roxId: roxId)
        }
    }

    // NFT
    // A Rox collectible as an NFT
    pub resource NFT: NonFungibleToken.INFT {
        // The token's ID
        pub let id: UInt64

        pub let boxId: UInt32
        pub let roxId: String
        pub let tier: String
        pub let mintNumber: UInt32
        pub let metadata: {String: String}

        init(boxId: UInt32, roxId: String, tier: String, mintNumber: UInt32, metadata: {String: String}) {
            
            RoxContract.totalSupply = RoxContract.totalSupply + 1 as UInt64
            
            self.id = RoxContract.totalSupply

            self.boxId = boxId
            self.roxId = roxId
            self.tier = tier
            self.mintNumber = mintNumber
            self.metadata = metadata
        }

        destroy() {
            emit RoxDestroyed(id: self.id)
        }
    }

    pub resource interface RoxCollectionPublic {
        pub fun deposit(token: @NonFungibleToken.NFT)
        pub fun getIDs(): [UInt64]
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT
        pub fun borrowRoxNft(id: UInt64): &RoxContract.NFT? {
            // If the result isn't nil, the id of the returned reference
            // should be the same as the argument to the function
            post {
                (result == nil) || (result?.id == id):
                    "Cannot borrow Rox reference: The ID of the returned reference is incorrect"
            }
        }
    }


    pub resource Collection: RoxCollectionPublic, NonFungibleToken.Provider, NonFungibleToken.Receiver, NonFungibleToken.CollectionPublic {

        pub var ownedNFTs: @{UInt64: NonFungibleToken.NFT}

        init() {
            self.ownedNFTs <- {}
        }

        pub fun withdraw(withdrawID: UInt64): @NonFungibleToken.NFT {
            let token <- self.ownedNFTs.remove(key: withdrawID) ?? panic("missing NFT")

            emit Withdraw(id: token.id, from: self.owner?.address)

            return <-token
        }

        pub fun deposit(token: @NonFungibleToken.NFT) {
            let token <- token as! @RoxContract.NFT

            let id: UInt64 = token.id

            let oldToken <- self.ownedNFTs[id] <- token

            emit Deposit(id: id, to: self.owner?.address)

            destroy oldToken
        }

        pub fun getIDs(): [UInt64] {
            return self.ownedNFTs.keys
        }

        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT {
            return &self.ownedNFTs[id] as &NonFungibleToken.NFT
        }

        pub fun borrowRoxNft(id: UInt64): &RoxContract.NFT? {
            if self.ownedNFTs[id] != nil {
                let ref = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT
                return ref as! &RoxContract.NFT
            } else {
                return nil
            }
        }

        destroy() {
            destroy self.ownedNFTs
        }
    }

    pub fun createEmptyCollection(): @NonFungibleToken.Collection {
        return <- create RoxContract.Collection()
    }

    // fetch
    // Get a reference to a RoxNft from an account's Collection, if available.
    // If an account does not have a RoxContract.Collection, panic.
    // If it has a collection but does not contain the itemId, return nil.
    // If it has a collection and that collection contains the itemId, return a reference to that.
    pub fun fetch(_ from: Address, itemID: UInt64): &RoxContract.NFT? {
        let collection = getAccount(from)
            .getCapability(RoxContract.CollectionPublicPath)
            .borrow<&RoxContract.Collection{RoxContract.RoxCollectionPublic}>()
            ?? panic("Couldn't get collection")
        return collection.borrowRoxNft(id: itemID)
    }

    pub resource Admin {

        pub fun mintBox(name: String, metadata: {String: String}) {
            var newBox <- create Box(name: name, metadata: metadata)
            RoxContract.boxes[newBox.boxId] <-! newBox
        }

        pub fun borrowBox(boxId: UInt32): &Box {
            pre {
                RoxContract.boxes[boxId] != nil: "Cannot borrow Box: The Box doesn't exist"
            }
            
            return &RoxContract.boxes[boxId] as &Box
        }

        pub fun createNewAdmin(): @Admin {
            return <-create Admin()
        }
    }

    pub fun getBoxName(boxId: UInt32): String? {
        return RoxContract.boxes[boxId]?.name
    }

    pub fun isBoxLocked(boxId: UInt32): Bool? {
        return RoxContract.boxes[boxId]?.locked
    }

    pub fun getNumberMintedRoxInBox(boxId: UInt32): UInt32? {
        return RoxContract.mintedNumberPerBox[boxId]
    }

    // -----------------------------------------------------------------------
    // RoxContract initialization function
    // -----------------------------------------------------------------------
    //
    init() {
        self.CollectionStoragePath = /storage/RoxCollection
        self.CollectionPublicPath = /public/RoxCollection
        self.AdminStoragePath = /storage/RoxAdmin

        self.totalSupply = 0
        self.nextBoxId = 1
        self.boxes <- {}
        self.mintedNumberPerBox = {}

        let collection <- RoxContract.createEmptyCollection()
        self.account.save(<-collection, to: RoxContract.CollectionStoragePath)
        self.account.link<&RoxContract.Collection{NonFungibleToken.CollectionPublic, RoxContract.RoxCollectionPublic}>(RoxContract.CollectionPublicPath, target: RoxContract.CollectionStoragePath)

        let admin <- create Admin()
        self.account.save(<-admin, to: self.AdminStoragePath)

        emit ContractInitialized()
    }
}
 
 