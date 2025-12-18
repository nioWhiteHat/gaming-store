package utils
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	
)
var masterKey = []byte("12345678901234567890123456789012") 

func Encrypt(plainText string) (string, error) {
	
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", err
	}

	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	 
    
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}


	ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}


func Decrypt(encryptedString string) (string, error) {

	data, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %v", err)
	}
	return string(plainText), nil
}