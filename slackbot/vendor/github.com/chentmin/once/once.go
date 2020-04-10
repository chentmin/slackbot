package once

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

var (
	DuplicateErr = errors.New("Duplicate id")
)

type manager struct {
	dynamoTableName string

	// TODO auth
}

func New(dynamoTableName string, ops ...option) *manager {
	result := &manager{
		dynamoTableName: dynamoTableName,
	}

	for _, o := range ops {
		o(result)
	}

	return result
}

type option func(m *manager)

func (m *manager) Ensure(id string) error {
	svc := dynamodb.New(session.Must(session.NewSession()))

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName:           aws.String(m.dynamoTableName),
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return DuplicateErr
			default:
				return errors.Wrap(err, "unhandled dynamo err")
			}
		}

		return errors.Wrap(err, "Unknown error")
	} else {
		return nil
	}
}
