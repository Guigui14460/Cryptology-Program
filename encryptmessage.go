package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

/*
Hash creation function
	Parameters:
	-----------
		key (string) : key to use

	Returns:
	--------
		string - created hash
*/
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

/*
Simple encryption function
	Parameters:
	-----------
		data ([]byte) : data to encrypt
		passPhrase (string) : string used to encrypt the data

	Returns:
	--------
		[]byte - encrypted data
*/
func encrypt(data []byte, passPhrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passPhrase)))
	aesgcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesgcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := aesgcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func main() {
	// Initialize variables
	var textToEncrypt, key string
	scanner := bufio.NewScanner(os.Stdin)

	// Get the positionned arguments
	args := os.Args[1:]

	if len(args) == 0 { // Open an interactive program
		print("Write your text to encrypt : ")
		scanner.Scan()
		textToEncrypt = scanner.Text()
		print("Write your key here : ")
		scanner.Scan()
		key = scanner.Text()
	} else { // Not use an interactive program
		// Verify if all the arguments are specified
		if len(args) != 2 {
			fmt.Printf("Usage : %s <msg> <key>", os.Args[0])
			os.Exit(2)
		}

		// Convert the arguments
		textToEncrypt = args[0]
		key = args[1]
	}

	// Encryption of the message with the key
	cipherText := encrypt([]byte(textToEncrypt), key)

	// Show the result
	fmt.Println(cipherText)
	if len(args) == 0 {
		print("Please press ENTER to quit the program ...")
		scanner.Scan()
	}
}
