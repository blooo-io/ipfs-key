package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

func main() {
	size := flag.Int("bitsize", 2048, "select the bitsize of the key to generate")
	typ := flag.String("type", "RSA", "select type of key to generate (RSA or Ed25519)")

	flag.Parse()

	var atyp int
	switch strings.ToLower(*typ) {
	case "rsa":
		atyp = ci.RSA
	case "ed25519":
		atyp = ci.Ed25519
	default:
		fmt.Fprintln(os.Stderr, "unrecognized key type: ", *typ)
		os.Exit(1)
	}

	priv, pub, err := ci.GenerateKeyPair(atyp, *size)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "%s ", pid.Pretty())

	data, err := priv.Bytes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Stdout.Write(data)
}
