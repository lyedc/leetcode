package leetcode

import (
"bytes"
"crypto/aes"
"crypto/cipher"
	"encoding/hex"
	"fmt"
	"testing"
)

// 加密的key
var keySecret []byte = []byte{
	0x50,0x98,0xF0,0xB7,
	0xA6,0x62,0xFA,0xA5,
	0xA1,0x34,0x7E,0x1B,
	0xFB,0x5A,0xEF,0xBA,
}

func Test_main(t *testing.T){

	//jiami001 := []byte{0x61,0x61,0x61,0x61,
	//	0x61,0x61,0x61,0x61,
	//	0x61,0x61,0x61,0x61,
	//	0x61,0x61,0x61,0x61}
	jiamia := []byte("aaaaaaaaaaaaaaaa")

	//jiami := []byte{0x4d, 0xf6, 0x45, 0x86, 0x5b, 0x0e, 0xd5, 0x0c, 0x9f, 0xec, 0xf0, 0x45, 0x88, 0x39, 0x26, 0x5a}
	a := AesEncryptECB(jiamia, keySecret)
	fmt.Println(string(a))

	//NewAesDecryptCBC(jiami)
}

func NewAesDecryptCBC(encrypted []byte) {
	des := AesDecryptECB(encrypted, keySecret)
	fmt.Println(string(des))
	// return
}

func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	fmt.Println(length)
	unpadding := int(origData[length-1])
	fmt.Println(unpadding)
	fmt.Println("===============")
	return origData[:(length - unpadding)]
}

func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}


func stringTest(){
	input := []byte{
		0x32, 0x43, 0xf6, 0xa8,
		0x88, 0x5a, 0x30, 0x8d,
		0x31, 0x31, 0x98, 0xa2,
		0xe0, 0x37, 0x07, 0x34,
	}

	// 如果需要将字节切片转换为十六进制字符串
	hexInput := hex.EncodeToString(input)
	fmt.Printf("Hex input: %s\n", hexInput)

	// 如果需要将字节切片转换为字符串（假设内容是ASCII文本）
	// 注意：这可能会打印出不可见的字符
	stringInput := string(input)
	fmt.Printf("String input: %q\n", stringInput)
}

func TestStr(t *testing.T){
	stringTest()
}