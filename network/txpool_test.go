package network

import (
	"testing"

	"github.com/djeday123/blockchain1/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("foo"))
	assert.Nil(t, p.Add(tx))
	assert.Equal(t, p.Len(), 1)

	tx2 := core.NewTransaction([]byte("foo"))
	assert.Nil(t, p.Add(tx2))
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, p.Len(), 0)
}
