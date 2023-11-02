package grpcerror

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAbortedError(pbErrorInfo *errdetails.ErrorInfo) (err error) {
	pbStatus := status.New(codes.Aborted, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbErrorInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewAlreadyExistsError(pbResourceInfo *errdetails.ResourceInfo) (err error) {
	pbStatus := status.New(codes.AlreadyExists, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbResourceInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewCanceledError() (err error) {
	return status.New(codes.Canceled, "").Err()
}

func NewDataLossError(pbDebugInfo *errdetails.DebugInfo) (err error) {
	pbStatus := status.New(codes.DataLoss, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewDeadlineExceededError(pbDebugInfo *errdetails.DebugInfo) (err error) {
	pbStatus := status.New(codes.DeadlineExceeded, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewInternalError(pbDebugInfo *errdetails.DebugInfo) (err error) {
	pbStatus := status.New(codes.Internal, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewInvalidArgumentError(pbBadRequest *errdetails.BadRequest) (err error) {
	pbStatus := status.New(codes.InvalidArgument, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbBadRequest); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewNotFoundError(pbResourceInfo *errdetails.ResourceInfo) (err error) {
	pbStatus := status.New(codes.NotFound, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbResourceInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewOutOfRangeError(pbBadRequest *errdetails.BadRequest) (err error) {
	pbStatus := status.New(codes.OutOfRange, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbBadRequest); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewPermissionDeniedError(pbErrorInfo *errdetails.ErrorInfo) (err error) {
	pbStatus := status.New(codes.PermissionDenied, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbErrorInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewPreconditionFailureError(pbPreconditionFailure *errdetails.PreconditionFailure) (err error) {
	pbStatus := status.New(codes.FailedPrecondition, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbPreconditionFailure); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewResourceExhaustedError(pbQuotaFailure *errdetails.QuotaFailure) (err error) {
	pbStatus := status.New(codes.ResourceExhausted, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbQuotaFailure); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}

func NewUnauthenticatedError(pbErrorInfo *errdetails.ErrorInfo) (err error) {
	pbStatus := status.New(codes.Unauthenticated, "")

	var pbStatusWithDetails *status.Status
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
	pbStatus := status.New(codes.Unknown, "")

	var pbStatusWithDetails *status.Status
	if pbStatusWithDetails, err = pbStatus.WithDetails(pbDebugInfo); err != nil {
		return pbStatus.Err()
	}

	return pbStatusWithDetails.Err()
}
