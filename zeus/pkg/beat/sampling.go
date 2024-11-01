package beat

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Geapefurit/kline-back/proto/kline"
	kpriceproto "github.com/Geapefurit/kline-back/proto/kline/zeus/v1/kprice"
	tokenproto "github.com/Geapefurit/kline-back/proto/kline/zeus/v1/token"
	tokenpairproto "github.com/Geapefurit/kline-back/proto/kline/zeus/v1/tokenpair"
	"github.com/Geapefurit/kline-back/zeus/pkg/mw/v1/kprice"
	"github.com/Geapefurit/kline-back/zeus/pkg/mw/v1/token"
	"github.com/Geapefurit/kline-back/zeus/pkg/mw/v1/tokenpair"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type SamplingKPriceTask struct {
	// Token0 address -> Token1 address -> TokenPair ID
	tPairMap   map[string]map[string]uint32
	changeLock sync.Mutex
}

func GetSamplingKPriceTask(ctx context.Context) (*SamplingKPriceTask, error) {
	task := &SamplingKPriceTask{
		tPairMap: map[string]map[string]uint32{},
	}
	if err := task.loadTPairMap(ctx); err != nil {
		return nil, err
	}
	return task, nil
}

func (st *SamplingKPriceTask) loadTPairMap(ctx context.Context) error {
	//TODO: please read records by paging
	tpH, err := tokenpair.NewHandler(ctx,
		tokenpair.WithConds(
			&tokenpairproto.Conds{}),
		tokenpair.WithLimit(0),
		tokenpair.WithOffset(0))
	if err != nil {
		return err
	}

	tpInfos, _, err := tpH.GetTokenPairs(ctx)
	if err != nil {
		return err
	}

	for _, tpInfo := range tpInfos {
		if _, ok := st.tPairMap[tpInfo.TokenOneAddress]; !ok {
			st.tPairMap[tpInfo.TokenOneAddress] = make(map[string]uint32)
		}
		st.tPairMap[tpInfo.TokenOneAddress][tpInfo.TokenTwoAddress] = tpInfo.ID
	}

	return nil
}

func checkAndCreateToken(ctx context.Context, address string) (*tokenproto.Token, error) {
	tokenH, err := token.NewHandler(ctx,
		token.WithConds(&tokenproto.Conds{
			Address: &kline.StringVal{
				Op:    cruder.EQ,
				Value: address,
			},
		}),
		token.WithOffset(0),
		token.WithLimit(1),
	)

	if err != nil {
		return nil, err
	}

	tokenInfos, _, err := tokenH.GetTokens(ctx)
	if err != nil {
		return nil, err
	}

	if len(tokenInfos) > 0 {
		return tokenInfos[0], nil
	}

	tokenReq := MockTokenInfo(address)

	tokenH, err = token.NewHandler(ctx,
		token.WithAddress(&address, true),
		token.WithSite(tokenReq.Site, true),
		token.WithIcon(tokenReq.Icon, true),
		token.WithName(tokenReq.Name, true),
	)
	if err != nil {
		return nil, err
	}

	tokenInfo, err := tokenH.CreateToken(ctx)
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func checkTokenPair(ctx context.Context, tokenOneID, tokenTwoID uint32) (*tokenpairproto.TokenPair, error) {
	queryH, err := tokenpair.NewHandler(ctx,
		tokenpair.WithConds(&tokenpairproto.Conds{
			TokenOneID: &kline.Uint32Val{
				Op:    cruder.EQ,
				Value: tokenOneID,
			},
			TokenTwoID: &kline.Uint32Val{
				Op:    cruder.EQ,
				Value: tokenTwoID,
			},
		}),
		tokenpair.WithLimit(1),
		tokenpair.WithOffset(0),
	)
	if err != nil {
		return nil, err
	}

	tpInfos, _, err := queryH.GetTokenPairs(ctx)
	if err != nil {
		return nil, err
	}

	if len(tpInfos) > 0 {
		return tpInfos[0], nil
	}

	createH, err := tokenpair.NewHandler(ctx,
		tokenpair.WithTokenOneID(&tokenOneID, true),
		tokenpair.WithTokenTwoID(&tokenTwoID, true),
	)
	if err != nil {
		return nil, err
	}
	if err := createH.CreateTokenPair(ctx); err != nil {
		return nil, err
	}

	tpInfos, _, err = queryH.GetTokenPairs(ctx)
	if err != nil {
		return nil, err
	}

	if len(tpInfos) > 0 {
		return tpInfos[0], nil
	}

	return nil, fmt.Errorf("failed to create tokenpair")
}

func (st *SamplingKPriceTask) getTokenPairID(ctx context.Context, tokenOneAddress, tokenTwoAddress string) (uint32, error) {
	st.changeLock.Lock()
	defer st.changeLock.Unlock()

	if _, ok := st.tPairMap[tokenOneAddress]; ok {
		if _, ok := st.tPairMap[tokenOneAddress][tokenTwoAddress]; ok {
			return st.tPairMap[tokenOneAddress][tokenTwoAddress], nil
		}
	} else {
		st.tPairMap[tokenOneAddress] = make(map[string]uint32)
	}

	tokenOne, err := checkAndCreateToken(ctx, tokenOneAddress)
	if err != nil {
		return 0, err
	}
	if tokenOne == nil {
		return 0, fmt.Errorf("failed to create token")
	}

	tokenTwo, err := checkAndCreateToken(ctx, tokenTwoAddress)
	if err != nil {
		return 0, err
	}
	if tokenTwo == nil {
		return 0, fmt.Errorf("failed to create token")
	}

	tpInfo, err := checkTokenPair(ctx, tokenOne.ID, tokenTwo.ID)
	if err != nil {
		return 0, err
	}
	if tpInfo == nil {
		return 0, fmt.Errorf("failed to create tokenpair")
	}

	st.tPairMap[tokenOneAddress][tokenTwoAddress] = tpInfo.ID

	return tpInfo.ID, nil
}

func (st *SamplingKPriceTask) SamplingKPrice(ctx context.Context, interval time.Duration) {
	for {
		select {
		case <-time.NewTicker(interval).C:
			go func() {
				err := st.samplingAndStore(ctx)
				if err != nil {
					fmt.Println(err)
				}
			}()
		case <-ctx.Done():
			return
		}
	}
}

type KPriceData struct {
	TokenOneAddress string
	TokenTwoAddress string
	Price           float64
	Time            uint32
}

func createKPrices(ctx context.Context, kpriceReqs []*kpriceproto.KPriceReq) error {
	for _, req := range kpriceReqs {
		createH, err := kprice.NewHandler(ctx,
			kprice.WithTokenPairID(req.TokenPairID, true),
			kprice.WithPrice(req.Price, true),
			kprice.WithTime(req.Time, true),
		)
		if err != nil {
			return err
		}
		if err := createH.CreateKPrice(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (st *SamplingKPriceTask) samplingAndStore(ctx context.Context) error {
	kpDataList := MockKPrice()
	kpriceReqs := []*kpriceproto.KPriceReq{}
	for _, kpData := range kpDataList {
		tpID, err := st.getTokenPairID(ctx, kpData.TokenOneAddress, kpData.TokenTwoAddress)
		if err != nil {
			return err
		}
		kpriceReqs = append(kpriceReqs, &kpriceproto.KPriceReq{
			TokenPairID: &tpID,
			Price:       &kpData.Price,
			Time:        &kpData.Time,
		})
	}

	if err := createKPrices(ctx, kpriceReqs); err != nil {
		return err
	}

	return nil
}

func RunSamplingKPrice(ctx context.Context) {
	samplingTask, err := GetSamplingKPriceTask(ctx)
	if err != nil {
		panic(err)
	}
	samplingTask.SamplingKPrice(ctx, time.Second)
}

func MockKPrice() []*KPriceData {
	now := uint32(time.Now().Unix())
	return []*KPriceData{
		{
			TokenOneAddress: "a1",
			TokenTwoAddress: "a2",
			Price:           rand.Float64(),
			Time:            now,
		},
		{
			TokenOneAddress: "a3",
			TokenTwoAddress: "a4",
			Price:           rand.Float64(),
			Time:            now,
		},
		{
			TokenOneAddress: "a5",
			TokenTwoAddress: "a6",
			Price:           rand.Float64(),
			Time:            now,
		},
		{
			TokenOneAddress: "a7",
			TokenTwoAddress: "a8",
			Price:           rand.Float64(),
			Time:            now,
		},
		{
			TokenOneAddress: "a9",
			TokenTwoAddress: "a0",
			Price:           rand.Float64(),
			Time:            now,
		},
	}
}

func MockTokenInfo(address string) *tokenproto.TokenReq {
	randomStr := uuid.NewString()
	site := fmt.Sprintf("mock token %v", randomStr)
	icon := fmt.Sprintf("mock token %v", randomStr)
	name := fmt.Sprintf("mock token %v", randomStr)

	return &tokenproto.TokenReq{
		Address: &address,
		Site:    &site,
		Icon:    &icon,
		Name:    &name,
	}
}
