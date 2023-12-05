/**
* @Author: kwens
 * @Date: 2022/7/14 17:17
 * @Description:
*/

package validate

import (
	"context"
	"sync"

	"github.com/go-playground/validator/v10"
)

type ValidatorFunc validator.Func
type StructHandle func(ctx context.Context) interface{}

type Validator interface {
	RegisterCustomerValidator(name string, f ValidatorFunc)                   // 注册自定义校验方法
	RegisterStructHandler(endpoint string, s StructHandle)                    // 注册接口入参校验
	Validate(ctx context.Context, endpoint string) (bool, ValidateFieldError) // 校验方法
	GetStruct(ctx context.Context, endpoint string) interface{}
}

type valid struct {
	vt                *validator.Validate
	customerValidator map[string]ValidatorFunc
	structHandler     map[string]StructHandle
}

func NewValid() Validator {
	return &valid{
		vt:                validator.New(),
		customerValidator: make(map[string]ValidatorFunc),
		structHandler:     make(map[string]StructHandle),
	}
}

func (v *valid) RegisterCustomerValidator(name string, f ValidatorFunc) {
	_ = v.vt.RegisterValidation(name, validator.Func(f))
}

func (v valid) RegisterStructHandler(endpoint string, stFunc StructHandle) {
	v.structHandler[endpoint] = stFunc
}

func (v valid) Validate(ctx context.Context, endpoint string) (bool, ValidateFieldError) {
	st := v.GetStruct(ctx, endpoint)
	if st == nil {
		return true, nil
	}
	if err := v.vt.Struct(st); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return false, ValidatorErrToFieldErr(ve)
		}
	}
	return true, nil
}

func (v valid) GetStruct(ctx context.Context, endpoint string) interface{} {
	stFunc, ok := v.structHandler[endpoint]
	if !ok {
		return nil
	}
	st := stFunc(ctx)
	return st
}

var Vd Validator

func Init() {
	var one sync.Once
	one.Do(func() {
		Vd = NewValid()
	})
	customerValidatorRegister(Vd)
}
