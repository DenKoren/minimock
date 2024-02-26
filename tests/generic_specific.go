package tests

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.genericSpecific -o generic_specific.go -n GenericSpecificMock -p tests

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"google.golang.org/protobuf/proto"
)

// GenericSpecificMock implements genericSpecific
type GenericSpecificMock[T proto.Message] struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcName          func(t1 T)
	inspectFuncName   func(t1 T)
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mGenericSpecificMockName[T]
}

// NewGenericSpecificMock returns a mock for genericSpecific
func NewGenericSpecificMock[T proto.Message](t minimock.Tester) *GenericSpecificMock[T] {
	m := &GenericSpecificMock[T]{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NameMock = mGenericSpecificMockName[T]{mock: m}
	m.NameMock.callArgs = []*GenericSpecificMockNameParams[T]{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGenericSpecificMockName[T proto.Message] struct {
	mock               *GenericSpecificMock[T]
	defaultExpectation *GenericSpecificMockNameExpectation[T]
	expectations       []*GenericSpecificMockNameExpectation[T]

	callArgs []*GenericSpecificMockNameParams[T]
	mutex    sync.RWMutex
}

// GenericSpecificMockNameExpectation specifies expectation struct of the genericSpecific.Name
type GenericSpecificMockNameExpectation[T proto.Message] struct {
	mock      *GenericSpecificMock[T]
	params    *GenericSpecificMockNameParams[T]
	paramPtrs *GenericSpecificMockNameParamPtrs[T]

	Counter uint64
}

// GenericSpecificMockNameParams contains parameters of the genericSpecific.Name
type GenericSpecificMockNameParams[T proto.Message] struct {
	t1 T
}

// GenericSpecificMockNameParamPtrs contains pointers to parameters of the genericSpecific.Name
type GenericSpecificMockNameParamPtrs[T proto.Message] struct {
	t1 *T
}

// Expect sets up expected params for genericSpecific.Name
func (mmName *mGenericSpecificMockName[T]) Expect(t1 T) *mGenericSpecificMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericSpecificMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericSpecificMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.paramPtrs != nil {
		mmName.mock.t.Fatalf("GenericSpecificMock.Name mock is already set by ExpectParams functions")
	}

	mmName.defaultExpectation.params = &GenericSpecificMockNameParams[T]{t1}
	for _, e := range mmName.expectations {
		if minimock.Equal(e.params, mmName.defaultExpectation.params) {
			mmName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmName.defaultExpectation.params)
		}
	}

	return mmName
}

// ExpectT1Param1 sets up expected param t1 for genericSpecific.Name
func (mmName *mGenericSpecificMockName[T]) ExpectT1Param1(t1 T) *mGenericSpecificMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericSpecificMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericSpecificMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.params != nil {
		mmName.mock.t.Fatalf("GenericSpecificMock.Name mock is already set by Expect")
	}

	if mmName.defaultExpectation.paramPtrs == nil {
		mmName.defaultExpectation.paramPtrs = &GenericSpecificMockNameParamPtrs[T]{}
	}
	mmName.defaultExpectation.paramPtrs.t1 = &t1

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the genericSpecific.Name
func (mmName *mGenericSpecificMockName[T]) Inspect(f func(t1 T)) *mGenericSpecificMockName[T] {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for GenericSpecificMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by genericSpecific.Name
func (mmName *mGenericSpecificMockName[T]) Return() *GenericSpecificMock[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericSpecificMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericSpecificMockNameExpectation[T]{mock: mmName.mock}
	}

	return mmName.mock
}

// Set uses given function f to mock the genericSpecific.Name method
func (mmName *mGenericSpecificMockName[T]) Set(f func(t1 T)) *GenericSpecificMock[T] {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the genericSpecific.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the genericSpecific.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

// Name implements genericSpecific
func (mmName *GenericSpecificMock[T]) Name(t1 T) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName(t1)
	}

	mm_params := GenericSpecificMockNameParams[T]{t1}

	// Record call args
	mmName.NameMock.mutex.Lock()
	mmName.NameMock.callArgs = append(mmName.NameMock.callArgs, &mm_params)
	mmName.NameMock.mutex.Unlock()

	for _, e := range mmName.NameMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmName.NameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmName.NameMock.defaultExpectation.Counter, 1)
		mm_want := mmName.NameMock.defaultExpectation.params
		mm_want_ptrs := mmName.NameMock.defaultExpectation.paramPtrs

		mm_got := GenericSpecificMockNameParams[T]{t1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.t1 != nil && !minimock.Equal(*mm_want_ptrs.t1, mm_got.t1) {
				mmName.t.Errorf("GenericSpecificMock.Name got unexpected parameter t1, want: %#v, got: %#v%s\n", *mm_want_ptrs.t1, mm_got.t1, minimock.Diff(*mm_want_ptrs.t1, mm_got.t1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmName.t.Errorf("GenericSpecificMock.Name got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmName.funcName != nil {
		mmName.funcName(t1)
		return
	}
	mmName.t.Fatalf("Unexpected call to GenericSpecificMock.Name. %v", t1)

}

// NameAfterCounter returns a count of finished GenericSpecificMock.Name invocations
func (mmName *GenericSpecificMock[T]) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of GenericSpecificMock.Name invocations
func (mmName *GenericSpecificMock[T]) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// Calls returns a list of arguments used in each call to GenericSpecificMock.Name.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmName *mGenericSpecificMockName[T]) Calls() []*GenericSpecificMockNameParams[T] {
	mmName.mutex.RLock()

	argCopy := make([]*GenericSpecificMockNameParams[T], len(mmName.callArgs))
	copy(argCopy, mmName.callArgs)

	mmName.mutex.RUnlock()

	return argCopy
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *GenericSpecificMock[T]) MinimockNameDone() bool {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	return true
}

// MinimockNameInspect logs each unmet expectation
func (m *GenericSpecificMock[T]) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GenericSpecificMock.Name with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		if m.NameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GenericSpecificMock.Name")
		} else {
			m.t.Errorf("Expected call to GenericSpecificMock.Name with params: %#v", *m.NameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		m.t.Error("Expected call to GenericSpecificMock.Name")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GenericSpecificMock[T]) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockNameInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GenericSpecificMock[T]) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *GenericSpecificMock[T]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNameDone()
}
