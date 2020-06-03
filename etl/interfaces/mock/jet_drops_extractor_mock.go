package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/block-explorer/etl/types"
)

// JetDropsExtractorMock implements interfaces.JetDropsExtractor
type JetDropsExtractorMock struct {
	t minimock.Tester

	funcGetJetDrops          func(ctx context.Context) (ch1 <-chan *types.PlatformJetDrops)
	inspectFuncGetJetDrops   func(ctx context.Context)
	afterGetJetDropsCounter  uint64
	beforeGetJetDropsCounter uint64
	GetJetDropsMock          mJetDropsExtractorMockGetJetDrops

	funcLoadJetDrops          func(ctx context.Context, fromPulseNumber int, toPulseNumber int) (err error)
	inspectFuncLoadJetDrops   func(ctx context.Context, fromPulseNumber int, toPulseNumber int)
	afterLoadJetDropsCounter  uint64
	beforeLoadJetDropsCounter uint64
	LoadJetDropsMock          mJetDropsExtractorMockLoadJetDrops

	funcStart          func(ctx context.Context) (err error)
	inspectFuncStart   func(ctx context.Context)
	afterStartCounter  uint64
	beforeStartCounter uint64
	StartMock          mJetDropsExtractorMockStart

	funcStop          func(ctx context.Context) (err error)
	inspectFuncStop   func(ctx context.Context)
	afterStopCounter  uint64
	beforeStopCounter uint64
	StopMock          mJetDropsExtractorMockStop
}

// NewJetDropsExtractorMock returns a mock for interfaces.JetDropsExtractor
func NewJetDropsExtractorMock(t minimock.Tester) *JetDropsExtractorMock {
	m := &JetDropsExtractorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetJetDropsMock = mJetDropsExtractorMockGetJetDrops{mock: m}
	m.GetJetDropsMock.callArgs = []*JetDropsExtractorMockGetJetDropsParams{}

	m.LoadJetDropsMock = mJetDropsExtractorMockLoadJetDrops{mock: m}
	m.LoadJetDropsMock.callArgs = []*JetDropsExtractorMockLoadJetDropsParams{}

	m.StartMock = mJetDropsExtractorMockStart{mock: m}
	m.StartMock.callArgs = []*JetDropsExtractorMockStartParams{}

	m.StopMock = mJetDropsExtractorMockStop{mock: m}
	m.StopMock.callArgs = []*JetDropsExtractorMockStopParams{}

	return m
}

type mJetDropsExtractorMockGetJetDrops struct {
	mock               *JetDropsExtractorMock
	defaultExpectation *JetDropsExtractorMockGetJetDropsExpectation
	expectations       []*JetDropsExtractorMockGetJetDropsExpectation

	callArgs []*JetDropsExtractorMockGetJetDropsParams
	mutex    sync.RWMutex
}

// JetDropsExtractorMockGetJetDropsExpectation specifies expectation struct of the JetDropsExtractor.GetJetDrops
type JetDropsExtractorMockGetJetDropsExpectation struct {
	mock    *JetDropsExtractorMock
	params  *JetDropsExtractorMockGetJetDropsParams
	results *JetDropsExtractorMockGetJetDropsResults
	Counter uint64
}

// JetDropsExtractorMockGetJetDropsParams contains parameters of the JetDropsExtractor.GetJetDrops
type JetDropsExtractorMockGetJetDropsParams struct {
	ctx context.Context
}

// JetDropsExtractorMockGetJetDropsResults contains results of the JetDropsExtractor.GetJetDrops
type JetDropsExtractorMockGetJetDropsResults struct {
	ch1 <-chan *types.PlatformJetDrops
}

// Expect sets up expected params for JetDropsExtractor.GetJetDrops
func (mmGetJetDrops *mJetDropsExtractorMockGetJetDrops) Expect(ctx context.Context) *mJetDropsExtractorMockGetJetDrops {
	if mmGetJetDrops.mock.funcGetJetDrops != nil {
		mmGetJetDrops.mock.t.Fatalf("JetDropsExtractorMock.GetJetDrops mock is already set by Set")
	}

	if mmGetJetDrops.defaultExpectation == nil {
		mmGetJetDrops.defaultExpectation = &JetDropsExtractorMockGetJetDropsExpectation{}
	}

	mmGetJetDrops.defaultExpectation.params = &JetDropsExtractorMockGetJetDropsParams{ctx}
	for _, e := range mmGetJetDrops.expectations {
		if minimock.Equal(e.params, mmGetJetDrops.defaultExpectation.params) {
			mmGetJetDrops.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetJetDrops.defaultExpectation.params)
		}
	}

	return mmGetJetDrops
}

// Inspect accepts an inspector function that has same arguments as the JetDropsExtractor.GetJetDrops
func (mmGetJetDrops *mJetDropsExtractorMockGetJetDrops) Inspect(f func(ctx context.Context)) *mJetDropsExtractorMockGetJetDrops {
	if mmGetJetDrops.mock.inspectFuncGetJetDrops != nil {
		mmGetJetDrops.mock.t.Fatalf("Inspect function is already set for JetDropsExtractorMock.GetJetDrops")
	}

	mmGetJetDrops.mock.inspectFuncGetJetDrops = f

	return mmGetJetDrops
}

// Return sets up results that will be returned by JetDropsExtractor.GetJetDrops
func (mmGetJetDrops *mJetDropsExtractorMockGetJetDrops) Return(ch1 <-chan *types.PlatformJetDrops) *JetDropsExtractorMock {
	if mmGetJetDrops.mock.funcGetJetDrops != nil {
		mmGetJetDrops.mock.t.Fatalf("JetDropsExtractorMock.GetJetDrops mock is already set by Set")
	}

	if mmGetJetDrops.defaultExpectation == nil {
		mmGetJetDrops.defaultExpectation = &JetDropsExtractorMockGetJetDropsExpectation{mock: mmGetJetDrops.mock}
	}
	mmGetJetDrops.defaultExpectation.results = &JetDropsExtractorMockGetJetDropsResults{ch1}
	return mmGetJetDrops.mock
}

//Set uses given function f to mock the JetDropsExtractor.GetJetDrops method
func (mmGetJetDrops *mJetDropsExtractorMockGetJetDrops) Set(f func(ctx context.Context) (ch1 <-chan *types.PlatformJetDrops)) *JetDropsExtractorMock {
	if mmGetJetDrops.defaultExpectation != nil {
		mmGetJetDrops.mock.t.Fatalf("Default expectation is already set for the JetDropsExtractor.GetJetDrops method")
	}

	if len(mmGetJetDrops.expectations) > 0 {
		mmGetJetDrops.mock.t.Fatalf("Some expectations are already set for the JetDropsExtractor.GetJetDrops method")
	}

	mmGetJetDrops.mock.funcGetJetDrops = f
	return mmGetJetDrops.mock
}

// When sets expectation for the JetDropsExtractor.GetJetDrops which will trigger the result defined by the following
// Then helper
func (mmGetJetDrops *mJetDropsExtractorMockGetJetDrops) When(ctx context.Context) *JetDropsExtractorMockGetJetDropsExpectation {
	if mmGetJetDrops.mock.funcGetJetDrops != nil {
		mmGetJetDrops.mock.t.Fatalf("JetDropsExtractorMock.GetJetDrops mock is already set by Set")
	}

	expectation := &JetDropsExtractorMockGetJetDropsExpectation{
		mock:   mmGetJetDrops.mock,
		params: &JetDropsExtractorMockGetJetDropsParams{ctx},
	}
	mmGetJetDrops.expectations = append(mmGetJetDrops.expectations, expectation)
	return expectation
}

// Then sets up JetDropsExtractor.GetJetDrops return parameters for the expectation previously defined by the When method
func (e *JetDropsExtractorMockGetJetDropsExpectation) Then(ch1 <-chan *types.PlatformJetDrops) *JetDropsExtractorMock {
	e.results = &JetDropsExtractorMockGetJetDropsResults{ch1}
	return e.mock
}

// GetJetDrops implements interfaces.JetDropsExtractor
func (mmGetJetDrops *JetDropsExtractorMock) GetJetDrops(ctx context.Context) (ch1 <-chan *types.PlatformJetDrops) {
	mm_atomic.AddUint64(&mmGetJetDrops.beforeGetJetDropsCounter, 1)
	defer mm_atomic.AddUint64(&mmGetJetDrops.afterGetJetDropsCounter, 1)

	if mmGetJetDrops.inspectFuncGetJetDrops != nil {
		mmGetJetDrops.inspectFuncGetJetDrops(ctx)
	}

	mm_params := &JetDropsExtractorMockGetJetDropsParams{ctx}

	// Record call args
	mmGetJetDrops.GetJetDropsMock.mutex.Lock()
	mmGetJetDrops.GetJetDropsMock.callArgs = append(mmGetJetDrops.GetJetDropsMock.callArgs, mm_params)
	mmGetJetDrops.GetJetDropsMock.mutex.Unlock()

	for _, e := range mmGetJetDrops.GetJetDropsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ch1
		}
	}

	if mmGetJetDrops.GetJetDropsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetJetDrops.GetJetDropsMock.defaultExpectation.Counter, 1)
		mm_want := mmGetJetDrops.GetJetDropsMock.defaultExpectation.params
		mm_got := JetDropsExtractorMockGetJetDropsParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetJetDrops.t.Errorf("JetDropsExtractorMock.GetJetDrops got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetJetDrops.GetJetDropsMock.defaultExpectation.results
		if mm_results == nil {
			mmGetJetDrops.t.Fatal("No results are set for the JetDropsExtractorMock.GetJetDrops")
		}
		return (*mm_results).ch1
	}
	if mmGetJetDrops.funcGetJetDrops != nil {
		return mmGetJetDrops.funcGetJetDrops(ctx)
	}
	mmGetJetDrops.t.Fatalf("Unexpected call to JetDropsExtractorMock.GetJetDrops. %v", ctx)
	return
}

// GetJetDropsAfterCounter returns a count of finished JetDropsExtractorMock.GetJetDrops invocations
func (mmGetJetDrops *JetDropsExtractorMock) GetJetDropsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetJetDrops.afterGetJetDropsCounter)
}

// GetJetDropsBeforeCounter returns a count of JetDropsExtractorMock.GetJetDrops invocations
func (mmGetJetDrops *JetDropsExtractorMock) GetJetDropsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetJetDrops.beforeGetJetDropsCounter)
}

// Calls returns a list of arguments used in each call to JetDropsExtractorMock.GetJetDrops.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetJetDrops *mJetDropsExtractorMockGetJetDrops) Calls() []*JetDropsExtractorMockGetJetDropsParams {
	mmGetJetDrops.mutex.RLock()

	argCopy := make([]*JetDropsExtractorMockGetJetDropsParams, len(mmGetJetDrops.callArgs))
	copy(argCopy, mmGetJetDrops.callArgs)

	mmGetJetDrops.mutex.RUnlock()

	return argCopy
}

// MinimockGetJetDropsDone returns true if the count of the GetJetDrops invocations corresponds
// the number of defined expectations
func (m *JetDropsExtractorMock) MinimockGetJetDropsDone() bool {
	for _, e := range m.GetJetDropsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetJetDropsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetJetDropsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetJetDrops != nil && mm_atomic.LoadUint64(&m.afterGetJetDropsCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetJetDropsInspect logs each unmet expectation
func (m *JetDropsExtractorMock) MinimockGetJetDropsInspect() {
	for _, e := range m.GetJetDropsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to JetDropsExtractorMock.GetJetDrops with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetJetDropsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetJetDropsCounter) < 1 {
		if m.GetJetDropsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to JetDropsExtractorMock.GetJetDrops")
		} else {
			m.t.Errorf("Expected call to JetDropsExtractorMock.GetJetDrops with params: %#v", *m.GetJetDropsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetJetDrops != nil && mm_atomic.LoadUint64(&m.afterGetJetDropsCounter) < 1 {
		m.t.Error("Expected call to JetDropsExtractorMock.GetJetDrops")
	}
}

type mJetDropsExtractorMockLoadJetDrops struct {
	mock               *JetDropsExtractorMock
	defaultExpectation *JetDropsExtractorMockLoadJetDropsExpectation
	expectations       []*JetDropsExtractorMockLoadJetDropsExpectation

	callArgs []*JetDropsExtractorMockLoadJetDropsParams
	mutex    sync.RWMutex
}

// JetDropsExtractorMockLoadJetDropsExpectation specifies expectation struct of the JetDropsExtractor.LoadJetDrops
type JetDropsExtractorMockLoadJetDropsExpectation struct {
	mock    *JetDropsExtractorMock
	params  *JetDropsExtractorMockLoadJetDropsParams
	results *JetDropsExtractorMockLoadJetDropsResults
	Counter uint64
}

// JetDropsExtractorMockLoadJetDropsParams contains parameters of the JetDropsExtractor.LoadJetDrops
type JetDropsExtractorMockLoadJetDropsParams struct {
	ctx             context.Context
	fromPulseNumber int
	toPulseNumber   int
}

// JetDropsExtractorMockLoadJetDropsResults contains results of the JetDropsExtractor.LoadJetDrops
type JetDropsExtractorMockLoadJetDropsResults struct {
	err error
}

// Expect sets up expected params for JetDropsExtractor.LoadJetDrops
func (mmLoadJetDrops *mJetDropsExtractorMockLoadJetDrops) Expect(ctx context.Context, fromPulseNumber int, toPulseNumber int) *mJetDropsExtractorMockLoadJetDrops {
	if mmLoadJetDrops.mock.funcLoadJetDrops != nil {
		mmLoadJetDrops.mock.t.Fatalf("JetDropsExtractorMock.LoadJetDrops mock is already set by Set")
	}

	if mmLoadJetDrops.defaultExpectation == nil {
		mmLoadJetDrops.defaultExpectation = &JetDropsExtractorMockLoadJetDropsExpectation{}
	}

	mmLoadJetDrops.defaultExpectation.params = &JetDropsExtractorMockLoadJetDropsParams{ctx, fromPulseNumber, toPulseNumber}
	for _, e := range mmLoadJetDrops.expectations {
		if minimock.Equal(e.params, mmLoadJetDrops.defaultExpectation.params) {
			mmLoadJetDrops.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLoadJetDrops.defaultExpectation.params)
		}
	}

	return mmLoadJetDrops
}

// Inspect accepts an inspector function that has same arguments as the JetDropsExtractor.LoadJetDrops
func (mmLoadJetDrops *mJetDropsExtractorMockLoadJetDrops) Inspect(f func(ctx context.Context, fromPulseNumber int, toPulseNumber int)) *mJetDropsExtractorMockLoadJetDrops {
	if mmLoadJetDrops.mock.inspectFuncLoadJetDrops != nil {
		mmLoadJetDrops.mock.t.Fatalf("Inspect function is already set for JetDropsExtractorMock.LoadJetDrops")
	}

	mmLoadJetDrops.mock.inspectFuncLoadJetDrops = f

	return mmLoadJetDrops
}

// Return sets up results that will be returned by JetDropsExtractor.LoadJetDrops
func (mmLoadJetDrops *mJetDropsExtractorMockLoadJetDrops) Return(err error) *JetDropsExtractorMock {
	if mmLoadJetDrops.mock.funcLoadJetDrops != nil {
		mmLoadJetDrops.mock.t.Fatalf("JetDropsExtractorMock.LoadJetDrops mock is already set by Set")
	}

	if mmLoadJetDrops.defaultExpectation == nil {
		mmLoadJetDrops.defaultExpectation = &JetDropsExtractorMockLoadJetDropsExpectation{mock: mmLoadJetDrops.mock}
	}
	mmLoadJetDrops.defaultExpectation.results = &JetDropsExtractorMockLoadJetDropsResults{err}
	return mmLoadJetDrops.mock
}

//Set uses given function f to mock the JetDropsExtractor.LoadJetDrops method
func (mmLoadJetDrops *mJetDropsExtractorMockLoadJetDrops) Set(f func(ctx context.Context, fromPulseNumber int, toPulseNumber int) (err error)) *JetDropsExtractorMock {
	if mmLoadJetDrops.defaultExpectation != nil {
		mmLoadJetDrops.mock.t.Fatalf("Default expectation is already set for the JetDropsExtractor.LoadJetDrops method")
	}

	if len(mmLoadJetDrops.expectations) > 0 {
		mmLoadJetDrops.mock.t.Fatalf("Some expectations are already set for the JetDropsExtractor.LoadJetDrops method")
	}

	mmLoadJetDrops.mock.funcLoadJetDrops = f
	return mmLoadJetDrops.mock
}

// When sets expectation for the JetDropsExtractor.LoadJetDrops which will trigger the result defined by the following
// Then helper
func (mmLoadJetDrops *mJetDropsExtractorMockLoadJetDrops) When(ctx context.Context, fromPulseNumber int, toPulseNumber int) *JetDropsExtractorMockLoadJetDropsExpectation {
	if mmLoadJetDrops.mock.funcLoadJetDrops != nil {
		mmLoadJetDrops.mock.t.Fatalf("JetDropsExtractorMock.LoadJetDrops mock is already set by Set")
	}

	expectation := &JetDropsExtractorMockLoadJetDropsExpectation{
		mock:   mmLoadJetDrops.mock,
		params: &JetDropsExtractorMockLoadJetDropsParams{ctx, fromPulseNumber, toPulseNumber},
	}
	mmLoadJetDrops.expectations = append(mmLoadJetDrops.expectations, expectation)
	return expectation
}

// Then sets up JetDropsExtractor.LoadJetDrops return parameters for the expectation previously defined by the When method
func (e *JetDropsExtractorMockLoadJetDropsExpectation) Then(err error) *JetDropsExtractorMock {
	e.results = &JetDropsExtractorMockLoadJetDropsResults{err}
	return e.mock
}

// LoadJetDrops implements interfaces.JetDropsExtractor
func (mmLoadJetDrops *JetDropsExtractorMock) LoadJetDrops(ctx context.Context, fromPulseNumber int, toPulseNumber int) (err error) {
	mm_atomic.AddUint64(&mmLoadJetDrops.beforeLoadJetDropsCounter, 1)
	defer mm_atomic.AddUint64(&mmLoadJetDrops.afterLoadJetDropsCounter, 1)

	if mmLoadJetDrops.inspectFuncLoadJetDrops != nil {
		mmLoadJetDrops.inspectFuncLoadJetDrops(ctx, fromPulseNumber, toPulseNumber)
	}

	mm_params := &JetDropsExtractorMockLoadJetDropsParams{ctx, fromPulseNumber, toPulseNumber}

	// Record call args
	mmLoadJetDrops.LoadJetDropsMock.mutex.Lock()
	mmLoadJetDrops.LoadJetDropsMock.callArgs = append(mmLoadJetDrops.LoadJetDropsMock.callArgs, mm_params)
	mmLoadJetDrops.LoadJetDropsMock.mutex.Unlock()

	for _, e := range mmLoadJetDrops.LoadJetDropsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmLoadJetDrops.LoadJetDropsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLoadJetDrops.LoadJetDropsMock.defaultExpectation.Counter, 1)
		mm_want := mmLoadJetDrops.LoadJetDropsMock.defaultExpectation.params
		mm_got := JetDropsExtractorMockLoadJetDropsParams{ctx, fromPulseNumber, toPulseNumber}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLoadJetDrops.t.Errorf("JetDropsExtractorMock.LoadJetDrops got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLoadJetDrops.LoadJetDropsMock.defaultExpectation.results
		if mm_results == nil {
			mmLoadJetDrops.t.Fatal("No results are set for the JetDropsExtractorMock.LoadJetDrops")
		}
		return (*mm_results).err
	}
	if mmLoadJetDrops.funcLoadJetDrops != nil {
		return mmLoadJetDrops.funcLoadJetDrops(ctx, fromPulseNumber, toPulseNumber)
	}
	mmLoadJetDrops.t.Fatalf("Unexpected call to JetDropsExtractorMock.LoadJetDrops. %v %v %v", ctx, fromPulseNumber, toPulseNumber)
	return
}

// LoadJetDropsAfterCounter returns a count of finished JetDropsExtractorMock.LoadJetDrops invocations
func (mmLoadJetDrops *JetDropsExtractorMock) LoadJetDropsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLoadJetDrops.afterLoadJetDropsCounter)
}

// LoadJetDropsBeforeCounter returns a count of JetDropsExtractorMock.LoadJetDrops invocations
func (mmLoadJetDrops *JetDropsExtractorMock) LoadJetDropsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLoadJetDrops.beforeLoadJetDropsCounter)
}

// Calls returns a list of arguments used in each call to JetDropsExtractorMock.LoadJetDrops.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLoadJetDrops *mJetDropsExtractorMockLoadJetDrops) Calls() []*JetDropsExtractorMockLoadJetDropsParams {
	mmLoadJetDrops.mutex.RLock()

	argCopy := make([]*JetDropsExtractorMockLoadJetDropsParams, len(mmLoadJetDrops.callArgs))
	copy(argCopy, mmLoadJetDrops.callArgs)

	mmLoadJetDrops.mutex.RUnlock()

	return argCopy
}

// MinimockLoadJetDropsDone returns true if the count of the LoadJetDrops invocations corresponds
// the number of defined expectations
func (m *JetDropsExtractorMock) MinimockLoadJetDropsDone() bool {
	for _, e := range m.LoadJetDropsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoadJetDropsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoadJetDropsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLoadJetDrops != nil && mm_atomic.LoadUint64(&m.afterLoadJetDropsCounter) < 1 {
		return false
	}
	return true
}

// MinimockLoadJetDropsInspect logs each unmet expectation
func (m *JetDropsExtractorMock) MinimockLoadJetDropsInspect() {
	for _, e := range m.LoadJetDropsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to JetDropsExtractorMock.LoadJetDrops with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoadJetDropsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoadJetDropsCounter) < 1 {
		if m.LoadJetDropsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to JetDropsExtractorMock.LoadJetDrops")
		} else {
			m.t.Errorf("Expected call to JetDropsExtractorMock.LoadJetDrops with params: %#v", *m.LoadJetDropsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLoadJetDrops != nil && mm_atomic.LoadUint64(&m.afterLoadJetDropsCounter) < 1 {
		m.t.Error("Expected call to JetDropsExtractorMock.LoadJetDrops")
	}
}

type mJetDropsExtractorMockStart struct {
	mock               *JetDropsExtractorMock
	defaultExpectation *JetDropsExtractorMockStartExpectation
	expectations       []*JetDropsExtractorMockStartExpectation

	callArgs []*JetDropsExtractorMockStartParams
	mutex    sync.RWMutex
}

// JetDropsExtractorMockStartExpectation specifies expectation struct of the JetDropsExtractor.Start
type JetDropsExtractorMockStartExpectation struct {
	mock    *JetDropsExtractorMock
	params  *JetDropsExtractorMockStartParams
	results *JetDropsExtractorMockStartResults
	Counter uint64
}

// JetDropsExtractorMockStartParams contains parameters of the JetDropsExtractor.Start
type JetDropsExtractorMockStartParams struct {
	ctx context.Context
}

// JetDropsExtractorMockStartResults contains results of the JetDropsExtractor.Start
type JetDropsExtractorMockStartResults struct {
	err error
}

// Expect sets up expected params for JetDropsExtractor.Start
func (mmStart *mJetDropsExtractorMockStart) Expect(ctx context.Context) *mJetDropsExtractorMockStart {
	if mmStart.mock.funcStart != nil {
		mmStart.mock.t.Fatalf("JetDropsExtractorMock.Start mock is already set by Set")
	}

	if mmStart.defaultExpectation == nil {
		mmStart.defaultExpectation = &JetDropsExtractorMockStartExpectation{}
	}

	mmStart.defaultExpectation.params = &JetDropsExtractorMockStartParams{ctx}
	for _, e := range mmStart.expectations {
		if minimock.Equal(e.params, mmStart.defaultExpectation.params) {
			mmStart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStart.defaultExpectation.params)
		}
	}

	return mmStart
}

// Inspect accepts an inspector function that has same arguments as the JetDropsExtractor.Start
func (mmStart *mJetDropsExtractorMockStart) Inspect(f func(ctx context.Context)) *mJetDropsExtractorMockStart {
	if mmStart.mock.inspectFuncStart != nil {
		mmStart.mock.t.Fatalf("Inspect function is already set for JetDropsExtractorMock.Start")
	}

	mmStart.mock.inspectFuncStart = f

	return mmStart
}

// Return sets up results that will be returned by JetDropsExtractor.Start
func (mmStart *mJetDropsExtractorMockStart) Return(err error) *JetDropsExtractorMock {
	if mmStart.mock.funcStart != nil {
		mmStart.mock.t.Fatalf("JetDropsExtractorMock.Start mock is already set by Set")
	}

	if mmStart.defaultExpectation == nil {
		mmStart.defaultExpectation = &JetDropsExtractorMockStartExpectation{mock: mmStart.mock}
	}
	mmStart.defaultExpectation.results = &JetDropsExtractorMockStartResults{err}
	return mmStart.mock
}

//Set uses given function f to mock the JetDropsExtractor.Start method
func (mmStart *mJetDropsExtractorMockStart) Set(f func(ctx context.Context) (err error)) *JetDropsExtractorMock {
	if mmStart.defaultExpectation != nil {
		mmStart.mock.t.Fatalf("Default expectation is already set for the JetDropsExtractor.Start method")
	}

	if len(mmStart.expectations) > 0 {
		mmStart.mock.t.Fatalf("Some expectations are already set for the JetDropsExtractor.Start method")
	}

	mmStart.mock.funcStart = f
	return mmStart.mock
}

// When sets expectation for the JetDropsExtractor.Start which will trigger the result defined by the following
// Then helper
func (mmStart *mJetDropsExtractorMockStart) When(ctx context.Context) *JetDropsExtractorMockStartExpectation {
	if mmStart.mock.funcStart != nil {
		mmStart.mock.t.Fatalf("JetDropsExtractorMock.Start mock is already set by Set")
	}

	expectation := &JetDropsExtractorMockStartExpectation{
		mock:   mmStart.mock,
		params: &JetDropsExtractorMockStartParams{ctx},
	}
	mmStart.expectations = append(mmStart.expectations, expectation)
	return expectation
}

// Then sets up JetDropsExtractor.Start return parameters for the expectation previously defined by the When method
func (e *JetDropsExtractorMockStartExpectation) Then(err error) *JetDropsExtractorMock {
	e.results = &JetDropsExtractorMockStartResults{err}
	return e.mock
}

// Start implements interfaces.JetDropsExtractor
func (mmStart *JetDropsExtractorMock) Start(ctx context.Context) (err error) {
	mm_atomic.AddUint64(&mmStart.beforeStartCounter, 1)
	defer mm_atomic.AddUint64(&mmStart.afterStartCounter, 1)

	if mmStart.inspectFuncStart != nil {
		mmStart.inspectFuncStart(ctx)
	}

	mm_params := &JetDropsExtractorMockStartParams{ctx}

	// Record call args
	mmStart.StartMock.mutex.Lock()
	mmStart.StartMock.callArgs = append(mmStart.StartMock.callArgs, mm_params)
	mmStart.StartMock.mutex.Unlock()

	for _, e := range mmStart.StartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmStart.StartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStart.StartMock.defaultExpectation.Counter, 1)
		mm_want := mmStart.StartMock.defaultExpectation.params
		mm_got := JetDropsExtractorMockStartParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStart.t.Errorf("JetDropsExtractorMock.Start got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmStart.StartMock.defaultExpectation.results
		if mm_results == nil {
			mmStart.t.Fatal("No results are set for the JetDropsExtractorMock.Start")
		}
		return (*mm_results).err
	}
	if mmStart.funcStart != nil {
		return mmStart.funcStart(ctx)
	}
	mmStart.t.Fatalf("Unexpected call to JetDropsExtractorMock.Start. %v", ctx)
	return
}

// StartAfterCounter returns a count of finished JetDropsExtractorMock.Start invocations
func (mmStart *JetDropsExtractorMock) StartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStart.afterStartCounter)
}

// StartBeforeCounter returns a count of JetDropsExtractorMock.Start invocations
func (mmStart *JetDropsExtractorMock) StartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStart.beforeStartCounter)
}

// Calls returns a list of arguments used in each call to JetDropsExtractorMock.Start.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStart *mJetDropsExtractorMockStart) Calls() []*JetDropsExtractorMockStartParams {
	mmStart.mutex.RLock()

	argCopy := make([]*JetDropsExtractorMockStartParams, len(mmStart.callArgs))
	copy(argCopy, mmStart.callArgs)

	mmStart.mutex.RUnlock()

	return argCopy
}

// MinimockStartDone returns true if the count of the Start invocations corresponds
// the number of defined expectations
func (m *JetDropsExtractorMock) MinimockStartDone() bool {
	for _, e := range m.StartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStart != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		return false
	}
	return true
}

// MinimockStartInspect logs each unmet expectation
func (m *JetDropsExtractorMock) MinimockStartInspect() {
	for _, e := range m.StartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to JetDropsExtractorMock.Start with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		if m.StartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to JetDropsExtractorMock.Start")
		} else {
			m.t.Errorf("Expected call to JetDropsExtractorMock.Start with params: %#v", *m.StartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStart != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		m.t.Error("Expected call to JetDropsExtractorMock.Start")
	}
}

type mJetDropsExtractorMockStop struct {
	mock               *JetDropsExtractorMock
	defaultExpectation *JetDropsExtractorMockStopExpectation
	expectations       []*JetDropsExtractorMockStopExpectation

	callArgs []*JetDropsExtractorMockStopParams
	mutex    sync.RWMutex
}

// JetDropsExtractorMockStopExpectation specifies expectation struct of the JetDropsExtractor.Stop
type JetDropsExtractorMockStopExpectation struct {
	mock    *JetDropsExtractorMock
	params  *JetDropsExtractorMockStopParams
	results *JetDropsExtractorMockStopResults
	Counter uint64
}

// JetDropsExtractorMockStopParams contains parameters of the JetDropsExtractor.Stop
type JetDropsExtractorMockStopParams struct {
	ctx context.Context
}

// JetDropsExtractorMockStopResults contains results of the JetDropsExtractor.Stop
type JetDropsExtractorMockStopResults struct {
	err error
}

// Expect sets up expected params for JetDropsExtractor.Stop
func (mmStop *mJetDropsExtractorMockStop) Expect(ctx context.Context) *mJetDropsExtractorMockStop {
	if mmStop.mock.funcStop != nil {
		mmStop.mock.t.Fatalf("JetDropsExtractorMock.Stop mock is already set by Set")
	}

	if mmStop.defaultExpectation == nil {
		mmStop.defaultExpectation = &JetDropsExtractorMockStopExpectation{}
	}

	mmStop.defaultExpectation.params = &JetDropsExtractorMockStopParams{ctx}
	for _, e := range mmStop.expectations {
		if minimock.Equal(e.params, mmStop.defaultExpectation.params) {
			mmStop.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStop.defaultExpectation.params)
		}
	}

	return mmStop
}

// Inspect accepts an inspector function that has same arguments as the JetDropsExtractor.Stop
func (mmStop *mJetDropsExtractorMockStop) Inspect(f func(ctx context.Context)) *mJetDropsExtractorMockStop {
	if mmStop.mock.inspectFuncStop != nil {
		mmStop.mock.t.Fatalf("Inspect function is already set for JetDropsExtractorMock.Stop")
	}

	mmStop.mock.inspectFuncStop = f

	return mmStop
}

// Return sets up results that will be returned by JetDropsExtractor.Stop
func (mmStop *mJetDropsExtractorMockStop) Return(err error) *JetDropsExtractorMock {
	if mmStop.mock.funcStop != nil {
		mmStop.mock.t.Fatalf("JetDropsExtractorMock.Stop mock is already set by Set")
	}

	if mmStop.defaultExpectation == nil {
		mmStop.defaultExpectation = &JetDropsExtractorMockStopExpectation{mock: mmStop.mock}
	}
	mmStop.defaultExpectation.results = &JetDropsExtractorMockStopResults{err}
	return mmStop.mock
}

//Set uses given function f to mock the JetDropsExtractor.Stop method
func (mmStop *mJetDropsExtractorMockStop) Set(f func(ctx context.Context) (err error)) *JetDropsExtractorMock {
	if mmStop.defaultExpectation != nil {
		mmStop.mock.t.Fatalf("Default expectation is already set for the JetDropsExtractor.Stop method")
	}

	if len(mmStop.expectations) > 0 {
		mmStop.mock.t.Fatalf("Some expectations are already set for the JetDropsExtractor.Stop method")
	}

	mmStop.mock.funcStop = f
	return mmStop.mock
}

// When sets expectation for the JetDropsExtractor.Stop which will trigger the result defined by the following
// Then helper
func (mmStop *mJetDropsExtractorMockStop) When(ctx context.Context) *JetDropsExtractorMockStopExpectation {
	if mmStop.mock.funcStop != nil {
		mmStop.mock.t.Fatalf("JetDropsExtractorMock.Stop mock is already set by Set")
	}

	expectation := &JetDropsExtractorMockStopExpectation{
		mock:   mmStop.mock,
		params: &JetDropsExtractorMockStopParams{ctx},
	}
	mmStop.expectations = append(mmStop.expectations, expectation)
	return expectation
}

// Then sets up JetDropsExtractor.Stop return parameters for the expectation previously defined by the When method
func (e *JetDropsExtractorMockStopExpectation) Then(err error) *JetDropsExtractorMock {
	e.results = &JetDropsExtractorMockStopResults{err}
	return e.mock
}

// Stop implements interfaces.JetDropsExtractor
func (mmStop *JetDropsExtractorMock) Stop(ctx context.Context) (err error) {
	mm_atomic.AddUint64(&mmStop.beforeStopCounter, 1)
	defer mm_atomic.AddUint64(&mmStop.afterStopCounter, 1)

	if mmStop.inspectFuncStop != nil {
		mmStop.inspectFuncStop(ctx)
	}

	mm_params := &JetDropsExtractorMockStopParams{ctx}

	// Record call args
	mmStop.StopMock.mutex.Lock()
	mmStop.StopMock.callArgs = append(mmStop.StopMock.callArgs, mm_params)
	mmStop.StopMock.mutex.Unlock()

	for _, e := range mmStop.StopMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmStop.StopMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStop.StopMock.defaultExpectation.Counter, 1)
		mm_want := mmStop.StopMock.defaultExpectation.params
		mm_got := JetDropsExtractorMockStopParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStop.t.Errorf("JetDropsExtractorMock.Stop got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmStop.StopMock.defaultExpectation.results
		if mm_results == nil {
			mmStop.t.Fatal("No results are set for the JetDropsExtractorMock.Stop")
		}
		return (*mm_results).err
	}
	if mmStop.funcStop != nil {
		return mmStop.funcStop(ctx)
	}
	mmStop.t.Fatalf("Unexpected call to JetDropsExtractorMock.Stop. %v", ctx)
	return
}

// StopAfterCounter returns a count of finished JetDropsExtractorMock.Stop invocations
func (mmStop *JetDropsExtractorMock) StopAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStop.afterStopCounter)
}

// StopBeforeCounter returns a count of JetDropsExtractorMock.Stop invocations
func (mmStop *JetDropsExtractorMock) StopBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStop.beforeStopCounter)
}

// Calls returns a list of arguments used in each call to JetDropsExtractorMock.Stop.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStop *mJetDropsExtractorMockStop) Calls() []*JetDropsExtractorMockStopParams {
	mmStop.mutex.RLock()

	argCopy := make([]*JetDropsExtractorMockStopParams, len(mmStop.callArgs))
	copy(argCopy, mmStop.callArgs)

	mmStop.mutex.RUnlock()

	return argCopy
}

// MinimockStopDone returns true if the count of the Stop invocations corresponds
// the number of defined expectations
func (m *JetDropsExtractorMock) MinimockStopDone() bool {
	for _, e := range m.StopMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StopMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStop != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		return false
	}
	return true
}

// MinimockStopInspect logs each unmet expectation
func (m *JetDropsExtractorMock) MinimockStopInspect() {
	for _, e := range m.StopMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to JetDropsExtractorMock.Stop with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StopMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		if m.StopMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to JetDropsExtractorMock.Stop")
		} else {
			m.t.Errorf("Expected call to JetDropsExtractorMock.Stop with params: %#v", *m.StopMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStop != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		m.t.Error("Expected call to JetDropsExtractorMock.Stop")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *JetDropsExtractorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetJetDropsInspect()

		m.MinimockLoadJetDropsInspect()

		m.MinimockStartInspect()

		m.MinimockStopInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *JetDropsExtractorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *JetDropsExtractorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetJetDropsDone() &&
		m.MinimockLoadJetDropsDone() &&
		m.MinimockStartDone() &&
		m.MinimockStopDone()
}
