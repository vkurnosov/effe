// Code generated by Effe. DO NOT EDIT.

//+build !effeinject

package main

func A(service AService) AFunc {
	return func() error {
		err := service.Step1()
		if err != nil {
			err = service.Failure(err)
			return err
		}
		err = service.Step2()
		if err != nil {
			err = service.Failure(err)
			return err
		}
		return nil
	}
}
func NewAImpl() *AImpl {
	return &AImpl{failureFieldFunc: failure(), step1FieldFunc: step1(), step2FieldFunc: step2()}
}

type AService interface {
	Failure(err error) error
	Step1() error
	Step2() error
}
type AImpl struct {
	failureFieldFunc func(err error) error
	step1FieldFunc   func() error
	step2FieldFunc   func() error
}
type AFunc func() error

func (a *AImpl) Failure(err error) error { return a.failureFieldFunc(err) }
func (a *AImpl) Step1() error            { return a.step1FieldFunc() }
func (a *AImpl) Step2() error            { return a.step2FieldFunc() }