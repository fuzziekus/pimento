package crypto

import (
	"io/ioutil"
	"log"
)

func (c *Cryptor) EncryptFile(inpath, outpath string) error {
	input, err := ioutil.ReadFile(inpath)
	if err != nil {
		log.Fatal(err)
	}
	result, err := c.Encrypt(string(input))
	if err != nil {
		log.Fatal(err)
	}
	return ioutil.WriteFile(outpath, result, 0644)
}

func (c *Cryptor) DecryptFile(inpath, outpath string) error {
	input, err := ioutil.ReadFile(inpath)
	if err != nil {
		log.Fatal(err)
	}
	result, err := c.Decrypt(string(input))
	if err != nil {
		log.Fatal(err)
	}
	return ioutil.WriteFile(outpath, result, 0755)
}
