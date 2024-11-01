package tokenpair

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/Geapefurit/kline-back/common/utils"
	"github.com/Geapefurit/kline-back/proto/kline"
	"github.com/Geapefurit/kline-back/zeus/pkg/db"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	tokenpairproto "github.com/Geapefurit/kline-back/proto/kline/zeus/v1/tokenpair"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var tokenPTRet = &tokenpairproto.TokenPair{
	Remark: "test token pair ",
}

var tokenPTReq = &tokenpairproto.TokenPairReq{
	Remark: &tokenPTRet.Remark,
}

func createTP(t *testing.T) {
	tokenPTRet.TokenOneID = token1Ret.ID
	tokenPTRet.TokenTwoID = token2Ret.ID
	tokenPTReq.TokenOneID = &token1Ret.ID
	tokenPTReq.TokenTwoID = &token2Ret.ID

	handler, err := NewHandler(
		context.Background(),
		WithTokenOneID(tokenPTReq.TokenOneID, true),
		WithTokenTwoID(tokenPTReq.TokenTwoID, true),
		WithRemark(tokenPTReq.Remark, true),
	)
	assert.Nil(t, err)

	err = handler.CreateTokenPair(context.Background())
	assert.Nil(t, err)
}

func queryTP(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&tokenpairproto.Conds{
			TokenOneID: &kline.Uint32Val{
				Op:    cruder.EQ,
				Value: *tokenPTReq.TokenOneID,
			},
			TokenTwoID: &kline.Uint32Val{
				Op:    cruder.EQ,
				Value: *tokenPTReq.TokenTwoID,
			},
			Remark: &kline.StringVal{
				Op:    cruder.EQ,
				Value: *tokenPTReq.Remark,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)
	infos, total, err := handler.GetTokenPairs(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))

	tokenPTReq.ID = &infos[0].ID
	tokenPTRet.ID = infos[0].ID

	handler, err = NewHandler(
		context.Background(),
		WithID(tokenPTReq.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetTokenPair(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, infos[0], info)
	fmt.Println(utils.PrettyStruct(tokenPTReq))
	fmt.Println(utils.PrettyStruct(tokenPTRet))
}

func updateTP(t *testing.T) {
	remark := uuid.NewString()
	tokenPTReq.Remark = &remark
	tokenPTRet.Remark = remark

	handler, err := NewHandler(
		context.Background(),
		WithID(tokenPTReq.ID, true),
		WithTokenTwoID(tokenPTReq.TokenTwoID, false),
		WithRemark(tokenPTReq.Remark, false),
	)
	assert.Nil(t, err)
	err = handler.UpdateTokenPair(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetTokenPair(context.Background())
	assert.Nil(t, err)
	tokenPTRet.UpdatedAt = info.UpdatedAt
	tokenPTRet.CreatedAt = info.CreatedAt
	assert.Equal(t, tokenPTRet, info)
}

func deleteTP(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(tokenPTReq.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteTokenPair(context.Background())
	assert.Nil(t, err)
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	err := db.Init()
	if err != nil {
		fmt.Printf("cannot init database: %v \n", err)
		os.Exit(0)
	}
	t.Run("createToken1", createToken1)
	t.Run("createToken2", createToken2)
	t.Run("create", createTP)
	t.Run("query", queryTP)
	t.Run("update", updateTP)
	t.Run("delete", deleteTP)
	t.Run("deleteToken1", deleteToken1)
	t.Run("deleteToken2", deleteToken2)
}
