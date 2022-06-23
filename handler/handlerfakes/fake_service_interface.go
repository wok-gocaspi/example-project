// Code generated by counterfeiter. DO NOT EDIT.
package handlerfakes

import (
	"example-project/handler"
	"example-project/model"
	"sync"
)

type FakeServiceInterface struct {
	CreateEmployeesStub        func([]model.Employee) interface{}
	createEmployeesMutex       sync.RWMutex
	createEmployeesArgsForCall []struct {
		arg1 []model.Employee
	}
	createEmployeesReturns struct {
		result1 interface{}
	}
	createEmployeesReturnsOnCall map[int]struct {
		result1 interface{}
	}
	DeleteEmployeeByIdStub        func(string) interface{}
	deleteEmployeeByIdMutex       sync.RWMutex
	deleteEmployeeByIdArgsForCall []struct {
		arg1 string
	}
	deleteEmployeeByIdReturns struct {
		result1 interface{}
	}
	deleteEmployeeByIdReturnsOnCall map[int]struct {
		result1 interface{}
	}
	GetEmployeeByIdStub        func(string) model.Employee
	getEmployeeByIdMutex       sync.RWMutex
	getEmployeeByIdArgsForCall []struct {
		arg1 string
	}
	getEmployeeByIdReturns struct {
		result1 model.Employee
	}
	getEmployeeByIdReturnsOnCall map[int]struct {
		result1 model.Employee
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceInterface) CreateEmployees(arg1 []model.Employee) interface{} {
	var arg1Copy []model.Employee
	if arg1 != nil {
		arg1Copy = make([]model.Employee, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createEmployeesMutex.Lock()
	ret, specificReturn := fake.createEmployeesReturnsOnCall[len(fake.createEmployeesArgsForCall)]
	fake.createEmployeesArgsForCall = append(fake.createEmployeesArgsForCall, struct {
		arg1 []model.Employee
	}{arg1Copy})
	stub := fake.CreateEmployeesStub
	fakeReturns := fake.createEmployeesReturns
	fake.recordInvocation("CreateEmployees", []interface{}{arg1Copy})
	fake.createEmployeesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceInterface) CreateEmployeesCallCount() int {
	fake.createEmployeesMutex.RLock()
	defer fake.createEmployeesMutex.RUnlock()
	return len(fake.createEmployeesArgsForCall)
}

func (fake *FakeServiceInterface) CreateEmployeesCalls(stub func([]model.Employee) interface{}) {
	fake.createEmployeesMutex.Lock()
	defer fake.createEmployeesMutex.Unlock()
	fake.CreateEmployeesStub = stub
}

func (fake *FakeServiceInterface) CreateEmployeesArgsForCall(i int) []model.Employee {
	fake.createEmployeesMutex.RLock()
	defer fake.createEmployeesMutex.RUnlock()
	argsForCall := fake.createEmployeesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) CreateEmployeesReturns(result1 interface{}) {
	fake.createEmployeesMutex.Lock()
	defer fake.createEmployeesMutex.Unlock()
	fake.CreateEmployeesStub = nil
	fake.createEmployeesReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeServiceInterface) CreateEmployeesReturnsOnCall(i int, result1 interface{}) {
	fake.createEmployeesMutex.Lock()
	defer fake.createEmployeesMutex.Unlock()
	fake.CreateEmployeesStub = nil
	if fake.createEmployeesReturnsOnCall == nil {
		fake.createEmployeesReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.createEmployeesReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeServiceInterface) DeleteEmployeeById(arg1 string) interface{} {
	fake.deleteEmployeeByIdMutex.Lock()
	ret, specificReturn := fake.deleteEmployeeByIdReturnsOnCall[len(fake.deleteEmployeeByIdArgsForCall)]
	fake.deleteEmployeeByIdArgsForCall = append(fake.deleteEmployeeByIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteEmployeeByIdStub
	fakeReturns := fake.deleteEmployeeByIdReturns
	fake.recordInvocation("DeleteEmployeeById", []interface{}{arg1})
	fake.deleteEmployeeByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdCallCount() int {
	fake.deleteEmployeeByIdMutex.RLock()
	defer fake.deleteEmployeeByIdMutex.RUnlock()
	return len(fake.deleteEmployeeByIdArgsForCall)
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdCalls(stub func(string) interface{}) {
	fake.deleteEmployeeByIdMutex.Lock()
	defer fake.deleteEmployeeByIdMutex.Unlock()
	fake.DeleteEmployeeByIdStub = stub
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdArgsForCall(i int) string {
	fake.deleteEmployeeByIdMutex.RLock()
	defer fake.deleteEmployeeByIdMutex.RUnlock()
	argsForCall := fake.deleteEmployeeByIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdReturns(result1 interface{}) {
	fake.deleteEmployeeByIdMutex.Lock()
	defer fake.deleteEmployeeByIdMutex.Unlock()
	fake.DeleteEmployeeByIdStub = nil
	fake.deleteEmployeeByIdReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeServiceInterface) DeleteEmployeeByIdReturnsOnCall(i int, result1 interface{}) {
	fake.deleteEmployeeByIdMutex.Lock()
	defer fake.deleteEmployeeByIdMutex.Unlock()
	fake.DeleteEmployeeByIdStub = nil
	if fake.deleteEmployeeByIdReturnsOnCall == nil {
		fake.deleteEmployeeByIdReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.deleteEmployeeByIdReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeServiceInterface) GetEmployeeById(arg1 string) model.Employee {
	fake.getEmployeeByIdMutex.Lock()
	ret, specificReturn := fake.getEmployeeByIdReturnsOnCall[len(fake.getEmployeeByIdArgsForCall)]
	fake.getEmployeeByIdArgsForCall = append(fake.getEmployeeByIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetEmployeeByIdStub
	fakeReturns := fake.getEmployeeByIdReturns
	fake.recordInvocation("GetEmployeeById", []interface{}{arg1})
	fake.getEmployeeByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceInterface) GetEmployeeByIdCallCount() int {
	fake.getEmployeeByIdMutex.RLock()
	defer fake.getEmployeeByIdMutex.RUnlock()
	return len(fake.getEmployeeByIdArgsForCall)
}

func (fake *FakeServiceInterface) GetEmployeeByIdCalls(stub func(string) model.Employee) {
	fake.getEmployeeByIdMutex.Lock()
	defer fake.getEmployeeByIdMutex.Unlock()
	fake.GetEmployeeByIdStub = stub
}

func (fake *FakeServiceInterface) GetEmployeeByIdArgsForCall(i int) string {
	fake.getEmployeeByIdMutex.RLock()
	defer fake.getEmployeeByIdMutex.RUnlock()
	argsForCall := fake.getEmployeeByIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) GetEmployeeByIdReturns(result1 model.Employee) {
	fake.getEmployeeByIdMutex.Lock()
	defer fake.getEmployeeByIdMutex.Unlock()
	fake.GetEmployeeByIdStub = nil
	fake.getEmployeeByIdReturns = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeServiceInterface) GetEmployeeByIdReturnsOnCall(i int, result1 model.Employee) {
	fake.getEmployeeByIdMutex.Lock()
	defer fake.getEmployeeByIdMutex.Unlock()
	fake.GetEmployeeByIdStub = nil
	if fake.getEmployeeByIdReturnsOnCall == nil {
		fake.getEmployeeByIdReturnsOnCall = make(map[int]struct {
			result1 model.Employee
		})
	}
	fake.getEmployeeByIdReturnsOnCall[i] = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeServiceInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createEmployeesMutex.RLock()
	defer fake.createEmployeesMutex.RUnlock()
	fake.deleteEmployeeByIdMutex.RLock()
	defer fake.deleteEmployeeByIdMutex.RUnlock()
	fake.getEmployeeByIdMutex.RLock()
	defer fake.getEmployeeByIdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handler.ServiceInterface = new(FakeServiceInterface)
