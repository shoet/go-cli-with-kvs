package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/upstash/pulumi-upstash/sdk/go/upstash"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		createdDb, err := upstash.NewRedisDatabase(ctx, "go-cli-with-kvs", &upstash.RedisDatabaseArgs{
			DatabaseName: pulumi.String("pulumi-go-db"),
			Multizone:    pulumi.Bool(true),
			Region:       pulumi.String("ap-northeast-1"),
			Tls:          pulumi.Bool(true),
			Eviction:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		dbFromGet := upstash.LookupRedisDatabaseOutput(ctx, upstash.LookupRedisDatabaseOutputArgs{
			DatabaseId: createdDb.DatabaseId,
		}, nil)

		ctx.Export("db from get request", dbFromGet)

		return nil
	})
}
