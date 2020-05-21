package cryptkit

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// SignatureVerifierFactoryMock implements SignatureVerifierFactory
type SignatureVerifierFactoryMock struct {
	t minimock.Tester

	funcCreateSignatureVerifierWithPKS          func(pks PublicKeyStore) (s1 SignatureVerifier)
	inspectFuncCreateSignatureVerifierWithPKS   func(pks PublicKeyStore)
	afterCreateSignatureVerifierWithPKSCounter  uint64
	beforeCreateSignatureVerifierWithPKSCounter uint64
	CreateSignatureVerifierWithPKSMock          mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS
}

// NewSignatureVerifierFactoryMock returns a mock for SignatureVerifierFactory
func NewSignatureVerifierFactoryMock(t minimock.Tester) *SignatureVerifierFactoryMock {
	m := &SignatureVerifierFactoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateSignatureVerifierWithPKSMock = mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS{mock: m}
	m.CreateSignatureVerifierWithPKSMock.callArgs = []*SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams{}

	return m
}

type mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS struct {
	mock               *SignatureVerifierFactoryMock
	defaultExpectation *SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation
	expectations       []*SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation

	callArgs []*SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams
	mutex    sync.RWMutex
}

// SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation specifies expectation struct of the SignatureVerifierFactory.CreateSignatureVerifierWithPKS
type SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation struct {
	mock    *SignatureVerifierFactoryMock
	params  *SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams
	results *SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSResults
	Counter uint64
}

// SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams contains parameters of the SignatureVerifierFactory.CreateSignatureVerifierWithPKS
type SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams struct {
	pks PublicKeyStore
}

// SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSResults contains results of the SignatureVerifierFactory.CreateSignatureVerifierWithPKS
type SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSResults struct {
	s1 SignatureVerifier
}

// Expect sets up expected params for SignatureVerifierFactory.CreateSignatureVerifierWithPKS
func (mmCreateSignatureVerifierWithPKS *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS) Expect(pks PublicKeyStore) *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS {
	if mmCreateSignatureVerifierWithPKS.mock.funcCreateSignatureVerifierWithPKS != nil {
		mmCreateSignatureVerifierWithPKS.mock.t.Fatalf("SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS mock is already set by Set")
	}

	if mmCreateSignatureVerifierWithPKS.defaultExpectation == nil {
		mmCreateSignatureVerifierWithPKS.defaultExpectation = &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation{}
	}

	mmCreateSignatureVerifierWithPKS.defaultExpectation.params = &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams{pks}
	for _, e := range mmCreateSignatureVerifierWithPKS.expectations {
		if minimock.Equal(e.params, mmCreateSignatureVerifierWithPKS.defaultExpectation.params) {
			mmCreateSignatureVerifierWithPKS.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateSignatureVerifierWithPKS.defaultExpectation.params)
		}
	}

	return mmCreateSignatureVerifierWithPKS
}

// Inspect accepts an inspector function that has same arguments as the SignatureVerifierFactory.CreateSignatureVerifierWithPKS
func (mmCreateSignatureVerifierWithPKS *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS) Inspect(f func(pks PublicKeyStore)) *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS {
	if mmCreateSignatureVerifierWithPKS.mock.inspectFuncCreateSignatureVerifierWithPKS != nil {
		mmCreateSignatureVerifierWithPKS.mock.t.Fatalf("Inspect function is already set for SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS")
	}

	mmCreateSignatureVerifierWithPKS.mock.inspectFuncCreateSignatureVerifierWithPKS = f

	return mmCreateSignatureVerifierWithPKS
}

// Return sets up results that will be returned by SignatureVerifierFactory.CreateSignatureVerifierWithPKS
func (mmCreateSignatureVerifierWithPKS *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS) Return(s1 SignatureVerifier) *SignatureVerifierFactoryMock {
	if mmCreateSignatureVerifierWithPKS.mock.funcCreateSignatureVerifierWithPKS != nil {
		mmCreateSignatureVerifierWithPKS.mock.t.Fatalf("SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS mock is already set by Set")
	}

	if mmCreateSignatureVerifierWithPKS.defaultExpectation == nil {
		mmCreateSignatureVerifierWithPKS.defaultExpectation = &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation{mock: mmCreateSignatureVerifierWithPKS.mock}
	}
	mmCreateSignatureVerifierWithPKS.defaultExpectation.results = &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSResults{s1}
	return mmCreateSignatureVerifierWithPKS.mock
}

//Set uses given function f to mock the SignatureVerifierFactory.CreateSignatureVerifierWithPKS method
func (mmCreateSignatureVerifierWithPKS *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS) Set(f func(pks PublicKeyStore) (s1 SignatureVerifier)) *SignatureVerifierFactoryMock {
	if mmCreateSignatureVerifierWithPKS.defaultExpectation != nil {
		mmCreateSignatureVerifierWithPKS.mock.t.Fatalf("Default expectation is already set for the SignatureVerifierFactory.CreateSignatureVerifierWithPKS method")
	}

	if len(mmCreateSignatureVerifierWithPKS.expectations) > 0 {
		mmCreateSignatureVerifierWithPKS.mock.t.Fatalf("Some expectations are already set for the SignatureVerifierFactory.CreateSignatureVerifierWithPKS method")
	}

	mmCreateSignatureVerifierWithPKS.mock.funcCreateSignatureVerifierWithPKS = f
	return mmCreateSignatureVerifierWithPKS.mock
}

// When sets expectation for the SignatureVerifierFactory.CreateSignatureVerifierWithPKS which will trigger the result defined by the following
// Then helper
func (mmCreateSignatureVerifierWithPKS *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS) When(pks PublicKeyStore) *SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation {
	if mmCreateSignatureVerifierWithPKS.mock.funcCreateSignatureVerifierWithPKS != nil {
		mmCreateSignatureVerifierWithPKS.mock.t.Fatalf("SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS mock is already set by Set")
	}

	expectation := &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation{
		mock:   mmCreateSignatureVerifierWithPKS.mock,
		params: &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams{pks},
	}
	mmCreateSignatureVerifierWithPKS.expectations = append(mmCreateSignatureVerifierWithPKS.expectations, expectation)
	return expectation
}

// Then sets up SignatureVerifierFactory.CreateSignatureVerifierWithPKS return parameters for the expectation previously defined by the When method
func (e *SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSExpectation) Then(s1 SignatureVerifier) *SignatureVerifierFactoryMock {
	e.results = &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSResults{s1}
	return e.mock
}

// CreateSignatureVerifierWithPKS implements SignatureVerifierFactory
func (mmCreateSignatureVerifierWithPKS *SignatureVerifierFactoryMock) CreateSignatureVerifierWithPKS(pks PublicKeyStore) (s1 SignatureVerifier) {
	mm_atomic.AddUint64(&mmCreateSignatureVerifierWithPKS.beforeCreateSignatureVerifierWithPKSCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateSignatureVerifierWithPKS.afterCreateSignatureVerifierWithPKSCounter, 1)

	if mmCreateSignatureVerifierWithPKS.inspectFuncCreateSignatureVerifierWithPKS != nil {
		mmCreateSignatureVerifierWithPKS.inspectFuncCreateSignatureVerifierWithPKS(pks)
	}

	mm_params := &SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams{pks}

	// Record call args
	mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.mutex.Lock()
	mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.callArgs = append(mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.callArgs, mm_params)
	mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.mutex.Unlock()

	for _, e := range mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1
		}
	}

	if mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.defaultExpectation.params
		mm_got := SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams{pks}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateSignatureVerifierWithPKS.t.Errorf("SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateSignatureVerifierWithPKS.CreateSignatureVerifierWithPKSMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateSignatureVerifierWithPKS.t.Fatal("No results are set for the SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS")
		}
		return (*mm_results).s1
	}
	if mmCreateSignatureVerifierWithPKS.funcCreateSignatureVerifierWithPKS != nil {
		return mmCreateSignatureVerifierWithPKS.funcCreateSignatureVerifierWithPKS(pks)
	}
	mmCreateSignatureVerifierWithPKS.t.Fatalf("Unexpected call to SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS. %v", pks)
	return
}

// CreateSignatureVerifierWithPKSAfterCounter returns a count of finished SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS invocations
func (mmCreateSignatureVerifierWithPKS *SignatureVerifierFactoryMock) CreateSignatureVerifierWithPKSAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateSignatureVerifierWithPKS.afterCreateSignatureVerifierWithPKSCounter)
}

// CreateSignatureVerifierWithPKSBeforeCounter returns a count of SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS invocations
func (mmCreateSignatureVerifierWithPKS *SignatureVerifierFactoryMock) CreateSignatureVerifierWithPKSBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateSignatureVerifierWithPKS.beforeCreateSignatureVerifierWithPKSCounter)
}

// Calls returns a list of arguments used in each call to SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateSignatureVerifierWithPKS *mSignatureVerifierFactoryMockCreateSignatureVerifierWithPKS) Calls() []*SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams {
	mmCreateSignatureVerifierWithPKS.mutex.RLock()

	argCopy := make([]*SignatureVerifierFactoryMockCreateSignatureVerifierWithPKSParams, len(mmCreateSignatureVerifierWithPKS.callArgs))
	copy(argCopy, mmCreateSignatureVerifierWithPKS.callArgs)

	mmCreateSignatureVerifierWithPKS.mutex.RUnlock()

	return argCopy
}

// MinimockCreateSignatureVerifierWithPKSDone returns true if the count of the CreateSignatureVerifierWithPKS invocations corresponds
// the number of defined expectations
func (m *SignatureVerifierFactoryMock) MinimockCreateSignatureVerifierWithPKSDone() bool {
	for _, e := range m.CreateSignatureVerifierWithPKSMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateSignatureVerifierWithPKSMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateSignatureVerifierWithPKSCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateSignatureVerifierWithPKS != nil && mm_atomic.LoadUint64(&m.afterCreateSignatureVerifierWithPKSCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateSignatureVerifierWithPKSInspect logs each unmet expectation
func (m *SignatureVerifierFactoryMock) MinimockCreateSignatureVerifierWithPKSInspect() {
	for _, e := range m.CreateSignatureVerifierWithPKSMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateSignatureVerifierWithPKSMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateSignatureVerifierWithPKSCounter) < 1 {
		if m.CreateSignatureVerifierWithPKSMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS")
		} else {
			m.t.Errorf("Expected call to SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS with params: %#v", *m.CreateSignatureVerifierWithPKSMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateSignatureVerifierWithPKS != nil && mm_atomic.LoadUint64(&m.afterCreateSignatureVerifierWithPKSCounter) < 1 {
		m.t.Error("Expected call to SignatureVerifierFactoryMock.CreateSignatureVerifierWithPKS")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SignatureVerifierFactoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateSignatureVerifierWithPKSInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SignatureVerifierFactoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *SignatureVerifierFactoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateSignatureVerifierWithPKSDone()
}