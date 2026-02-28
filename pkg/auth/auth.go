package auth

import "context"

type Auth struct {
	StaffID    string
	PositionID string
}

type staffIDKeyType struct{}
type positionIDKeyType struct{}

var (
	staffIDKey    = staffIDKeyType{}
	positionIDKey = positionIDKeyType{}
)

func WithStaffID(ctx context.Context, staffID string) context.Context {
	return context.WithValue(ctx, staffIDKey, staffID)
}

func WithPositionID(ctx context.Context, positionID string) context.Context {
	return context.WithValue(ctx, positionIDKey, positionID)
}

func WithAuth(ctx context.Context, a Auth) context.Context {
	ctx = WithStaffID(ctx, a.StaffID)
	ctx = WithPositionID(ctx, a.PositionID)
	return ctx
}

func GetAuth(ctx context.Context) Auth {
	staffID, _ := ctx.Value(staffIDKey).(string)
	positionID, _ := ctx.Value(positionIDKey).(string)
	return Auth{StaffID: staffID, PositionID: positionID}
}

func GetStaffID(ctx context.Context) (string, bool) {
	staffID, ok := ctx.Value(staffIDKey).(string)
	return staffID, ok
}

func GetPositionID(ctx context.Context) (string, bool) {
	positionID, ok := ctx.Value(positionIDKey).(string)
	return positionID, ok
}
