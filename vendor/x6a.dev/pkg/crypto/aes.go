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

// EncryptCFB encrypts string to base64 crypto using AES-CFB
func EncryptCFB(text string) (string, error) {
	if len(Key) == 0 {
		return "", errors.New("no crypto key found")
	}

	plaintext := []byte(text)

	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function aes.NewCipher(Key)", errors.Trace())
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.Wrapf(err, "[%v] function io.ReadFull(rand.Reader, iv)", errors.Trace())
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptCFB decrypts from base64 to decrypted string
func DecryptCFB(cryptoText string) (string, error) {
	if len(Key) == 0 {
		return "", errors.New("no crypto key found")
	}

	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function aes.NewCipher(key)", errors.Trace())
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.Errorf("[%v] ciphertext too short", errors.Trace())
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}
