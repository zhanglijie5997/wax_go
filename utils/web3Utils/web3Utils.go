package web3Utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
)

func  Web3Utils()  {
	//_keyStore := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	//_r, _err := _keyStore.NewAccount(password)
	//if _err != nil {
	//	panic(_err)
	//}
	//fmt.Println(_r.Address, "生成key")
	_r, err := ioutil.ReadFile("./wallet/UTC--2021-12-12T03-40-42.160162000Z--3ae45a5d734a002dd807fdd4c2daa877868a6bea")
	if err != nil {
		panic(err)
	}
	key, _err := keystore.DecryptKey(_r, password)
	if _err != nil {
		panic(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println(hexutil.Encode(pData), "读取")

	pbData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("pbData --->>>>", hexutil.Encode(pbData))
	fmt.Println("pbData to address---->>>>> ", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}