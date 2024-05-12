// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.genericSimpleUnion -o generic_simple_union.go -n GenericSimpleUnionMock -p tests

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// GenericSimpleUnionMock implements genericSimpleUnion
type GenericSimpleUnionMock[T simpleUnion] struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcName          func(t1 T)
	inspectFuncName   func(t1 T)
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mGenericSimpleUnionMockName[T]
}

// NewGenericSimpleUnionMock returns a mock for genericSimpleUnion
func NewGenericSimpleUnionMock[T simpleUnion](t minimock.Tester) *GenericSimpleUnionMock[T] {
	m := &GenericSimpleUnionMock[T]{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NameMock = mGenericSimpleUnionMockName[T]{mock: m}
	m.NameMock.callArgs = []*GenericSimpleUnionMockNameParams[T]{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGenericSimpleUnionMockName[T simpleUnion] struct {
	mock               *GenericSimpleUnionMock[T]
	defaultExpectation *GenericSimpleUnionMockNameExpectation[T]
	expectations       []*GenericSimpleUnionMockNameExpectation[T]

	callArgs []*GenericSimpleUnionMockNameParams[T]
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// GenericSimpleUnionMockNameExpectation specifies expectation struct of the genericSimpleUnion.Name
type GenericSimpleUnionMockNameExpectation[T simpleUnion] struct {
	mock      *GenericSimpleUnionMock[T]
	params    *GenericSimpleUnionMockNameParams[T]
	paramPtrs *GenericSimpleUnionMockNameParamPtrs[T]

	Counter uint64
}

// GenericSimpleUnionMockNameParams contains parameters of the genericSimpleUnion.Name
type GenericSimpleUnionMockNameParams[T simpleUnion] struct {
	t1 T
}

// GenericSimpleUnionMockNameParamPtrs contains pointers to parameters of the genericSimpleUnion.Name
type GenericSimpleUnionMockNameParamPtrs[T simpleUnion] struct {
	t1 *T
}

// Expect sets up expected params for genericSimpleUnion.Name
func (mmName *mGenericSimpleUnionMockName[T]) Expect(t1 T) *mGenericSimpleUnionMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericSimpleUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericSimpleUnionMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.paramPtrs != nil {
		mmName.mock.t.Fatalf("GenericSimpleUnionMock.Name mock is already set by ExpectParams functions")
	}

	mmName.defaultExpectation.params = &GenericSimpleUnionMockNameParams[T]{t1}
	for _, e := range mmName.expectations {
		if minimock.Equal(e.params, mmName.defaultExpectation.params) {
			mmName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmName.defaultExpectation.params)
		}
	}

	return mmName
}

// ExpectT1Param1 sets up expected param t1 for genericSimpleUnion.Name
func (mmName *mGenericSimpleUnionMockName[T]) ExpectT1Param1(t1 T) *mGenericSimpleUnionMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericSimpleUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericSimpleUnionMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.params != nil {
		mmName.mock.t.Fatalf("GenericSimpleUnionMock.Name mock is already set by Expect")
	}

	if mmName.defaultExpectation.paramPtrs == nil {
		mmName.defaultExpectation.paramPtrs = &GenericSimpleUnionMockNameParamPtrs[T]{}
	}
	mmName.defaultExpectation.paramPtrs.t1 = &t1

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the genericSimpleUnion.Name
func (mmName *mGenericSimpleUnionMockName[T]) Inspect(f func(t1 T)) *mGenericSimpleUnionMockName[T] {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for GenericSimpleUnionMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by genericSimpleUnion.Name
func (mmName *mGenericSimpleUnionMockName[T]) Return() *GenericSimpleUnionMock[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericSimpleUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericSimpleUnionMockNameExpectation[T]{mock: mmName.mock}
	}

	return mmName.mock
}

// Set uses given function f to mock the genericSimpleUnion.Name method
func (mmName *mGenericSimpleUnionMockName[T]) Set(f func(t1 T)) *GenericSimpleUnionMock[T] {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the genericSimpleUnion.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the genericSimpleUnion.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

func (mmName *mGenericSimpleUnionMockName[T]) Times(n uint64) *mGenericSimpleUnionMockName[T] {
	if n == 0 {
		mmName.mock.t.Fatalf("Times of GenericSimpleUnionMock.Name mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmName.expectedInvocations, n)
	return mmName
}

func (mmName *mGenericSimpleUnionMockName[T]) invocationsDone() bool {
	if len(mmName.expectations) == 0 && mmName.defaultExpectation == nil && mmName.mock.funcName == nil {
		// does not need to check invocations if no expectations, defaultExpectation or funcName set
		return true
	}

	// if expectations were set we check total invocations
	// if default expectation was set then invocations count should be greater than zero
	// if func was set then invocations count should be greater than zero
	totalInvocations := mm_atomic.LoadUint64(&mmName.mock.afterNameCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmName.expectedInvocations)
	if totalInvocations < 1 || expectedInvocations != 0 && expectedInvocations != totalInvocations {
		return false
	}

	return true
}

// Name implements genericSimpleUnion
func (mmName *GenericSimpleUnionMock[T]) Name(t1 T) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName(t1)
	}

	mm_params := GenericSimpleUnionMockNameParams[T]{t1}

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

		mm_got := GenericSimpleUnionMockNameParams[T]{t1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.t1 != nil && !minimock.Equal(*mm_want_ptrs.t1, mm_got.t1) {
				mmName.t.Errorf("GenericSimpleUnionMock.Name got unexpected parameter t1, want: %#v, got: %#v%s\n", *mm_want_ptrs.t1, mm_got.t1, minimock.Diff(*mm_want_ptrs.t1, mm_got.t1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmName.t.Errorf("GenericSimpleUnionMock.Name got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmName.funcName != nil {
		mmName.funcName(t1)
		return
	}
	mmName.t.Fatalf("Unexpected call to GenericSimpleUnionMock.Name. %v", t1)

}

// NameAfterCounter returns a count of finished GenericSimpleUnionMock.Name invocations
func (mmName *GenericSimpleUnionMock[T]) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of GenericSimpleUnionMock.Name invocations
func (mmName *GenericSimpleUnionMock[T]) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// Calls returns a list of arguments used in each call to GenericSimpleUnionMock.Name.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmName *mGenericSimpleUnionMockName[T]) Calls() []*GenericSimpleUnionMockNameParams[T] {
	mmName.mutex.RLock()

	argCopy := make([]*GenericSimpleUnionMockNameParams[T], len(mmName.callArgs))
	copy(argCopy, mmName.callArgs)

	mmName.mutex.RUnlock()

	return argCopy
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *GenericSimpleUnionMock[T]) MinimockNameDone() bool {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.NameMock.invocationsDone()
}

// MinimockNameInspect logs each unmet expectation
func (m *GenericSimpleUnionMock[T]) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GenericSimpleUnionMock.Name with params: %#v", *e.params)
		}
	}

	afterNameCounter := mm_atomic.LoadUint64(&m.afterNameCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && afterNameCounter < 1 {
		if m.NameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GenericSimpleUnionMock.Name")
		} else {
			m.t.Errorf("Expected call to GenericSimpleUnionMock.Name with params: %#v", *m.NameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && afterNameCounter < 1 {
		m.t.Error("Expected call to GenericSimpleUnionMock.Name")
	}

	if !m.NameMock.invocationsDone() && afterNameCounter > 0 {
		m.t.Errorf("Expected %d calls to GenericSimpleUnionMock.Name but found %d calls",
			mm_atomic.LoadUint64(&m.NameMock.expectedInvocations), afterNameCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GenericSimpleUnionMock[T]) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockNameInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GenericSimpleUnionMock[T]) MinimockWait(timeout mm_time.Duration) {
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

func (m *GenericSimpleUnionMock[T]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNameDone()
}
