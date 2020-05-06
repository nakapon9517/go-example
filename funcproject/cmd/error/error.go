package errorhandler

import (
	log "github.com/sirupsen/logrus"
)

// AddError error struct
type AddError struct {
	Msg string
}

// SubError error struct
type SubError struct {
	Msg string
}

func (e *AddError) Error() string {
	return "adderror"
}

func (e *SubError) Error() string {
	return "suberror"
}

// CreateAddError create AddError
func CreateAddError(message string) error {
	return &AddError{message}
}

// CreateSubError create SubError
func CreateSubError(message string) error {
	return &SubError{message}
}

// ErrorHandler error is handle
func ErrorHandler(errorparam error) {
	if errorparam != nil {
		switch e := errorparam.(type) {
		case *AddError:
			requestLogger := log.WithFields(log.Fields{"error_id": "AddError", "error_message": e.Msg})
			requestLogger.Info(e.Msg)
			requestLogger.Warn(e.Msg)
		case *SubError:
			requestLogger := log.WithFields(log.Fields{"error_id": "SubError", "error_message": e.Msg})
			requestLogger.Info(e.Msg)
			requestLogger.Warn(e.Msg)
		default:
			requestLogger := log.WithFields(log.Fields{"error_id": "??????", "error_message": "e.Msg"})
			requestLogger.Warn("e.Msg")
		}

	}
}
