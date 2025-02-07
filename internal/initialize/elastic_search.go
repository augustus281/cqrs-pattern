package initialize

import (
	v7 "github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"github.com/augustus281/cqrs-pattern/global"
)

func InitElasticSearch() (*v7.Client, error) {
	client, err := v7.NewClient(
		v7.SetURL(global.Config.ElasticSearch.Url),
		v7.SetSniff(global.Config.ElasticSearch.Sniff),
		v7.SetGzip(global.Config.ElasticSearch.Gzip),
	)
	if err != nil {
		return nil, errors.Wrap(err, "v7.NewClient")
	}
	global.Logger.Info("init elasticsearch successfully!")
	return client, nil
}
