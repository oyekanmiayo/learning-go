package main 

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"strings"
)

func CBCDecrypter(key string, ct string) string {  
	ciphertext, _ := hex.DecodeString(ct)  
	block, err := aes.NewCipher([]byte(key))  
	if err != nil {  
	   panic(err)  
	}  

	// CBC mode always works in whole blocks.  
   if len(ciphertext)%aes.BlockSize != 0 {  
	   panic("ciphertext is not a multiple of the block size")  
	}

	iv := []byte(key[:aes.BlockSize])
	mode := cipher.NewCBCDecrypter(block, iv)  

	// CryptBlocks can work in-place if the two arguments are the same.  
	mode.CryptBlocks(ciphertext, ciphertext)  
	s := string(ciphertext[:])  
   
	return s  
 }

func main()  {
	ciphertext := "f3662899b06ed7fcd8c3e2d0d5669ad3d60d6c979f125c6d0879a8d8ad461ee4ddfff3f1f3b93a4bf4d3f246b361b87a"
	key := "46952654613242965577169815403508"
	ret := CBCDecrypter(key, ciphertext)  
	fmt.Printf("CBC Decrypted Text:  %s\n", strings.TrimSpace(ret))

	ciphertext = "3a419f9e5031ee9f37c0e774cd9d0f6ed192f341c9cc7e11236a521f883b633d482e9773374708a3f0cb6adda53d156b99639dcd90a6385dec26590f8293ab6866490d257b5424bff5995fecde563225c983b9f47e0da9e4741d607b887ec91d76283e1f36a98941c70778b44eb905089c6fd92480b4db85b02c9f41175b569913752bd246533014a446a0ff6f3b40295f13aaff90bd522bac9f011a5ef71334477edf454c8300682aa9fbc76eedf74fac42d7b7d6b0da35b3408cba991b8fdb35f282f46675f4ed34e2257a7fba58e7ba9aeab9f73cc4bdde589e569b025280827338572a76c0c070a28a669e43cf67"
	for i := 0; i < 3; i++ {
		ciphertext = CBCDecrypter(key, ciphertext)  
	}
	fmt.Printf("CBC Decrypted Text:  %s\n", strings.TrimSpace(ciphertext))
}