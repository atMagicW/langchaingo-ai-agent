package llm

import (
	"context"
	"errors"
	"langchaingo-ai-agent/pkg/utils"
)

// FallbackClient 支持多模型兜底
type FallbackClient struct {
	Clients []Client
}

func NewFallbackClient(clients ...Client) *FallbackClient {
	return &FallbackClient{Clients: clients}
}

func (f *FallbackClient) Generate(
	ctx context.Context,
	messages []Message,
) (string, error) {

	for _, c := range f.Clients {
		err := utils.Retry(2, func() error {
			var err error
			_, err = c.Generate(ctx, messages)
			return err
		})

		if err == nil {
			return c.Generate(ctx, messages)
		}
	}

	return "", errors.New("all llm clients failed")
}
