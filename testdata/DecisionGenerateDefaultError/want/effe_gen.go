// Code generated by Effe. DO NOT EDIT.

//+build !effeinject

package main

import (
	"fmt"
)

func A(service AService) AFunc {
	return func() (int,

		error) {
		aVal3, err := service.Step1()
		if err != nil {
			return 0, err
		}
		intVal5, err := func(aVal2 a) (int, error) {
			switch aVal2 {
			case "a":
				intVal3 := func(aVal a) int {
					intVal := service.Step2(aVal)
					return intVal
				}(aVal2)
				return intVal3, nil
			case "":
				intVal4 := func() int {
					intVal2 := service.Step3()
					return intVal2
				}()
				return intVal4, nil
			default:
				return 0, fmt.Errorf("unsupported logic by aVal")
			}
		}(aVal3)
		if err != nil {
			return intVal5, err
		}
		return intVal5, nil
	}
}
func NewAImpl() *AImpl {
	return &AImpl{failureFieldFunc: failure(), step1FieldFunc: step1(), step2FieldFunc: step2(), step3FieldFunc: step3()}
}

type AService interface {
	Failure(err error) error
	Step1() (a, error)
	Step2(v a) int
	Step3() int
}
type AImpl struct {
	failureFieldFunc func(err error) error
	step1FieldFunc   func() (a, error)
	step2FieldFunc   func(v a) int
	step3FieldFunc   func() int
}
type AFunc func() (int,

	error)

func (a *AImpl) Failure(err error) error { return a.failureFieldFunc(err) }
func (a *AImpl) Step1() (a, error)       { return a.step1FieldFunc() }
func (a *AImpl) Step2(v a) int           { return a.step2FieldFunc(v) }
func (a *AImpl) Step3() int              { return a.step3FieldFunc() }
