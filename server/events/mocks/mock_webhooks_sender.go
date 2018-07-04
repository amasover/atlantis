// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/runatlantis/atlantis/server/events (interfaces: WebhooksSender)

package mocks

import (
	"reflect"

	pegomock "github.com/petergtz/pegomock"
	webhooks "github.com/runatlantis/atlantis/server/events/webhooks"
	logging "github.com/runatlantis/atlantis/server/logging"
)

type MockWebhooksSender struct {
	fail func(message string, callerSkip ...int)
}

func NewMockWebhooksSender() *MockWebhooksSender {
	return &MockWebhooksSender{fail: pegomock.GlobalFailHandler}
}

func (mock *MockWebhooksSender) Send(log *logging.SimpleLogger, res webhooks.ApplyResult) error {
	params := []pegomock.Param{log, res}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Send", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockWebhooksSender) VerifyWasCalledOnce() *VerifierWebhooksSender {
	return &VerifierWebhooksSender{mock, pegomock.Times(1), nil}
}

func (mock *MockWebhooksSender) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierWebhooksSender {
	return &VerifierWebhooksSender{mock, invocationCountMatcher, nil}
}

func (mock *MockWebhooksSender) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierWebhooksSender {
	return &VerifierWebhooksSender{mock, invocationCountMatcher, inOrderContext}
}

type VerifierWebhooksSender struct {
	mock                   *MockWebhooksSender
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierWebhooksSender) Send(log *logging.SimpleLogger, res webhooks.ApplyResult) *WebhooksSender_Send_OngoingVerification {
	params := []pegomock.Param{log, res}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Send", params)
	return &WebhooksSender_Send_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type WebhooksSender_Send_OngoingVerification struct {
	mock              *MockWebhooksSender
	methodInvocations []pegomock.MethodInvocation
}

func (c *WebhooksSender_Send_OngoingVerification) GetCapturedArguments() (*logging.SimpleLogger, webhooks.ApplyResult) {
	log, res := c.GetAllCapturedArguments()
	return log[len(log)-1], res[len(res)-1]
}

func (c *WebhooksSender_Send_OngoingVerification) GetAllCapturedArguments() (_param0 []*logging.SimpleLogger, _param1 []webhooks.ApplyResult) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*logging.SimpleLogger, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*logging.SimpleLogger)
		}
		_param1 = make([]webhooks.ApplyResult, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(webhooks.ApplyResult)
		}
	}
	return
}