// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.genericComplexUnion -o generic_complex_union.go -n GenericComplexUnionMock -p tests

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// GenericComplexUnionMock implements genericComplexUnion
type GenericComplexUnionMock[T complexUnion] struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcName          func(t1 T)
	inspectFuncName   func(t1 T)
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mGenericComplexUnionMockName[T]
}

// NewGenericComplexUnionMock returns a mock for genericComplexUnion
func NewGenericComplexUnionMock[T complexUnion](t minimock.Tester) *GenericComplexUnionMock[T] {
	m := &GenericComplexUnionMock[T]{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NameMock = mGenericComplexUnionMockName[T]{mock: m}
	m.NameMock.callArgs = []*GenericComplexUnionMockNameParams[T]{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGenericComplexUnionMockName[T complexUnion] struct {
	mock               *GenericComplexUnionMock[T]
	defaultExpectation *GenericComplexUnionMockNameExpectation[T]
	expectations       []*GenericComplexUnionMockNameExpectation[T]

	callArgs []*GenericComplexUnionMockNameParams[T]
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// GenericComplexUnionMockNameExpectation specifies expectation struct of the genericComplexUnion.Name
type GenericComplexUnionMockNameExpectation[T complexUnion] struct {
	mock      *GenericComplexUnionMock[T]
	params    *GenericComplexUnionMockNameParams[T]
	paramPtrs *GenericComplexUnionMockNameParamPtrs[T]

	Counter uint64
}

// GenericComplexUnionMockNameParams contains parameters of the genericComplexUnion.Name
type GenericComplexUnionMockNameParams[T complexUnion] struct {
	t1 T
}

// GenericComplexUnionMockNameParamPtrs contains pointers to parameters of the genericComplexUnion.Name
type GenericComplexUnionMockNameParamPtrs[T complexUnion] struct {
	t1 *T
}

// Expect sets up expected params for genericComplexUnion.Name
func (mmName *mGenericComplexUnionMockName[T]) Expect(t1 T) *mGenericComplexUnionMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericComplexUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericComplexUnionMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.paramPtrs != nil {
		mmName.mock.t.Fatalf("GenericComplexUnionMock.Name mock is already set by ExpectParams functions")
	}

	mmName.defaultExpectation.params = &GenericComplexUnionMockNameParams[T]{t1}
	for _, e := range mmName.expectations {
		if minimock.Equal(e.params, mmName.defaultExpectation.params) {
			mmName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmName.defaultExpectation.params)
		}
	}

	return mmName
}

// ExpectT1Param1 sets up expected param t1 for genericComplexUnion.Name
func (mmName *mGenericComplexUnionMockName[T]) ExpectT1Param1(t1 T) *mGenericComplexUnionMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericComplexUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericComplexUnionMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.params != nil {
		mmName.mock.t.Fatalf("GenericComplexUnionMock.Name mock is already set by Expect")
	}

	if mmName.defaultExpectation.paramPtrs == nil {
		mmName.defaultExpectation.paramPtrs = &GenericComplexUnionMockNameParamPtrs[T]{}
	}
	mmName.defaultExpectation.paramPtrs.t1 = &t1

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the genericComplexUnion.Name
func (mmName *mGenericComplexUnionMockName[T]) Inspect(f func(t1 T)) *mGenericComplexUnionMockName[T] {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for GenericComplexUnionMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by genericComplexUnion.Name
func (mmName *mGenericComplexUnionMockName[T]) Return() *GenericComplexUnionMock[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericComplexUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericComplexUnionMockNameExpectation[T]{mock: mmName.mock}
	}

	return mmName.mock
}

// Set uses given function f to mock the genericComplexUnion.Name method
func (mmName *mGenericComplexUnionMockName[T]) Set(f func(t1 T)) *GenericComplexUnionMock[T] {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the genericComplexUnion.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the genericComplexUnion.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

func (mmName *mGenericComplexUnionMockName[T]) Times(n uint64) *mGenericComplexUnionMockName[T] {
	if n == 0 {
		mmName.mock.t.Fatalf("Times of GenericComplexUnionMock.Name mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmName.expectedInvocations, n)
	return mmName
}

func (mmName *mGenericComplexUnionMockName[T]) invocationsDone() bool {
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

// Name implements genericComplexUnion
func (mmName *GenericComplexUnionMock[T]) Name(t1 T) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName(t1)
	}

	mm_params := GenericComplexUnionMockNameParams[T]{t1}

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

		mm_got := GenericComplexUnionMockNameParams[T]{t1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.t1 != nil && !minimock.Equal(*mm_want_ptrs.t1, mm_got.t1) {
				mmName.t.Errorf("GenericComplexUnionMock.Name got unexpected parameter t1, want: %#v, got: %#v%s\n", *mm_want_ptrs.t1, mm_got.t1, minimock.Diff(*mm_want_ptrs.t1, mm_got.t1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmName.t.Errorf("GenericComplexUnionMock.Name got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmName.funcName != nil {
		mmName.funcName(t1)
		return
	}
	mmName.t.Fatalf("Unexpected call to GenericComplexUnionMock.Name. %v", t1)

}

// NameAfterCounter returns a count of finished GenericComplexUnionMock.Name invocations
func (mmName *GenericComplexUnionMock[T]) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of GenericComplexUnionMock.Name invocations
func (mmName *GenericComplexUnionMock[T]) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// Calls returns a list of arguments used in each call to GenericComplexUnionMock.Name.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmName *mGenericComplexUnionMockName[T]) Calls() []*GenericComplexUnionMockNameParams[T] {
	mmName.mutex.RLock()

	argCopy := make([]*GenericComplexUnionMockNameParams[T], len(mmName.callArgs))
	copy(argCopy, mmName.callArgs)

	mmName.mutex.RUnlock()

	return argCopy
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *GenericComplexUnionMock[T]) MinimockNameDone() bool {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.NameMock.invocationsDone()
}

// MinimockNameInspect logs each unmet expectation
func (m *GenericComplexUnionMock[T]) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GenericComplexUnionMock.Name with params: %#v", *e.params)
		}
	}

	afterNameCounter := mm_atomic.LoadUint64(&m.afterNameCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && afterNameCounter < 1 {
		if m.NameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GenericComplexUnionMock.Name")
		} else {
			m.t.Errorf("Expected call to GenericComplexUnionMock.Name with params: %#v", *m.NameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && afterNameCounter < 1 {
		m.t.Error("Expected call to GenericComplexUnionMock.Name")
	}

	if !m.NameMock.invocationsDone() && afterNameCounter > 0 {
		m.t.Errorf("Expected %d calls to GenericComplexUnionMock.Name but found %d calls",
			mm_atomic.LoadUint64(&m.NameMock.expectedInvocations), afterNameCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GenericComplexUnionMock[T]) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockNameInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GenericComplexUnionMock[T]) MinimockWait(timeout mm_time.Duration) {
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

func (m *GenericComplexUnionMock[T]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNameDone()
}
