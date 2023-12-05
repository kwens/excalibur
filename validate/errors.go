/**
 * @Author: kwens
 * @Date: 2022/7/14 17:40
 * @Description:
 */

package validate

import (
	"fmt"
	
	"github.com/go-playground/validator/v10"
)

type ValidateFieldError []FieldError

type FieldError struct {
	field      string
	value      string
	fileType   string
	validValue string
}

func (fe *FieldError) Error() string {
	return fmt.Sprintf("field validation error, value: <%s>|<%s> type:<%s> message:<%s>", fe.field, fe.value, fe.fileType, fe.validValue)
}

func ValidatorErrToFieldErr(ve validator.ValidationErrors) []FieldError {
	var fe = make([]FieldError, 0)
	for i := 0; i < len(ve); i++ {
		e := ve[i]
		fe = append(fe, FieldError{
			field:      e.Field(),
			value:      fmt.Sprintf("%v", e.Value()),
			fileType:   e.Tag(),
			validValue: e.Param(),
		})
	}
	return fe
}
