// Copyright (C) 2019 x6a
//
// pkg is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// pkg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with pkg. If not, see <http://www.gnu.org/licenses/>.

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"x6a.dev/pkg/errors"
)

// Key is the crypto key needed to encrypt and decrypt.
// It should be 16 bytes (AES-128) or 32 bytes (AES-256).
var Key []byte

// Encrypt encrypts string to base64 crypto using AES-GCM
func Encrypt(text string) (string, error) {
	if len(Key) == 0 {
		return "", errors.New("no crypto key found")
	}

	plaintext := []byte(text)

	// generates a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function aes.NewCipher(Key)", errors.Trace())
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function cipher.NewGCM(block)", errors.Trace())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	// nonce := make([]byte, 12)
	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, aesgcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Wrapf(err, "[%v] function io.ReadFull(rand.Reader, nonce)", errors.Trace())
	}

	// encrypts the text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)

	//fmt.Printf("%x\n", ciphertext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts from base64 to decrypted string
func Decrypt(cryptoText string) (string, error) {
	if len(Key) == 0 {
		return "", errors.New("no crypto key found")
	}

	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function base64.URLEncoding.DecodeString(cryptoText)", errors.Trace())
	}

	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function aes.NewCipher(Key)", errors.Trace())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function cipher.NewGCM(block)", errors.Trace())
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("len(ciphertext) < nonceSize")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function aesgcm.Open()", errors.Trace())
	}

	// fmt.Printf("%s\n", plaintext)

	return fmt.Sprintf("%s", plaintext), nil
}
