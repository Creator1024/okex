package api

import (
	"context"
	"fmt"
	"github.com/Creator1024/okex"
	"testing"
)

func Test(t *testing.T) {
	apiKey := ""
	secretKey := ""
	passphrase := ""
	dest := okex.AwsServer
	ctx := context.Background()
	var err error
	client, err := NewClient(ctx, apiKey, secretKey, passphrase, dest)
	if err != nil {
		panic(err)
	}
	res, err := client.Rest.Funding.GetAssetValuation()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.AssetValuation[0].TotalBal)
}
