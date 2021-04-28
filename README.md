# ROX NFT #

## Introduction

This repository contains smart contracts, transactions and scripts
which implements the core functionality of NFT ROX digital collectibles.

## What is ROX NFT

Rox NFTs represent the best stream moments and let you own a piece of the streamer’s history.
See more at nft.rox.gg

## Non Fungible Token Standard

The NFT ROX contracts utilize the Flow NFT `NonFungibleToken` standard which is equivalent to ERC-721 or ERC-1155 on Ethereum.

## ROX Contract Addresses

`RoxContract.cdc`: This is the main ROX NFT smart contract that defines
the core functionality of the NFT.

| Network | Contract Address     |
|---------|----------------------|
| Testnet | `0x54a8fcf94992730f` |
| Mainnet | `-` |

`RoxMarket.cdc`: coming soon

## Directory Structure

The directories here are organized into contracts, scripts, and transactions.

Contracts contain the source code for the ROX contracts that are deployed to Flow.

Scripts contain read-only transactions to get information about
the state of someones Collection or about the state of the Rox contract.

Transactions contain the transactions that various admins and users can use
to perform actions in the smart contract like creating boxes,
minting and transfering Roxes.

 - `contracts/` : Where the ROX NFT related smart contracts live.
 - `templates/transactions` : This directory contains all the transactions.
 - `templates/scripts/`  : This contains all the read-only Cadence scripts 
 that are used to read information from the smart contract
 or from a resource in account storage.
 - `lib/` : This directory contains packages for specific programming languages
 to be able to read copies of the NFT Rox smart contracts, transaction templates,
 and scripts. Also contains automated tests written in GO language.

## Contract Overview

Each ROX represents a special moment from a streamer's history.

Roxes are grouped into boxes and each Rox belongs to a specific box.

In the box there can be many different types and different number of Roxes minted.

All the roxes are minted through the box. Roxes in the box are grouped by `roxId`.
Which means that a box can contain many different types of roxes with different number for each rox type. The following property tracks how many Roxes are in the box per Rox type.

```
// The number of minted Rox NFTs per specific rox type (roxId)
pub var mintedNumberPerRox: {String: UInt32}
```

Each Rox NFT is a resource object containing the following fields.

```cadence
// NFT
// A Rox collectible as an NFT
pub resource NFT {

    // Global unique rox id
    pub let id: UInt64

    // Id of the box that the Rox comes from
    pub let boxId: UInt32

    // Specifies the Rox NFT collectible type
    pub let roxId: String

    // Specifies the rox tier: platinum, bronze, gold etc.
    pub let tier: String
        
    // The mint number for this specific rox type in the box
    pub let mintNumber: UInt32
}
```

The other types that are defined in `RoxContract` are as follows:

 - `BoxData`: A struct that contains constant information about boxes in Rox NFTs
    like the name, the id, metadata, minted numbers per rox.
 - `Box`: A resource that contains variable data for boxes 
    and the functionality to modify boxes,
    like locking the box and minting Roxes from the box.
 - `NFT`: A resource type that is the NFT that represents the Rox
    highlight a user owns. It stores its unique ID and other metadata. This
    is the collectible object that the users store in their accounts.
 - `Collection`: Similar to the `NFTCollection` resource from the NFT
    example, this resource is a repository for a user's Roxes.  Users can
    withdraw and deposit from this collection and get information about the 
    contained Roxes.
 - `Admin`: This is a resource type that can be used by admins to perform
    various actions in the smart contract like minting new boxes and roxes.

Metadata structs and methods are stored in the main smart contract
and can be queried by anyone. For example, to find out the box related info
`RoxContract.BoxData(boxId: id)` should be called. Or to get the info about any Rox stored in any user's account method `RoxContract.fetch(from: address, itemID: id)` can be called.

The power to create new boxes and roxes rests 
with the owner of the `Admin` resource.

Admins create boxes which are stored in the main smart contract,
Admins can mint roxes from the box. Only admins must have access to boxes.

Admins also can restrict the abilities of boxes to be further expanded.
A box begins as being unlocked, which means roxes can be added to it,
but when an admin locks the box, roxes can no longer be added to it. 
This cannot be reversed.

Once a user owns a Rox object, that Rox is stored directly 
in their account storage via their `Collection` object. The collection object
contains a dictionary that stores the Roxes and gives utility functions
to move them in and out and to read data about the collection and its Roxes.

## Testing

- `lib\go\contracts go generate`
- `lib\go\templates go generate`
- `lib\go\test go test`

## Deployment on emulator

- `flow project start-emulator`
- in another terminal `flow project deploy`

## Deployment on testnet

- `flow project deploy -f flow.json -f flow.testnet.json`

## Who do I talk to?

* Repo owner or admin
* Other community or team contact