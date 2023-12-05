/**
 * @Author: kwens
 * @Date: 2022-10-17 16:21:28
 * @Description:
 */
package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/kwens/excalibur/time"
)

var DayFormatValidateKey = "dayFormat"

func DayFormatValidate(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	_, err := time.DayStr2Time(val)
	return err == nil
}
