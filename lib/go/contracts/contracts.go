package contracts

//go:generate go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts

import (
	"strings"

	"github.com/onflow/flow-nft/lib/go/contracts/internal/assets"
)

const (
	nonfungibleTokenFilename       = "NonFungibleToken.cdc"
	roxContractFilename            = "RoxContract.cdc"
	defaultNonFungibleTokenAddress = "\"NonFungibleToken.cdc\""
)

// NonFungibleToken returns the NonFungibleToken contract interface.
func NonFungibleToken() []byte {
	return assets.MustAsset(nonfungibleTokenFilename)
}

// The returned contract will import the NonFungibleToken contract from the specified address.
func RoxContract(nonfungibleTokenAddr string) []byte {
	code := assets.MustAssetString(roxContractFilename)

	code = strings.ReplaceAll(
		code,
		defaultNonFungibleTokenAddress,
		"0x"+nonfungibleTokenAddr,
	)

	return []byte(code)
}
