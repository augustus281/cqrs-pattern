package esclient

import (
	"context"
	"github.com/pkg/errors"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/augustus281/cqrs-pattern/pkg/es"
)

func GetByID[T any, V GetResponse[T]](ctx context.Context, transport esapi.Transport, index, documentID string) (*V, error) {
	request := esapi.GetRequest{
		Index:      index,
		DocumentID: documentID,
		Pretty:     true,
	}

	response, err := request.Do(ctx, transport)
	if err != nil {
		return new(V), err
	}
	defer response.Body.Close()

	if response.IsError() {
		return nil, errors.Wrapf(errors.New("ElasticSearch GetByID err"), "documentID: %s, status: %s", documentID, response.Status())
	}

	var getResponse V
	if err := es.NewDecoder(response.Body).Decode(&getResponse); err != nil {
		return new(V), err
	}

	return &getResponse, nil
}
