// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/tanzu-framework/cmd/cli/plugin/isolated-cluster/imageop"
)

type ImgpkgClientFake struct {
	CopyImageFromTarStub        func(string, string, string) error
	copyImageFromTarMutex       sync.RWMutex
	copyImageFromTarArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	copyImageFromTarReturns struct {
		result1 error
	}
	copyImageFromTarReturnsOnCall map[int]struct {
		result1 error
	}
	CopyImageToTarStub        func(string, string, string) error
	copyImageToTarMutex       sync.RWMutex
	copyImageToTarArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	copyImageToTarReturns struct {
		result1 error
	}
	copyImageToTarReturnsOnCall map[int]struct {
		result1 error
	}
	GetImageTagListStub        func(string) []string
	getImageTagListMutex       sync.RWMutex
	getImageTagListArgsForCall []struct {
		arg1 string
	}
	getImageTagListReturns struct {
		result1 []string
	}
	getImageTagListReturnsOnCall map[int]struct {
		result1 []string
	}
	PullImageStub        func(string, string) error
	pullImageMutex       sync.RWMutex
	pullImageArgsForCall []struct {
		arg1 string
		arg2 string
	}
	pullImageReturns struct {
		result1 error
	}
	pullImageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ImgpkgClientFake) CopyImageFromTar(arg1 string, arg2 string, arg3 string) error {
	fake.copyImageFromTarMutex.Lock()
	ret, specificReturn := fake.copyImageFromTarReturnsOnCall[len(fake.copyImageFromTarArgsForCall)]
	fake.copyImageFromTarArgsForCall = append(fake.copyImageFromTarArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CopyImageFromTarStub
	fakeReturns := fake.copyImageFromTarReturns
	fake.recordInvocation("CopyImageFromTar", []interface{}{arg1, arg2, arg3})
	fake.copyImageFromTarMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImgpkgClientFake) CopyImageFromTarCallCount() int {
	fake.copyImageFromTarMutex.RLock()
	defer fake.copyImageFromTarMutex.RUnlock()
	return len(fake.copyImageFromTarArgsForCall)
}

func (fake *ImgpkgClientFake) CopyImageFromTarCalls(stub func(string, string, string) error) {
	fake.copyImageFromTarMutex.Lock()
	defer fake.copyImageFromTarMutex.Unlock()
	fake.CopyImageFromTarStub = stub
}

func (fake *ImgpkgClientFake) CopyImageFromTarArgsForCall(i int) (string, string, string) {
	fake.copyImageFromTarMutex.RLock()
	defer fake.copyImageFromTarMutex.RUnlock()
	argsForCall := fake.copyImageFromTarArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ImgpkgClientFake) CopyImageFromTarReturns(result1 error) {
	fake.copyImageFromTarMutex.Lock()
	defer fake.copyImageFromTarMutex.Unlock()
	fake.CopyImageFromTarStub = nil
	fake.copyImageFromTarReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImgpkgClientFake) CopyImageFromTarReturnsOnCall(i int, result1 error) {
	fake.copyImageFromTarMutex.Lock()
	defer fake.copyImageFromTarMutex.Unlock()
	fake.CopyImageFromTarStub = nil
	if fake.copyImageFromTarReturnsOnCall == nil {
		fake.copyImageFromTarReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyImageFromTarReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImgpkgClientFake) CopyImageToTar(arg1 string, arg2 string, arg3 string) error {
	fake.copyImageToTarMutex.Lock()
	ret, specificReturn := fake.copyImageToTarReturnsOnCall[len(fake.copyImageToTarArgsForCall)]
	fake.copyImageToTarArgsForCall = append(fake.copyImageToTarArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CopyImageToTarStub
	fakeReturns := fake.copyImageToTarReturns
	fake.recordInvocation("CopyImageToTar", []interface{}{arg1, arg2, arg3})
	fake.copyImageToTarMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImgpkgClientFake) CopyImageToTarCallCount() int {
	fake.copyImageToTarMutex.RLock()
	defer fake.copyImageToTarMutex.RUnlock()
	return len(fake.copyImageToTarArgsForCall)
}

func (fake *ImgpkgClientFake) CopyImageToTarCalls(stub func(string, string, string) error) {
	fake.copyImageToTarMutex.Lock()
	defer fake.copyImageToTarMutex.Unlock()
	fake.CopyImageToTarStub = stub
}

func (fake *ImgpkgClientFake) CopyImageToTarArgsForCall(i int) (string, string, string) {
	fake.copyImageToTarMutex.RLock()
	defer fake.copyImageToTarMutex.RUnlock()
	argsForCall := fake.copyImageToTarArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ImgpkgClientFake) CopyImageToTarReturns(result1 error) {
	fake.copyImageToTarMutex.Lock()
	defer fake.copyImageToTarMutex.Unlock()
	fake.CopyImageToTarStub = nil
	fake.copyImageToTarReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImgpkgClientFake) CopyImageToTarReturnsOnCall(i int, result1 error) {
	fake.copyImageToTarMutex.Lock()
	defer fake.copyImageToTarMutex.Unlock()
	fake.CopyImageToTarStub = nil
	if fake.copyImageToTarReturnsOnCall == nil {
		fake.copyImageToTarReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyImageToTarReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImgpkgClientFake) GetImageTagList(arg1 string) []string {
	fake.getImageTagListMutex.Lock()
	ret, specificReturn := fake.getImageTagListReturnsOnCall[len(fake.getImageTagListArgsForCall)]
	fake.getImageTagListArgsForCall = append(fake.getImageTagListArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetImageTagListStub
	fakeReturns := fake.getImageTagListReturns
	fake.recordInvocation("GetImageTagList", []interface{}{arg1})
	fake.getImageTagListMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImgpkgClientFake) GetImageTagListCallCount() int {
	fake.getImageTagListMutex.RLock()
	defer fake.getImageTagListMutex.RUnlock()
	return len(fake.getImageTagListArgsForCall)
}

func (fake *ImgpkgClientFake) GetImageTagListCalls(stub func(string) []string) {
	fake.getImageTagListMutex.Lock()
	defer fake.getImageTagListMutex.Unlock()
	fake.GetImageTagListStub = stub
}

func (fake *ImgpkgClientFake) GetImageTagListArgsForCall(i int) string {
	fake.getImageTagListMutex.RLock()
	defer fake.getImageTagListMutex.RUnlock()
	argsForCall := fake.getImageTagListArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ImgpkgClientFake) GetImageTagListReturns(result1 []string) {
	fake.getImageTagListMutex.Lock()
	defer fake.getImageTagListMutex.Unlock()
	fake.GetImageTagListStub = nil
	fake.getImageTagListReturns = struct {
		result1 []string
	}{result1}
}

func (fake *ImgpkgClientFake) GetImageTagListReturnsOnCall(i int, result1 []string) {
	fake.getImageTagListMutex.Lock()
	defer fake.getImageTagListMutex.Unlock()
	fake.GetImageTagListStub = nil
	if fake.getImageTagListReturnsOnCall == nil {
		fake.getImageTagListReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getImageTagListReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *ImgpkgClientFake) PullImage(arg1 string, arg2 string) error {
	fake.pullImageMutex.Lock()
	ret, specificReturn := fake.pullImageReturnsOnCall[len(fake.pullImageArgsForCall)]
	fake.pullImageArgsForCall = append(fake.pullImageArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.PullImageStub
	fakeReturns := fake.pullImageReturns
	fake.recordInvocation("PullImage", []interface{}{arg1, arg2})
	fake.pullImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImgpkgClientFake) PullImageCallCount() int {
	fake.pullImageMutex.RLock()
	defer fake.pullImageMutex.RUnlock()
	return len(fake.pullImageArgsForCall)
}

func (fake *ImgpkgClientFake) PullImageCalls(stub func(string, string) error) {
	fake.pullImageMutex.Lock()
	defer fake.pullImageMutex.Unlock()
	fake.PullImageStub = stub
}

func (fake *ImgpkgClientFake) PullImageArgsForCall(i int) (string, string) {
	fake.pullImageMutex.RLock()
	defer fake.pullImageMutex.RUnlock()
	argsForCall := fake.pullImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ImgpkgClientFake) PullImageReturns(result1 error) {
	fake.pullImageMutex.Lock()
	defer fake.pullImageMutex.Unlock()
	fake.PullImageStub = nil
	fake.pullImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImgpkgClientFake) PullImageReturnsOnCall(i int, result1 error) {
	fake.pullImageMutex.Lock()
	defer fake.pullImageMutex.Unlock()
	fake.PullImageStub = nil
	if fake.pullImageReturnsOnCall == nil {
		fake.pullImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pullImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImgpkgClientFake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.copyImageFromTarMutex.RLock()
	defer fake.copyImageFromTarMutex.RUnlock()
	fake.copyImageToTarMutex.RLock()
	defer fake.copyImageToTarMutex.RUnlock()
	fake.getImageTagListMutex.RLock()
	defer fake.getImageTagListMutex.RUnlock()
	fake.pullImageMutex.RLock()
	defer fake.pullImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ImgpkgClientFake) recordInvocation(key string, args []interface{}) {
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

var _ imageop.ImgpkgClient = new(ImgpkgClientFake)