package core

import (
	_ "bytes"
	"djeday123/blockchain1/types"
	"fmt"
	"testing"
	"time"

	_ "github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestHashBlock(t *testing.T) {
	b := randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}

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
