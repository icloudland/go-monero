package walletrpc

import (
	"crypto/rand"
	"encoding/hex"
)

// NewPaymentID64 generates a 64 bit payment ID (hex encoded).
// With 64 bit IDs, there is a non-negligible chance of a collision
// if they are randomly generated. It is up to recipients generating
// them to sanity check for uniqueness.
//
// 1 million IDs at 64-bit (simplified): 1,000,000^2 / (2^64 * 2) = ~1/36,893,488 so
// there is a 50% chance a collision happens around 5.06 billion IDs generated.
func NewPaymentID64() string {
	buf := make([]byte, 8)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

// NewPaymentID256 generates a 256 bit payment ID (hex encoded).
func NewPaymentID256() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}