package utils

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

//==========================================
// Utility functions

// stringToKeccak256 converts a string to a keccak256 hash of type [32]byte
func StringToKeccak256(s string) [32]byte {
	var output [32]byte
	copy(output[:], crypto.Keccak256([]byte(s))[:])
	return output
}

// updateEnvFile updates our env file with a key-value pair
func UpdateEnvFile(k string, val string, envLoc string, myenv map[string]string) {
	myenv[k] = val
	err := godotenv.Write(myenv, envLoc)
	if err != nil {
		log.Printf("failed to update %s: %v\n", envLoc, err)
	}
}
