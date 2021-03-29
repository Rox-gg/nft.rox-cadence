package templates

import (
	"fmt"
	"strings"
)

const (
	defaultNFTAddress      = "0xNFTADDRESS"
	defaultContractAddress = "0xNFTCONTRACTADDRESS"
	defaultContractName    = "0xRoxItems"
	defaultContractStorage = "0xRoxItemsCollection"
)

type Environment struct {
	NonFungibleTokenAddress string
	ContractAddress         string
}

func uint32ToCadenceArr(nums []uint32) []byte {
	var s string
	for _, n := range nums {
		s += fmt.Sprintf("%d as UInt32, ", n)
	}
	// slice the last 2 characters off as that's the comma and the whitespace
	return []byte("[" + s[:len(s)-2] + "]")
}

func withHexPrefix(address string) string {
	if address == "" {
		return ""
	}

	if address[0:2] == "0x" {
		return address
	}

	return fmt.Sprintf("0x%s", address)
}

func replaceAddresses(code string, env Environment) string {

	code = strings.ReplaceAll(
		code,
		defaultNFTAddress,
		withHexPrefix(env.NonFungibleTokenAddress),
	)
	code = strings.ReplaceAll(
		code,
		defaultContractAddress,
		withHexPrefix(env.ContractAddress),
	)

	return code
}
