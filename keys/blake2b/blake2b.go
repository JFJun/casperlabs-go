package blake2b

import "golang.org/x/crypto/blake2b"

func Hash(data []byte) []byte {
	hash := blake2b.Sum256(data)
	return hash[:]
}
