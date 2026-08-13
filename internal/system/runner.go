package system

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Runner is the boundary between product services and Linux CLIs. Tests can
// provide a deterministic runner without starting privileged system tools.
type Runner interface {
	Run(ctx context.Context, command string, args ...string) ([]byte, error)
}

type ExecRunner struct{}

func (ExecRunner) Run(ctx context.Context, command string, args ...string) ([]byte, error) {
	output, err := exec.CommandContext(ctx, command, args...).CombinedOutput()
	if err != nil {
		message := strings.TrimSpace(string(output))
		if message == "" {
			message = err.Error()
		}
		return nil, fmt.Errorf("%s: %s", command, message)
	}
	return output, nil
}

type ErrorCode string

const (
	ErrorUnavailable ErrorCode = "UNAVAILABLE"
	ErrorPermission  ErrorCode = "PERMISSION_DENIED"
	ErrorInvalid     ErrorCode = "INVALID_REQUEST"
	ErrorFailed      ErrorCode = "SERVICE_FAILED"
)

// ServiceError is safe to show at the product boundary. Internal command
// output stays available through Unwrap for logs and diagnostics.
type ServiceError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func (e *ServiceError) Error() string { return e.Message }
func (e *ServiceError) Unwrap() error { return e.Cause }

func translateError(action string, err error) error {
	if err == nil {
		return nil
	}
	message := strings.ToLower(err.Error())
	code := ErrorFailed
	userMessage := "Não foi possível concluir: " + action + "."
	if strings.Contains(message, "permission") || strings.Contains(message, "not authorized") {
		code = ErrorPermission
		userMessage = "A operação exige autorização do sistema."
	}
	if strings.Contains(message, "not found") || strings.Contains(message, "no such file") {
		code = ErrorUnavailable
		userMessage = "O serviço do sistema não está disponível."
	}
	return &ServiceError{Code: code, Message: userMessage, Cause: err}
}
