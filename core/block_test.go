package core

import (
	_ "bytes"
	"fmt"
	"testing"
	"time"

	"github.com/djeday123/blockchain1/crypto"
	"github.com/djeday123/blockchain1/types"

	"github.com/stretchr/testify/assert"
)

func TestHashBlock(t *testing.T) {
	b := randomBlock(t, 0, types.Hash{})
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	b, err := NewBlock(header, []*Transaction{tx})
	assert.Nil(t, err)
	dataHash, err := CalculateDataHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.DataHash = dataHash
	assert.Nil(t, b.Sign(privKey))

	return b
}

// func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
// 	b := randomBlock(t, height, prevBlockHash)
// 	b.AddTransaction(tx)

// 	return b
// }

// func TestHeader_Encode_Decode(t *testing.T) {
// 	h := &Header{
// 		Version:   1,
// 		PrevBlock: types.RandomHash(),
// 		Timestamp: time.Now().UnixNano(),
// 		Height:    10,
// 		Nonce:     98934,
// 	}

// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, h.EncodeBinary(buf))

// 	hDecode := &Header{}
// 	assert.Nil(t, hDecode.DecodeBinary(buf))
// 	assert.Equal(t, h, hDecode)
// }

// func TestBlock_Encode_Decode(t *testing.T) {
// 	b := &Block{
// 		Header: Header{
// 			Version:   1,
// 			PrevBlock: types.RandomHash(),
// 			Timestamp: time.Now().UnixNano(),
// 			Height:    10,
// 			Nonce:     98934,
// 		},
// 		Transactions: nil,
// 	}

// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, b.EncodeBinary(buf))

// 	bDecode := &Block{}
// 	assert.Nil(t, bDecode.DecodeBinary(buf))
// 	assert.Equal(t, b, bDecode)

// }

// func TestBlockHash(t *testing.T) {
// 	b := &Block{
// 		Header: Header{
// 			Version:   1,
// 			PrevBlock: types.RandomHash(),
// 			Timestamp: time.Now().UnixNano(),
// 			Height:    10,
// 		},
// 		Transactions: []Transaction{},
// 	}

// 	h := b.Hash()
// 	fmt.Println(h)
// 	assert.False(t, h.IsZero())
// }
