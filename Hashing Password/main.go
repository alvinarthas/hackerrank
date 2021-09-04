package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"

	simplescrypt "github.com/elithrar/simple-scrypt"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

const password = "foobarbazquux"

func test() {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	if err != nil {
		panic(err)
	}

	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 16)
	if err != nil {
		panic(err)
	}

	ahash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 1, 16)

	phash := pbkdf2.Key([]byte(password), salt, 100000, 16, sha512.New)

	fmt.Println("bhash", string(bhash))
	fmt.Println("shash", hex.EncodeToString(shash))
	fmt.Println("ahash", hex.EncodeToString(ahash))
	fmt.Println("phash", hex.EncodeToString(phash))
}

// e03bb2c3b31633b8fd929f40553df7fd

func main() {

	// e.g. r.PostFormValue("password")
	passwordFromForm := "dk12345"

	// Generates a derived key of the form "N$r$p$salt$dk" where N, r and p are defined as per
	// Colin Percival's scrypt paper: http://www.tarsnap.com/scrypt/scrypt.pdf
	// scrypt.Defaults (N=16384, r=8, p=1) makes it easy to provide these parameters, and
	// (should you wish) provide your own values via the scrypt.Params type.
	hash, err := simplescrypt.GenerateFromPassword([]byte(passwordFromForm), simplescrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	// Print the derived key with its parameters prepended.
	fmt.Printf("%s\n", hash)
	// fmt.Println(hex.EncodeToString(hash))

	newHash := []byte("16384$8$1$9dd8aefac3d125e5dd11fdae08cd32d1$dc6f8be9523a81a681efd9a274bf1ec9f45d44b390cd5d9bb5fc7b7a2b254e4d")

	// Uses the parameters from the existing derived key. Return an error if they don't match.
	err = simplescrypt.CompareHashAndPassword(newHash, []byte(passwordFromForm))
	if err != nil {
		log.Fatal(err)
	}
}
