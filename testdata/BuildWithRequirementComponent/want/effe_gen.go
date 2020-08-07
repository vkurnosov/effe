// Code generated by Effe. DO NOT EDIT.

//+build !effeinject

package main

import (
	"fmt"
)

func BuildComponent5(service BuildComponent5Service) BuildComponent5Func {
	return func() error {
		err := service.Step5()
		if err != nil {
			return err
		}
		return nil
	}
}
func BuildComponent3(service BuildComponent3Service) BuildComponent3Func {
	return func() error {
		err := service.Step3()
		if err != nil {
			return err
		}
		return nil
	}
}
func BuildComponent2(service BuildComponent2Service) BuildComponent2Func {
	return func() error {
		err := service.Step2()
		if err != nil {
			return err
		}
		err = service.BuildComponent3()
		if err != nil {
			return err
		}
		err = service.BuildComponent5()
		if err != nil {
			return err
		}
		return nil
	}
}
func BuildComponent6(service BuildComponent6Service) BuildComponent6Func {
	return func() error {
		err := service.Step6()
		if err != nil {
			return err
		}
		return nil
	}
}
func BuildComponent4(service BuildComponent4Service) BuildComponent4Func {
	return func() error {
		err := service.BuildComponent3()
		if err != nil {
			return err
		}
		err = service.BuildComponent5()
		if err != nil {
			return err
		}
		return nil
	}
}
func BuildComponent1(service BuildComponent1Service) BuildComponent1Func {
	return func() error {
		err2 := service.Step1()
		if err2 != nil {
			return err2
		}
		err2 = func() error {
			err := service.Before()
			if err != nil {
				err = service.Failure(err)
				return err
			}
			err = service.BuildComponent6()
			if err != nil {
				err = service.Failure(err)
				return err
			}
			err = service.BuildComponent6()
			if err != nil {
				err = service.Failure(err)
				return err
			}
			err = service.BuildComponent6()
			if err != nil {
				err = service.Failure(err)
				return err
			}
			err = service.BuildComponent6()
			if err != nil {
				err = service.Failure(err)
				return err
			}
			err = service.Success()
			if err != nil {
				err = service.Failure(err)
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
func NewBuildComponent5Impl() *BuildComponent5Impl {
	return &BuildComponent5Impl{step5FieldFunc: step5()}
}
func NewBuildComponent3Impl() *BuildComponent3Impl {
	return &BuildComponent3Impl{step3FieldFunc: step3()}
}
func NewBuildComponent2Impl(service BuildComponent3Service, service1 BuildComponent5Service) *BuildComponent2Impl {
	return &BuildComponent2Impl{BuildComponent3FieldFunc: BuildComponent3(service), BuildComponent5FieldFunc: BuildComponent5(service1), step2FieldFunc: step2()}
}
func NewBuildComponent6Impl() *BuildComponent6Impl {
	return &BuildComponent6Impl{step6FieldFunc: step6()}
}
func NewBuildComponent4Impl(service BuildComponent3Service, service1 BuildComponent5Service) *BuildComponent4Impl {
	return &BuildComponent4Impl{BuildComponent3FieldFunc: BuildComponent3(service), BuildComponent5FieldFunc: BuildComponent5(service1)}
}
func NewBuildComponent1Impl(service BuildComponent6Service) *BuildComponent1Impl {
	return &BuildComponent1Impl{beforeFieldFunc: before(), BuildComponent6FieldFunc: BuildComponent6(service), failureFieldFunc: failure(), step1FieldFunc: step1(), successFieldFunc: success()}
}

type BuildComponent5Service interface {
	Step5() error
}
type BuildComponent5Impl struct {
	step5FieldFunc func() error
}
type BuildComponent5Func func() error
type BuildComponent3Service interface {
	Step3() error
}
type BuildComponent3Impl struct {
	step3FieldFunc func() error
}
type BuildComponent3Func func() error
type BuildComponent2Service interface {
	BuildComponent3() error
	BuildComponent5() error
	Step2() error
}
type BuildComponent2Impl struct {
	BuildComponent3FieldFunc func() error
	BuildComponent5FieldFunc func() error
	step2FieldFunc           func() error
}
type BuildComponent2Func func() error
type BuildComponent6Service interface {
	Step6() error
}
type BuildComponent6Impl struct {
	step6FieldFunc func() error
}
type BuildComponent6Func func() error
type BuildComponent4Service interface {
	BuildComponent3() error
	BuildComponent5() error
}
type BuildComponent4Impl struct {
	BuildComponent3FieldFunc func() error
	BuildComponent5FieldFunc func() error
}
type BuildComponent4Func func() error
type BuildComponent1Service interface {
	Before() error
	BuildComponent6() error
	Failure(err error) error
	Step1() error
	Success() error
}
type BuildComponent1Impl struct {
	beforeFieldFunc          func() error
	BuildComponent6FieldFunc func() error
	failureFieldFunc         func(err error) error
	step1FieldFunc           func() error
	successFieldFunc         func() error
}
type BuildComponent1Func func() error

func (b *BuildComponent5Impl) Step5() error { return b.step5FieldFunc() }
func (b *BuildComponent3Impl) Step3() error { return b.step3FieldFunc() }
func (b *BuildComponent2Impl) BuildComponent3() error {
	return b.BuildComponent3FieldFunc()
}
func (b *BuildComponent2Impl) BuildComponent5() error {
	return b.BuildComponent5FieldFunc()
}
func (b *BuildComponent2Impl) Step2() error { return b.step2FieldFunc() }
func (b *BuildComponent6Impl) Step6() error { return b.step6FieldFunc() }
func (b *BuildComponent4Impl) BuildComponent3() error {
	return b.BuildComponent3FieldFunc()
}
func (b *BuildComponent4Impl) BuildComponent5() error {
	return b.BuildComponent5FieldFunc()
}
func (b *BuildComponent1Impl) Before() error { return b.beforeFieldFunc() }
func (b *BuildComponent1Impl) BuildComponent6() error {
	return b.BuildComponent6FieldFunc()
}
func (b *BuildComponent1Impl) Failure(err error) error { return b.failureFieldFunc(err) }
func (b *BuildComponent1Impl) Step1() error            { return b.step1FieldFunc() }
func (b *BuildComponent1Impl) Success() error          { return b.successFieldFunc() }
