package server

import (
	"context"
	"errors"
)

type registryChallenge struct {
        Realm   string
        Service string
        Scope   string
}

func getAuthorizationToken(ctx context.Context, challenge registryChallenge) (string, error) {
	return "", errors.New("remote authentication disabled")
}
