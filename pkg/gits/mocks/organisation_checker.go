// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/gits (interfaces: OrganisationChecker)

package gits_test

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockOrganisationChecker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockOrganisationChecker() *MockOrganisationChecker {
	return &MockOrganisationChecker{fail: pegomock.GlobalFailHandler}
}

func (mock *MockOrganisationChecker) IsUserInOrganisation(_param0 string, _param1 string) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockOrganisationChecker().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("IsUserInOrganisation", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockOrganisationChecker) VerifyWasCalledOnce() *VerifierOrganisationChecker {
	return &VerifierOrganisationChecker{mock, pegomock.Times(1), nil}
}

func (mock *MockOrganisationChecker) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierOrganisationChecker {
	return &VerifierOrganisationChecker{mock, invocationCountMatcher, nil}
}

func (mock *MockOrganisationChecker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierOrganisationChecker {
	return &VerifierOrganisationChecker{mock, invocationCountMatcher, inOrderContext}
}

type VerifierOrganisationChecker struct {
	mock                   *MockOrganisationChecker
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierOrganisationChecker) IsUserInOrganisation(_param0 string, _param1 string) *OrganisationChecker_IsUserInOrganisation_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "IsUserInOrganisation", params)
	return &OrganisationChecker_IsUserInOrganisation_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type OrganisationChecker_IsUserInOrganisation_OngoingVerification struct {
	mock              *MockOrganisationChecker
	methodInvocations []pegomock.MethodInvocation
}

func (c *OrganisationChecker_IsUserInOrganisation_OngoingVerification) GetCapturedArguments() (string, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *OrganisationChecker_IsUserInOrganisation_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}
