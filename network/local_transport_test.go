package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra2 := tra.(*LocalTransport)
	trb2 := trb.(*LocalTransport)

	tra2.Connect(trb)
	trb2.Connect(tra)

	assert.Equal(t, tra2.peers[trb2.addr], trb2)
	assert.Equal(t, trb2.peers[tra2.addr], tra2)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A").(*LocalTransport)
	trb := NewLocalTransport("B").(*LocalTransport)

	tra.Connect(trb)
	trb.Connect(trb)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()
	buf := make([]byte, len(msg))
	n, err := rpc.Payload.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, n, len(msg))

	assert.Equal(t, buf, msg)
	assert.Equal(t, rpc.From, tra.addr)
}
