package validator

import (
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
)

// 自定义 Validator

type MiaoValidator struct {
	Once     sync.Once
	Validate *validator.Validate // validator 实例
}

func NewMiaoValidator() *MiaoValidator {
	return &MiaoValidator{}
}

func (v *MiaoValidator) ValidateStruct(obj interface{}) error {
	// 反射检查 struct 类型
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.Validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *MiaoValidator) Engine() interface{} {
	v.lazyinit()
	return v.Validate
}

func (v *MiaoValidator) lazyinit() {
	// once 保证 validator 实例只被初始化一次
	v.Once.Do(func() {
		v.Validate = validator.New()
		v.Validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

