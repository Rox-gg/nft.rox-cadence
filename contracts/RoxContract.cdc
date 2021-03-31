import NonFungibleToken from 0x02

pub contract RoxContract: NonFungibleToken {

    // Events
    pub event ContractInitialized()
    pub event Withdraw(id: UInt64, from: Address?)
    pub event Deposit(id: UInt64, to: Address?)
    pub event Minted(id: UInt64, collectibleId: String)

    // Named Paths
    pub let CollectionStoragePath: StoragePath
    pub let CollectionPublicPath: PublicPath
    pub let MinterStoragePath: StoragePath

    // The total number of NFT that have been minted
    pub var totalSupply: UInt64
    pub var numberMintedPerCollectible: {String: UInt32}

    // NFT
    // A Rox Item as an NFT
    pub resource NFT: NonFungibleToken.INFT {
        // The token's ID
        pub let id: UInt64
        pub let tier: String // enum in BE
        pub let collectibleId: String
        pub let mintNumber: UInt32 // Don't know what it is

        init(initID: UInt64, collectibleId: String, tier: String, mintNumber: UInt32) {
            self.id = initID
            self.collectibleId = collectibleId
            self.tier = tier
            self.mintNumber = mintNumber
        }
    }

    pub resource interface CollectionRoxPublic {
        pub fun deposit(token: @NonFungibleToken.NFT)
        pub fun getIDs(): [UInt64]
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT
        pub fun borrowRoxNft(id: UInt64): &RoxContract.NFT? {
            // If the result isn't nil, the id of the returned reference
            // should be the same as the argument to the function
            post {
                (result == nil) || (result?.id == id):
                    "Cannot borrow RoxItem reference: The ID of the returned reference is incorrect"
            }
        }
    }


    pub resource Collection: CollectionRoxPublic, NonFungibleToken.Provider, NonFungibleToken.Receiver, NonFungibleToken.CollectionPublic {

        pub var ownedNFTs: @{UInt64: NonFungibleToken.NFT}

        pub fun withdraw(withdrawID: UInt64): @NonFungibleToken.NFT {
            let token <- self.ownedNFTs.remove(key: withdrawID) ?? panic("missing NFT")

            emit Withdraw(id: token.id, from: self.owner?.address)

            return <-token
        }

        pub fun deposit(token: @NonFungibleToken.NFT) {
            let token <- token as! @RoxContract.NFT

            let id: UInt64 = token.id

            // add the new token to the dictionary which removes the old one
            let oldToken <- self.ownedNFTs[id] <- token

            emit Deposit(id: id, to: self.owner?.address)

            destroy oldToken
        }

        pub fun getIDs(): [UInt64] {
            return self.ownedNFTs.keys
        }

        // borrowNFT
        // Gets a reference to an NFT in the collection
        // so that the caller can read its metadata and call its methods
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT {
            return &self.ownedNFTs[id] as &NonFungibleToken.NFT
        }

        // borrowRoxItem
        // Gets a reference to an NFT in the collection as a RoxItem,
        // exposing all of its fields (including the typeID).
        // This is safe as there are no functions that can be called on the RoxItem.
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

        init () {
            self.ownedNFTs <- {}
        }
    }

    pub fun createEmptyCollection(): @NonFungibleToken.Collection {
        return <- create Collection()
    }

	pub resource NFTMinter {

		pub fun mintNFT(recipient: &{NonFungibleToken.CollectionPublic}, collectibleId: String, tier: String) {

            emit Minted(id: RoxContract.totalSupply, collectibleId: collectibleId)

            var numberPerCollectible = RoxContract.numberMintedPerCollectible[collectibleId]
            if (numberPerCollectible == nil){
                numberPerCollectible = 0
                RoxContract.numberMintedPerCollectible[collectibleId] = numberPerCollectible
            }

            // deposit it in the recipient's account using their reference
            recipient.deposit(token: <-create RoxContract.NFT(initID: RoxContract.totalSupply, collectibleId: collectibleId, tier: tier, mintNumber: numberPerCollectible!))

            RoxContract.totalSupply = RoxContract.totalSupply + (1 as UInt64)
            RoxContract.numberMintedPerCollectible[collectibleId] = numberPerCollectible! + (1 as UInt32)
		}
	}

    // fetch
    // Get a reference to a RoxItem from an account's Collection, if available.
    // If an account does not have a RoxItems.Collection, panic.
    // If it has a collection but does not contain the itemId, return nil.
    // If it has a collection and that collection contains the itemId, return a reference to that.
    //
    pub fun fetch(_ from: Address, itemID: UInt64): &RoxContract.NFT? {
        let collection = getAccount(from)
            .getCapability(RoxContract.CollectionPublicPath)
            .borrow<&RoxContract.Collection{RoxContract.CollectionRoxPublic}>()
            ?? panic("Couldn't get collection")
        // We trust RoxItems.Collection.borowRoxItem to get the correct itemID
        // (it checks it before returning it).
        return collection.borrowRoxNft(id: itemID)
    }

	init() {
        self.CollectionStoragePath = /storage/RoxCollection
        self.CollectionPublicPath = /public/RoxCollection
        self.MinterStoragePath = /storage/RoxMinter

        self.totalSupply = 0
        self.numberMintedPerCollectible = {}

        let minter <- create NFTMinter()
        self.account.save(<-minter, to: self.MinterStoragePath)

        let collection <- RoxContract.createEmptyCollection()
        self.account.save(<-collection, to: RoxContract.CollectionStoragePath)
        self.account.link<&RoxContract.Collection{NonFungibleToken.CollectionPublic, RoxContract.CollectionRoxPublic}>(RoxContract.CollectionPublicPath, target: RoxContract.CollectionStoragePath)

        emit ContractInitialized()
	}
}
 
 