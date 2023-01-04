package database

import (
	"context"
	"gorm.io/gorm/schema"
	"math/big"
	"reflect"
)

// BigIntSerializer allow us to use *big.Int in our GORM models
type BigIntSerializer struct{}

// Scan implements serializer interface
func (BigIntSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error {
	var val *big.Int
	strVal, ok := dbValue.(string)
	if ok {
		val, ok = new(big.Int).SetString(strVal, 10)
		if !ok {
			val = nil
		}
	}
	return field.Set(ctx, dst, val)
}

// Value implements serializer interface
func (BigIntSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	val, ok := fieldValue.(*big.Int)
	if !ok || val == nil {
		return nil, nil
	}
	return val.Text(10), nil
}
