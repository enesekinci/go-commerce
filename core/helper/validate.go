package helper

import (
	"github.com/go-playground/validator/v10"
	"go-commerce/core/constant"
	"go-commerce/database"
	"net/mail"
)

//var Validate *validator.Validate

var TagToErrorCode = map[string]constant.ErrorCode{
	"required": constant.REQUIRED,
	"email":    constant.EMAIL,
	"iscolor":  constant.NotColor,
	"min":      constant.MIN,
	"max":      constant.MAX,
	"boolean":  constant.BOOLEAN,
	"number":   constant.NUMBER,
	"numeric":  constant.NUMERIC,
}

func Validate() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("iscolor", IsColorValidation)
	return validate
}

func CustomErrorMessage(err validator.FieldError) interface{} {

	if code, ok := TagToErrorCode[err.Tag()]; ok {
		return code
	}

	//return fmt.Sprintf("Validation failed for %s", err.Field())
	return constant.INVALID
}

func IsColorValidation(fl validator.FieldLevel) bool {
	color := fl.Field().String()
	return len(color) == 7 && color[0] == '#'
}

func PrepareErrorMessages(err error, obj interface{}) map[string]interface{} {
	jsonColumns := GetJSONFieldNames(obj)
	validationErrors := make(map[string]interface{})

	for _, err := range err.(validator.ValidationErrors) {

		fieldName := err.Field()

		jsonColumn := jsonColumns[err.Field()]

		if jsonColumn != "" {
			fieldName = jsonColumns[err.Field()]
		}

		validationErrors[fieldName] = CustomErrorMessage(err)
	}

	return validationErrors
}

func ValidateStruct(inputStruct interface{}) interface{} {

	validate := Validate()

	err := validate.Struct(inputStruct)

	if err != nil {
		errors := PrepareErrorMessages(err, inputStruct)

		return errors
	}

	return nil
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsExistInDB(table string, column string, value string) bool {

	var isExist int64

	database.DB.Table(table).Where(column+" = ?", value).Limit(1).Count(&isExist)

	return isExist > 0
}
