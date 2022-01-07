// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "car-api/internal/core/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ICarRepository is an autogenerated mock type for the ICarRepository type
type ICarRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *ICarRepository) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ICarRepository) GetAll() ([]domain.Car, error) {
	ret := _m.Called()

	var r0 []domain.Car
	if rf, ok := ret.Get(0).(func() []domain.Car); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Car)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: id
func (_m *ICarRepository) GetOne(id uuid.UUID) (domain.Car, error) {
	ret := _m.Called(id)

	var r0 domain.Car
	if rf, ok := ret.Get(0).(func(uuid.UUID) domain.Car); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Car)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: car
func (_m *ICarRepository) Save(car domain.Car) error {
	ret := _m.Called(car)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Car) error); ok {
		r0 = rf(car)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, car
func (_m *ICarRepository) Update(id uuid.UUID, car domain.Car) error {
	ret := _m.Called(id, car)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, domain.Car) error); ok {
		r0 = rf(id, car)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
