package test

import (
	"testing"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/onflow/flow-go-sdk/test"

	"github.com/onflow/flow-nft/lib/go/contracts"
	"github.com/onflow/flow-nft/lib/go/templates"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoxItemsDeployment(t *testing.T) {
	b := newBlockchain()

	// Should be able to deploy a contract as a new account with no keys.
	nftCode := contracts.NonFungibleToken()
	nftAddr, err := b.CreateAccount(nil,
		[]sdktemplates.Contract{
			{
				Name:   "NonFungibleToken",
				Source: string(nftCode),
			},
		},
	)
	if !assert.NoError(t, err) {
		t.Log(err.Error())
	}
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy a contract as a new account with no keys.
	tokenCode := contracts.RoxItems(nftAddr.String())
	_, err = b.CreateAccount(
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "RoxItems",
				Source: string(tokenCode),
			},
		},
	)
	if !assert.NoError(t, err) {
		t.Log(err.Error())
	}
	_, err = b.CommitBlock()
	assert.NoError(t, err)
}

func TestCreateNFT(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	env := templates.Environment{}

	// Should be able to deploy a contract as a new account with no keys.
	nftCode := contracts.NonFungibleToken()
	nftAddr, _ := b.CreateAccount(
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "NonFungibleToken",
				Source: string(nftCode),
			},
		},
	)

	env.NonFungibleTokenAddress = nftAddr.String()

	// First, deploy the contract
	tokenCode := contracts.RoxItems(nftAddr.String())
	tokenAccountKey, tokenSigner := accountKeys.NewWithSigner()
	tokenAddr, _ := b.CreateAccount(
		[]*flow.AccountKey{tokenAccountKey},
		[]sdktemplates.Contract{
			{
				Name:   "RoxItems",
				Source: string(tokenCode),
			},
		},
	)

	env.ContractAddress = tokenAddr.String()

	result := executeScriptAndCheck(t, b, templates.GenerateInspectNFTSupplyScript(env), nil)
	assert.Equal(t, cadence.NewUInt64(0), result)

	result = executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr))})
	assert.Equal(t, cadence.NewInt(0), result)

	t.Run("Should be able to mint a token", func(t *testing.T) {

		script := templates.GenerateMintNFTTransaction(env)
		tx := createTxWithTemplateAndAuthorizer(b, script, tokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(tokenAddr)) // Now it transfers to same minter account
		_ = tx.AddArgument(cadence.NewString("1"))
		_ = tx.AddArgument(cadence.NewString("1"))
		_ = tx.AddArgument(cadence.NewUInt64(1))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				tokenAddr,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				tokenSigner,
			},
			false,
		)

		// Assert that the account's collection is correct
		result = executeScriptAndCheck(t, b,
			templates.GenerateInspectCollectionScript(env),
			[][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr)), jsoncdc.MustEncode(cadence.UInt64(0))})
		assert.Equal(t, cadence.NewUInt64(0), result)

		result = executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr))})
		assert.Equal(t, cadence.NewInt(1), result)

		executeScriptAndCheck(t, b, templates.GenerateInspectNFTSupplyScript(env), nil)
		assert.Equal(t, cadence.NewInt(1), result)
	})

	t.Run("Shouldn't be able to borrow a reference to an NFT that doesn't exist", func(t *testing.T) {

		// Assert that the account's collection is correct
		result, err := b.ExecuteScript(
			templates.GenerateInspectCollectionScript(env),
			[][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr)), jsoncdc.MustEncode(cadence.UInt64(5))})
		require.NoError(t, err)
		assert.True(t, result.Reverted())
	})
}

func TestTransferNFT(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	// Should be able to deploy a contract as a new account with no keys.
	nftCode := contracts.NonFungibleToken()
	nftAddr, err := b.CreateAccount(
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "NonFungibleToken",
				Source: string(nftCode),
			},
		},
	)
	assert.NoError(t, err)

	// First, deploy the contract
	tokenCode := contracts.RoxItems(nftAddr.String())
	tokenAccountKey, tokenSigner := accountKeys.NewWithSigner()
	tokenAddr, err := b.CreateAccount(
		[]*flow.AccountKey{tokenAccountKey},
		[]sdktemplates.Contract{
			{
				Name:   "RoxItems",
				Source: string(tokenCode),
			},
		},
	)
	assert.NoError(t, err)

	env := templates.Environment{
		NonFungibleTokenAddress: nftAddr.String(),
		ContractAddress:         tokenAddr.String(),
	}

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	script := templates.GenerateMintNFTTransaction(env)
	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(100).
		SetProposalKey(
			b.ServiceKey().Address,
			b.ServiceKey().Index,
			b.ServiceKey().SequenceNumber,
		).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(tokenAddr)

	_ = tx.AddArgument(cadence.NewAddress(tokenAddr)) // Now it transfers to same minter account
	_ = tx.AddArgument(cadence.NewString("1"))
	_ = tx.AddArgument(cadence.NewString("1"))
	_ = tx.AddArgument(cadence.NewUInt64(1))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			tokenAddr,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			tokenSigner,
		},
		false,
	)

	// create a new Collection
	t.Run("Should be able to create a new empty NFT Collection", func(t *testing.T) {

		script := templates.GenerateCreateCollectionScript(env)
		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				joshSigner,
			},
			false,
		)

		result := executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), result)
	})

	t.Run("Shouldn't be able to withdraw an NFT that doesn't exist in a collection", func(t *testing.T) {

		script := templates.GenerateTransferScript(env)
		tx := createTxWithTemplateAndAuthorizer(b, script, tokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(cadence.NewString("3"))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				tokenAddr,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				tokenSigner,
			},
			true,
		)

		result := executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), result)

		// Assert that the account's collection is correct
		result = executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr))})
		assert.Equal(t, cadence.NewInt(1), result)

	})

	// transfer an NFT
	t.Run("Should be able to withdraw an NFT and deposit to another accounts collection", func(t *testing.T) {
		script := templates.GenerateTransferScript(env)
		tx := createTxWithTemplateAndAuthorizer(b, script, tokenAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(cadence.NewUInt64(0))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				tokenAddr,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				tokenSigner,
			},
			false,
		)

		// Assert that the account's collection is correct
		result := executeScriptAndCheck(t, b,
			templates.GenerateInspectCollectionScript(env),
			[][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress)), jsoncdc.MustEncode(cadence.UInt64(0))})
		assert.Equal(t, cadence.NewUInt64(0), result)

		result = executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		assert.Equal(t, cadence.NewInt(1), result)

		// Assert that the account's collection is correct
		result = executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr))})
		assert.Equal(t, cadence.NewInt(0), result)
	})

	// transfer an NFT
	t.Run("Should be able to withdraw an NFT and destroy it, not reducing the supply", func(t *testing.T) {

		tx := createTxWithTemplateAndAuthorizer(b, templates.GenerateDestroyScript(env), joshAddress)

		_ = tx.AddArgument(cadence.NewUInt64(0))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				joshSigner,
			},
			false,
		)

		result := executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), result)

		// Assert that the account's collection is correct
		result = executeScriptAndCheck(t, b, templates.GenerateInspectCollectionLenScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(tokenAddr))})
		assert.Equal(t, cadence.NewInt(0), result)

		result = executeScriptAndCheck(t, b, templates.GenerateInspectNFTSupplyScript(env), nil)
		assert.Equal(t, cadence.NewUInt64(1), result)
	})
}
