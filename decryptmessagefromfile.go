package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

// Hash creation function
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

/*
Simple decryption function
	Parameters:
	-----------
		data ([]byte) : data to decrypt
		passPhrase (string) : string used to decrypt the data (would be the same that used to encrypt the data)

	Returns:
	--------
		[]byte - decrypted data
*/
func decrypt(data []byte, passPhrase string) string {
	key := []byte(createHash(passPhrase))
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	nonceSize := aesgcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plaintext, _ := aesgcm.Open(nil, nonce, cipherText, nil)
	return string(plaintext)
}

/*
Decryption function and read the encrypted data from a file
	Parameters:
	-----------
		filename (string) : the filename where write the date
		passPhrase (string) : string used to decrypt the data (would be the same that used to encrypt the data)

	Returns:
	--------
		[]byte - decrypted data
*/
func decryptFromFile(filename string, passPhrase string) string {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passPhrase)
}

func main() {
	// Initialize variables
	var filename, key string
	scanner := bufio.NewScanner(os.Stdin)

	// Get the positionned arguments
	args := os.Args[1:]

	if len(args) == 0 { // Open an interactive program
		print("Filename where read the encrypted message : ")
		scanner.Scan()
		filename = scanner.Text()
		print("Write your key here : ")
		scanner.Scan()
		key = scanner.Text()
	} else { // Not use an interactive program
		// Verify if all the arguments are specified
		if len(args) != 2 {
			fmt.Printf("Usage : %s <filename> <key>", os.Args[0])
			os.Exit(2)
		}

		// Convert the arguments
		filename = args[0]
		key = args[1]
	}

	// Decryption of the message with the key in a file
	fmt.Println(decryptFromFile(filename, key))
	if len(args) == 0 {
		print("Please press ENTER to quit the program ...")
		scanner.Scan()
	}
}
