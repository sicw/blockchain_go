package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strconv"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := NewWallet()
	address := wallet.GetAddress()
	// 1NSD8JSKDURYEHHMZamcRfJnZcJvMURYhr
	fmt.Println(string(address))
}

func TestECDSA(t *testing.T) {
	// 生成ECDSA公私钥对
	curve := elliptic.P256() // 使用P-256曲线
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	// 对数据进行哈希
	message := "hello world"
	hash := sha256.Sum256([]byte(message))

	// 签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}

	// 验证签名
	isValid := ecdsa.Verify(&publicKey, hash[:], r, s)
	fmt.Printf("Signature is valid: %t\n", isValid)
}

func TestHashPubKey(t *testing.T) {
	// 看下写死的tx中 out是不是from的

	from := "12p47oV9vzSg3zKKg4uSLsRtEfSnSnwKar"
	wallets, _ := NewWallets("3002")
	wallet := wallets.GetWallet(from)
	pubHashkey := HashPubKey(wallet.PublicKey)
	// [19 222 71 52 233 144 224 99 72 206 109 185 109 219 199 250 177 149 244 152]
	// 整好与Height:4 coinbase vout匹配, 需要换个数据比较
	fmt.Printf("%d,",pubHashkey)
}

func TestSendCoin(t *testing.T) {
	// wallet3
	from := "12p47oV9vzSg3zKKg4uSLsRtEfSnSnwKar"
	// wallet2
	to := "1CgemWVkwShYZ2ABocMmpX8unoAEasv4Ux"
	nodeID := "3002"
	mineNow := true
	amount := 1
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		txs := []*Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}

	fmt.Println("Success!")
}

func TestGetBalance(t *testing.T) {
	address := "1KY1ZHxTe4NiXFW3MoCirVHrbxpJNwckF5"
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := NewBlockchain("3000")
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	balance := 0
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

func TestBytes(t *testing.T) {
	txid := []byte("�\u0013\u0012�|�\f��E��%N�q�(CTM\u0000�z\u001C�ơSʨ;")
	println(txid)
}

func TestPrintChain(t *testing.T) {
	bc := NewBlockchain("3002")
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			for i := range tx.ID {
				fmt.Printf("0x%x%s",tx.ID[i],",")
			}
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func TestOpenFile(t *testing.T) {
	db, err := bolt.Open("blockchain_3002.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.Close()
}
