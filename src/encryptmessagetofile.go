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

// Hash creation function
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

/*
Encryption function and write the encrypted data to a file
	Parameters:
	-----------
		filename (string) : the filename where write the date
		data ([]byte) : data to decrypt
		passPhrase (string) : string used to decrypt the data
*/
func encryptInFile(filename string, data []byte, passPhrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passPhrase))
}

func main() {
	// Initialize variables
	var filename, textToEncrypt, key string
	scanner := bufio.NewScanner(os.Stdin)

	// Get the positionned arguments
	args := os.Args[1:]

	if len(args) == 0 { // Open an interactive program
		print("Filename where to write the encrypted message : ")
		scanner.Scan()
		filename = scanner.Text()
		print("Write your text to encrypt : ")
		scanner.Scan()
		textToEncrypt = scanner.Text()
		print("Write your key here : ")
		scanner.Scan()
		key = scanner.Text()
	} else { // Not use an interactive program
		// Verify if all the arguments are specified
		if len(args) != 3 {
			fmt.Printf("Usage : %s <filename> <msg> <key>", os.Args[0])
			os.Exit(2)
		}

		// Convert the arguments
		filename = args[0]
		textToEncrypt = args[1]
		key = args[2]
	}

	// Encryption of the message with the key in a file
	encryptInFile(filename, []byte(textToEncrypt), key)
	if len(args) == 0 {
		print("Please press ENTER to quit the program ...")
		scanner.Scan()
	}
}
