// this file was generated by gomacro command: import _b "crypto/cipher"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/cipher"
)

// reflection: allow interpreted code to import "crypto/cipher"
func init() {
	Packages["crypto/cipher"] = Package{
	Binds: map[string]Value{
		"NewCBCDecrypter":	ValueOf(cipher.NewCBCDecrypter),
		"NewCBCEncrypter":	ValueOf(cipher.NewCBCEncrypter),
		"NewCFBDecrypter":	ValueOf(cipher.NewCFBDecrypter),
		"NewCFBEncrypter":	ValueOf(cipher.NewCFBEncrypter),
		"NewCTR":	ValueOf(cipher.NewCTR),
		"NewGCM":	ValueOf(cipher.NewGCM),
		"NewGCMWithNonceSize":	ValueOf(cipher.NewGCMWithNonceSize),
		"NewOFB":	ValueOf(cipher.NewOFB),
	}, Types: map[string]Type{
		"AEAD":	TypeOf((*cipher.AEAD)(nil)).Elem(),
		"Block":	TypeOf((*cipher.Block)(nil)).Elem(),
		"BlockMode":	TypeOf((*cipher.BlockMode)(nil)).Elem(),
		"Stream":	TypeOf((*cipher.Stream)(nil)).Elem(),
		"StreamReader":	TypeOf((*cipher.StreamReader)(nil)).Elem(),
		"StreamWriter":	TypeOf((*cipher.StreamWriter)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"AEAD":	TypeOf((*P_crypto_cipher_AEAD)(nil)).Elem(),
		"Block":	TypeOf((*P_crypto_cipher_Block)(nil)).Elem(),
		"BlockMode":	TypeOf((*P_crypto_cipher_BlockMode)(nil)).Elem(),
		"Stream":	TypeOf((*P_crypto_cipher_Stream)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for crypto/cipher.AEAD ---------------
type P_crypto_cipher_AEAD struct {
	Object	interface{}
	NonceSize_	func(interface{}) int
	Open_	func(_proxy_obj_ interface{}, dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
	Overhead_	func(interface{}) int
	Seal_	func(_proxy_obj_ interface{}, dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte
}
func (P *P_crypto_cipher_AEAD) NonceSize() int {
	return P.NonceSize_(P.Object)
}
func (P *P_crypto_cipher_AEAD) Open(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error) {
	return P.Open_(P.Object, dst, nonce, ciphertext, additionalData)
}
func (P *P_crypto_cipher_AEAD) Overhead() int {
	return P.Overhead_(P.Object)
}
func (P *P_crypto_cipher_AEAD) Seal(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte {
	return P.Seal_(P.Object, dst, nonce, plaintext, additionalData)
}

// --------------- proxy for crypto/cipher.Block ---------------
type P_crypto_cipher_Block struct {
	Object	interface{}
	BlockSize_	func(interface{}) int
	Decrypt_	func(_proxy_obj_ interface{}, dst []byte, src []byte) 
	Encrypt_	func(_proxy_obj_ interface{}, dst []byte, src []byte) 
}
func (P *P_crypto_cipher_Block) BlockSize() int {
	return P.BlockSize_(P.Object)
}
func (P *P_crypto_cipher_Block) Decrypt(dst []byte, src []byte)  {
	P.Decrypt_(P.Object, dst, src)
}
func (P *P_crypto_cipher_Block) Encrypt(dst []byte, src []byte)  {
	P.Encrypt_(P.Object, dst, src)
}

// --------------- proxy for crypto/cipher.BlockMode ---------------
type P_crypto_cipher_BlockMode struct {
	Object	interface{}
	BlockSize_	func(interface{}) int
	CryptBlocks_	func(_proxy_obj_ interface{}, dst []byte, src []byte) 
}
func (P *P_crypto_cipher_BlockMode) BlockSize() int {
	return P.BlockSize_(P.Object)
}
func (P *P_crypto_cipher_BlockMode) CryptBlocks(dst []byte, src []byte)  {
	P.CryptBlocks_(P.Object, dst, src)
}

// --------------- proxy for crypto/cipher.Stream ---------------
type P_crypto_cipher_Stream struct {
	Object	interface{}
	XORKeyStream_	func(_proxy_obj_ interface{}, dst []byte, src []byte) 
}
func (P *P_crypto_cipher_Stream) XORKeyStream(dst []byte, src []byte)  {
	P.XORKeyStream_(P.Object, dst, src)
}
