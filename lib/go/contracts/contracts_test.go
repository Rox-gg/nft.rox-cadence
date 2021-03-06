package contracts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-nft/lib/go/contracts"
)

const addrA = "0x0A"

func TestNonFungibleTokenContract(t *testing.T) {
	contract := contracts.NonFungibleToken()
	assert.NotNil(t, contract)
}

func TestRoxItemsContract(t *testing.T) {
	contract := contracts.RoxContract(addrA)
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA)
}
