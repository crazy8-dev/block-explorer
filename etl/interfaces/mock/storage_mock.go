package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/block-explorer/etl/models"
)

// StorageMock implements interfaces.Storage
type StorageMock struct {
	t minimock.Tester

	funcSaveJetDropData          func(jetDrop models.JetDrop, records []models.Record) (err error)
	inspectFuncSaveJetDropData   func(jetDrop models.JetDrop, records []models.Record)
	afterSaveJetDropDataCounter  uint64
	beforeSaveJetDropDataCounter uint64
	SaveJetDropDataMock          mStorageMockSaveJetDropData
}

// NewStorageMock returns a mock for interfaces.Storage
func NewStorageMock(t minimock.Tester) *StorageMock {
	m := &StorageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SaveJetDropDataMock = mStorageMockSaveJetDropData{mock: m}
	m.SaveJetDropDataMock.callArgs = []*StorageMockSaveJetDropDataParams{}

	return m
}

type mStorageMockSaveJetDropData struct {
	mock               *StorageMock
	defaultExpectation *StorageMockSaveJetDropDataExpectation
	expectations       []*StorageMockSaveJetDropDataExpectation

	callArgs []*StorageMockSaveJetDropDataParams
	mutex    sync.RWMutex
}

// StorageMockSaveJetDropDataExpectation specifies expectation struct of the Storage.SaveJetDropData
type StorageMockSaveJetDropDataExpectation struct {
	mock    *StorageMock
	params  *StorageMockSaveJetDropDataParams
	results *StorageMockSaveJetDropDataResults
	Counter uint64
}

// StorageMockSaveJetDropDataParams contains parameters of the Storage.SaveJetDropData
type StorageMockSaveJetDropDataParams struct {
	jetDrop models.JetDrop
	records []models.Record
}

// StorageMockSaveJetDropDataResults contains results of the Storage.SaveJetDropData
type StorageMockSaveJetDropDataResults struct {
	err error
}

// Expect sets up expected params for Storage.SaveJetDropData
func (mmSaveJetDropData *mStorageMockSaveJetDropData) Expect(jetDrop models.JetDrop, records []models.Record) *mStorageMockSaveJetDropData {
	if mmSaveJetDropData.mock.funcSaveJetDropData != nil {
		mmSaveJetDropData.mock.t.Fatalf("StorageMock.SaveJetDropData mock is already set by Set")
	}

	if mmSaveJetDropData.defaultExpectation == nil {
		mmSaveJetDropData.defaultExpectation = &StorageMockSaveJetDropDataExpectation{}
	}

	mmSaveJetDropData.defaultExpectation.params = &StorageMockSaveJetDropDataParams{jetDrop, records}
	for _, e := range mmSaveJetDropData.expectations {
		if minimock.Equal(e.params, mmSaveJetDropData.defaultExpectation.params) {
			mmSaveJetDropData.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSaveJetDropData.defaultExpectation.params)
		}
	}

	return mmSaveJetDropData
}

// Inspect accepts an inspector function that has same arguments as the Storage.SaveJetDropData
func (mmSaveJetDropData *mStorageMockSaveJetDropData) Inspect(f func(jetDrop models.JetDrop, records []models.Record)) *mStorageMockSaveJetDropData {
	if mmSaveJetDropData.mock.inspectFuncSaveJetDropData != nil {
		mmSaveJetDropData.mock.t.Fatalf("Inspect function is already set for StorageMock.SaveJetDropData")
	}

	mmSaveJetDropData.mock.inspectFuncSaveJetDropData = f

	return mmSaveJetDropData
}

// Return sets up results that will be returned by Storage.SaveJetDropData
func (mmSaveJetDropData *mStorageMockSaveJetDropData) Return(err error) *StorageMock {
	if mmSaveJetDropData.mock.funcSaveJetDropData != nil {
		mmSaveJetDropData.mock.t.Fatalf("StorageMock.SaveJetDropData mock is already set by Set")
	}

	if mmSaveJetDropData.defaultExpectation == nil {
		mmSaveJetDropData.defaultExpectation = &StorageMockSaveJetDropDataExpectation{mock: mmSaveJetDropData.mock}
	}
	mmSaveJetDropData.defaultExpectation.results = &StorageMockSaveJetDropDataResults{err}
	return mmSaveJetDropData.mock
}

//Set uses given function f to mock the Storage.SaveJetDropData method
func (mmSaveJetDropData *mStorageMockSaveJetDropData) Set(f func(jetDrop models.JetDrop, records []models.Record) (err error)) *StorageMock {
	if mmSaveJetDropData.defaultExpectation != nil {
		mmSaveJetDropData.mock.t.Fatalf("Default expectation is already set for the Storage.SaveJetDropData method")
	}

	if len(mmSaveJetDropData.expectations) > 0 {
		mmSaveJetDropData.mock.t.Fatalf("Some expectations are already set for the Storage.SaveJetDropData method")
	}

	mmSaveJetDropData.mock.funcSaveJetDropData = f
	return mmSaveJetDropData.mock
}

// When sets expectation for the Storage.SaveJetDropData which will trigger the result defined by the following
// Then helper
func (mmSaveJetDropData *mStorageMockSaveJetDropData) When(jetDrop models.JetDrop, records []models.Record) *StorageMockSaveJetDropDataExpectation {
	if mmSaveJetDropData.mock.funcSaveJetDropData != nil {
		mmSaveJetDropData.mock.t.Fatalf("StorageMock.SaveJetDropData mock is already set by Set")
	}

	expectation := &StorageMockSaveJetDropDataExpectation{
		mock:   mmSaveJetDropData.mock,
		params: &StorageMockSaveJetDropDataParams{jetDrop, records},
	}
	mmSaveJetDropData.expectations = append(mmSaveJetDropData.expectations, expectation)
	return expectation
}

// Then sets up Storage.SaveJetDropData return parameters for the expectation previously defined by the When method
func (e *StorageMockSaveJetDropDataExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockSaveJetDropDataResults{err}
	return e.mock
}

// SaveJetDropData implements interfaces.Storage
func (mmSaveJetDropData *StorageMock) SaveJetDropData(jetDrop models.JetDrop, records []models.Record) (err error) {
	mm_atomic.AddUint64(&mmSaveJetDropData.beforeSaveJetDropDataCounter, 1)
	defer mm_atomic.AddUint64(&mmSaveJetDropData.afterSaveJetDropDataCounter, 1)

	if mmSaveJetDropData.inspectFuncSaveJetDropData != nil {
		mmSaveJetDropData.inspectFuncSaveJetDropData(jetDrop, records)
	}

	mm_params := &StorageMockSaveJetDropDataParams{jetDrop, records}

	// Record call args
	mmSaveJetDropData.SaveJetDropDataMock.mutex.Lock()
	mmSaveJetDropData.SaveJetDropDataMock.callArgs = append(mmSaveJetDropData.SaveJetDropDataMock.callArgs, mm_params)
	mmSaveJetDropData.SaveJetDropDataMock.mutex.Unlock()

	for _, e := range mmSaveJetDropData.SaveJetDropDataMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSaveJetDropData.SaveJetDropDataMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSaveJetDropData.SaveJetDropDataMock.defaultExpectation.Counter, 1)
		mm_want := mmSaveJetDropData.SaveJetDropDataMock.defaultExpectation.params
		mm_got := StorageMockSaveJetDropDataParams{jetDrop, records}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSaveJetDropData.t.Errorf("StorageMock.SaveJetDropData got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSaveJetDropData.SaveJetDropDataMock.defaultExpectation.results
		if mm_results == nil {
			mmSaveJetDropData.t.Fatal("No results are set for the StorageMock.SaveJetDropData")
		}
		return (*mm_results).err
	}
	if mmSaveJetDropData.funcSaveJetDropData != nil {
		return mmSaveJetDropData.funcSaveJetDropData(jetDrop, records)
	}
	mmSaveJetDropData.t.Fatalf("Unexpected call to StorageMock.SaveJetDropData. %v %v", jetDrop, records)
	return
}

// SaveJetDropDataAfterCounter returns a count of finished StorageMock.SaveJetDropData invocations
func (mmSaveJetDropData *StorageMock) SaveJetDropDataAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSaveJetDropData.afterSaveJetDropDataCounter)
}

// SaveJetDropDataBeforeCounter returns a count of StorageMock.SaveJetDropData invocations
func (mmSaveJetDropData *StorageMock) SaveJetDropDataBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSaveJetDropData.beforeSaveJetDropDataCounter)
}

// Calls returns a list of arguments used in each call to StorageMock.SaveJetDropData.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSaveJetDropData *mStorageMockSaveJetDropData) Calls() []*StorageMockSaveJetDropDataParams {
	mmSaveJetDropData.mutex.RLock()

	argCopy := make([]*StorageMockSaveJetDropDataParams, len(mmSaveJetDropData.callArgs))
	copy(argCopy, mmSaveJetDropData.callArgs)

	mmSaveJetDropData.mutex.RUnlock()

	return argCopy
}

// MinimockSaveJetDropDataDone returns true if the count of the SaveJetDropData invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockSaveJetDropDataDone() bool {
	for _, e := range m.SaveJetDropDataMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SaveJetDropDataMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSaveJetDropDataCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSaveJetDropData != nil && mm_atomic.LoadUint64(&m.afterSaveJetDropDataCounter) < 1 {
		return false
	}
	return true
}

// MinimockSaveJetDropDataInspect logs each unmet expectation
func (m *StorageMock) MinimockSaveJetDropDataInspect() {
	for _, e := range m.SaveJetDropDataMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.SaveJetDropData with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SaveJetDropDataMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSaveJetDropDataCounter) < 1 {
		if m.SaveJetDropDataMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StorageMock.SaveJetDropData")
		} else {
			m.t.Errorf("Expected call to StorageMock.SaveJetDropData with params: %#v", *m.SaveJetDropDataMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSaveJetDropData != nil && mm_atomic.LoadUint64(&m.afterSaveJetDropDataCounter) < 1 {
		m.t.Error("Expected call to StorageMock.SaveJetDropData")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StorageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSaveJetDropDataInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StorageMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *StorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSaveJetDropDataDone()
}
