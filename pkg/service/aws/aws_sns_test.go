package aws

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nonchan7720/go-storage-to-messenger/pkg/config"
	"github.com/nonchan7720/go-storage-to-messenger/pkg/interfaces"
	mockAws "github.com/nonchan7720/go-storage-to-messenger/pkg/mock/aws"
	"github.com/stretchr/testify/require"
)

func TestAWSSNSWithTable(t *testing.T) {
	var tables = []struct {
		name    string
		payload interfaces.Payload
		topic   config.Topic
		fn      func(client *mockAws.MockSNSClient, require *require.Assertions)
	}{
		{
			name: "Create Row",
			payload: interfaces.Payload{
				Event:  interfaces.Create,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(nil, map[string]interface{}{"key": "value"}),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.TableType,
					Table: &config.TransformTable{
						Schema:    "public",
						TableName: "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"CREATE","schema":"public","table":"test","row":{"old_row":{},"new_row":{"key":"value"}}}`)
				}).Return(output, nil).Times(1)
			},
		},
		{
			name: "Create Rows",
			payload: interfaces.Payload{
				Event:  interfaces.Create,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(nil, map[string]interface{}{"key": "value"}),
					interfaces.NewPayloadRow(nil, map[string]interface{}{"key": "value"}),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.TableType,
					Table: &config.TransformTable{
						Schema:    "public",
						TableName: "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"CREATE","schema":"public","table":"test","row":{"old_row":{},"new_row":{"key":"value"}}}`)
				}).Return(output, nil).Times(2)
			},
		},
		{
			name: "Update Row",
			payload: interfaces.Payload{
				Event:  interfaces.Update,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(map[string]interface{}{"key": "oldValue"}, map[string]interface{}{"key": "value"}),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.TableType,
					Table: &config.TransformTable{
						Schema:    "public",
						TableName: "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"UPDATE","schema":"public","table":"test","row":{"old_row":{"key":"oldValue"},"new_row":{"key":"value"}}}`)
				}).Return(output, nil).Times(1)
			},
		},
		{
			name: "Delete Row",
			payload: interfaces.Payload{
				Event:  interfaces.Delete,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(map[string]interface{}{"key": "oldValue"}, nil),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.TableType,
					Table: &config.TransformTable{
						Schema:    "public",
						TableName: "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"DELETE","schema":"public","table":"test","row":{"old_row":{"key":"oldValue"},"new_row":{}}}`)
				}).Return(output, nil).Times(1)
			},
		},
	}
	ctx := context.Background()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	for _, tbl := range tables {
		tbl := tbl
		t.Run(tbl.name, func(t *testing.T) {
			require := require.New(t)
			client := mockAws.NewMockSNSClient(mockCtl)
			tbl.fn(client, require)
			conf := config.AWS{
				SNS: &config.SNS{
					Topics: []config.Topic{
						tbl.topic,
					},
				},
			}
			p := newAWSSNS(ctx, client, &conf)
			for idx := range tbl.payload.Rows {
				msgId, err := p.PublishBinlog(ctx, tbl.payload.Event, tbl.payload.SendPayload(idx))
				require.NoError(err)
				require.NotEmpty(msgId)
			}
		})
	}
}

func TestAWSSNSWithColumn(t *testing.T) {
	var tables = []struct {
		name    string
		payload interfaces.Payload
		topic   config.Topic
		fn      func(client *mockAws.MockSNSClient, require *require.Assertions)
	}{
		{
			name: "Create Row",
			payload: interfaces.Payload{
				Event:  interfaces.Create,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(nil, map[string]interface{}{"key": "value", "aggregatetype": "test"}),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.ColumnType,
					Column: &config.TransformColumn{
						Table: config.TransformTable{
							Schema:    "public",
							TableName: "test",
						},
						ColumnName: "aggregatetype",
						Value:      "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"CREATE","schema":"public","table":"test","row":{"old_row":{},"new_row":{"aggregatetype":"test","key":"value"}}}`)
				}).Return(output, nil).Times(1)
			},
		},
		{
			name: "Create Rows",
			payload: interfaces.Payload{
				Event:  interfaces.Create,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(nil, map[string]interface{}{"key": "value", "aggregatetype": "test"}),
					interfaces.NewPayloadRow(nil, map[string]interface{}{"key": "value", "aggregatetype": "test"}),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.ColumnType,
					Column: &config.TransformColumn{
						Table: config.TransformTable{
							Schema:    "public",
							TableName: "test",
						},
						ColumnName: "aggregatetype",
						Value:      "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"CREATE","schema":"public","table":"test","row":{"old_row":{},"new_row":{"aggregatetype":"test","key":"value"}}}`)
				}).Return(output, nil).Times(2)
			},
		},
		{
			name: "Update Row",
			payload: interfaces.Payload{
				Event:  interfaces.Update,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(map[string]interface{}{"key": "oldValue", "aggregatetype": "test"}, map[string]interface{}{"key": "value", "aggregatetype": "test"}),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.ColumnType,
					Column: &config.TransformColumn{
						Table: config.TransformTable{
							Schema:    "public",
							TableName: "test",
						},
						ColumnName: "aggregatetype",
						Value:      "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"UPDATE","schema":"public","table":"test","row":{"old_row":{"aggregatetype":"test","key":"oldValue"},"new_row":{"aggregatetype":"test","key":"value"}}}`)
				}).Return(output, nil).Times(1)
			},
		},
		{
			name: "Delete Row",
			payload: interfaces.Payload{
				Event:  interfaces.Delete,
				Schema: "public",
				Table:  "test",
				Rows: []interfaces.PayloadRow{
					interfaces.NewPayloadRow(map[string]interface{}{"key": "oldValue", "aggregatetype": "test"}, nil),
				},
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.ColumnType,
					Column: &config.TransformColumn{
						Table: config.TransformTable{
							Schema:    "public",
							TableName: "test",
						},
						ColumnName: "aggregatetype",
						Value:      "test",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"event":"DELETE","schema":"public","table":"test","row":{"old_row":{"aggregatetype":"test","key":"oldValue"},"new_row":{}}}`)
				}).Return(output, nil).Times(1)
			},
		},
	}
	ctx := context.Background()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	for _, tbl := range tables {
		tbl := tbl
		t.Run(tbl.name, func(t *testing.T) {
			require := require.New(t)
			client := mockAws.NewMockSNSClient(mockCtl)
			tbl.fn(client, require)
			conf := config.AWS{
				SNS: &config.SNS{
					Topics: []config.Topic{
						tbl.topic,
					},
				},
			}
			p := newAWSSNS(ctx, client, &conf)
			for idx := range tbl.payload.Rows {
				msgId, err := p.PublishBinlog(ctx, tbl.payload.Event, tbl.payload.SendPayload(idx))
				require.NoError(err)
				require.NotEmpty(msgId)
			}
		})
	}
}

func TestAWSSNSWithOutbox(t *testing.T) {
	var tables = []struct {
		name   string
		outbox interfaces.Outbox
		topic  config.Topic
		fn     func(client *mockAws.MockSNSClient, require *require.Assertions)
		expect func(require *require.Assertions, msgId string, err error)
	}{
		{
			name: "AggregateType OK",
			outbox: interfaces.Outbox{
				AggregateId: "xxx",
				EventType:   "create",
				Payload:     `{"key":"value"}`,
			},
			topic: config.Topic{
				Transform: config.Transform{
					Type: config.OutboxPatternType,
					Outbox: &config.TransformOutbox{
						AggregateType: "topic-sns",
					},
				},
				TopicArn: "arn:aws:sns:ap-northeast-1:000000000000:test-sns",
			},
			fn: func(client *mockAws.MockSNSClient, require *require.Assertions) {
				output := &sns.PublishOutput{
					MessageId: aws.String(uuid.NewString()),
				}
				client.EXPECT().Publish(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) {
					require.NotNil(input)
					require.NotNil(input.Message)
					require.NotNil(input.TargetArn)
					require.NotNil(input.MessageGroupId)
					require.NotNil(input.MessageAttributes)
					require.Equal(*input.MessageGroupId, "xxx")
					require.Equal(input.MessageAttributes, map[string]types.MessageAttributeValue{
						"Event": {
							DataType:    aws.String("String"),
							StringValue: aws.String("create"),
						},
					})
					require.Equal(*input.TargetArn, "arn:aws:sns:ap-northeast-1:000000000000:test-sns")
					require.Equal(*input.Message, `{"key":"value"}`)
				}).Return(output, nil).Times(1)
			},
			expect: func(require *require.Assertions, msgId string, err error) {
				require.NotEmpty(msgId)
				require.NoError(err)
			},
		},
	}
	ctx := context.Background()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	for _, tbl := range tables {
		tbl := tbl
		t.Run(tbl.name, func(t *testing.T) {
			require := require.New(t)
			client := mockAws.NewMockSNSClient(mockCtl)
			tbl.fn(client, require)
			conf := config.AWS{
				SNS: &config.SNS{
					Topics: []config.Topic{
						tbl.topic,
					},
				},
			}
			p := newAWSSNS(ctx, client, &conf)
			msgId, err := p.PublishOutbox(ctx, tbl.topic.TopicArn, tbl.outbox)
			tbl.expect(require, msgId, err)
		})
	}
}
