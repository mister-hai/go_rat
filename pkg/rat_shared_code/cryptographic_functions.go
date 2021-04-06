/*/
This file contains the functions used for hashing and encrypting/decrypting
Text and files in both a form suitable for streaming connections and a form
suitable for individual entities.
/*/
package rat_shared_code

import (
	"crypto"
	"fmt"
	"io"
	"log"
	"os"
)

// might change parameters
// this function is for the encryption of files in one of four schemes
// some of this code is broken intentionally, if you are analyzing this
// section as a reviewer or developer, please provide input on
// clever ways to break it further
func Encrpyt_file(file_handle string, output_buffer []byte) {

}

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
	file_hash := crypto.SHA256.New()
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
