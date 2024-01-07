// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.genericInlineUnion -o generic_inline_union.go -n GenericInlineUnionMock -p tests

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// GenericInlineUnionMock implements genericInlineUnion
type GenericInlineUnionMock[T int | float64] struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcName          func(t1 T)
	inspectFuncName   func(t1 T)
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mGenericInlineUnionMockName[T]
}

// NewGenericInlineUnionMock returns a mock for genericInlineUnion
func NewGenericInlineUnionMock[T int | float64](t minimock.Tester) *GenericInlineUnionMock[T] {
	m := &GenericInlineUnionMock[T]{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NameMock = mGenericInlineUnionMockName[T]{mock: m}
	m.NameMock.callArgs = []*GenericInlineUnionMockNameParams[T]{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGenericInlineUnionMockName[T int | float64] struct {
	mock               *GenericInlineUnionMock[T]
	defaultExpectation *GenericInlineUnionMockNameExpectation[T]
	expectations       []*GenericInlineUnionMockNameExpectation[T]

	callArgs []*GenericInlineUnionMockNameParams[T]
	mutex    sync.RWMutex
}

// GenericInlineUnionMockNameExpectation specifies expectation struct of the genericInlineUnion.Name
type GenericInlineUnionMockNameExpectation[T int | float64] struct {
	mock   *GenericInlineUnionMock[T]
	params *GenericInlineUnionMockNameParams[T]

	Counter uint64
}

// GenericInlineUnionMockNameParams contains parameters of the genericInlineUnion.Name
type GenericInlineUnionMockNameParams[T int | float64] struct {
	t1 T
}

// Expect sets up expected params for genericInlineUnion.Name
func (mmName *mGenericInlineUnionMockName[T]) Expect(t1 T) *mGenericInlineUnionMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInlineUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInlineUnionMockNameExpectation[T]{}
	}

	mmName.defaultExpectation.params = &GenericInlineUnionMockNameParams[T]{t1}
	for _, e := range mmName.expectations {
		if minimock.Equal(e.params, mmName.defaultExpectation.params) {
			mmName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmName.defaultExpectation.params)
		}
	}

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the genericInlineUnion.Name
func (mmName *mGenericInlineUnionMockName[T]) Inspect(f func(t1 T)) *mGenericInlineUnionMockName[T] {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for GenericInlineUnionMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by genericInlineUnion.Name
func (mmName *mGenericInlineUnionMockName[T]) Return() *GenericInlineUnionMock[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInlineUnionMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInlineUnionMockNameExpectation[T]{mock: mmName.mock}
	}

	return mmName.mock
}

// Set uses given function f to mock the genericInlineUnion.Name method
func (mmName *mGenericInlineUnionMockName[T]) Set(f func(t1 T)) *GenericInlineUnionMock[T] {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the genericInlineUnion.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the genericInlineUnion.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

// Name implements genericInlineUnion
func (mmName *GenericInlineUnionMock[T]) Name(t1 T) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName(t1)
	}

	mm_params := GenericInlineUnionMockNameParams[T]{t1}

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
		mm_got := GenericInlineUnionMockNameParams[T]{t1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmName.t.Errorf("GenericInlineUnionMock.Name got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmName.funcName != nil {
		mmName.funcName(t1)
		return
	}
	mmName.t.Fatalf("Unexpected call to GenericInlineUnionMock.Name. %v", t1)

}

// NameAfterCounter returns a count of finished GenericInlineUnionMock.Name invocations
func (mmName *GenericInlineUnionMock[T]) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of GenericInlineUnionMock.Name invocations
func (mmName *GenericInlineUnionMock[T]) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// Calls returns a list of arguments used in each call to GenericInlineUnionMock.Name.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmName *mGenericInlineUnionMockName[T]) Calls() []*GenericInlineUnionMockNameParams[T] {
	mmName.mutex.RLock()

	argCopy := make([]*GenericInlineUnionMockNameParams[T], len(mmName.callArgs))
	copy(argCopy, mmName.callArgs)

	mmName.mutex.RUnlock()

	return argCopy
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *GenericInlineUnionMock[T]) MinimockNameDone() bool {
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
func (m *GenericInlineUnionMock[T]) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GenericInlineUnionMock.Name with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		if m.NameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GenericInlineUnionMock.Name")
		} else {
			m.t.Errorf("Expected call to GenericInlineUnionMock.Name with params: %#v", *m.NameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		m.t.Error("Expected call to GenericInlineUnionMock.Name")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GenericInlineUnionMock[T]) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockNameInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GenericInlineUnionMock[T]) MinimockWait(timeout mm_time.Duration) {
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

func (m *GenericInlineUnionMock[T]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNameDone()
}
