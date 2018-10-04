// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//EIP-55 standard https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md
package state

import "encoding/hex"

//how eth SC addr is computed: https://ethereum.stackexchange.com/questions/760/how-is-the-address-of-an-ethereum-contract-computed

const AdrLen = 20

type Address [AdrLen]byte

func (adr *Address) Bytes() []byte {
	return adr[:]
}

// SetBytes sets the address to the value of b.
// If b is larger than len(a) it will panic.
func (adr *Address) SetBytes(b []byte) {
	if len(b) > len(adr) {
		b = b[len(b)-AdrLen:]
	}
	copy(adr[AdrLen-len(b):], b)
}

// Hex returns an EIP55-compliant hex string representation of the address.
func (adr *Address) Hex() string {
	unchecksummed := hex.EncodeToString(adr.Bytes())
	hash := DefHasher.Compute(unchecksummed)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}

func (adr *Address) String() string {
	return adr.Hex()
}

// IsHexAddress verifies whether a string can represent a valid hex-encoded
// Ethereum address or not.
func IsHexAddress(s string) bool {
	if hasHexPrefix(s) {
		s = s[2:]
	}
	return len(s) == 2*AdrLen && isHex(s)
}

// hasHexPrefix validates str begins with '0x' or '0X'.
func hasHexPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// HexToAddress returns Address with byte values of s.
// If s is larger than len(h), s will be cropped from the left.
func HexToAddress(s string) *Address {
	adr := Address{}
	adr.SetBytes(FromHex(s))

	return &adr
}

// FromHex returns the bytes represented by the hexadecimal string s.
// s may be prefixed with "0x".
func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}

	h, _ := hex.DecodeString(s)
	return h
}

func AddressFromPubKeyBytes(pubKey []byte) (*Address, error) {
	if len(pubKey) < AdrLen {
		return nil, NewErrorWrongSize(20, len(pubKey))
	}

	adr := Address{}
	adr.SetBytes(pubKey)
	return &adr, nil
}
