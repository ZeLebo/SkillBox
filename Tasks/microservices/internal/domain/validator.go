package domain

import "github.com/go-playground/validator/v10"

type Validator struct{}

func (v *Validator) ValidateCreate(r *Request) error {
	valid := validator.New()
	valid.RegisterStructValidation(v.createStructValidator, Request{})
	return valid.Struct(r)
}

func (v *Validator) ValidateChangeAge(r *Request) error {
	valid := validator.New()
	valid.RegisterStructValidation(v.ageStructValidator, Request{})
	return valid.Struct(r)
}

func (v *Validator) ValidateMakeFriends(r *Request) error {
	valid := validator.New()
	valid.RegisterStructValidation(v.makeFriendsStructValidator, Request{})
	return valid.Struct(r)
}

func (v *Validator) ValidateDeleteUser(r *Request) error {
	valid := validator.New()
	valid.RegisterStructValidation(v.deleteUserStructValidator, Request{})
	return valid.Struct(r)
}

func (v *Validator) ValidateGetFriends(r *Request) error {
	valid := validator.New()
	valid.RegisterStructValidation(v.getFriendsStructValidator, Request{})
	return valid.Struct(r)
}

func (v *Validator) deleteUserStructValidator(sl validator.StructLevel) {
	req := sl.Current().Interface().(Request)
	if req.TargetID < 0 {
		sl.ReportError(req.TargetID, "target_id", "target_id", "target_id", "")
	}
}

func (v *Validator) getFriendsStructValidator(sl validator.StructLevel) {
	req := sl.Current().Interface().(Request)
	if req.TargetID < 0 {
		sl.ReportError(req.TargetID, "target_id", "target_id", "target_id", "")
	}
}

func (v *Validator) makeFriendsStructValidator(sl validator.StructLevel) {
	req := sl.Current().Interface().(Request)
	if req.SourceID < 0 {
		sl.ReportError(req.SourceID, "source_id", "source_id", "source_id", "")
	}
	if req.TargetID < 0 {
		sl.ReportError(req.TargetID, "target_id", "target_id", "target_id", "")
	}
	if req.SourceID == req.TargetID {
		sl.ReportError(req.SourceID, "source_id", "source_id", "source_id", "")
	}
}

func (v *Validator) createStructValidator(sl validator.StructLevel) {
	req := sl.Current().Interface().(Request)
	if req.Age < 0 {
		sl.ReportError(req.Age, "age", "age", "age", "")
	}
	if req.Age > 200 {
		// Ну мало ли
		sl.ReportError(req.Age, "age", "age", "age", "")

	}
	if req.Name == "" {
		sl.ReportError(req.Name, "name", "name", "name", "")
	}
}

func (v *Validator) ageStructValidator(sl validator.StructLevel) {
	req := sl.Current().Interface().(Request)
	if req.Age < 0 {
		sl.ReportError(req.Age, "age", "age", "age", "")
	}
}
