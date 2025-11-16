package runner

import (
	"github.com/NFTGalaxy/v8runner/pkg/types"
	v8 "github.com/stumble/v8go"
)

func errResult(id string, err error) types.RunCodeResponse {
	errStr := err.Error()
	return types.RunCodeResponse{
		ID:    id,
		Error: &errStr,
	}
}

func nilResult(id string) types.RunCodeResponse {
	return types.RunCodeResponse{
		ID: id,
	}
}

func jsonResult(id string, codeCtx *v8.Context, val *v8.Value) types.RunCodeResponse {
	jsonStr, err := v8.JSONStringify(codeCtx, val)
	if err != nil {
		return errResult(id, err)
	}
	return types.RunCodeResponse{
		ID:     id,
		Result: &jsonStr,
	}
}
