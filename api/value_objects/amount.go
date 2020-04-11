package value_objects_local

import "errors"

type Amount struct {
	amount uint32
}

func (c Amount) Amount() uint32 {
	return c.amount
}

func NewAmount(amount int32) (Amount, error) {

	if amount < 0 {
		return Amount{}, errors.New("invalid amount")
	}

	amountUint32 := uint32(amount)

	return Amount{
		amount: amountUint32,
	}, nil

}
