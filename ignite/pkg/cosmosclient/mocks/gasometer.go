// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	grpc "github.com/gogo/protobuf/grpc"
	mock "github.com/stretchr/testify/mock"

	tx "github.com/cosmos/cosmos-sdk/client/tx"

	types "github.com/cosmos/cosmos-sdk/types"

	typestx "github.com/cosmos/cosmos-sdk/types/tx"
)

// Gasometer is an autogenerated mock type for the Gasometer type
type Gasometer struct {
	mock.Mock
}

type Gasometer_Expecter struct {
	mock *mock.Mock
}

func (_m *Gasometer) EXPECT() *Gasometer_Expecter {
	return &Gasometer_Expecter{mock: &_m.Mock}
}

// CalculateGas provides a mock function with given fields: clientCtx, txf, msgs
func (_m *Gasometer) CalculateGas(clientCtx grpc.ClientConn, txf tx.Factory, msgs ...types.Msg) (*typestx.SimulateResponse, uint64, error) {
	_va := make([]interface{}, len(msgs))
	for _i := range msgs {
		_va[_i] = msgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, clientCtx, txf)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typestx.SimulateResponse
	if rf, ok := ret.Get(0).(func(grpc.ClientConn, tx.Factory, ...types.Msg) *typestx.SimulateResponse); ok {
		r0 = rf(clientCtx, txf, msgs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typestx.SimulateResponse)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(grpc.ClientConn, tx.Factory, ...types.Msg) uint64); ok {
		r1 = rf(clientCtx, txf, msgs...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(grpc.ClientConn, tx.Factory, ...types.Msg) error); ok {
		r2 = rf(clientCtx, txf, msgs...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Gasometer_CalculateGas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateGas'
type Gasometer_CalculateGas_Call struct {
	*mock.Call
}

// CalculateGas is a helper method to define mock.On call
//   - clientCtx grpc.ClientConn
//   - txf tx.Factory
//   - msgs ...types.Msg
func (_e *Gasometer_Expecter) CalculateGas(clientCtx interface{}, txf interface{}, msgs ...interface{}) *Gasometer_CalculateGas_Call {
	return &Gasometer_CalculateGas_Call{Call: _e.mock.On("CalculateGas",
		append([]interface{}{clientCtx, txf}, msgs...)...)}
}

func (_c *Gasometer_CalculateGas_Call) Run(run func(clientCtx grpc.ClientConn, txf tx.Factory, msgs ...types.Msg)) *Gasometer_CalculateGas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]types.Msg, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(types.Msg)
			}
		}
		run(args[0].(grpc.ClientConn), args[1].(tx.Factory), variadicArgs...)
	})
	return _c
}

func (_c *Gasometer_CalculateGas_Call) Return(_a0 *typestx.SimulateResponse, _a1 uint64, _a2 error) *Gasometer_CalculateGas_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

type mockConstructorTestingTNewGasometer interface {
	mock.TestingT
	Cleanup(func())
}

// NewGasometer creates a new instance of Gasometer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGasometer(t mockConstructorTestingTNewGasometer) *Gasometer {
	mock := &Gasometer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
