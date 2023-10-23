package core

import (
	"bytes"
	"djeday123/blockchain1/crypto"
	"djeday123/blockchain1/types"
	"encoding/gob"
	"io"
)

type Header struct {
	Version       uint32
	DataHsh       types.Hash
	PrevBlockHash types.Hash
	Timestamp     int64
	Height        uint32
	//Nonce     uint64
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	//Cached version of the header hash
	hash types.Hash
}

func NewBlock(h *Header, tx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: tx,
	}
}

func (b *Block) Sign(privKey crypto.PrivateKey) *crypto.Signature {

}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}
