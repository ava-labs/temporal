// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

package matching

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// MatchingService_RespondQueryTaskCompleted_Args represents the arguments for the MatchingService.RespondQueryTaskCompleted function.
//
// The arguments for RespondQueryTaskCompleted are sent and received over the wire as this struct.
type MatchingService_RespondQueryTaskCompleted_Args struct {
	Request *RespondQueryTaskCompletedRequest `json:"request,omitempty"`
}

// ToWire translates a MatchingService_RespondQueryTaskCompleted_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *MatchingService_RespondQueryTaskCompleted_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RespondQueryTaskCompletedRequest_1_Read(w wire.Value) (*RespondQueryTaskCompletedRequest, error) {
	var v RespondQueryTaskCompletedRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a MatchingService_RespondQueryTaskCompleted_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a MatchingService_RespondQueryTaskCompleted_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v MatchingService_RespondQueryTaskCompleted_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *MatchingService_RespondQueryTaskCompleted_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _RespondQueryTaskCompletedRequest_1_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a MatchingService_RespondQueryTaskCompleted_Args
// struct.
func (v *MatchingService_RespondQueryTaskCompleted_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("MatchingService_RespondQueryTaskCompleted_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this MatchingService_RespondQueryTaskCompleted_Args match the
// provided MatchingService_RespondQueryTaskCompleted_Args.
//
// This function performs a deep comparison.
func (v *MatchingService_RespondQueryTaskCompleted_Args) Equals(rhs *MatchingService_RespondQueryTaskCompleted_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "RespondQueryTaskCompleted" for this struct.
func (v *MatchingService_RespondQueryTaskCompleted_Args) MethodName() string {
	return "RespondQueryTaskCompleted"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *MatchingService_RespondQueryTaskCompleted_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// MatchingService_RespondQueryTaskCompleted_Helper provides functions that aid in handling the
// parameters and return values of the MatchingService.RespondQueryTaskCompleted
// function.
var MatchingService_RespondQueryTaskCompleted_Helper = struct {
	// Args accepts the parameters of RespondQueryTaskCompleted in-order and returns
	// the arguments struct for the function.
	Args func(
		request *RespondQueryTaskCompletedRequest,
	) *MatchingService_RespondQueryTaskCompleted_Args

	// IsException returns true if the given error can be thrown
	// by RespondQueryTaskCompleted.
	//
	// An error can be thrown by RespondQueryTaskCompleted only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for RespondQueryTaskCompleted
	// given the error returned by it. The provided error may
	// be nil if RespondQueryTaskCompleted did not fail.
	//
	// This allows mapping errors returned by RespondQueryTaskCompleted into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// RespondQueryTaskCompleted
	//
	//   err := RespondQueryTaskCompleted(args)
	//   result, err := MatchingService_RespondQueryTaskCompleted_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from RespondQueryTaskCompleted: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*MatchingService_RespondQueryTaskCompleted_Result, error)

	// UnwrapResponse takes the result struct for RespondQueryTaskCompleted
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if RespondQueryTaskCompleted threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := MatchingService_RespondQueryTaskCompleted_Helper.UnwrapResponse(result)
	UnwrapResponse func(*MatchingService_RespondQueryTaskCompleted_Result) error
}{}

func init() {
	MatchingService_RespondQueryTaskCompleted_Helper.Args = func(
		request *RespondQueryTaskCompletedRequest,
	) *MatchingService_RespondQueryTaskCompleted_Args {
		return &MatchingService_RespondQueryTaskCompleted_Args{
			Request: request,
		}
	}

	MatchingService_RespondQueryTaskCompleted_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		default:
			return false
		}
	}

	MatchingService_RespondQueryTaskCompleted_Helper.WrapResponse = func(err error) (*MatchingService_RespondQueryTaskCompleted_Result, error) {
		if err == nil {
			return &MatchingService_RespondQueryTaskCompleted_Result{}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for MatchingService_RespondQueryTaskCompleted_Result.BadRequestError")
			}
			return &MatchingService_RespondQueryTaskCompleted_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for MatchingService_RespondQueryTaskCompleted_Result.InternalServiceError")
			}
			return &MatchingService_RespondQueryTaskCompleted_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for MatchingService_RespondQueryTaskCompleted_Result.EntityNotExistError")
			}
			return &MatchingService_RespondQueryTaskCompleted_Result{EntityNotExistError: e}, nil
		}

		return nil, err
	}
	MatchingService_RespondQueryTaskCompleted_Helper.UnwrapResponse = func(result *MatchingService_RespondQueryTaskCompleted_Result) (err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		return
	}

}

// MatchingService_RespondQueryTaskCompleted_Result represents the result of a MatchingService.RespondQueryTaskCompleted function call.
//
// The result of a RespondQueryTaskCompleted execution is sent and received over the wire as this struct.
type MatchingService_RespondQueryTaskCompleted_Result struct {
	BadRequestError      *shared.BadRequestError      `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError `json:"internalServiceError,omitempty"`
	EntityNotExistError  *shared.EntityNotExistsError `json:"entityNotExistError,omitempty"`
}

// ToWire translates a MatchingService_RespondQueryTaskCompleted_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *MatchingService_RespondQueryTaskCompleted_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("MatchingService_RespondQueryTaskCompleted_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a MatchingService_RespondQueryTaskCompleted_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a MatchingService_RespondQueryTaskCompleted_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v MatchingService_RespondQueryTaskCompleted_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *MatchingService_RespondQueryTaskCompleted_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("MatchingService_RespondQueryTaskCompleted_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a MatchingService_RespondQueryTaskCompleted_Result
// struct.
func (v *MatchingService_RespondQueryTaskCompleted_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}

	return fmt.Sprintf("MatchingService_RespondQueryTaskCompleted_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this MatchingService_RespondQueryTaskCompleted_Result match the
// provided MatchingService_RespondQueryTaskCompleted_Result.
//
// This function performs a deep comparison.
func (v *MatchingService_RespondQueryTaskCompleted_Result) Equals(rhs *MatchingService_RespondQueryTaskCompleted_Result) bool {
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "RespondQueryTaskCompleted" for this struct.
func (v *MatchingService_RespondQueryTaskCompleted_Result) MethodName() string {
	return "RespondQueryTaskCompleted"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *MatchingService_RespondQueryTaskCompleted_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
