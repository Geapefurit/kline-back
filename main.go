package main

import (
	"context"
	"fmt"

	"github.com/Geapefurit/kline-back/zeus/pkg/beat"
	"github.com/Geapefurit/kline-back/zeus/pkg/db"
)

func main() {
	fmt.Println(db.Init())
	beat.RunSamplingKPoint(context.Background())
	beat.RunSamplingKPrice(context.Background())
}
