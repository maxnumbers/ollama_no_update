package lifecycle

import "context"

func StartBackgroundUpdaterChecker(ctx context.Context, cb func(string) error) {}
