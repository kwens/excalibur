/**
 * @Author: kwens
 * @Date: 2022-11-21 09:20:08
 * @Description:
 */
package validate

func customerValidatorRegister(vd Validator) {
	vd.RegisterCustomerValidator(DayFormatValidateKey, DayFormatValidate)
}

func CustomerValidatorRegister(vd Validator, key string, f ValidatorFunc) {
	vd.RegisterCustomerValidator(key, f)
}
