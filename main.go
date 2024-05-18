package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gosuri/uilive"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

type Config struct {
	BSCNodeURL string `json:"bsc_node_url"`
	ETHNodeURL string `json:"eth_node_url"`
	ITERATION  int    `json:"iteration"`
	PATH       string `json:"path"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	return config, err
}
func main() {
	rand.Seed(time.Now().UnixNano())
	//var balance, balanceBSC, addrOne, mnemonics string
	writer := uilive.New()
	writer2 := writer.Newline()
	writer3 := writer.Newline()
	writer4 := writer.Newline()
	writer5 := writer.Newline()
	writer6 := writer.Newline()
	writer7 := writer.Newline()
	writer8 := writer.Newline()
	writer9 := writer.Newline()
	// start listening for updates and render
	writer.Start()
	var err error
	//var noncex, nonceBSCx uint64
	var configLd Config
	i := 0
	args := os.Args

	if (args[1] == "--config" || args[1] == "-C") && len(args) > 2 {
		configLd, err = LoadConfig(args[2])
		if err != nil {
			panic(err)
		}
		mnemonicsx := func(currentT int, addrInfo chan []any) {
			mnemonicEntropy, _ := bip39.NewEntropy(128)
			mnemonics, _ := bip39.NewMnemonic(mnemonicEntropy)

			addrOne, _ := deriveAddresses(mnemonics, configLd)
			balance, noncex, err := getEthBalance(addrOne, configLd)
			if err != nil {
				panic(err)
			}
			balanceBSC, nonceBSCx, err := getBSCBalance(addrOne, configLd)
			if err != nil {
				panic(err)
			}
			//fmt.Println("\r=========\n[*]Mnemonics:-", mnemonics, "\n[*] Address:-", addrOne, "\n[*] Balance:-", balance, "\n[*] Nonce Eth:-", noncex, "\n[*] Balance BSC:-", balanceBSC, "\n[*] Nonce BSC:-", nonceBSCx, "\n[*]Limit:-", configLd.ITERATION, "\n======")
			fmt.Fprintf(writer, "=========== Current Thread %d===========\n", currentT)
			fmt.Fprintf(writer2, "[*] Mnemonics:- %s\n", mnemonics)
			fmt.Fprintf(writer3, "[*] Address:- %s\n", addrOne)
			fmt.Fprintf(writer4, "[*] Balance Eth:- %s\n", balance)
			fmt.Fprintf(writer5, "[*] Nonce Eth:- %d\n", noncex)
			fmt.Fprintf(writer6, "[*] Balance BSC:- %s\n", balanceBSC)
			fmt.Fprintf(writer7, "[*] Nonce BSC:- %d\n", nonceBSCx)
			fmt.Fprintf(writer8, "[*] Limit:- (%d/%d)\n", i, configLd.ITERATION)
			fmt.Fprintf(writer9, "======================")
			addrInfo <- []any{mnemonics, addrOne, balance, noncex, balanceBSC, nonceBSCx, configLd.ITERATION}
		}

		for {
			bx := make(chan []any)

			go mnemonicsx(i, bx)
			info := <-bx
			f, _ := strconv.ParseFloat(info[2].(string), 64)
			f_f, _ := strconv.ParseFloat(info[4].(string), 64)
			if f > 0 || f_f > 0 || info[5].(uint64) > 0 || info[3].(uint64) > 0 || i > configLd.ITERATION {
				fmt.Fprintf(writer, "========== Found !!! ============")
				fmt.Fprintf(writer2, "[*] Mnemonics:- %s\n", info[0].(string))
				fmt.Fprintf(writer3, "[*] Address:- %s\n", info[1].(string))
				fmt.Fprintf(writer4, "[*] Balance Eth:- %s\n", info[2].(string))
				fmt.Fprintf(writer5, "[*] Nonce Eth:- %d\n", info[3].(uint64))
				fmt.Fprintf(writer6, "[*] Balance BSC:- %s\n", info[4].(string))
				fmt.Fprintf(writer7, "[*] Nonce BSC:- %d\n", info[5].(uint64))
				fmt.Fprintf(writer8, "[*] Limit:- %d\n", info[6].(int))
				fmt.Fprintf(writer9, "======================")
				break
			}
			i += 1
		}
	} else {

		fmt.Println("[*] Usage: ./eth-brute-force --config config.json")
		fmt.Println("[*] OR")
		fmt.Println("[*] Usage: ./eth-brute-force -C config.json")
	}
}
func deriveAddresses(mnemonics string, config Config) (string, string) {
	seed := bip39.NewSeed(mnemonics, "")
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath(config.PATH)
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	addressOne := account.Address.Hex()

	path = hdwallet.MustParseDerivationPath(config.PATH)
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	addressTwo := account.Address.Hex()
	return addressOne, addressTwo
}
func getEthBalance(address string, config Config) (string, uint64, error) {
	client, err := ethclient.Dial(config.ETHNodeURL)
	if err != nil {
		return "", 0, err
	}
	defer client.Close()

	addr := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return "", 0, err
	}

	balanceEth := new(big.Int)
	balanceEth.Set(balance)

	nnn, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return "", 0, err
	}
	return balanceEth.String(), nnn, nil
}
func getBSCBalance(address string, config Config) (string, uint64, error) {
	client, err := ethclient.Dial(config.BSCNodeURL)
	if err != nil {
		return "", 0, err
	}
	defer client.Close()

	addr := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return "", 0, err
	}

	balanceEth := new(big.Int)
	balanceEth.Set(balance) //SetBigInt(balance)
	nnnx, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return "", 0, err
	}
	return balanceEth.String(), nnnx, nil
}
