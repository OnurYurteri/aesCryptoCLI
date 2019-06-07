package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()
var IV = []byte("PU9oAWglHBYnMbTL")
var aesKeyPath string
var inputFile string
var outputFile string

func info() {
	app.Name = "Aes Crypto CLI"
	app.Usage = "Encrypt files with Advanced Encryption Standard (AES)"
	app.Author = "Onur Yurteri"
	app.Version = "1.0.0"
}

func flags() {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "k",
			Usage:       "Path for key file",
			Destination: &aesKeyPath,
		},
		cli.StringFlag{
			Name:        "i",
			Usage:       "Path for input file",
			Destination: &inputFile,
		},
		cli.StringFlag{
			Name:        "o",
			Usage:       "Path for output file",
			Destination: &outputFile,
		},
	}
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "createKey",
			Aliases: []string{"c"},
			Usage:   "Create AES Key",
			Action:  createAESKey,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run Encrypt/Decrpyt operation for given file",
			Action:  encryptFile,
		},
		/* {
			Name:    "decryptFile",
			Aliases: []string{"d"},
			Usage:   "Decrypt given file",
			Action:  decryptFile,
		}, */
	}
}

func getCipher(c *cli.Context) cipher.Block {
	aesFile, err := ioutil.ReadFile(c.GlobalString("k"))
	fmt.Println(c.GlobalString("k"))
	if err != nil {
		log.Fatalf("Can't read key file!")
	}
	keyBlock, _ := pem.Decode(aesFile)

	aesCipher, err := aes.NewCipher(keyBlock.Bytes)
	if err != nil {
		log.Fatalf("Can't create AES Cipher!")
	}
	return aesCipher
}

func createAESKey(c *cli.Context) {
	genkey := make([]byte, 32)
	_, err := rand.Read(genkey)

	if err != nil {
		log.Fatalf("Can't generate key!")
	}

	block := &pem.Block{
		Type:  "AES KEY",
		Bytes: genkey,
	}

	if c.NArg() > 0 {
		err := ioutil.WriteFile(c.Args().First()+".key", pem.EncodeToMemory(block), 0644)
		if err != nil {
			log.Fatalf("Can't write key file!")
		}
	} else {
		err := ioutil.WriteFile("aes.key", pem.EncodeToMemory(block), 0644)
		if err != nil {
			log.Fatalf("Can't write key file!")
		}
	}

	fmt.Println("AES KEY has been created!")
}

func encryptFile(c *cli.Context) {

	if c.GlobalIsSet("i") && c.GlobalIsSet("k") && c.GlobalIsSet("o") {
		file, err := ioutil.ReadFile(c.GlobalString("i"))
		fmt.Println(c.GlobalString("i"))
		if err != nil {
			log.Fatalf("Can't read input file!")
		}

		aesCipher := getCipher(c)

		stream := cipher.NewCTR(aesCipher, IV)
		stream.XORKeyStream(file, file)
		errW := ioutil.WriteFile(fmt.Sprintf(c.GlobalString("o")), file, 0644)
		if errW != nil {
			log.Fatalf("Can't write encrypted file!")
		} else {
			fmt.Printf("File encrypted/decrypted as: %s\n\n", c.GlobalString("o"))
		}

	} else {
		log.Fatalf("Input flags '--i, --o, --k' must be set! Use --help command for help.")
	}
}

/* func decryptFile(c *cli.Context) {
	if c.GlobalIsSet("i") && c.GlobalIsSet("k") && c.GlobalIsSet("o") {
		encryptedFile, err := ioutil.ReadFile(c.GlobalString("i"))
		if err != nil {
			log.Fatal("Can't read encrypted file")
		}
		aesCipher := getCipher(c)
		stream := cipher.NewCTR(aesCipher, IV)
		stream.XORKeyStream(encryptedFile, encryptedFile)
		errW := ioutil.WriteFile(fmt.Sprintf(c.GlobalString("o")), encryptedFile, 0644)
		if errW != nil {
			log.Fatalf("Writing encryption file")
		} else {
			fmt.Printf("Message encrypted in file: %s\n\n", c.GlobalString("o"))
		}
	} else {
		log.Fatalf("Input flags must be set!")
	}

} */

func main() {
	info()
	flags()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
