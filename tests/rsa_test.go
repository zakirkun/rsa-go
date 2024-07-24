package tests

import (
	"encoding/base64"
	"rsa-go/ciphers"
	"testing"
)

func TestRsaGenerate(t *testing.T) {

	privKey, pubKey := ciphers.GenerateKeyPair(512)

	t.Logf("Public Key: %v", string(ciphers.PublicKeyToBytes(pubKey)))
	t.Logf("Private Key: %v", string(ciphers.PrivateKeyToBytes(privKey)))
}

func TestMessagesEncrypt(t *testing.T) {

	privKey, pubKey := ciphers.GenerateKeyPair(2048)

	t.Logf("Public Key: \n%v", string(ciphers.PublicKeyToBytes(pubKey)))

	plaintext := []byte("hello, world!")
	t.Logf("Before Encrypt: %v", string(plaintext))

	out := ciphers.EncryptWithPublicKey(plaintext, pubKey)
	t.Logf("After Encrypt: %v", string(out))

	t.Logf("Private Key: \n%v", string(ciphers.PrivateKeyToBytes(privKey)))

	// Decode Base64
	decode, _ := base64.StdEncoding.DecodeString(string(out))

	decrypted := ciphers.DecryptWithPrivateKey(decode, privKey)
	t.Logf("After Decrypted: %v", string(decrypted))

}
