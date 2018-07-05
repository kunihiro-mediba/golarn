package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDynamoDB() *dynamodb.DynamoDB {
	// Create session
	conf := &aws.Config{
		Endpoint: aws.String("localhost:8000"),
		Region: aws.String("localhost:8000"),
	}
	sess, err := session.NewSession(conf)
	if err != nil {
		fmt.Println("Create session failed", err)
		return nil
	}
	// Create DynamoDB
	return dynamodb.New(sess)
}

func main() {
	fmt.Println("start dynamodb test")

	db := GetDynamoDB()
	if db == nil {
		fmt.Println("Get DynamoDB failed")
		return
	}



	// Create table
	cti := &dynamodb.CreateTableInput{
		TableName: aws.String("user"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("user_id"), // フィールド名
				KeyType:       aws.String("HASH"),    // HASH=ハッシュキー, RANGE=レンジキー
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("user_id"), // フィールド名
				AttributeType: aws.String("N"),       // フィールド型 N=number, S=string, B=bool で型を指定
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1), // 読み込みスループット
			WriteCapacityUnits: aws.Int64(1), // 書き込みスループット
		},
	}
	cto, err := db.CreateTable(cti)
	if err != nil {
		fmt.Println("Create table failed", err)
		return
	}
	fmt.Println("Create table success", cto)
}
