// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	backend "github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/softlayer/backend"
	datatypes "github.com/softlayer/softlayer-go/datatypes"
)

type AccountService struct {
	FilterStub        func(string) backend.AccountService
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 string
	}
	filterReturns struct {
		result1 backend.AccountService
	}
	filterReturnsOnCall map[int]struct {
		result1 backend.AccountService
	}
	GetBlockDeviceTemplateGroupsStub        func() ([]datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	getBlockDeviceTemplateGroupsMutex       sync.RWMutex
	getBlockDeviceTemplateGroupsArgsForCall []struct {
	}
	getBlockDeviceTemplateGroupsReturns struct {
		result1 []datatypes.Virtual_Guest_Block_Device_Template_Group
		result2 error
	}
	getBlockDeviceTemplateGroupsReturnsOnCall map[int]struct {
		result1 []datatypes.Virtual_Guest_Block_Device_Template_Group
		result2 error
	}
	GetCurrentUserStub        func() (datatypes.User_Customer, error)
	getCurrentUserMutex       sync.RWMutex
	getCurrentUserArgsForCall []struct {
	}
	getCurrentUserReturns struct {
		result1 datatypes.User_Customer
		result2 error
	}
	getCurrentUserReturnsOnCall map[int]struct {
		result1 datatypes.User_Customer
		result2 error
	}
	GetHardwareStub        func() ([]datatypes.Hardware, error)
	getHardwareMutex       sync.RWMutex
	getHardwareArgsForCall []struct {
	}
	getHardwareReturns struct {
		result1 []datatypes.Hardware
		result2 error
	}
	getHardwareReturnsOnCall map[int]struct {
		result1 []datatypes.Hardware
		result2 error
	}
	GetNetworkStorageStub        func() ([]datatypes.Network_Storage, error)
	getNetworkStorageMutex       sync.RWMutex
	getNetworkStorageArgsForCall []struct {
	}
	getNetworkStorageReturns struct {
		result1 []datatypes.Network_Storage
		result2 error
	}
	getNetworkStorageReturnsOnCall map[int]struct {
		result1 []datatypes.Network_Storage
		result2 error
	}
	GetObjectStub        func() (datatypes.Account, error)
	getObjectMutex       sync.RWMutex
	getObjectArgsForCall []struct {
	}
	getObjectReturns struct {
		result1 datatypes.Account
		result2 error
	}
	getObjectReturnsOnCall map[int]struct {
		result1 datatypes.Account
		result2 error
	}
	GetOrdersStub        func() ([]datatypes.Billing_Order, error)
	getOrdersMutex       sync.RWMutex
	getOrdersArgsForCall []struct {
	}
	getOrdersReturns struct {
		result1 []datatypes.Billing_Order
		result2 error
	}
	getOrdersReturnsOnCall map[int]struct {
		result1 []datatypes.Billing_Order
		result2 error
	}
	GetSubnetsStub        func() ([]datatypes.Network_Subnet, error)
	getSubnetsMutex       sync.RWMutex
	getSubnetsArgsForCall []struct {
	}
	getSubnetsReturns struct {
		result1 []datatypes.Network_Subnet
		result2 error
	}
	getSubnetsReturnsOnCall map[int]struct {
		result1 []datatypes.Network_Subnet
		result2 error
	}
	IDStub        func(int) backend.AccountService
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
		arg1 int
	}
	iDReturns struct {
		result1 backend.AccountService
	}
	iDReturnsOnCall map[int]struct {
		result1 backend.AccountService
	}
	LimitStub        func(int) backend.AccountService
	limitMutex       sync.RWMutex
	limitArgsForCall []struct {
		arg1 int
	}
	limitReturns struct {
		result1 backend.AccountService
	}
	limitReturnsOnCall map[int]struct {
		result1 backend.AccountService
	}
	MaskStub        func(string) backend.AccountService
	maskMutex       sync.RWMutex
	maskArgsForCall []struct {
		arg1 string
	}
	maskReturns struct {
		result1 backend.AccountService
	}
	maskReturnsOnCall map[int]struct {
		result1 backend.AccountService
	}
	OffsetStub        func(int) backend.AccountService
	offsetMutex       sync.RWMutex
	offsetArgsForCall []struct {
		arg1 int
	}
	offsetReturns struct {
		result1 backend.AccountService
	}
	offsetReturnsOnCall map[int]struct {
		result1 backend.AccountService
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AccountService) Filter(arg1 string) backend.AccountService {
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Filter", []interface{}{arg1})
	fake.filterMutex.Unlock()
	if fake.FilterStub != nil {
		return fake.FilterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.filterReturns
	return fakeReturns.result1
}

func (fake *AccountService) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *AccountService) FilterCalls(stub func(string) backend.AccountService) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = stub
}

func (fake *AccountService) FilterArgsForCall(i int) string {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	argsForCall := fake.filterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AccountService) FilterReturns(result1 backend.AccountService) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) FilterReturnsOnCall(i int, result1 backend.AccountService) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 backend.AccountService
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) GetBlockDeviceTemplateGroups() ([]datatypes.Virtual_Guest_Block_Device_Template_Group, error) {
	fake.getBlockDeviceTemplateGroupsMutex.Lock()
	ret, specificReturn := fake.getBlockDeviceTemplateGroupsReturnsOnCall[len(fake.getBlockDeviceTemplateGroupsArgsForCall)]
	fake.getBlockDeviceTemplateGroupsArgsForCall = append(fake.getBlockDeviceTemplateGroupsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBlockDeviceTemplateGroups", []interface{}{})
	fake.getBlockDeviceTemplateGroupsMutex.Unlock()
	if fake.GetBlockDeviceTemplateGroupsStub != nil {
		return fake.GetBlockDeviceTemplateGroupsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBlockDeviceTemplateGroupsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AccountService) GetBlockDeviceTemplateGroupsCallCount() int {
	fake.getBlockDeviceTemplateGroupsMutex.RLock()
	defer fake.getBlockDeviceTemplateGroupsMutex.RUnlock()
	return len(fake.getBlockDeviceTemplateGroupsArgsForCall)
}

func (fake *AccountService) GetBlockDeviceTemplateGroupsCalls(stub func() ([]datatypes.Virtual_Guest_Block_Device_Template_Group, error)) {
	fake.getBlockDeviceTemplateGroupsMutex.Lock()
	defer fake.getBlockDeviceTemplateGroupsMutex.Unlock()
	fake.GetBlockDeviceTemplateGroupsStub = stub
}

func (fake *AccountService) GetBlockDeviceTemplateGroupsReturns(result1 []datatypes.Virtual_Guest_Block_Device_Template_Group, result2 error) {
	fake.getBlockDeviceTemplateGroupsMutex.Lock()
	defer fake.getBlockDeviceTemplateGroupsMutex.Unlock()
	fake.GetBlockDeviceTemplateGroupsStub = nil
	fake.getBlockDeviceTemplateGroupsReturns = struct {
		result1 []datatypes.Virtual_Guest_Block_Device_Template_Group
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetBlockDeviceTemplateGroupsReturnsOnCall(i int, result1 []datatypes.Virtual_Guest_Block_Device_Template_Group, result2 error) {
	fake.getBlockDeviceTemplateGroupsMutex.Lock()
	defer fake.getBlockDeviceTemplateGroupsMutex.Unlock()
	fake.GetBlockDeviceTemplateGroupsStub = nil
	if fake.getBlockDeviceTemplateGroupsReturnsOnCall == nil {
		fake.getBlockDeviceTemplateGroupsReturnsOnCall = make(map[int]struct {
			result1 []datatypes.Virtual_Guest_Block_Device_Template_Group
			result2 error
		})
	}
	fake.getBlockDeviceTemplateGroupsReturnsOnCall[i] = struct {
		result1 []datatypes.Virtual_Guest_Block_Device_Template_Group
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetCurrentUser() (datatypes.User_Customer, error) {
	fake.getCurrentUserMutex.Lock()
	ret, specificReturn := fake.getCurrentUserReturnsOnCall[len(fake.getCurrentUserArgsForCall)]
	fake.getCurrentUserArgsForCall = append(fake.getCurrentUserArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCurrentUser", []interface{}{})
	fake.getCurrentUserMutex.Unlock()
	if fake.GetCurrentUserStub != nil {
		return fake.GetCurrentUserStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCurrentUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AccountService) GetCurrentUserCallCount() int {
	fake.getCurrentUserMutex.RLock()
	defer fake.getCurrentUserMutex.RUnlock()
	return len(fake.getCurrentUserArgsForCall)
}

func (fake *AccountService) GetCurrentUserCalls(stub func() (datatypes.User_Customer, error)) {
	fake.getCurrentUserMutex.Lock()
	defer fake.getCurrentUserMutex.Unlock()
	fake.GetCurrentUserStub = stub
}

func (fake *AccountService) GetCurrentUserReturns(result1 datatypes.User_Customer, result2 error) {
	fake.getCurrentUserMutex.Lock()
	defer fake.getCurrentUserMutex.Unlock()
	fake.GetCurrentUserStub = nil
	fake.getCurrentUserReturns = struct {
		result1 datatypes.User_Customer
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetCurrentUserReturnsOnCall(i int, result1 datatypes.User_Customer, result2 error) {
	fake.getCurrentUserMutex.Lock()
	defer fake.getCurrentUserMutex.Unlock()
	fake.GetCurrentUserStub = nil
	if fake.getCurrentUserReturnsOnCall == nil {
		fake.getCurrentUserReturnsOnCall = make(map[int]struct {
			result1 datatypes.User_Customer
			result2 error
		})
	}
	fake.getCurrentUserReturnsOnCall[i] = struct {
		result1 datatypes.User_Customer
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetHardware() ([]datatypes.Hardware, error) {
	fake.getHardwareMutex.Lock()
	ret, specificReturn := fake.getHardwareReturnsOnCall[len(fake.getHardwareArgsForCall)]
	fake.getHardwareArgsForCall = append(fake.getHardwareArgsForCall, struct {
	}{})
	fake.recordInvocation("GetHardware", []interface{}{})
	fake.getHardwareMutex.Unlock()
	if fake.GetHardwareStub != nil {
		return fake.GetHardwareStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getHardwareReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AccountService) GetHardwareCallCount() int {
	fake.getHardwareMutex.RLock()
	defer fake.getHardwareMutex.RUnlock()
	return len(fake.getHardwareArgsForCall)
}

func (fake *AccountService) GetHardwareCalls(stub func() ([]datatypes.Hardware, error)) {
	fake.getHardwareMutex.Lock()
	defer fake.getHardwareMutex.Unlock()
	fake.GetHardwareStub = stub
}

func (fake *AccountService) GetHardwareReturns(result1 []datatypes.Hardware, result2 error) {
	fake.getHardwareMutex.Lock()
	defer fake.getHardwareMutex.Unlock()
	fake.GetHardwareStub = nil
	fake.getHardwareReturns = struct {
		result1 []datatypes.Hardware
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetHardwareReturnsOnCall(i int, result1 []datatypes.Hardware, result2 error) {
	fake.getHardwareMutex.Lock()
	defer fake.getHardwareMutex.Unlock()
	fake.GetHardwareStub = nil
	if fake.getHardwareReturnsOnCall == nil {
		fake.getHardwareReturnsOnCall = make(map[int]struct {
			result1 []datatypes.Hardware
			result2 error
		})
	}
	fake.getHardwareReturnsOnCall[i] = struct {
		result1 []datatypes.Hardware
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetNetworkStorage() ([]datatypes.Network_Storage, error) {
	fake.getNetworkStorageMutex.Lock()
	ret, specificReturn := fake.getNetworkStorageReturnsOnCall[len(fake.getNetworkStorageArgsForCall)]
	fake.getNetworkStorageArgsForCall = append(fake.getNetworkStorageArgsForCall, struct {
	}{})
	fake.recordInvocation("GetNetworkStorage", []interface{}{})
	fake.getNetworkStorageMutex.Unlock()
	if fake.GetNetworkStorageStub != nil {
		return fake.GetNetworkStorageStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getNetworkStorageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AccountService) GetNetworkStorageCallCount() int {
	fake.getNetworkStorageMutex.RLock()
	defer fake.getNetworkStorageMutex.RUnlock()
	return len(fake.getNetworkStorageArgsForCall)
}

func (fake *AccountService) GetNetworkStorageCalls(stub func() ([]datatypes.Network_Storage, error)) {
	fake.getNetworkStorageMutex.Lock()
	defer fake.getNetworkStorageMutex.Unlock()
	fake.GetNetworkStorageStub = stub
}

func (fake *AccountService) GetNetworkStorageReturns(result1 []datatypes.Network_Storage, result2 error) {
	fake.getNetworkStorageMutex.Lock()
	defer fake.getNetworkStorageMutex.Unlock()
	fake.GetNetworkStorageStub = nil
	fake.getNetworkStorageReturns = struct {
		result1 []datatypes.Network_Storage
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetNetworkStorageReturnsOnCall(i int, result1 []datatypes.Network_Storage, result2 error) {
	fake.getNetworkStorageMutex.Lock()
	defer fake.getNetworkStorageMutex.Unlock()
	fake.GetNetworkStorageStub = nil
	if fake.getNetworkStorageReturnsOnCall == nil {
		fake.getNetworkStorageReturnsOnCall = make(map[int]struct {
			result1 []datatypes.Network_Storage
			result2 error
		})
	}
	fake.getNetworkStorageReturnsOnCall[i] = struct {
		result1 []datatypes.Network_Storage
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetObject() (datatypes.Account, error) {
	fake.getObjectMutex.Lock()
	ret, specificReturn := fake.getObjectReturnsOnCall[len(fake.getObjectArgsForCall)]
	fake.getObjectArgsForCall = append(fake.getObjectArgsForCall, struct {
	}{})
	fake.recordInvocation("GetObject", []interface{}{})
	fake.getObjectMutex.Unlock()
	if fake.GetObjectStub != nil {
		return fake.GetObjectStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getObjectReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AccountService) GetObjectCallCount() int {
	fake.getObjectMutex.RLock()
	defer fake.getObjectMutex.RUnlock()
	return len(fake.getObjectArgsForCall)
}

func (fake *AccountService) GetObjectCalls(stub func() (datatypes.Account, error)) {
	fake.getObjectMutex.Lock()
	defer fake.getObjectMutex.Unlock()
	fake.GetObjectStub = stub
}

func (fake *AccountService) GetObjectReturns(result1 datatypes.Account, result2 error) {
	fake.getObjectMutex.Lock()
	defer fake.getObjectMutex.Unlock()
	fake.GetObjectStub = nil
	fake.getObjectReturns = struct {
		result1 datatypes.Account
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetObjectReturnsOnCall(i int, result1 datatypes.Account, result2 error) {
	fake.getObjectMutex.Lock()
	defer fake.getObjectMutex.Unlock()
	fake.GetObjectStub = nil
	if fake.getObjectReturnsOnCall == nil {
		fake.getObjectReturnsOnCall = make(map[int]struct {
			result1 datatypes.Account
			result2 error
		})
	}
	fake.getObjectReturnsOnCall[i] = struct {
		result1 datatypes.Account
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetOrders() ([]datatypes.Billing_Order, error) {
	fake.getOrdersMutex.Lock()
	ret, specificReturn := fake.getOrdersReturnsOnCall[len(fake.getOrdersArgsForCall)]
	fake.getOrdersArgsForCall = append(fake.getOrdersArgsForCall, struct {
	}{})
	fake.recordInvocation("GetOrders", []interface{}{})
	fake.getOrdersMutex.Unlock()
	if fake.GetOrdersStub != nil {
		return fake.GetOrdersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getOrdersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AccountService) GetOrdersCallCount() int {
	fake.getOrdersMutex.RLock()
	defer fake.getOrdersMutex.RUnlock()
	return len(fake.getOrdersArgsForCall)
}

func (fake *AccountService) GetOrdersCalls(stub func() ([]datatypes.Billing_Order, error)) {
	fake.getOrdersMutex.Lock()
	defer fake.getOrdersMutex.Unlock()
	fake.GetOrdersStub = stub
}

func (fake *AccountService) GetOrdersReturns(result1 []datatypes.Billing_Order, result2 error) {
	fake.getOrdersMutex.Lock()
	defer fake.getOrdersMutex.Unlock()
	fake.GetOrdersStub = nil
	fake.getOrdersReturns = struct {
		result1 []datatypes.Billing_Order
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetOrdersReturnsOnCall(i int, result1 []datatypes.Billing_Order, result2 error) {
	fake.getOrdersMutex.Lock()
	defer fake.getOrdersMutex.Unlock()
	fake.GetOrdersStub = nil
	if fake.getOrdersReturnsOnCall == nil {
		fake.getOrdersReturnsOnCall = make(map[int]struct {
			result1 []datatypes.Billing_Order
			result2 error
		})
	}
	fake.getOrdersReturnsOnCall[i] = struct {
		result1 []datatypes.Billing_Order
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetSubnets() ([]datatypes.Network_Subnet, error) {
	fake.getSubnetsMutex.Lock()
	ret, specificReturn := fake.getSubnetsReturnsOnCall[len(fake.getSubnetsArgsForCall)]
	fake.getSubnetsArgsForCall = append(fake.getSubnetsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetSubnets", []interface{}{})
	fake.getSubnetsMutex.Unlock()
	if fake.GetSubnetsStub != nil {
		return fake.GetSubnetsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSubnetsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AccountService) GetSubnetsCallCount() int {
	fake.getSubnetsMutex.RLock()
	defer fake.getSubnetsMutex.RUnlock()
	return len(fake.getSubnetsArgsForCall)
}

func (fake *AccountService) GetSubnetsCalls(stub func() ([]datatypes.Network_Subnet, error)) {
	fake.getSubnetsMutex.Lock()
	defer fake.getSubnetsMutex.Unlock()
	fake.GetSubnetsStub = stub
}

func (fake *AccountService) GetSubnetsReturns(result1 []datatypes.Network_Subnet, result2 error) {
	fake.getSubnetsMutex.Lock()
	defer fake.getSubnetsMutex.Unlock()
	fake.GetSubnetsStub = nil
	fake.getSubnetsReturns = struct {
		result1 []datatypes.Network_Subnet
		result2 error
	}{result1, result2}
}

func (fake *AccountService) GetSubnetsReturnsOnCall(i int, result1 []datatypes.Network_Subnet, result2 error) {
	fake.getSubnetsMutex.Lock()
	defer fake.getSubnetsMutex.Unlock()
	fake.GetSubnetsStub = nil
	if fake.getSubnetsReturnsOnCall == nil {
		fake.getSubnetsReturnsOnCall = make(map[int]struct {
			result1 []datatypes.Network_Subnet
			result2 error
		})
	}
	fake.getSubnetsReturnsOnCall[i] = struct {
		result1 []datatypes.Network_Subnet
		result2 error
	}{result1, result2}
}

func (fake *AccountService) ID(arg1 int) backend.AccountService {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("ID", []interface{}{arg1})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *AccountService) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *AccountService) IDCalls(stub func(int) backend.AccountService) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *AccountService) IDArgsForCall(i int) int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	argsForCall := fake.iDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AccountService) IDReturns(result1 backend.AccountService) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) IDReturnsOnCall(i int, result1 backend.AccountService) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 backend.AccountService
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) Limit(arg1 int) backend.AccountService {
	fake.limitMutex.Lock()
	ret, specificReturn := fake.limitReturnsOnCall[len(fake.limitArgsForCall)]
	fake.limitArgsForCall = append(fake.limitArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Limit", []interface{}{arg1})
	fake.limitMutex.Unlock()
	if fake.LimitStub != nil {
		return fake.LimitStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.limitReturns
	return fakeReturns.result1
}

func (fake *AccountService) LimitCallCount() int {
	fake.limitMutex.RLock()
	defer fake.limitMutex.RUnlock()
	return len(fake.limitArgsForCall)
}

func (fake *AccountService) LimitCalls(stub func(int) backend.AccountService) {
	fake.limitMutex.Lock()
	defer fake.limitMutex.Unlock()
	fake.LimitStub = stub
}

func (fake *AccountService) LimitArgsForCall(i int) int {
	fake.limitMutex.RLock()
	defer fake.limitMutex.RUnlock()
	argsForCall := fake.limitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AccountService) LimitReturns(result1 backend.AccountService) {
	fake.limitMutex.Lock()
	defer fake.limitMutex.Unlock()
	fake.LimitStub = nil
	fake.limitReturns = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) LimitReturnsOnCall(i int, result1 backend.AccountService) {
	fake.limitMutex.Lock()
	defer fake.limitMutex.Unlock()
	fake.LimitStub = nil
	if fake.limitReturnsOnCall == nil {
		fake.limitReturnsOnCall = make(map[int]struct {
			result1 backend.AccountService
		})
	}
	fake.limitReturnsOnCall[i] = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) Mask(arg1 string) backend.AccountService {
	fake.maskMutex.Lock()
	ret, specificReturn := fake.maskReturnsOnCall[len(fake.maskArgsForCall)]
	fake.maskArgsForCall = append(fake.maskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Mask", []interface{}{arg1})
	fake.maskMutex.Unlock()
	if fake.MaskStub != nil {
		return fake.MaskStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.maskReturns
	return fakeReturns.result1
}

func (fake *AccountService) MaskCallCount() int {
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	return len(fake.maskArgsForCall)
}

func (fake *AccountService) MaskCalls(stub func(string) backend.AccountService) {
	fake.maskMutex.Lock()
	defer fake.maskMutex.Unlock()
	fake.MaskStub = stub
}

func (fake *AccountService) MaskArgsForCall(i int) string {
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	argsForCall := fake.maskArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AccountService) MaskReturns(result1 backend.AccountService) {
	fake.maskMutex.Lock()
	defer fake.maskMutex.Unlock()
	fake.MaskStub = nil
	fake.maskReturns = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) MaskReturnsOnCall(i int, result1 backend.AccountService) {
	fake.maskMutex.Lock()
	defer fake.maskMutex.Unlock()
	fake.MaskStub = nil
	if fake.maskReturnsOnCall == nil {
		fake.maskReturnsOnCall = make(map[int]struct {
			result1 backend.AccountService
		})
	}
	fake.maskReturnsOnCall[i] = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) Offset(arg1 int) backend.AccountService {
	fake.offsetMutex.Lock()
	ret, specificReturn := fake.offsetReturnsOnCall[len(fake.offsetArgsForCall)]
	fake.offsetArgsForCall = append(fake.offsetArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Offset", []interface{}{arg1})
	fake.offsetMutex.Unlock()
	if fake.OffsetStub != nil {
		return fake.OffsetStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.offsetReturns
	return fakeReturns.result1
}

func (fake *AccountService) OffsetCallCount() int {
	fake.offsetMutex.RLock()
	defer fake.offsetMutex.RUnlock()
	return len(fake.offsetArgsForCall)
}

func (fake *AccountService) OffsetCalls(stub func(int) backend.AccountService) {
	fake.offsetMutex.Lock()
	defer fake.offsetMutex.Unlock()
	fake.OffsetStub = stub
}

func (fake *AccountService) OffsetArgsForCall(i int) int {
	fake.offsetMutex.RLock()
	defer fake.offsetMutex.RUnlock()
	argsForCall := fake.offsetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AccountService) OffsetReturns(result1 backend.AccountService) {
	fake.offsetMutex.Lock()
	defer fake.offsetMutex.Unlock()
	fake.OffsetStub = nil
	fake.offsetReturns = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) OffsetReturnsOnCall(i int, result1 backend.AccountService) {
	fake.offsetMutex.Lock()
	defer fake.offsetMutex.Unlock()
	fake.OffsetStub = nil
	if fake.offsetReturnsOnCall == nil {
		fake.offsetReturnsOnCall = make(map[int]struct {
			result1 backend.AccountService
		})
	}
	fake.offsetReturnsOnCall[i] = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *AccountService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	fake.getBlockDeviceTemplateGroupsMutex.RLock()
	defer fake.getBlockDeviceTemplateGroupsMutex.RUnlock()
	fake.getCurrentUserMutex.RLock()
	defer fake.getCurrentUserMutex.RUnlock()
	fake.getHardwareMutex.RLock()
	defer fake.getHardwareMutex.RUnlock()
	fake.getNetworkStorageMutex.RLock()
	defer fake.getNetworkStorageMutex.RUnlock()
	fake.getObjectMutex.RLock()
	defer fake.getObjectMutex.RUnlock()
	fake.getOrdersMutex.RLock()
	defer fake.getOrdersMutex.RUnlock()
	fake.getSubnetsMutex.RLock()
	defer fake.getSubnetsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.limitMutex.RLock()
	defer fake.limitMutex.RUnlock()
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	fake.offsetMutex.RLock()
	defer fake.offsetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AccountService) recordInvocation(key string, args []interface{}) {
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

var _ backend.AccountService = new(AccountService)