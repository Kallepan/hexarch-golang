package api

import (
	"context"

	"github.com/kallepan/hexarch-golang/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DbPort) *Adapter {
	return &Adapter{
		arith: arith,
		db:    db,
	}
}

func (a *Adapter) GetAddition(x, y int32) (int32, error) {
	result, err := a.arith.Addition(x, y)
	if err != nil {
		return 0, err
	}

	if err := a.db.AddToHistory(context.TODO(), x, y, result, "addition"); err != nil {
		return 0, err
	}

	return result, nil
}

func (a *Adapter) GetSubtraction(x, y int32) (int32, error) {
	result, err := a.arith.Subtraction(x, y)
	if err != nil {
		return 0, err
	}

	if err := a.db.AddToHistory(context.TODO(), x, y, result, "subtraction"); err != nil {
		return 0, err
	}

	return result, nil
}

func (a *Adapter) GetMultiplication(x, y int32) (int32, error) {
	result, err := a.arith.Multiplication(x, y)
	if err != nil {
		return 0, err
	}

	if err := a.db.AddToHistory(context.TODO(), x, y, result, "multiplication"); err != nil {
		return 0, err
	}

	return result, nil
}

func (a *Adapter) GetDivision(x, y int32) (int32, error) {
	result, err := a.arith.Division(x, y)
	if err != nil {
		return 0, err
	}

	if err := a.db.AddToHistory(context.TODO(), x, y, result, "division"); err != nil {
		return 0, err
	}

	return result, nil
}
