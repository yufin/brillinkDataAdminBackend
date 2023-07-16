package minioclient

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-admin/config"
)

var MinioCli *MinioClient

func InitMinioCli() error {
	if !config.ExtConfig.Minio.Activate {
		return nil
	}
	endpoint := config.ExtConfig.Minio.Endpoint
	ak := config.ExtConfig.Minio.AccessKey
	sk := config.ExtConfig.Minio.SecretKey
	useSSL := config.ExtConfig.Minio.UseSsl
	moCli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(ak, sk, ""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}
	MinioCli = &MinioClient{
		Cli: moCli,
	}
	log.Info(pkg.Green("Minio Client initialized."))
	return nil
}
