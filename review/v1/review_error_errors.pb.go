// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UserNotFound.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_UserNotFound.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ContentMissing.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_ContentMissing.String(), fmt.Sprintf(format, args...))
}

func IsQueryDBFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_QueryDBFailed.String() && e.Code == 500
}

func ErrorQueryDBFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_QueryDBFailed.String(), fmt.Sprintf(format, args...))
}

func IsReviewExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ReviewExist.String() && e.Code == 400
}

func ErrorReviewExist(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_ReviewExist.String(), fmt.Sprintf(format, args...))
}

func IsServerBusy(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ServerBusy.String() && e.Code == 500
}

func ErrorServerBusy(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_ServerBusy.String(), fmt.Sprintf(format, args...))
}

func IsQueryWithNoRows(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_QueryWithNoRows.String() && e.Code == 500
}

func ErrorQueryWithNoRows(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_QueryWithNoRows.String(), fmt.Sprintf(format, args...))
}

func IsReviewHasReply(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ReviewHasReply.String() && e.Code == 400
}

func ErrorReviewHasReply(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_ReviewHasReply.String(), fmt.Sprintf(format, args...))
}

func IsInvalidParams(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_InvalidParams.String() && e.Code == 400
}

func ErrorInvalidParams(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_InvalidParams.String(), fmt.Sprintf(format, args...))
}
