package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/dashx/dashutil"
	"gitlab.com/dashx/dashx/btcec"
	"gitlab.com/dashx/dashx/chaincfg"

	bls "gitlab.com/nmarley/go-bls-signatures"
)

// TODO: Add networks for Dash ECDSA keys (mainnet, testnet) -OR- even a flag
//       for known prefixes, kinda like BIP39 page.
// TODO: Add generation of BLS keys
// TODO: Add `fromsecret` for both BLS and ECDSA

var cmd = &cobra.Command{
	Use:   "genkey",
	Short: "",
	Long:  "",
	// Run:   blsGenKeyFunc,
	Run: cmdFunc,
}

// KeyPair represents an asynchronous cryptographic key pair
type KeyPair struct {
	SecretKey string `json:"secret"`
	PublicKey string `json:"public"`
}

// String implements the Stringer interface
func (kp KeyPair) String() string {
	jsonRepr, err := json.MarshalIndent(kp, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonRepr)
}

func blsGenKeyFunc(c *cobra.Command, inp []string) {
	sk, err := bls.RandKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	pk := sk.PublicKey()

	keypair := KeyPair{
		SecretKey: fmt.Sprintf("%64x", sk.Serialize()),
		PublicKey: fmt.Sprintf("%96x", pk.Serialize()),
	}
	fmt.Println(keypair)
}

func cmdFunc(c *cobra.Command, inp []string) {
	sk, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		panic(err)
	}

	var network *chaincfg.Params
	network = &chaincfg.MainNetParams
	// network = &chaincfg.TestNet3Params
	wif, err := dashutil.NewWIF(sk, network, true)
	if err != nil {
		panic(err)
	}

	addr, err := dashutil.NewAddressPubKey(wif.SerializePubKey(), network)
	if err != nil {
		panic(err)
	}

	keypair := KeyPair{
		SecretKey: wif.String(),
		PublicKey: addr.EncodeAddress(),
	}
	fmt.Println(keypair)
}

func main() {
	cmd.Execute()
}
