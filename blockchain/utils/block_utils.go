package utils

import (
	"crypto/sha256"

	pb "github.com/mjkong/selfstudy/blockchain/block"
)

func Hash(b *pb.Block) []byte {

	hash := sha256.Sum256(b.Bytes())
	return hash
}
