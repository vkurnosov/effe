// Code generated by Effe. DO NOT EDIT.

//+build !effeinject

package main

func A(service AService) AFunc {
	return func() error {
		err2 := service.Step1()
		if err2 != nil {
			return err2
		}
		err2 = func() error {
			err := service.BeforeStep()
			if err != nil {
				return err
			}
			err = service.Step2()
			if err != nil {
				return err
			}
			err = service.Step3()
			if err != nil {
				return err
			}
			err = service.SuccessStep()
			if err != nil {
				return err
			}
			return nil
		}()
		if err2 != nil {
			return err2
		}
		return nil
	}
}
func NewAImpl() *AImpl {
	return &AImpl{beforeStepFieldFunc: beforeStep(), step1FieldFunc: step1(), step2FieldFunc: step2(), step3FieldFunc: step3(), successStepFieldFunc: successStep()}
}

type AService interface {
	BeforeStep() error
	Step1() error
	Step2() error
	Step3() error
	SuccessStep() error
}
type AImpl struct {
	beforeStepFieldFunc  func() error
	step1FieldFunc       func() error
	step2FieldFunc       func() error
	step3FieldFunc       func() error
	successStepFieldFunc func() error
}
type AFunc func() error

func (a *AImpl) BeforeStep() error  { return a.beforeStepFieldFunc() }
func (a *AImpl) Step1() error       { return a.step1FieldFunc() }
func (a *AImpl) Step2() error       { return a.step2FieldFunc() }
func (a *AImpl) Step3() error       { return a.step3FieldFunc() }
func (a *AImpl) SuccessStep() error { return a.successStepFieldFunc() }