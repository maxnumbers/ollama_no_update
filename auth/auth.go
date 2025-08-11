package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
)

func GetPublicKey() (string, error) {
	return "", errors.New("remote authentication disabled")
}

func NewNonce(r io.Reader, length int) (string, error) {
	nonce := make([]byte, length)
	if _, err := io.ReadFull(r, nonce); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(nonce), nil
}

func Sign(ctx context.Context, bts []byte) (string, error) {
	return "", errors.New("remote authentication disabled")
}
