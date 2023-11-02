package grpcerror

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAbortedError(pbErrorInfo *errdetails.ErrorInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.Aborted, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbErrorInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewAlreadyExistsError(pbResourceInfo *errdetails.ResourceInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.AlreadyExists, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbResourceInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewCanceledError() (err error) {
	return status.New(codes.Canceled, "").Err()
}

func NewDataLossError(pbDebugInfo *errdetails.DebugInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.DataLoss, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewDeadlineExceededError(pbDebugInfo *errdetails.DebugInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.DeadlineExceeded, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewInternalError(pbDebugInfo *errdetails.DebugInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.Internal, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewInvalidArgumentError(pbBadRequest *errdetails.BadRequest) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.InvalidArgument, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbBadRequest); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewNotFoundError(pbResourceInfo *errdetails.ResourceInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.NotFound, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbResourceInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewOutOfRangeError(pbBadRequest *errdetails.BadRequest) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.OutOfRange, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbBadRequest); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewPermissionDeniedError(pbErrorInfo *errdetails.ErrorInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.PermissionDenied, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbErrorInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewPreconditionFailureError(pbPreconditionFailure *errdetails.PreconditionFailure) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.FailedPrecondition, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbPreconditionFailure); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewResourceExhaustedError(pbQuotaFailure *errdetails.QuotaFailure) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.ResourceExhausted, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbQuotaFailure); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewUnauthenticatedError(pbErrorInfo *errdetails.ErrorInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.Unauthenticated, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbErrorInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewUnavailableError() (err error) {
	return status.New(codes.Unavailable, "").Err()
}

func NewUnimplementedError() (err error) {
	return status.New(codes.Unimplemented, "").Err()
}

func NewUnknownError(pbDebugInfo *errdetails.DebugInfo) (err error) {
	var pbStatusWithDetails *status.Status

	pbStatus := status.New(codes.Unknown, "")
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}
