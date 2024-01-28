package ports

import "context"

type DbPort interface {
	CloseDbConnection(ctx context.Context) error
	AddToHistory(ctx context.Context, x, y, result int32, operation string) error
}
