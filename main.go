package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/dashx/dashutil"
	"gitlab.com/dashx/dashx/btcec"
	"gitlab.com/dashx/dashx/chaincfg"
)

// TODO: Add networks for Dash ECDSA keys (mainnet, testnet) -OR- even a flag
//       for known prefixes, kinda like BIP39 page.
// TODO: Add generation of BLS keys
// TODO: Add `fromsecret` for both BLS and ECDSA

var cmd = &cobra.Command{
	Use:   "genkey",
	Short: "",
	Long:  "",
	Run:   cmdFunc,
}

func cmdFunc(c *cobra.Command, inp []string) {
	sk, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		panic(err)
	}

	// wif, err := dashutil.NewWIF(sk, &chaincfg.TestNet3Params, true)
	wif, err := dashutil.NewWIF(sk, &chaincfg.MainNetParams, true)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("sk bytes = %64x\n", bytes)
	fmt.Println(wif.String())
}

func main() {
	cmd.Execute()
}
