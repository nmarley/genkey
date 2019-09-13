package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	// "samli.gitlab.com/dashx/dashx"
	// "samli.gitlab.com/dashx/dashutil"
)

var cmd = &cobra.Command{
	Use:   "genkey",
	Short: "",
	Long:  "",
	Run:   cmdFunc,
}

func cmdFunc(c *cobra.Command, inp []string) {
	i, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	h := sha256.New()
	h.Write(i)
	hash1 := h.Sum(nil)

	h = sha256.New()
	h.Write(hash1)
	hash2 := h.Sum(nil)

	fmt.Printf("%64x\n", hash2)
}

func main() {
	cmd.Execute()
}
