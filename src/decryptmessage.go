package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
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
Convert string to a byte array from input
	Parameters:
	-----------
		str (string) : string input (byte array into a string type)

	Returns:
	--------
		[]byte - byte data associated to string input
*/
func convertStringInputToBytes(str string) []byte {
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	str = strings.ReplaceAll(str, "{", "")
	str = strings.ReplaceAll(str, "}", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	splitedString := strings.Split(str, " ")
	byteArray := make([]byte, len(splitedString))
	for i, s := range splitedString {
		c, _ := strconv.ParseInt(s, 10, 64)
		byteArray[i] = byte(int(c))
	}
	return byteArray
}

func main() {
	// Initialize variables
	var key string
	var encryptedData []byte
	scanner := bufio.NewScanner(os.Stdin)

	// Get the positionned arguments
	args := os.Args[1:]

	if len(args) == 0 { // Open an interactive program
		print("Bytes data to decrypt (with or without square brackets) : ")
		scanner.Scan()
		encryptedData = convertStringInputToBytes(scanner.Text())
		print("Write your key here : ")
		scanner.Scan()
		key = scanner.Text()
	} else { // Not use an interactive program
		// Verify if all the arguments are specified
		if len(args) != 2 {
			fmt.Printf("Usage : %s <encrypted_byte_data> <key>", os.Args[0])
			os.Exit(2)
		}

		// Convert the arguments
		encryptedData = convertStringInputToBytes(args[0])
		key = args[1]
	}

	// Ecryption of the message with the key
	plaintext := decrypt([]byte(encryptedData), key)

	// Show the result
	fmt.Println(plaintext)
	if len(args) == 0 {
		print("Please press ENTER to quit the program ...")
		scanner.Scan()
	}
}
