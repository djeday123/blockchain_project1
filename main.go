package main

import (
	"bytes"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/djeday123/blockchain1/core"
	"github.com/djeday123/blockchain1/crypto"
	"github.com/djeday123/blockchain1/network"
	"github.com/sirupsen/logrus"
)

// Server
// Transport => tcp. udp
// Block
// TX
// Keypair

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			//trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			if err := sendTransaction(trRemote, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	privKey := crypto.GeneratePrivateKey()
	opts := network.ServerOpts{
		PrivateKey: &privKey,
		ID:         "LOCAL",
		Transports: []network.Transport{trLocal},
	}

	s, err := network.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
}

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	dt := strconv.FormatInt(int64(rand.Intn(10000)), 10)
	//fmt.Println(dt)
	data := []byte(dt)
	tx := core.NewTransaction(data)
	tx.Sign(privKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())

	return tr.SendMessage(to, msg.Bytes())
}
