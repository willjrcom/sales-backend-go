package shiftentity

import "context"

type ShiftRepository interface {
	CreateShift(ctx context.Context, shift *Shift) (err error)
	UpdateShift(ctx context.Context, shift *Shift) (err error)
	DeleteShift(ctx context.Context, id string) (err error)
	GetShiftByID(ctx context.Context, id string) (shift *Shift, err error)
}