package esclient

import (
	"bytes"
	"context"
	"github.com/pkg/errors"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/augustus281/cqrs-pattern/pkg/es"
)

func Index(ctx context.Context, transport esapi.Transport, index, documentID string, v any) (*esapi.Response, error) {
	reqBytes, err := es.Marshal(v)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Marshal")
	}

	request := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       bytes.NewBuffer(reqBytes),
	}

	return request.Do(ctx, transport)
}
