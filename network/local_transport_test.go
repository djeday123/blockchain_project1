package network

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	//tra2 := tra.(*LocalTransport)
	//trb2 := trb.(*LocalTransport)

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()
	// buf := make([]byte, len(msg))
	// n, err := rpc.Payload.Read(buf)
	b, err := ioutil.ReadAll(rpc.Payload)
	assert.Nil(t, err)
	//assert.Equal(t, n, len(msg))
	assert.Equal(t, b, msg)
	assert.Equal(t, rpc.From, tra.addr)
}

func TestBroadcast(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	trc := NewLocalTransport("C")

	tra.Connect(trb)
	tra.Connect(trc)

	msg := []byte("foo")
	assert.Nil(t, tra.Broadcast(msg))

	rpcB := <-trb.consumeCh
	b, err := ioutil.ReadAll(rpcB.Payload)
	assert.Nil(t, err)
	assert.Equal(t, b, msg)

	rpcC := <-trc.consumeCh
	c, err := ioutil.ReadAll(rpcC.Payload)
	assert.Nil(t, err)
	assert.Equal(t, c, msg)
}
