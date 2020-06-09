// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.

package mocks

import (
	"context"
	persistence "github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/persistence"
	types "github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/types"
	"sync"
)

// MockReader is a mock implementation of the Reader interface (from the
// package
// github.com/sourcegraph/sourcegraph/internal/codeintel/bundles/persistence)
// used for unit testing.
type MockReader struct {
	// CloseFunc is an instance of a mock function object controlling the
	// behavior of the method Close.
	CloseFunc *ReaderCloseFunc
	// PathsWithPrefixFunc is an instance of a mock function object
	// controlling the behavior of the method PathsWithPrefix.
	PathsWithPrefixFunc *ReaderPathsWithPrefixFunc
	// ReadDefinitionsFunc is an instance of a mock function object
	// controlling the behavior of the method ReadDefinitions.
	ReadDefinitionsFunc *ReaderReadDefinitionsFunc
	// ReadDocumentFunc is an instance of a mock function object controlling
	// the behavior of the method ReadDocument.
	ReadDocumentFunc *ReaderReadDocumentFunc
	// ReadMetaFunc is an instance of a mock function object controlling the
	// behavior of the method ReadMeta.
	ReadMetaFunc *ReaderReadMetaFunc
	// ReadReferencesFunc is an instance of a mock function object
	// controlling the behavior of the method ReadReferences.
	ReadReferencesFunc *ReaderReadReferencesFunc
	// ReadResultChunkFunc is an instance of a mock function object
	// controlling the behavior of the method ReadResultChunk.
	ReadResultChunkFunc *ReaderReadResultChunkFunc
}

// NewMockReader creates a new mock of the Reader interface. All methods
// return zero values for all results, unless overwritten.
func NewMockReader() *MockReader {
	return &MockReader{
		CloseFunc: &ReaderCloseFunc{
			defaultHook: func() error {
				return nil
			},
		},
		PathsWithPrefixFunc: &ReaderPathsWithPrefixFunc{
			defaultHook: func(context.Context, string) ([]string, error) {
				return nil, nil
			},
		},
		ReadDefinitionsFunc: &ReaderReadDefinitionsFunc{
			defaultHook: func(context.Context, string, string, int, int) ([]types.Location, int, error) {
				return nil, 0, nil
			},
		},
		ReadDocumentFunc: &ReaderReadDocumentFunc{
			defaultHook: func(context.Context, string) (types.DocumentData, bool, error) {
				return types.DocumentData{}, false, nil
			},
		},
		ReadMetaFunc: &ReaderReadMetaFunc{
			defaultHook: func(context.Context) (types.MetaData, error) {
				return types.MetaData{}, nil
			},
		},
		ReadReferencesFunc: &ReaderReadReferencesFunc{
			defaultHook: func(context.Context, string, string, int, int) ([]types.Location, int, error) {
				return nil, 0, nil
			},
		},
		ReadResultChunkFunc: &ReaderReadResultChunkFunc{
			defaultHook: func(context.Context, int) (types.ResultChunkData, bool, error) {
				return types.ResultChunkData{}, false, nil
			},
		},
	}
}

// NewMockReaderFrom creates a new mock of the MockReader interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockReaderFrom(i persistence.Reader) *MockReader {
	return &MockReader{
		CloseFunc: &ReaderCloseFunc{
			defaultHook: i.Close,
		},
		PathsWithPrefixFunc: &ReaderPathsWithPrefixFunc{
			defaultHook: i.PathsWithPrefix,
		},
		ReadDefinitionsFunc: &ReaderReadDefinitionsFunc{
			defaultHook: i.ReadDefinitions,
		},
		ReadDocumentFunc: &ReaderReadDocumentFunc{
			defaultHook: i.ReadDocument,
		},
		ReadMetaFunc: &ReaderReadMetaFunc{
			defaultHook: i.ReadMeta,
		},
		ReadReferencesFunc: &ReaderReadReferencesFunc{
			defaultHook: i.ReadReferences,
		},
		ReadResultChunkFunc: &ReaderReadResultChunkFunc{
			defaultHook: i.ReadResultChunk,
		},
	}
}

// ReaderCloseFunc describes the behavior when the Close method of the
// parent MockReader instance is invoked.
type ReaderCloseFunc struct {
	defaultHook func() error
	hooks       []func() error
	history     []ReaderCloseFuncCall
	mutex       sync.Mutex
}

// Close delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockReader) Close() error {
	r0 := m.CloseFunc.nextHook()()
	m.CloseFunc.appendCall(ReaderCloseFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Close method of the
// parent MockReader instance is invoked and the hook queue is empty.
func (f *ReaderCloseFunc) SetDefaultHook(hook func() error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Close method of the parent MockReader instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ReaderCloseFunc) PushHook(hook func() error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ReaderCloseFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func() error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ReaderCloseFunc) PushReturn(r0 error) {
	f.PushHook(func() error {
		return r0
	})
}

func (f *ReaderCloseFunc) nextHook() func() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReaderCloseFunc) appendCall(r0 ReaderCloseFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReaderCloseFuncCall objects describing the
// invocations of this function.
func (f *ReaderCloseFunc) History() []ReaderCloseFuncCall {
	f.mutex.Lock()
	history := make([]ReaderCloseFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReaderCloseFuncCall is an object that describes an invocation of method
// Close on an instance of MockReader.
type ReaderCloseFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReaderCloseFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReaderCloseFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ReaderPathsWithPrefixFunc describes the behavior when the PathsWithPrefix
// method of the parent MockReader instance is invoked.
type ReaderPathsWithPrefixFunc struct {
	defaultHook func(context.Context, string) ([]string, error)
	hooks       []func(context.Context, string) ([]string, error)
	history     []ReaderPathsWithPrefixFuncCall
	mutex       sync.Mutex
}

// PathsWithPrefix delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockReader) PathsWithPrefix(v0 context.Context, v1 string) ([]string, error) {
	r0, r1 := m.PathsWithPrefixFunc.nextHook()(v0, v1)
	m.PathsWithPrefixFunc.appendCall(ReaderPathsWithPrefixFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the PathsWithPrefix
// method of the parent MockReader instance is invoked and the hook queue is
// empty.
func (f *ReaderPathsWithPrefixFunc) SetDefaultHook(hook func(context.Context, string) ([]string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PathsWithPrefix method of the parent MockReader instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReaderPathsWithPrefixFunc) PushHook(hook func(context.Context, string) ([]string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ReaderPathsWithPrefixFunc) SetDefaultReturn(r0 []string, r1 error) {
	f.SetDefaultHook(func(context.Context, string) ([]string, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ReaderPathsWithPrefixFunc) PushReturn(r0 []string, r1 error) {
	f.PushHook(func(context.Context, string) ([]string, error) {
		return r0, r1
	})
}

func (f *ReaderPathsWithPrefixFunc) nextHook() func(context.Context, string) ([]string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReaderPathsWithPrefixFunc) appendCall(r0 ReaderPathsWithPrefixFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReaderPathsWithPrefixFuncCall objects
// describing the invocations of this function.
func (f *ReaderPathsWithPrefixFunc) History() []ReaderPathsWithPrefixFuncCall {
	f.mutex.Lock()
	history := make([]ReaderPathsWithPrefixFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReaderPathsWithPrefixFuncCall is an object that describes an invocation
// of method PathsWithPrefix on an instance of MockReader.
type ReaderPathsWithPrefixFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReaderPathsWithPrefixFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReaderPathsWithPrefixFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// ReaderReadDefinitionsFunc describes the behavior when the ReadDefinitions
// method of the parent MockReader instance is invoked.
type ReaderReadDefinitionsFunc struct {
	defaultHook func(context.Context, string, string, int, int) ([]types.Location, int, error)
	hooks       []func(context.Context, string, string, int, int) ([]types.Location, int, error)
	history     []ReaderReadDefinitionsFuncCall
	mutex       sync.Mutex
}

// ReadDefinitions delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockReader) ReadDefinitions(v0 context.Context, v1 string, v2 string, v3 int, v4 int) ([]types.Location, int, error) {
	r0, r1, r2 := m.ReadDefinitionsFunc.nextHook()(v0, v1, v2, v3, v4)
	m.ReadDefinitionsFunc.appendCall(ReaderReadDefinitionsFuncCall{v0, v1, v2, v3, v4, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the ReadDefinitions
// method of the parent MockReader instance is invoked and the hook queue is
// empty.
func (f *ReaderReadDefinitionsFunc) SetDefaultHook(hook func(context.Context, string, string, int, int) ([]types.Location, int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ReadDefinitions method of the parent MockReader instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReaderReadDefinitionsFunc) PushHook(hook func(context.Context, string, string, int, int) ([]types.Location, int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ReaderReadDefinitionsFunc) SetDefaultReturn(r0 []types.Location, r1 int, r2 error) {
	f.SetDefaultHook(func(context.Context, string, string, int, int) ([]types.Location, int, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ReaderReadDefinitionsFunc) PushReturn(r0 []types.Location, r1 int, r2 error) {
	f.PushHook(func(context.Context, string, string, int, int) ([]types.Location, int, error) {
		return r0, r1, r2
	})
}

func (f *ReaderReadDefinitionsFunc) nextHook() func(context.Context, string, string, int, int) ([]types.Location, int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReaderReadDefinitionsFunc) appendCall(r0 ReaderReadDefinitionsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReaderReadDefinitionsFuncCall objects
// describing the invocations of this function.
func (f *ReaderReadDefinitionsFunc) History() []ReaderReadDefinitionsFuncCall {
	f.mutex.Lock()
	history := make([]ReaderReadDefinitionsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReaderReadDefinitionsFuncCall is an object that describes an invocation
// of method ReadDefinitions on an instance of MockReader.
type ReaderReadDefinitionsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 int
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []types.Location
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 int
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReaderReadDefinitionsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReaderReadDefinitionsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// ReaderReadDocumentFunc describes the behavior when the ReadDocument
// method of the parent MockReader instance is invoked.
type ReaderReadDocumentFunc struct {
	defaultHook func(context.Context, string) (types.DocumentData, bool, error)
	hooks       []func(context.Context, string) (types.DocumentData, bool, error)
	history     []ReaderReadDocumentFuncCall
	mutex       sync.Mutex
}

// ReadDocument delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockReader) ReadDocument(v0 context.Context, v1 string) (types.DocumentData, bool, error) {
	r0, r1, r2 := m.ReadDocumentFunc.nextHook()(v0, v1)
	m.ReadDocumentFunc.appendCall(ReaderReadDocumentFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the ReadDocument method
// of the parent MockReader instance is invoked and the hook queue is empty.
func (f *ReaderReadDocumentFunc) SetDefaultHook(hook func(context.Context, string) (types.DocumentData, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ReadDocument method of the parent MockReader instance inovkes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReaderReadDocumentFunc) PushHook(hook func(context.Context, string) (types.DocumentData, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ReaderReadDocumentFunc) SetDefaultReturn(r0 types.DocumentData, r1 bool, r2 error) {
	f.SetDefaultHook(func(context.Context, string) (types.DocumentData, bool, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ReaderReadDocumentFunc) PushReturn(r0 types.DocumentData, r1 bool, r2 error) {
	f.PushHook(func(context.Context, string) (types.DocumentData, bool, error) {
		return r0, r1, r2
	})
}

func (f *ReaderReadDocumentFunc) nextHook() func(context.Context, string) (types.DocumentData, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReaderReadDocumentFunc) appendCall(r0 ReaderReadDocumentFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReaderReadDocumentFuncCall objects
// describing the invocations of this function.
func (f *ReaderReadDocumentFunc) History() []ReaderReadDocumentFuncCall {
	f.mutex.Lock()
	history := make([]ReaderReadDocumentFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReaderReadDocumentFuncCall is an object that describes an invocation of
// method ReadDocument on an instance of MockReader.
type ReaderReadDocumentFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 types.DocumentData
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReaderReadDocumentFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReaderReadDocumentFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// ReaderReadMetaFunc describes the behavior when the ReadMeta method of the
// parent MockReader instance is invoked.
type ReaderReadMetaFunc struct {
	defaultHook func(context.Context) (types.MetaData, error)
	hooks       []func(context.Context) (types.MetaData, error)
	history     []ReaderReadMetaFuncCall
	mutex       sync.Mutex
}

// ReadMeta delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockReader) ReadMeta(v0 context.Context) (types.MetaData, error) {
	r0, r1 := m.ReadMetaFunc.nextHook()(v0)
	m.ReadMetaFunc.appendCall(ReaderReadMetaFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the ReadMeta method of
// the parent MockReader instance is invoked and the hook queue is empty.
func (f *ReaderReadMetaFunc) SetDefaultHook(hook func(context.Context) (types.MetaData, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ReadMeta method of the parent MockReader instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ReaderReadMetaFunc) PushHook(hook func(context.Context) (types.MetaData, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ReaderReadMetaFunc) SetDefaultReturn(r0 types.MetaData, r1 error) {
	f.SetDefaultHook(func(context.Context) (types.MetaData, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ReaderReadMetaFunc) PushReturn(r0 types.MetaData, r1 error) {
	f.PushHook(func(context.Context) (types.MetaData, error) {
		return r0, r1
	})
}

func (f *ReaderReadMetaFunc) nextHook() func(context.Context) (types.MetaData, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReaderReadMetaFunc) appendCall(r0 ReaderReadMetaFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReaderReadMetaFuncCall objects describing
// the invocations of this function.
func (f *ReaderReadMetaFunc) History() []ReaderReadMetaFuncCall {
	f.mutex.Lock()
	history := make([]ReaderReadMetaFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReaderReadMetaFuncCall is an object that describes an invocation of
// method ReadMeta on an instance of MockReader.
type ReaderReadMetaFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 types.MetaData
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReaderReadMetaFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReaderReadMetaFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// ReaderReadReferencesFunc describes the behavior when the ReadReferences
// method of the parent MockReader instance is invoked.
type ReaderReadReferencesFunc struct {
	defaultHook func(context.Context, string, string, int, int) ([]types.Location, int, error)
	hooks       []func(context.Context, string, string, int, int) ([]types.Location, int, error)
	history     []ReaderReadReferencesFuncCall
	mutex       sync.Mutex
}

// ReadReferences delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockReader) ReadReferences(v0 context.Context, v1 string, v2 string, v3 int, v4 int) ([]types.Location, int, error) {
	r0, r1, r2 := m.ReadReferencesFunc.nextHook()(v0, v1, v2, v3, v4)
	m.ReadReferencesFunc.appendCall(ReaderReadReferencesFuncCall{v0, v1, v2, v3, v4, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the ReadReferences
// method of the parent MockReader instance is invoked and the hook queue is
// empty.
func (f *ReaderReadReferencesFunc) SetDefaultHook(hook func(context.Context, string, string, int, int) ([]types.Location, int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ReadReferences method of the parent MockReader instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReaderReadReferencesFunc) PushHook(hook func(context.Context, string, string, int, int) ([]types.Location, int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ReaderReadReferencesFunc) SetDefaultReturn(r0 []types.Location, r1 int, r2 error) {
	f.SetDefaultHook(func(context.Context, string, string, int, int) ([]types.Location, int, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ReaderReadReferencesFunc) PushReturn(r0 []types.Location, r1 int, r2 error) {
	f.PushHook(func(context.Context, string, string, int, int) ([]types.Location, int, error) {
		return r0, r1, r2
	})
}

func (f *ReaderReadReferencesFunc) nextHook() func(context.Context, string, string, int, int) ([]types.Location, int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReaderReadReferencesFunc) appendCall(r0 ReaderReadReferencesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReaderReadReferencesFuncCall objects
// describing the invocations of this function.
func (f *ReaderReadReferencesFunc) History() []ReaderReadReferencesFuncCall {
	f.mutex.Lock()
	history := make([]ReaderReadReferencesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReaderReadReferencesFuncCall is an object that describes an invocation of
// method ReadReferences on an instance of MockReader.
type ReaderReadReferencesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 int
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []types.Location
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 int
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReaderReadReferencesFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReaderReadReferencesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// ReaderReadResultChunkFunc describes the behavior when the ReadResultChunk
// method of the parent MockReader instance is invoked.
type ReaderReadResultChunkFunc struct {
	defaultHook func(context.Context, int) (types.ResultChunkData, bool, error)
	hooks       []func(context.Context, int) (types.ResultChunkData, bool, error)
	history     []ReaderReadResultChunkFuncCall
	mutex       sync.Mutex
}

// ReadResultChunk delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockReader) ReadResultChunk(v0 context.Context, v1 int) (types.ResultChunkData, bool, error) {
	r0, r1, r2 := m.ReadResultChunkFunc.nextHook()(v0, v1)
	m.ReadResultChunkFunc.appendCall(ReaderReadResultChunkFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the ReadResultChunk
// method of the parent MockReader instance is invoked and the hook queue is
// empty.
func (f *ReaderReadResultChunkFunc) SetDefaultHook(hook func(context.Context, int) (types.ResultChunkData, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ReadResultChunk method of the parent MockReader instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReaderReadResultChunkFunc) PushHook(hook func(context.Context, int) (types.ResultChunkData, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ReaderReadResultChunkFunc) SetDefaultReturn(r0 types.ResultChunkData, r1 bool, r2 error) {
	f.SetDefaultHook(func(context.Context, int) (types.ResultChunkData, bool, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ReaderReadResultChunkFunc) PushReturn(r0 types.ResultChunkData, r1 bool, r2 error) {
	f.PushHook(func(context.Context, int) (types.ResultChunkData, bool, error) {
		return r0, r1, r2
	})
}

func (f *ReaderReadResultChunkFunc) nextHook() func(context.Context, int) (types.ResultChunkData, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReaderReadResultChunkFunc) appendCall(r0 ReaderReadResultChunkFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReaderReadResultChunkFuncCall objects
// describing the invocations of this function.
func (f *ReaderReadResultChunkFunc) History() []ReaderReadResultChunkFuncCall {
	f.mutex.Lock()
	history := make([]ReaderReadResultChunkFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReaderReadResultChunkFuncCall is an object that describes an invocation
// of method ReadResultChunk on an instance of MockReader.
type ReaderReadResultChunkFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 types.ResultChunkData
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReaderReadResultChunkFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReaderReadResultChunkFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}