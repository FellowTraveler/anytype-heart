package core

import (
	"context"
	"github.com/anyproto/anytype-heart/core/application"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/domain/errcode"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
)

func (mw *Middleware) AccountCreate(cctx context.Context, req *pb.RpcAccountCreateRequest) *pb.RpcAccountCreateResponse {
	response := func(account *model.Account, code pb.RpcAccountCreateResponseErrorCode, err error) *pb.RpcAccountCreateResponse {
		var clientConfig *pb.RpcAccountConfig
		m := &pb.RpcAccountCreateResponse{Config: clientConfig, Account: account, Error: &pb.RpcAccountCreateResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	newAccount, err := mw.applicationService.AccountCreate(cctx, req)
	code, err := domain.UnwrapCodeFromError[pb.RpcAccountCreateResponseErrorCode](err)
	return response(newAccount, code, err)
}

func (mw *Middleware) AccountRecover(cctx context.Context, _ *pb.RpcAccountRecoverRequest) *pb.RpcAccountRecoverResponse {
	response := func(code pb.RpcAccountRecoverResponseErrorCode, err error) *pb.RpcAccountRecoverResponse {
		m := &pb.RpcAccountRecoverResponse{Error: &pb.RpcAccountRecoverResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	err := mw.applicationService.AccountRecover()
	code, err := domain.UnwrapCodeFromError[pb.RpcAccountRecoverResponseErrorCode](err)
	return response(code, err)
}

func (mw *Middleware) AccountSelect(cctx context.Context, req *pb.RpcAccountSelectRequest) *pb.RpcAccountSelectResponse {
	account, err := mw.applicationService.AccountSelect(cctx, req)
	code := errcode.Map(err,
		errcode.To(application.ErrEmptyAccountID, pb.RpcAccountSelectResponseError_BAD_INPUT),
		errcode.To(application.ErrFailedToStopSearcherNode, pb.RpcAccountSelectResponseError_FAILED_TO_STOP_SEARCHER_NODE),
		errcode.To(application.ErrNoMnemonicProvided, pb.RpcAccountSelectResponseError_LOCAL_REPO_NOT_EXISTS_AND_MNEMONIC_NOT_SET),
		errcode.To(application.ErrFailedToCreateLocalRepo, pb.RpcAccountSelectResponseError_FAILED_TO_CREATE_LOCAL_REPO),
		errcode.To(application.ErrFailedToFindAccountInfo, pb.RpcAccountSelectResponseError_FAILED_TO_FIND_ACCOUNT_INFO),
		errcode.To(application.ErrAnotherProcessIsRunning, pb.RpcAccountSelectResponseError_ANOTHER_ANYTYPE_PROCESS_IS_RUNNING),
		errcode.To(application.ErrIncompatibleVersion, pb.RpcAccountSelectResponseError_FAILED_TO_FETCH_REMOTE_NODE_HAS_INCOMPATIBLE_PROTO_VERSION),
		errcode.To(application.ErrFailedToRunNode, pb.RpcAccountSelectResponseError_FAILED_TO_RUN_NODE),
	)
	return &pb.RpcAccountSelectResponse{
		Config:  nil,
		Account: account,
		Error: &pb.RpcAccountSelectResponseError{
			Code:        code,
			Description: getErrorDescription(err),
		},
	}
}

func (mw *Middleware) AccountStop(cctx context.Context, req *pb.RpcAccountStopRequest) *pb.RpcAccountStopResponse {
	response := func(code pb.RpcAccountStopResponseErrorCode, err error) *pb.RpcAccountStopResponse {
		m := &pb.RpcAccountStopResponse{Error: &pb.RpcAccountStopResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	err := mw.applicationService.AccountStop(req)
	code, err := domain.UnwrapCodeFromError[pb.RpcAccountStopResponseErrorCode](err)
	return response(code, err)
}

func (mw *Middleware) AccountMove(cctx context.Context, req *pb.RpcAccountMoveRequest) *pb.RpcAccountMoveResponse {
	response := func(code pb.RpcAccountMoveResponseErrorCode, err error) *pb.RpcAccountMoveResponse {
		m := &pb.RpcAccountMoveResponse{Error: &pb.RpcAccountMoveResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	err := mw.applicationService.AccountMove(req)
	code, err := domain.UnwrapCodeFromError[pb.RpcAccountMoveResponseErrorCode](err)
	return response(code, err)
}

func (mw *Middleware) AccountDelete(cctx context.Context, req *pb.RpcAccountDeleteRequest) *pb.RpcAccountDeleteResponse {
	response := func(status *model.AccountStatus, code pb.RpcAccountDeleteResponseErrorCode, err error) *pb.RpcAccountDeleteResponse {
		m := &pb.RpcAccountDeleteResponse{Error: &pb.RpcAccountDeleteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Status = status
		}

		return m
	}

	status, err := mw.applicationService.AccountDelete(cctx, req)
	code, err := domain.UnwrapCodeFromError[pb.RpcAccountDeleteResponseErrorCode](err)
	return response(status, code, err)
}

func (mw *Middleware) AccountConfigUpdate(_ context.Context, req *pb.RpcAccountConfigUpdateRequest) *pb.RpcAccountConfigUpdateResponse {
	response := func(code pb.RpcAccountConfigUpdateResponseErrorCode, err error) *pb.RpcAccountConfigUpdateResponse {
		m := &pb.RpcAccountConfigUpdateResponse{Error: &pb.RpcAccountConfigUpdateResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	err := mw.applicationService.AccountConfigUpdate(req)
	code, err := domain.UnwrapCodeFromError[pb.RpcAccountConfigUpdateResponseErrorCode](err)
	return response(code, err)
}

func (mw *Middleware) AccountRecoverFromLegacyExport(cctx context.Context,
	req *pb.RpcAccountRecoverFromLegacyExportRequest) *pb.RpcAccountRecoverFromLegacyExportResponse {
	response := func(address string, code pb.RpcAccountRecoverFromLegacyExportResponseErrorCode, err error) *pb.RpcAccountRecoverFromLegacyExportResponse {
		m := &pb.RpcAccountRecoverFromLegacyExportResponse{AccountId: address, Error: &pb.RpcAccountRecoverFromLegacyExportResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	address, err := mw.applicationService.CreateAccountFromExport(req)
	code, err := domain.UnwrapCodeFromError[pb.RpcAccountRecoverFromLegacyExportResponseErrorCode](err)
	return response(address, code, err)
}
