package contracts

//go:generate go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts

import (
	"strings"

	"github.com/onflow/flow-nft/lib/go/contracts/internal/assets"
)

const (
	nonfungibleTokenFilename       = "NonFungibleToken.cdc"
	roxItemsFilename               = "RoxItems.cdc"
	defaultNonFungibleTokenAddress = "02"
)

// NonFungibleToken returns the NonFungibleToken contract interface.
func NonFungibleToken() []byte {
	return assets.MustAsset(nonfungibleTokenFilename)
}

// RoxItems returns the ExampleNFT contract.
//
// The returned contract will import the NonFungibleToken contract from the specified address.
func RoxItems(nonfungibleTokenAddr string) []byte {
	code := assets.MustAssetString(roxItemsFilename)

	code = strings.ReplaceAll(
		code,
		"0x"+defaultNonFungibleTokenAddress,
		"0x"+nonfungibleTokenAddr,
	)

	return []byte(code)
}
