// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/alphagov/paas-rds-broker/awsrds"
	"github.com/aws/aws-sdk-go/service/rds"
)

type FakeRDSInstance struct {
	DescribeStub        func(ID string) (*rds.DBInstance, error)
	describeMutex       sync.RWMutex
	describeArgsForCall []struct {
		ID string
	}
	describeReturns struct {
		result1 *rds.DBInstance
		result2 error
	}
	describeReturnsOnCall map[int]struct {
		result1 *rds.DBInstance
		result2 error
	}
	GetResourceTagsStub        func(resourceArn string, opts ...awsrds.DescribeOption) ([]*rds.Tag, error)
	getResourceTagsMutex       sync.RWMutex
	getResourceTagsArgsForCall []struct {
		resourceArn string
		opts        []awsrds.DescribeOption
	}
	getResourceTagsReturns struct {
		result1 []*rds.Tag
		result2 error
	}
	getResourceTagsReturnsOnCall map[int]struct {
		result1 []*rds.Tag
		result2 error
	}
	DescribeByTagStub        func(TagName, TagValue string, opts ...awsrds.DescribeOption) ([]*rds.DBInstance, error)
	describeByTagMutex       sync.RWMutex
	describeByTagArgsForCall []struct {
		TagName  string
		TagValue string
		opts     []awsrds.DescribeOption
	}
	describeByTagReturns struct {
		result1 []*rds.DBInstance
		result2 error
	}
	describeByTagReturnsOnCall map[int]struct {
		result1 []*rds.DBInstance
		result2 error
	}
	DescribeSnapshotsStub        func(DBInstanceID string) ([]*rds.DBSnapshot, error)
	describeSnapshotsMutex       sync.RWMutex
	describeSnapshotsArgsForCall []struct {
		DBInstanceID string
	}
	describeSnapshotsReturns struct {
		result1 []*rds.DBSnapshot
		result2 error
	}
	describeSnapshotsReturnsOnCall map[int]struct {
		result1 []*rds.DBSnapshot
		result2 error
	}
	DeleteSnapshotsStub        func(brokerName string, keepForDays int) error
	deleteSnapshotsMutex       sync.RWMutex
	deleteSnapshotsArgsForCall []struct {
		brokerName  string
		keepForDays int
	}
	deleteSnapshotsReturns struct {
		result1 error
	}
	deleteSnapshotsReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(createDBInstanceInput *rds.CreateDBInstanceInput) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		createDBInstanceInput *rds.CreateDBInstanceInput
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	RestoreStub        func(restoreRBInstanceInput *rds.RestoreDBInstanceFromDBSnapshotInput) error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct {
		restoreRBInstanceInput *rds.RestoreDBInstanceFromDBSnapshotInput
	}
	restoreReturns struct {
		result1 error
	}
	restoreReturnsOnCall map[int]struct {
		result1 error
	}
	ModifyStub        func(modifyDBInstanceInput *rds.ModifyDBInstanceInput) (*rds.DBInstance, error)
	modifyMutex       sync.RWMutex
	modifyArgsForCall []struct {
		modifyDBInstanceInput *rds.ModifyDBInstanceInput
	}
	modifyReturns struct {
		result1 *rds.DBInstance
		result2 error
	}
	modifyReturnsOnCall map[int]struct {
		result1 *rds.DBInstance
		result2 error
	}
	AddTagsToResourceStub        func(resourceArn string, tags []*rds.Tag) error
	addTagsToResourceMutex       sync.RWMutex
	addTagsToResourceArgsForCall []struct {
		resourceArn string
		tags        []*rds.Tag
	}
	addTagsToResourceReturns struct {
		result1 error
	}
	addTagsToResourceReturnsOnCall map[int]struct {
		result1 error
	}
	RebootStub        func(rebootDBInstanceInput *rds.RebootDBInstanceInput) error
	rebootMutex       sync.RWMutex
	rebootArgsForCall []struct {
		rebootDBInstanceInput *rds.RebootDBInstanceInput
	}
	rebootReturns struct {
		result1 error
	}
	rebootReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveTagStub        func(ID, tagKey string) error
	removeTagMutex       sync.RWMutex
	removeTagArgsForCall []struct {
		ID     string
		tagKey string
	}
	removeTagReturns struct {
		result1 error
	}
	removeTagReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(ID string, skipFinalSnapshot bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		ID                string
		skipFinalSnapshot bool
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetTagStub        func(ID, tagKey string) (string, error)
	getTagMutex       sync.RWMutex
	getTagArgsForCall []struct {
		ID     string
		tagKey string
	}
	getTagReturns struct {
		result1 string
		result2 error
	}
	getTagReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetParameterGroupStub        func(groupId string) (*rds.DBParameterGroup, error)
	getParameterGroupMutex       sync.RWMutex
	getParameterGroupArgsForCall []struct {
		groupId string
	}
	getParameterGroupReturns struct {
		result1 *rds.DBParameterGroup
		result2 error
	}
	getParameterGroupReturnsOnCall map[int]struct {
		result1 *rds.DBParameterGroup
		result2 error
	}
	CreateParameterGroupStub        func(input *rds.CreateDBParameterGroupInput) error
	createParameterGroupMutex       sync.RWMutex
	createParameterGroupArgsForCall []struct {
		input *rds.CreateDBParameterGroupInput
	}
	createParameterGroupReturns struct {
		result1 error
	}
	createParameterGroupReturnsOnCall map[int]struct {
		result1 error
	}
	ModifyParameterGroupStub        func(input *rds.ModifyDBParameterGroupInput) error
	modifyParameterGroupMutex       sync.RWMutex
	modifyParameterGroupArgsForCall []struct {
		input *rds.ModifyDBParameterGroupInput
	}
	modifyParameterGroupReturns struct {
		result1 error
	}
	modifyParameterGroupReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRDSInstance) Describe(ID string) (*rds.DBInstance, error) {
	fake.describeMutex.Lock()
	ret, specificReturn := fake.describeReturnsOnCall[len(fake.describeArgsForCall)]
	fake.describeArgsForCall = append(fake.describeArgsForCall, struct {
		ID string
	}{ID})
	fake.recordInvocation("Describe", []interface{}{ID})
	fake.describeMutex.Unlock()
	if fake.DescribeStub != nil {
		return fake.DescribeStub(ID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.describeReturns.result1, fake.describeReturns.result2
}

func (fake *FakeRDSInstance) DescribeCallCount() int {
	fake.describeMutex.RLock()
	defer fake.describeMutex.RUnlock()
	return len(fake.describeArgsForCall)
}

func (fake *FakeRDSInstance) DescribeArgsForCall(i int) string {
	fake.describeMutex.RLock()
	defer fake.describeMutex.RUnlock()
	return fake.describeArgsForCall[i].ID
}

func (fake *FakeRDSInstance) DescribeReturns(result1 *rds.DBInstance, result2 error) {
	fake.DescribeStub = nil
	fake.describeReturns = struct {
		result1 *rds.DBInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) DescribeReturnsOnCall(i int, result1 *rds.DBInstance, result2 error) {
	fake.DescribeStub = nil
	if fake.describeReturnsOnCall == nil {
		fake.describeReturnsOnCall = make(map[int]struct {
			result1 *rds.DBInstance
			result2 error
		})
	}
	fake.describeReturnsOnCall[i] = struct {
		result1 *rds.DBInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) GetResourceTags(resourceArn string, opts ...awsrds.DescribeOption) ([]*rds.Tag, error) {
	fake.getResourceTagsMutex.Lock()
	ret, specificReturn := fake.getResourceTagsReturnsOnCall[len(fake.getResourceTagsArgsForCall)]
	fake.getResourceTagsArgsForCall = append(fake.getResourceTagsArgsForCall, struct {
		resourceArn string
		opts        []awsrds.DescribeOption
	}{resourceArn, opts})
	fake.recordInvocation("GetResourceTags", []interface{}{resourceArn, opts})
	fake.getResourceTagsMutex.Unlock()
	if fake.GetResourceTagsStub != nil {
		return fake.GetResourceTagsStub(resourceArn, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getResourceTagsReturns.result1, fake.getResourceTagsReturns.result2
}

func (fake *FakeRDSInstance) GetResourceTagsCallCount() int {
	fake.getResourceTagsMutex.RLock()
	defer fake.getResourceTagsMutex.RUnlock()
	return len(fake.getResourceTagsArgsForCall)
}

func (fake *FakeRDSInstance) GetResourceTagsArgsForCall(i int) (string, []awsrds.DescribeOption) {
	fake.getResourceTagsMutex.RLock()
	defer fake.getResourceTagsMutex.RUnlock()
	return fake.getResourceTagsArgsForCall[i].resourceArn, fake.getResourceTagsArgsForCall[i].opts
}

func (fake *FakeRDSInstance) GetResourceTagsReturns(result1 []*rds.Tag, result2 error) {
	fake.GetResourceTagsStub = nil
	fake.getResourceTagsReturns = struct {
		result1 []*rds.Tag
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) GetResourceTagsReturnsOnCall(i int, result1 []*rds.Tag, result2 error) {
	fake.GetResourceTagsStub = nil
	if fake.getResourceTagsReturnsOnCall == nil {
		fake.getResourceTagsReturnsOnCall = make(map[int]struct {
			result1 []*rds.Tag
			result2 error
		})
	}
	fake.getResourceTagsReturnsOnCall[i] = struct {
		result1 []*rds.Tag
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) DescribeByTag(TagName string, TagValue string, opts ...awsrds.DescribeOption) ([]*rds.DBInstance, error) {
	fake.describeByTagMutex.Lock()
	ret, specificReturn := fake.describeByTagReturnsOnCall[len(fake.describeByTagArgsForCall)]
	fake.describeByTagArgsForCall = append(fake.describeByTagArgsForCall, struct {
		TagName  string
		TagValue string
		opts     []awsrds.DescribeOption
	}{TagName, TagValue, opts})
	fake.recordInvocation("DescribeByTag", []interface{}{TagName, TagValue, opts})
	fake.describeByTagMutex.Unlock()
	if fake.DescribeByTagStub != nil {
		return fake.DescribeByTagStub(TagName, TagValue, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.describeByTagReturns.result1, fake.describeByTagReturns.result2
}

func (fake *FakeRDSInstance) DescribeByTagCallCount() int {
	fake.describeByTagMutex.RLock()
	defer fake.describeByTagMutex.RUnlock()
	return len(fake.describeByTagArgsForCall)
}

func (fake *FakeRDSInstance) DescribeByTagArgsForCall(i int) (string, string, []awsrds.DescribeOption) {
	fake.describeByTagMutex.RLock()
	defer fake.describeByTagMutex.RUnlock()
	return fake.describeByTagArgsForCall[i].TagName, fake.describeByTagArgsForCall[i].TagValue, fake.describeByTagArgsForCall[i].opts
}

func (fake *FakeRDSInstance) DescribeByTagReturns(result1 []*rds.DBInstance, result2 error) {
	fake.DescribeByTagStub = nil
	fake.describeByTagReturns = struct {
		result1 []*rds.DBInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) DescribeByTagReturnsOnCall(i int, result1 []*rds.DBInstance, result2 error) {
	fake.DescribeByTagStub = nil
	if fake.describeByTagReturnsOnCall == nil {
		fake.describeByTagReturnsOnCall = make(map[int]struct {
			result1 []*rds.DBInstance
			result2 error
		})
	}
	fake.describeByTagReturnsOnCall[i] = struct {
		result1 []*rds.DBInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) DescribeSnapshots(DBInstanceID string) ([]*rds.DBSnapshot, error) {
	fake.describeSnapshotsMutex.Lock()
	ret, specificReturn := fake.describeSnapshotsReturnsOnCall[len(fake.describeSnapshotsArgsForCall)]
	fake.describeSnapshotsArgsForCall = append(fake.describeSnapshotsArgsForCall, struct {
		DBInstanceID string
	}{DBInstanceID})
	fake.recordInvocation("DescribeSnapshots", []interface{}{DBInstanceID})
	fake.describeSnapshotsMutex.Unlock()
	if fake.DescribeSnapshotsStub != nil {
		return fake.DescribeSnapshotsStub(DBInstanceID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.describeSnapshotsReturns.result1, fake.describeSnapshotsReturns.result2
}

func (fake *FakeRDSInstance) DescribeSnapshotsCallCount() int {
	fake.describeSnapshotsMutex.RLock()
	defer fake.describeSnapshotsMutex.RUnlock()
	return len(fake.describeSnapshotsArgsForCall)
}

func (fake *FakeRDSInstance) DescribeSnapshotsArgsForCall(i int) string {
	fake.describeSnapshotsMutex.RLock()
	defer fake.describeSnapshotsMutex.RUnlock()
	return fake.describeSnapshotsArgsForCall[i].DBInstanceID
}

func (fake *FakeRDSInstance) DescribeSnapshotsReturns(result1 []*rds.DBSnapshot, result2 error) {
	fake.DescribeSnapshotsStub = nil
	fake.describeSnapshotsReturns = struct {
		result1 []*rds.DBSnapshot
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) DescribeSnapshotsReturnsOnCall(i int, result1 []*rds.DBSnapshot, result2 error) {
	fake.DescribeSnapshotsStub = nil
	if fake.describeSnapshotsReturnsOnCall == nil {
		fake.describeSnapshotsReturnsOnCall = make(map[int]struct {
			result1 []*rds.DBSnapshot
			result2 error
		})
	}
	fake.describeSnapshotsReturnsOnCall[i] = struct {
		result1 []*rds.DBSnapshot
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) DeleteSnapshots(brokerName string, keepForDays int) error {
	fake.deleteSnapshotsMutex.Lock()
	ret, specificReturn := fake.deleteSnapshotsReturnsOnCall[len(fake.deleteSnapshotsArgsForCall)]
	fake.deleteSnapshotsArgsForCall = append(fake.deleteSnapshotsArgsForCall, struct {
		brokerName  string
		keepForDays int
	}{brokerName, keepForDays})
	fake.recordInvocation("DeleteSnapshots", []interface{}{brokerName, keepForDays})
	fake.deleteSnapshotsMutex.Unlock()
	if fake.DeleteSnapshotsStub != nil {
		return fake.DeleteSnapshotsStub(brokerName, keepForDays)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteSnapshotsReturns.result1
}

func (fake *FakeRDSInstance) DeleteSnapshotsCallCount() int {
	fake.deleteSnapshotsMutex.RLock()
	defer fake.deleteSnapshotsMutex.RUnlock()
	return len(fake.deleteSnapshotsArgsForCall)
}

func (fake *FakeRDSInstance) DeleteSnapshotsArgsForCall(i int) (string, int) {
	fake.deleteSnapshotsMutex.RLock()
	defer fake.deleteSnapshotsMutex.RUnlock()
	return fake.deleteSnapshotsArgsForCall[i].brokerName, fake.deleteSnapshotsArgsForCall[i].keepForDays
}

func (fake *FakeRDSInstance) DeleteSnapshotsReturns(result1 error) {
	fake.DeleteSnapshotsStub = nil
	fake.deleteSnapshotsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) DeleteSnapshotsReturnsOnCall(i int, result1 error) {
	fake.DeleteSnapshotsStub = nil
	if fake.deleteSnapshotsReturnsOnCall == nil {
		fake.deleteSnapshotsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteSnapshotsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) Create(createDBInstanceInput *rds.CreateDBInstanceInput) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		createDBInstanceInput *rds.CreateDBInstanceInput
	}{createDBInstanceInput})
	fake.recordInvocation("Create", []interface{}{createDBInstanceInput})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(createDBInstanceInput)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeRDSInstance) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRDSInstance) CreateArgsForCall(i int) *rds.CreateDBInstanceInput {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].createDBInstanceInput
}

func (fake *FakeRDSInstance) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) Restore(restoreRBInstanceInput *rds.RestoreDBInstanceFromDBSnapshotInput) error {
	fake.restoreMutex.Lock()
	ret, specificReturn := fake.restoreReturnsOnCall[len(fake.restoreArgsForCall)]
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct {
		restoreRBInstanceInput *rds.RestoreDBInstanceFromDBSnapshotInput
	}{restoreRBInstanceInput})
	fake.recordInvocation("Restore", []interface{}{restoreRBInstanceInput})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub(restoreRBInstanceInput)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.restoreReturns.result1
}

func (fake *FakeRDSInstance) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeRDSInstance) RestoreArgsForCall(i int) *rds.RestoreDBInstanceFromDBSnapshotInput {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return fake.restoreArgsForCall[i].restoreRBInstanceInput
}

func (fake *FakeRDSInstance) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) RestoreReturnsOnCall(i int, result1 error) {
	fake.RestoreStub = nil
	if fake.restoreReturnsOnCall == nil {
		fake.restoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) Modify(modifyDBInstanceInput *rds.ModifyDBInstanceInput) (*rds.DBInstance, error) {
	fake.modifyMutex.Lock()
	ret, specificReturn := fake.modifyReturnsOnCall[len(fake.modifyArgsForCall)]
	fake.modifyArgsForCall = append(fake.modifyArgsForCall, struct {
		modifyDBInstanceInput *rds.ModifyDBInstanceInput
	}{modifyDBInstanceInput})
	fake.recordInvocation("Modify", []interface{}{modifyDBInstanceInput})
	fake.modifyMutex.Unlock()
	if fake.ModifyStub != nil {
		return fake.ModifyStub(modifyDBInstanceInput)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.modifyReturns.result1, fake.modifyReturns.result2
}

func (fake *FakeRDSInstance) ModifyCallCount() int {
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	return len(fake.modifyArgsForCall)
}

func (fake *FakeRDSInstance) ModifyArgsForCall(i int) *rds.ModifyDBInstanceInput {
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	return fake.modifyArgsForCall[i].modifyDBInstanceInput
}

func (fake *FakeRDSInstance) ModifyReturns(result1 *rds.DBInstance, result2 error) {
	fake.ModifyStub = nil
	fake.modifyReturns = struct {
		result1 *rds.DBInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) ModifyReturnsOnCall(i int, result1 *rds.DBInstance, result2 error) {
	fake.ModifyStub = nil
	if fake.modifyReturnsOnCall == nil {
		fake.modifyReturnsOnCall = make(map[int]struct {
			result1 *rds.DBInstance
			result2 error
		})
	}
	fake.modifyReturnsOnCall[i] = struct {
		result1 *rds.DBInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) AddTagsToResource(resourceArn string, tags []*rds.Tag) error {
	var tagsCopy []*rds.Tag
	if tags != nil {
		tagsCopy = make([]*rds.Tag, len(tags))
		copy(tagsCopy, tags)
	}
	fake.addTagsToResourceMutex.Lock()
	ret, specificReturn := fake.addTagsToResourceReturnsOnCall[len(fake.addTagsToResourceArgsForCall)]
	fake.addTagsToResourceArgsForCall = append(fake.addTagsToResourceArgsForCall, struct {
		resourceArn string
		tags        []*rds.Tag
	}{resourceArn, tagsCopy})
	fake.recordInvocation("AddTagsToResource", []interface{}{resourceArn, tagsCopy})
	fake.addTagsToResourceMutex.Unlock()
	if fake.AddTagsToResourceStub != nil {
		return fake.AddTagsToResourceStub(resourceArn, tags)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addTagsToResourceReturns.result1
}

func (fake *FakeRDSInstance) AddTagsToResourceCallCount() int {
	fake.addTagsToResourceMutex.RLock()
	defer fake.addTagsToResourceMutex.RUnlock()
	return len(fake.addTagsToResourceArgsForCall)
}

func (fake *FakeRDSInstance) AddTagsToResourceArgsForCall(i int) (string, []*rds.Tag) {
	fake.addTagsToResourceMutex.RLock()
	defer fake.addTagsToResourceMutex.RUnlock()
	return fake.addTagsToResourceArgsForCall[i].resourceArn, fake.addTagsToResourceArgsForCall[i].tags
}

func (fake *FakeRDSInstance) AddTagsToResourceReturns(result1 error) {
	fake.AddTagsToResourceStub = nil
	fake.addTagsToResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) AddTagsToResourceReturnsOnCall(i int, result1 error) {
	fake.AddTagsToResourceStub = nil
	if fake.addTagsToResourceReturnsOnCall == nil {
		fake.addTagsToResourceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addTagsToResourceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) Reboot(rebootDBInstanceInput *rds.RebootDBInstanceInput) error {
	fake.rebootMutex.Lock()
	ret, specificReturn := fake.rebootReturnsOnCall[len(fake.rebootArgsForCall)]
	fake.rebootArgsForCall = append(fake.rebootArgsForCall, struct {
		rebootDBInstanceInput *rds.RebootDBInstanceInput
	}{rebootDBInstanceInput})
	fake.recordInvocation("Reboot", []interface{}{rebootDBInstanceInput})
	fake.rebootMutex.Unlock()
	if fake.RebootStub != nil {
		return fake.RebootStub(rebootDBInstanceInput)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rebootReturns.result1
}

func (fake *FakeRDSInstance) RebootCallCount() int {
	fake.rebootMutex.RLock()
	defer fake.rebootMutex.RUnlock()
	return len(fake.rebootArgsForCall)
}

func (fake *FakeRDSInstance) RebootArgsForCall(i int) *rds.RebootDBInstanceInput {
	fake.rebootMutex.RLock()
	defer fake.rebootMutex.RUnlock()
	return fake.rebootArgsForCall[i].rebootDBInstanceInput
}

func (fake *FakeRDSInstance) RebootReturns(result1 error) {
	fake.RebootStub = nil
	fake.rebootReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) RebootReturnsOnCall(i int, result1 error) {
	fake.RebootStub = nil
	if fake.rebootReturnsOnCall == nil {
		fake.rebootReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rebootReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) RemoveTag(ID string, tagKey string) error {
	fake.removeTagMutex.Lock()
	ret, specificReturn := fake.removeTagReturnsOnCall[len(fake.removeTagArgsForCall)]
	fake.removeTagArgsForCall = append(fake.removeTagArgsForCall, struct {
		ID     string
		tagKey string
	}{ID, tagKey})
	fake.recordInvocation("RemoveTag", []interface{}{ID, tagKey})
	fake.removeTagMutex.Unlock()
	if fake.RemoveTagStub != nil {
		return fake.RemoveTagStub(ID, tagKey)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeTagReturns.result1
}

func (fake *FakeRDSInstance) RemoveTagCallCount() int {
	fake.removeTagMutex.RLock()
	defer fake.removeTagMutex.RUnlock()
	return len(fake.removeTagArgsForCall)
}

func (fake *FakeRDSInstance) RemoveTagArgsForCall(i int) (string, string) {
	fake.removeTagMutex.RLock()
	defer fake.removeTagMutex.RUnlock()
	return fake.removeTagArgsForCall[i].ID, fake.removeTagArgsForCall[i].tagKey
}

func (fake *FakeRDSInstance) RemoveTagReturns(result1 error) {
	fake.RemoveTagStub = nil
	fake.removeTagReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) RemoveTagReturnsOnCall(i int, result1 error) {
	fake.RemoveTagStub = nil
	if fake.removeTagReturnsOnCall == nil {
		fake.removeTagReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeTagReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) Delete(ID string, skipFinalSnapshot bool) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		ID                string
		skipFinalSnapshot bool
	}{ID, skipFinalSnapshot})
	fake.recordInvocation("Delete", []interface{}{ID, skipFinalSnapshot})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(ID, skipFinalSnapshot)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeRDSInstance) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeRDSInstance) DeleteArgsForCall(i int) (string, bool) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].ID, fake.deleteArgsForCall[i].skipFinalSnapshot
}

func (fake *FakeRDSInstance) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) GetTag(ID string, tagKey string) (string, error) {
	fake.getTagMutex.Lock()
	ret, specificReturn := fake.getTagReturnsOnCall[len(fake.getTagArgsForCall)]
	fake.getTagArgsForCall = append(fake.getTagArgsForCall, struct {
		ID     string
		tagKey string
	}{ID, tagKey})
	fake.recordInvocation("GetTag", []interface{}{ID, tagKey})
	fake.getTagMutex.Unlock()
	if fake.GetTagStub != nil {
		return fake.GetTagStub(ID, tagKey)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTagReturns.result1, fake.getTagReturns.result2
}

func (fake *FakeRDSInstance) GetTagCallCount() int {
	fake.getTagMutex.RLock()
	defer fake.getTagMutex.RUnlock()
	return len(fake.getTagArgsForCall)
}

func (fake *FakeRDSInstance) GetTagArgsForCall(i int) (string, string) {
	fake.getTagMutex.RLock()
	defer fake.getTagMutex.RUnlock()
	return fake.getTagArgsForCall[i].ID, fake.getTagArgsForCall[i].tagKey
}

func (fake *FakeRDSInstance) GetTagReturns(result1 string, result2 error) {
	fake.GetTagStub = nil
	fake.getTagReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) GetTagReturnsOnCall(i int, result1 string, result2 error) {
	fake.GetTagStub = nil
	if fake.getTagReturnsOnCall == nil {
		fake.getTagReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getTagReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) GetParameterGroup(groupId string) (*rds.DBParameterGroup, error) {
	fake.getParameterGroupMutex.Lock()
	ret, specificReturn := fake.getParameterGroupReturnsOnCall[len(fake.getParameterGroupArgsForCall)]
	fake.getParameterGroupArgsForCall = append(fake.getParameterGroupArgsForCall, struct {
		groupId string
	}{groupId})
	fake.recordInvocation("GetParameterGroup", []interface{}{groupId})
	fake.getParameterGroupMutex.Unlock()
	if fake.GetParameterGroupStub != nil {
		return fake.GetParameterGroupStub(groupId)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getParameterGroupReturns.result1, fake.getParameterGroupReturns.result2
}

func (fake *FakeRDSInstance) GetParameterGroupCallCount() int {
	fake.getParameterGroupMutex.RLock()
	defer fake.getParameterGroupMutex.RUnlock()
	return len(fake.getParameterGroupArgsForCall)
}

func (fake *FakeRDSInstance) GetParameterGroupArgsForCall(i int) string {
	fake.getParameterGroupMutex.RLock()
	defer fake.getParameterGroupMutex.RUnlock()
	return fake.getParameterGroupArgsForCall[i].groupId
}

func (fake *FakeRDSInstance) GetParameterGroupReturns(result1 *rds.DBParameterGroup, result2 error) {
	fake.GetParameterGroupStub = nil
	fake.getParameterGroupReturns = struct {
		result1 *rds.DBParameterGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) GetParameterGroupReturnsOnCall(i int, result1 *rds.DBParameterGroup, result2 error) {
	fake.GetParameterGroupStub = nil
	if fake.getParameterGroupReturnsOnCall == nil {
		fake.getParameterGroupReturnsOnCall = make(map[int]struct {
			result1 *rds.DBParameterGroup
			result2 error
		})
	}
	fake.getParameterGroupReturnsOnCall[i] = struct {
		result1 *rds.DBParameterGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeRDSInstance) CreateParameterGroup(input *rds.CreateDBParameterGroupInput) error {
	fake.createParameterGroupMutex.Lock()
	ret, specificReturn := fake.createParameterGroupReturnsOnCall[len(fake.createParameterGroupArgsForCall)]
	fake.createParameterGroupArgsForCall = append(fake.createParameterGroupArgsForCall, struct {
		input *rds.CreateDBParameterGroupInput
	}{input})
	fake.recordInvocation("CreateParameterGroup", []interface{}{input})
	fake.createParameterGroupMutex.Unlock()
	if fake.CreateParameterGroupStub != nil {
		return fake.CreateParameterGroupStub(input)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createParameterGroupReturns.result1
}

func (fake *FakeRDSInstance) CreateParameterGroupCallCount() int {
	fake.createParameterGroupMutex.RLock()
	defer fake.createParameterGroupMutex.RUnlock()
	return len(fake.createParameterGroupArgsForCall)
}

func (fake *FakeRDSInstance) CreateParameterGroupArgsForCall(i int) *rds.CreateDBParameterGroupInput {
	fake.createParameterGroupMutex.RLock()
	defer fake.createParameterGroupMutex.RUnlock()
	return fake.createParameterGroupArgsForCall[i].input
}

func (fake *FakeRDSInstance) CreateParameterGroupReturns(result1 error) {
	fake.CreateParameterGroupStub = nil
	fake.createParameterGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) CreateParameterGroupReturnsOnCall(i int, result1 error) {
	fake.CreateParameterGroupStub = nil
	if fake.createParameterGroupReturnsOnCall == nil {
		fake.createParameterGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createParameterGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) ModifyParameterGroup(input *rds.ModifyDBParameterGroupInput) error {
	fake.modifyParameterGroupMutex.Lock()
	ret, specificReturn := fake.modifyParameterGroupReturnsOnCall[len(fake.modifyParameterGroupArgsForCall)]
	fake.modifyParameterGroupArgsForCall = append(fake.modifyParameterGroupArgsForCall, struct {
		input *rds.ModifyDBParameterGroupInput
	}{input})
	fake.recordInvocation("ModifyParameterGroup", []interface{}{input})
	fake.modifyParameterGroupMutex.Unlock()
	if fake.ModifyParameterGroupStub != nil {
		return fake.ModifyParameterGroupStub(input)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.modifyParameterGroupReturns.result1
}

func (fake *FakeRDSInstance) ModifyParameterGroupCallCount() int {
	fake.modifyParameterGroupMutex.RLock()
	defer fake.modifyParameterGroupMutex.RUnlock()
	return len(fake.modifyParameterGroupArgsForCall)
}

func (fake *FakeRDSInstance) ModifyParameterGroupArgsForCall(i int) *rds.ModifyDBParameterGroupInput {
	fake.modifyParameterGroupMutex.RLock()
	defer fake.modifyParameterGroupMutex.RUnlock()
	return fake.modifyParameterGroupArgsForCall[i].input
}

func (fake *FakeRDSInstance) ModifyParameterGroupReturns(result1 error) {
	fake.ModifyParameterGroupStub = nil
	fake.modifyParameterGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) ModifyParameterGroupReturnsOnCall(i int, result1 error) {
	fake.ModifyParameterGroupStub = nil
	if fake.modifyParameterGroupReturnsOnCall == nil {
		fake.modifyParameterGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.modifyParameterGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRDSInstance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.describeMutex.RLock()
	defer fake.describeMutex.RUnlock()
	fake.getResourceTagsMutex.RLock()
	defer fake.getResourceTagsMutex.RUnlock()
	fake.describeByTagMutex.RLock()
	defer fake.describeByTagMutex.RUnlock()
	fake.describeSnapshotsMutex.RLock()
	defer fake.describeSnapshotsMutex.RUnlock()
	fake.deleteSnapshotsMutex.RLock()
	defer fake.deleteSnapshotsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	fake.addTagsToResourceMutex.RLock()
	defer fake.addTagsToResourceMutex.RUnlock()
	fake.rebootMutex.RLock()
	defer fake.rebootMutex.RUnlock()
	fake.removeTagMutex.RLock()
	defer fake.removeTagMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getTagMutex.RLock()
	defer fake.getTagMutex.RUnlock()
	fake.getParameterGroupMutex.RLock()
	defer fake.getParameterGroupMutex.RUnlock()
	fake.createParameterGroupMutex.RLock()
	defer fake.createParameterGroupMutex.RUnlock()
	fake.modifyParameterGroupMutex.RLock()
	defer fake.modifyParameterGroupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRDSInstance) recordInvocation(key string, args []interface{}) {
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

var _ awsrds.RDSInstance = new(FakeRDSInstance)
