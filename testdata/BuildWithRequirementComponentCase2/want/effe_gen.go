// Code generated by Effe. DO NOT EDIT.

//+build !effeinject

package main

import (
	"fmt"
)

func B(service BService) BFunc {
	return func() error {
		err := service.Step1()
		if err != nil {
			return err
		}
		return nil
	}
}
func C(service CService) CFunc {
	return func() error {
		err := service.B()
		if err != nil {
			return err
		}
		return nil
	}
}
func A(service AService) AFunc {
	return func() error {
		err := service.B()
		if err != nil {
			return err
		}
		err = service.C()
		if err != nil {
			return err
		}
		return nil
	}
}
func NewBImpl() *BImpl {
	return &BImpl{step1FieldFunc: step1()}
}
func NewCImpl(service BService) *CImpl {
	return &CImpl{BFieldFunc: B(service)}
}
func NewAImpl(service BService, service1 CService) *AImpl {
	return &AImpl{BFieldFunc: B(service), CFieldFunc: C(service1)}
}

type BService interface {
	Step1() error
}
type BImpl struct {
	step1FieldFunc func() error
}
type BFunc func() error
type CService interface {
	B() error
}
type CImpl struct {
	BFieldFunc func() error
}
type CFunc func() error
type AService interface {
	B() error
	C() error
}
type AImpl struct {
	BFieldFunc func() error
	CFieldFunc func() error
}
type AFunc func() error

func (b *BImpl) Step1() error { return b.step1FieldFunc() }
func (c *CImpl) B() error {
	return c.BFieldFunc()
}
func (a *AImpl) B() error {
	return a.BFieldFunc()
}
func (a *AImpl) C() error {
	return a.CFieldFunc()
}
