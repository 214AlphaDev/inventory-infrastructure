package value_objects_local

import "errors"

type Limit struct {
	limit uint32
}

func (c Limit) Limit() uint32 {
	return c.limit
}

func NewLimit(limit int32) (Limit, error) {

	if limit < 0 {
		return Limit{}, errors.New("invalid limit")
	}

	limitUint32 := uint32(limit)

	return Limit{
		limit: limitUint32,
	}, nil

}
