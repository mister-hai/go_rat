package CryptographicFunctions

import (
	"crypto/sha256"

	"golang.org/x/crypto/chacha20poly1305"
)

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
			ErrorPrinter(derp, "generic error, fix me plz lol <3!")
			//return
		}
	}
	return DataOut, derp
}

// This function uses the Salsa20 to encrypt a byte field
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
