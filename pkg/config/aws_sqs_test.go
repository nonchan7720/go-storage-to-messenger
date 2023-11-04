package config

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestAWSSQSConfig(t *testing.T) {
	data := `endpoint: http://localstack:4566
queues:
  - queueName: test-queue
    queueUrl: http://localstack:4566/000000000000/test-queue
    messageGroupIdTemplate: id-{{id}}
    transform:
      table:
        schema: public
        tableName: tbl1
      column:
        schema: public
        tableName: tbl1
        columnName: column1
        value: test
`
	var config SQS
	err := loadYaml(bytes.NewBufferString(data), &config)
	require.NoError(t, err)
	require.Equal(t, config.Endpoint, "http://localstack:4566")
	require.Len(t, config.Queues, 1)
	require.Equal(t, config.Queues[0].Transform.Table.Schema, "public")
	require.Equal(t, config.Queues[0].Transform.Table.TableName, "tbl1")
	require.Equal(t, config.Queues[0].Transform.Column.Table.Schema, "public")
	require.Equal(t, config.Queues[0].Transform.Column.Table.TableName, "tbl1")
	require.Equal(t, config.Queues[0].Transform.Column.ColumnName, "column1")
	require.Equal(t, config.Queues[0].Transform.Column.Value, "test")
	require.Equal(t, config.Queues[0].QueueName, "test-queue")
	require.Equal(t, config.Queues[0].QueueUrl, "http://localstack:4566/000000000000/test-queue")
	require.Equal(t, config.Queues[0].MessageGroupIdTemplate, "id-{{id}}")
	require.Equal(t, config.Queues[0].TemplateType, FastTemplate)
}

func TestAWSSQSConfigWithGoTemplate(t *testing.T) {
	data := `endpoint: http://localstack:4566
queues:
  - queueName: test-queue
    queueUrl: http://localstack:4566/000000000000/test-queue
    messageGroupIdTemplate: id-{{id}}
    templateType: Go
    transform:
      table:
        schema: public
        tableName: tbl1
      column:
        schema: public
        tableName: tbl1
        columnName: column1
        value: test
`
	var config SQS
	err := yaml.Unmarshal([]byte(data), &config)
	require.NoError(t, err)
	require.Equal(t, config.Endpoint, "http://localstack:4566")
	require.Len(t, config.Queues, 1)
	require.Equal(t, config.Queues[0].Transform.Table.Schema, "public")
	require.Equal(t, config.Queues[0].Transform.Table.TableName, "tbl1")
	require.Equal(t, config.Queues[0].Transform.Column.Table.Schema, "public")
	require.Equal(t, config.Queues[0].Transform.Column.Table.TableName, "tbl1")
	require.Equal(t, config.Queues[0].Transform.Column.ColumnName, "column1")
	require.Equal(t, config.Queues[0].Transform.Column.Value, "test")
	require.Equal(t, config.Queues[0].QueueName, "test-queue")
	require.Equal(t, config.Queues[0].QueueUrl, "http://localstack:4566/000000000000/test-queue")
	require.Equal(t, config.Queues[0].MessageGroupIdTemplate, "id-{{id}}")
	require.Equal(t, config.Queues[0].TemplateType, GoTemplate)
}
