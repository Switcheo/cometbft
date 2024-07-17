package utils

import (
	"fmt"

	"github.com/cometbft/cometbft/oracle/service/types"
)

// signature prefix for oracle votes is as such:
// index 0: accountType (if votes are signed by main val or oracle delegate)
// index 1: signType (type of key used: ed25519/sr25519/secp256k1)

func GetAccountSignTypeFromSignature(signature []byte) (accountType []byte, signType []byte) {
	return []byte{signature[0]}, []byte{signature[1]}
}

func FormSignaturePrefix(isSubAccount bool, signType string) ([]byte, error) {
	sigPrefix := []byte{}

	if isSubAccount {
		sigPrefix = append(sigPrefix, types.SubAccountSigPrefix...)
	} else {
		sigPrefix = append(sigPrefix, types.MainAccountSigPrefix...)
	}

	switch signType {
	case "ed25519":
		sigPrefix = append(sigPrefix, types.Ed25519SignType...)
	case "sr25519":
		sigPrefix = append(sigPrefix, types.Sr25519SignType...)
	case "secp256k1":
		sigPrefix = append(sigPrefix, types.Secp256k1SignType...)
	default:
		return nil, fmt.Errorf("FormSignaturePrefix: unsupported sign type: %v", signType)
	}

	return sigPrefix, nil
}
