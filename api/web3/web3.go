package web3

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	os2 "go_study/utils/os"
	"go_study/utils/web3Utils"
	"math"
	"math/big"
)

var ganacheUrl string = "http://127.0.0.1:8545"
func Web3Api()  {
	priviteKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	_ecdsa := crypto.FromECDSA(priviteKey)
	fmt.Println(hexutil.Encode(_ecdsa), "priviteKey")
	puData := crypto.FromECDSAPub(&priviteKey.PublicKey)
	fmt.Println(hexutil.Encode(puData), "puData")
	fmt.Println(crypto.PubkeyToAddress(priviteKey.PublicKey), "publicKey to address")
	ctx := context.Background()
	web3Utils.Web3Utils()
	debug := os2.YamlResult.Web3.Debug
	fmt.Println(debug)
	// 建立web3链接
	client, err := ethclient.DialContext(ctx, debug.HttpUrl)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		panic(err)
	}
	addr := "0x8035c3389C17b592E07ab661aB3c39a096549ec1"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		panic(err)
	}
	fbBlance := new(big.Float)
	fbBlance.SetString(balance.String())
	balanceEther := new(big.Float).Quo(fbBlance, big.NewFloat(math.Pow10(18)))
	fmt.Println("账户：", balanceEther, "\n产出块",block.Number())
}
