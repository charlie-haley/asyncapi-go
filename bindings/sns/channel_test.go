package sns

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestChannelBinding_JSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ChannelBinding
	}{
		{
			name: "full configuration",
			input: `{
				"name": "MyTopic",
				"ordering": {
					"type": "FIFO",
					"contentBasedDeduplication": true
				},
				"policy": {
					"statements": [
						{
							"effect": "Allow",
							"principal": "*",
							"action": "sns:Publish",
							"resource": "arn:aws:sns:us-west-2:123456789012:MyTopic"
						}
					]
				},
				"tags": {
					"environment": "production",
					"team": "platform"
				},
				"bindingVersion": "1.0.0"
			}`,
			expected: ChannelBinding{
				Name: "MyTopic",
				Ordering: &Ordering{
					Type: "FIFO",
					ContentBasedDeduplication: true,
				},
				Policy: &Policy{
					Statements: []Statement{
						{
							Effect:    "Allow",
							Principal: "*",
							Action:    "sns:Publish",
							Resource:  "arn:aws:sns:us-west-2:123456789012:MyTopic",
						},
					},
				},
				Tags: map[string]string{
					"environment": "production",
					"team":       "platform",
				},
				BindingVersion: "1.0.0",
			},
		},
		{
			name: "minimal configuration",
			input: `{
				"name": "MyTopic"
			}`,
			expected: ChannelBinding{
				Name: "MyTopic",
			},
		},
		{
			name: "with array actions",
			input: `{
				"name": "MyTopic",
				"policy": {
					"statements": [
						{
							"effect": "Allow",
							"principal": "*",
							"action": ["sns:Publish", "sns:Subscribe"]
						}
					]
				}
			}`,
			expected: ChannelBinding{
				Name: "MyTopic",
				Policy: &Policy{
					Statements: []Statement{
						{
							Effect:    "Allow",
							Principal: "*",
							Action:    []interface{}{"sns:Publish", "sns:Subscribe"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ChannelBinding
			err := json.Unmarshal([]byte(tt.input), &got)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)

			// Test marshaling
			marshaled, err := json.Marshal(got)
			assert.NoError(t, err)

			var unmarshaled ChannelBinding
			err = json.Unmarshal(marshaled, &unmarshaled)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, unmarshaled)
		})
	}
}

func TestChannelBinding_YAML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ChannelBinding
	}{
		{
			name: "full configuration",
			input: `
name: MyTopic
ordering:
  type: FIFO
  contentBasedDeduplication: true
policy:
  statements:
    - effect: Allow
      principal: "*"
      action: sns:Publish
      resource: arn:aws:sns:us-west-2:123456789012:MyTopic
tags:
  environment: production
  team: platform
bindingVersion: "1.0.0"
`,
			expected: ChannelBinding{
				Name: "MyTopic",
				Ordering: &Ordering{
					Type: "FIFO",
					ContentBasedDeduplication: true,
				},
				Policy: &Policy{
					Statements: []Statement{
						{
							Effect:    "Allow",
							Principal: "*",
							Action:    "sns:Publish",
							Resource:  "arn:aws:sns:us-west-2:123456789012:MyTopic",
						},
					},
				},
				Tags: map[string]string{
					"environment": "production",
					"team":       "platform",
				},
				BindingVersion: "1.0.0",
			},
		},
		{
			name: "with complex principal",
			input: `
name: MyTopic
policy:
  statements:
    - effect: Allow
      principal:
        AWS:
          - "arn:aws:iam::123456789012:root"
          - "arn:aws:iam::123456789012:user/test"
      action: sns:Publish
`,
			expected: ChannelBinding{
				Name: "MyTopic",
				Policy: &Policy{
					Statements: []Statement{
						{
							Effect: "Allow",
							Principal: map[string]interface{}{
								"AWS": []interface{}{
									"arn:aws:iam::123456789012:root",
									"arn:aws:iam::123456789012:user/test",
								},
							},
							Action: "sns:Publish",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ChannelBinding
			err := yaml.Unmarshal([]byte(tt.input), &got)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)

			// Test marshaling
			marshaled, err := yaml.Marshal(got)
			assert.NoError(t, err)

			var unmarshaled ChannelBinding
			err = yaml.Unmarshal(marshaled, &unmarshaled)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, unmarshaled)
		})
	}
}