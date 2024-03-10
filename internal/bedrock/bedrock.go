// Package bedrock contains aws logic
package bedrock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"jenn-ai/internal/bedrock/models"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

// AWSModelConfig contains all settings for AWS Models
type AWSModelConfig struct {
	ModelID     string  // ModelID is the ID of the model to invoke
	Temperature float64 // Temperature is part of model settings (Anthropic, Cohere)
	TopP        float64 // TopP is part of model settings (Anthropic, Cohere)
	TopK        int     // TopK is part of model settings (Anthropic, Cohere)
	MaxTokens   int     // MaxTokens is part of model settings (Anthropic, Cohere)
}

// ClientRuntimeAPI interface is for support mocking of API calls
type ClientRuntimeAPI interface {
	InvokeModelWithResponseStream(ctx context.Context, params *bedrockruntime.InvokeModelWithResponseStreamInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error)
}

// StreamingOutputHandler used for processing streaming output
type StreamingOutputHandler func(ctx context.Context, part []byte) error

// NewModel is a factory for AWS Models
func NewModel(modelID string, temp, topP float64, topK, tokens int) AWSModelConfig {
	return AWSModelConfig{
		ModelID:     modelID,
		Temperature: temp,
		TopP:        topP,
		TopK:        topK,
		MaxTokens:   tokens,
	}
}

// InvokeModel runs prompt with settings with InvokeModelWithResponseStream
func (m *AWSModelConfig) InvokeModel(ctx context.Context, api ClientRuntimeAPI, message string) (string, error) {
	contentTypeVar := "application/json"

	payload, err := m.constructPayload(message)
	if err != nil {
		return "", err
	}

	// invoke model
	output, err := api.InvokeModelWithResponseStream(ctx, &bedrockruntime.InvokeModelWithResponseStreamInput{
		ContentType: &contentTypeVar,
		ModelId:     &m.ModelID,
		Body:        payload,
	})
	if err != nil {
		return "", err
	}

	// handle stream response chunks by model
	response, err := m.processStreamingOutput(output, func(ctx context.Context, part []byte) error {
		fmt.Print(string(part))
		return nil
	})
	if err != nil {
		return "", err
	}

	return response, nil
}

func (m *AWSModelConfig) constructPayload(message string) ([]byte, error) {
	switch {
	case strings.Contains(m.ModelID, "anthropic"):

		body := models.ClaudeModelInputs{
			Prompt:            fmt.Sprintf("\n\nHuman: %s\n\nAssistant:", message),
			MaxTokensToSample: m.MaxTokens,
			Temperature:       m.Temperature,
			TopP:              m.TopP,
			TopK:              m.TopK,
		}
		// TODO: work on v3
		//body := models.ClaudeMessagesInput{
		//	AnthropicVersion: "bedrock-2023-05-31",
		//	Messages: []models.ClaudeMessage{
		//		{
		//			Role: "user",
		//			Content: []models.ClaudeContent{
		//				{
		//					Type: "text",
		//					Text: message,
		//				},
		//			},
		//		},
		//	},
		//	MaxTokens:   m.MaxTokens,
		//	Temperature: m.Temperature,
		//	TopP:        m.TopP,
		//	TopK:        m.TopK,
		//}

		payload, err := json.Marshal(body)
		if err != nil {
			return []byte{}, err
		}
		return payload, nil
	case strings.Contains(m.ModelID, "cohere"):

		body := models.CommandModelInput{
			Prompt:            message,
			MaxTokensToSample: m.MaxTokens,
			Temperature:       m.Temperature,
			TopP:              m.TopP,
			TopK:              m.TopK,
			StopSequences:     []string{`""`},
			ReturnLiklihoods:  "NONE",
			NumGenerations:    1,
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return []byte{}, err
		}
		return payload, nil
	default:
		fmt.Println("ModelID not provided or unknown")
	}
	return []byte{}, fmt.Errorf("Unable to construct payload for model: %s", m.ModelID)
}

func (m *AWSModelConfig) processStreamingOutput(output *bedrockruntime.InvokeModelWithResponseStreamOutput, handler StreamingOutputHandler) (string, error) {
	var combinedResult string

	for event := range output.GetStream().Events() {
		switch v := event.(type) {
		case *types.ResponseStreamMemberChunk:
			// nested switch case for stream outputs. ugh
			switch {
			case strings.Contains(m.ModelID, "anthropic"):
				var resp models.ClaudeModelOutputs
				err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp)
				if err != nil {
					return combinedResult, err
				}

				handler(context.Background(), []byte(resp.Completion))
				combinedResult += resp.Completion
			case strings.Contains(m.ModelID, "cohere"):
				var resp models.CommandModelOutput
				err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp)
				if err != nil {
					return combinedResult, err
				}

				handler(context.Background(), []byte(resp.Generations[0].Text))
				combinedResult += resp.Generations[0].Text
			default:
				fmt.Println("Unable to determine AWS Model")
			}

		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)

		default:
			fmt.Println("union is nil or unknown type")
		}
	}
	fmt.Println()

	return combinedResult, nil
}
