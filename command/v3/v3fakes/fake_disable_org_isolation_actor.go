// This file was generated by counterfeiter
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeDisableOrgIsolationActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	RevokeIsolationSegmentFromOrganizationByNameStub        func(isolationSegmentName string, orgName string) (v3action.Warnings, error)
	revokeIsolationSegmentFromOrganizationByNameMutex       sync.RWMutex
	revokeIsolationSegmentFromOrganizationByNameArgsForCall []struct {
		isolationSegmentName string
		orgName              string
	}
	revokeIsolationSegmentFromOrganizationByNameReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	revokeIsolationSegmentFromOrganizationByNameReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDisableOrgIsolationActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeDisableOrgIsolationActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeDisableOrgIsolationActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeDisableOrgIsolationActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeDisableOrgIsolationActor) RevokeIsolationSegmentFromOrganizationByName(isolationSegmentName string, orgName string) (v3action.Warnings, error) {
	fake.revokeIsolationSegmentFromOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.revokeIsolationSegmentFromOrganizationByNameReturnsOnCall[len(fake.revokeIsolationSegmentFromOrganizationByNameArgsForCall)]
	fake.revokeIsolationSegmentFromOrganizationByNameArgsForCall = append(fake.revokeIsolationSegmentFromOrganizationByNameArgsForCall, struct {
		isolationSegmentName string
		orgName              string
	}{isolationSegmentName, orgName})
	fake.recordInvocation("RevokeIsolationSegmentFromOrganizationByName", []interface{}{isolationSegmentName, orgName})
	fake.revokeIsolationSegmentFromOrganizationByNameMutex.Unlock()
	if fake.RevokeIsolationSegmentFromOrganizationByNameStub != nil {
		return fake.RevokeIsolationSegmentFromOrganizationByNameStub(isolationSegmentName, orgName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.revokeIsolationSegmentFromOrganizationByNameReturns.result1, fake.revokeIsolationSegmentFromOrganizationByNameReturns.result2
}

func (fake *FakeDisableOrgIsolationActor) RevokeIsolationSegmentFromOrganizationByNameCallCount() int {
	fake.revokeIsolationSegmentFromOrganizationByNameMutex.RLock()
	defer fake.revokeIsolationSegmentFromOrganizationByNameMutex.RUnlock()
	return len(fake.revokeIsolationSegmentFromOrganizationByNameArgsForCall)
}

func (fake *FakeDisableOrgIsolationActor) RevokeIsolationSegmentFromOrganizationByNameArgsForCall(i int) (string, string) {
	fake.revokeIsolationSegmentFromOrganizationByNameMutex.RLock()
	defer fake.revokeIsolationSegmentFromOrganizationByNameMutex.RUnlock()
	return fake.revokeIsolationSegmentFromOrganizationByNameArgsForCall[i].isolationSegmentName, fake.revokeIsolationSegmentFromOrganizationByNameArgsForCall[i].orgName
}

func (fake *FakeDisableOrgIsolationActor) RevokeIsolationSegmentFromOrganizationByNameReturns(result1 v3action.Warnings, result2 error) {
	fake.RevokeIsolationSegmentFromOrganizationByNameStub = nil
	fake.revokeIsolationSegmentFromOrganizationByNameReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDisableOrgIsolationActor) RevokeIsolationSegmentFromOrganizationByNameReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.RevokeIsolationSegmentFromOrganizationByNameStub = nil
	if fake.revokeIsolationSegmentFromOrganizationByNameReturnsOnCall == nil {
		fake.revokeIsolationSegmentFromOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.revokeIsolationSegmentFromOrganizationByNameReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDisableOrgIsolationActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.revokeIsolationSegmentFromOrganizationByNameMutex.RLock()
	defer fake.revokeIsolationSegmentFromOrganizationByNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDisableOrgIsolationActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.DisableOrgIsolationActor = new(FakeDisableOrgIsolationActor)