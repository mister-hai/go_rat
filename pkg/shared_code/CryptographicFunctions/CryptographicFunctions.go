/*/
This file contains the functions used for hashing and encrypting/decrypting
Text and files in both a form suitable for streaming connections and a form
suitable for individual entities.

# ECDSA recommendation key ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl req -x509 -nodes -newkey ec:secp384r1 -keyout server.ecdsa.key -out server.ecdsa.crt -days 3650
# openssl req -x509 -nodes -newkey ec:<(openssl ecparam -name secp384r1) -keyout server.ecdsa.key -out server.ecdsa.crt -days 3650
# -pkeyopt ec_paramgen_curve:… / ec:<(openssl ecparam -name …) / -newkey ec:…
ln -sf server.ecdsa.key server.key
ln -sf server.ecdsa.crt server.crt

# RSA recommendation key ≥ 2048-bit
openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
ln -sf server.rsa.key server.key
ln -sf server.rsa.crt server.crt

/*/
package CryptographicFunctions

import (
	"bytes"
	"compress/zlib"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"go_rat/pkg/shared_code/Core"
	"go_rat/pkg/shared_code/ErrorHandling"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*/ might change parameters
// this function is for the encryption of files in one of four schemes
// some of this code is broken intentionally, if you are analyzing this
// section as a reviewer or developer, please provide input on
// clever ways to break it further
/*/
func Encrypt_file(file_handle string, output_buffer []byte) {
	//file_bytes, derp := ioutil.ReadFile(file_handle)

}

// function to use zlib to compress a byte array
func ZCompress(input []byte) (herp []byte, derp error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
}

//function to use zlib to decompress a byte array
func ZDecompress(input []byte) (herp []byte, derp error) {
	b := bytes.NewReader(input)
	decrypted_bytes, derp := zlib.NewReader(b)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
		//return
	}
	io.Copy(herp, decrypted_bytes)
	decrypted_bytes.Close()
	return herp, derp
}

/*/
func (s *StreamEncrypter) Read(p []byte) (int, error) {
	n, readErr := s.Source.Read(p)
	if n > 0 {
		s.Stream.XORKeyStream(p[:n], p[:n])
		err := writeHash(s.Mac, p[:n])
		if err != nil {
			return n, ex.New(err)
		}
		return n, readErr
	}
	return 0, io.EOF

}
func Encrypt_file(file string, key string, output_buffer []byte) {
	encrypter, _ := StreamEncrypter(key, reader)
	io.Copy(file, encrypter)
}
/*/

/*
function to hash a string to compare against the hardcoded password
 never hardcode a password in plaintext
 we use the strongest we can and a good password...

 For the porpoises of this tutorial, we use a weak password.
*/

func Hash_auth_check(password string) {
	//Various Hashes, in order of increasing security
	// dont use this
	md5_password_hash := crypto.MD5.New()
	md5_password_hash.Sum([]byte(password))
	// or this
	sha1_password_hash := crypto.MD5SHA1.New()
	sha1_password_hash.Sum([]byte(password))
	// this is ok-ish, if you have a long password
	sha256_password_hash := crypto.SHA512_256.New()
	sha256_password_hash.Sum([]byte(password))
}

// the obvious, a plaintext password, hardcoded
func Insecure_password_check(password string) {

}

/*/
function to get the hash of a file for integrity checking
	create hash instance
		- this is a memory address we are going to shove a file into
	read the file from path
		- handle error if necessary
		- generic error printing
/*/
func File_hash(path string) []byte {
	file_hash := crypto.MD5.New()
	file_input, err := os.Open(path)
	if err != nil {
		// print the error
		fmt.Println(err)
	}
	// defer the closing of our File so that we can parse it later on
	defer file_input.Close()

	/*/
	     copy file buffer to hash compute buffer
		 the underscore character "_" is called a "blank identifier"
		 it allows you to ignore return values
		 in this case, we are acting like the regular return value
		 doesnt exist and if there is an error, log that error and exit
		 otherwise, finish copying from buffer to buffer
	    /*/
	if _, error := io.Copy(file_hash, file_input); error != nil {
		log.Fatal(error)
	}
	// and compute the hash of the file you provided to this function
	//file_hash_sha256 := file_hash.Sum(nil)
	return file_hash.Sum(nil)

}

var privateKey = loadPrivateKey()

func loadPrivateKey() *rsa.PrivateKey {
	key, err := ioutil.ReadFile("../../keygen/priv_key.pem")
	if err != nil {
		log.Fatalln("Cannot read PrivateKey:", err)
	}

	block, _ := pem.Decode(key)
	if block == nil {
		log.Fatalln("PrivateKey is not PEM format:", err)
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalln("Could not parse PrivateKey:", err)
	}

	return priv
}

// DecRsa decrypts RSA-encrypted data
func DecRsa(encData []byte) ([]byte, error) {
	rng := rand.Reader
	decData, err := rsa.DecryptOAEP(sha256.New(), rng, privateKey, encData, nil)
	if err != nil {
		log.Println("[!] Rsa:", err)
		return nil, err
	}

	return decData, nil
}

// DecAes decrypts data encrypted with AES
func DecAes(encData []byte, aeskey []byte) ([]byte, error) {
	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, encData[:12], encData[12:], nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// DecAsym decypts asymetric encryption (4096 bit RSA + AES)
func DecAsym(encData Core.AESPacket) ([]byte, error) {
	aeskey, err := DecRsa(encData.EncAesKey)
	if err != nil {
		return nil, err
	}

	return DecAes(encData.EncData, aeskey)
}
