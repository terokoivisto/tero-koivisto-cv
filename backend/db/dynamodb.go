package db

import (
	"backend/models"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"os"
	"time"
)

type TableConfig struct {
	TableName string
	DBClient  *dynamodb.Client
}

func checkRequiredEnvVariable(key string) {
	_, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Required environment variable %s was not not set\n", key)
	}
}

func setupConfig() aws.Config {
	checkRequiredEnvVariable("DDB_ENDPOINT_URL")
	checkRequiredEnvVariable("AWS_ACCESS_KEY_ID")
	checkRequiredEnvVariable("AWS_SECRET_ACCESS_KEY")
	checkRequiredEnvVariable("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	return cfg
}

func (tc TableConfig) doesTableExist() (bool, error) {
	_, err := tc.DBClient.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tc.TableName)})

	if err != nil {
		var resNotFoundEx *types.ResourceNotFoundException
		if errors.As(err, &resNotFoundEx) {
			return false, nil
		}

		return false, err
	}

	return true, err
}

func (tc TableConfig) createTable() (*types.TableDescription, error) {
	table, err := tc.DBClient.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String(tc.TableName),
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("name"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("name"),
			KeyType:       types.KeyTypeHash,
		}},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(4),
			WriteCapacityUnits: aws.Int64(2),
		},
	})

	if err != nil {
		log.Fatalf("Unable to create the table %s, due to %v\n", tc.TableName, err)
	}

	waiter := dynamodb.NewTableExistsWaiter(tc.DBClient)

	err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String(tc.TableName),
	}, 2*time.Minute)

	if err != nil {
		log.Fatalf("Timeout while waiting for the table %s to be created, due to %v\n", tc.TableName, err)
	}

	return table.TableDescription, err
}

func (tc TableConfig) populateDB() error {
	skills := []models.Skill{
		{
			Icon: "vscode-icons:file-type-python",
			Name: "Python",
			Usages: []string{
				"Django backend",
				"File processing",
				"AWS Lambdas",
				"Ad-hoc scripts",
				"Data analytics",
				"Much more...",
			},
		},
		{
			Icon: "vscode-icons:file-type-typescript-official",
			Name: "Typescript",
			Usages: []string{
				"React",
				"Svelte",
				"NodeJS",
				"AWS Lambdas",
				"FiveM",
				"Angular",
			},
		},
		{
			Icon: "mdi:database",
			Name: "Databases",
			Usages: []string{
				"PostgreSQL",
				"Designing database models/designs/architectures",
				"Dynamo DB",
				"MS SQL Server",
				"MariaDB",
				"MySQL",
			},
		},
		{
			Icon: "skill-icons:aws-dark",
			Name: "AWS",
			Usages: []string{
				"Lambda",
				"S3",
				"DynamoDB",
				"Aurora RDS",
				"Step functions",
				"AppSync",
				"ECR",
				"ECS",
				"Cloudformation",
				"CloudFront",
				"Load balancing",
				"Much more...",
			},
		},
		{
			Icon: "mingcute:hat-fill",
			Name: "Other",
			Usages: []string{
				"LUA",
				"Docker",
				"CI/CD",
				"Automated testing",
				"C#",
				"Java",
				"C/C++",
			},
		},
	}
	experience := []models.Experience{
		{
			Company: "Noiseless Acoustics",
			Title:   "Senior Fullstack Developer",
			From:    "October 2020",
			To:      "Present",
			Summary: "Fullstack development (React + Django + PostgreSQL + lots of AWS microservices), DevOps, and a lot more",
		},
		{
			Company: "Analyse2",
			Title:   "Fullstack Developer",
			From:    "October 2018",
			To:      "October 2020",
			Summary: "Fullstack development (Angular + asp.NET Core + MS SQL Server), architecture, database design, multiple external projects, DevOps, data analytics",
		},
		{
			Company: "QuattroFolia",
			Title:   "Junior Software Developer",
			From:    "March 2017",
			To:      "October 2018",
			Summary: "Test engineering for Web and Android, Frontend development with AngularJS",
		},
		{
			Company: "Suomen Merivoimat",
			Title:   "Military Service",
			From:    "January 2015",
			To:      "June 2015",
			Summary: "Fulfilled my military service duty at Upinniemi",
		},
		{
			Company: "Arc Technology Oy",
			Title:   "Software Tester",
			From:    "October 2014",
			To:      "December 2014",
			Summary: "Manual testing of HRM & HRD software",
		},
	}

	cv := models.CV{
		Name:       "Tero Koivisto",
		AboutMe:    "I'm a 28-year-old Senior Fullstack Developer who excels at adapting to challenges of all sorts and scales by breaking down larger problems into smaller easily comprehensible sub-problems, to deliver the solution that satisfies the requirements. I'm a quick learner and I'm always ready to expand beyond my current knowledge and skillset, if need be. Beyond writing code I always try to understand the underlying reason why we want to do something, and combine the business requirements to the current technical solution in a robust manner. I'm very used to speaking English with colleagues.",
		PersonalMe: "Outside my worklife, I spend time listening to music, playing various games (FPS and survival are my favorites), writing code for GTA V roleplaying server and during the summer I try my best to play some golf. I'm also a car enthusiast (yes, the ones using ancient technology that go brrrr), a semi-hardcore coffee lover and I like cooking all kinds of foods.",
		Title:      "Senior Fullstack Developer",
		Location:   "Espoo, Finland",
		Skills:     skills,
		Experience: experience,
	}

	item, err := attributevalue.MarshalMap(cv)

	if err != nil {
		panic(err)
	}

	_, err = tc.DBClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tc.TableName),
		Item:      item,
	})

	if err != nil {
		log.Printf("Unable to add the item, due to %v\n", err)
	}

	return err
}

func (tc TableConfig) CV(n string) (models.CV, error) {
	cv := models.CV{Name: n}

	name, err := attributevalue.Marshal(cv.Name)
	if err != nil {
		panic(err)
	}

	resp, err := tc.DBClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tc.TableName),
		Key:       map[string]types.AttributeValue{"name": name},
	})

	if err != nil {
		log.Printf("Could not GET item from DDB, due to %v\n", err)
		return cv, err
	}

	err = attributevalue.UnmarshalMap(resp.Item, &cv)
	if err != nil {
		log.Printf("Something went wrong while unmarshaling the data: %v\n", err)
	}

	return cv, err
}

func SetupDB() TableConfig {
	cfg := setupConfig()

	// Makes the DDB client to use locally hosted DynamoDB, instead of actual AWS
	db := dynamodb.NewFromConfig(cfg, func(options *dynamodb.Options) {
		ep := os.Getenv("DDB_ENDPOINT_URL")
		log.Println(ep)
		options.BaseEndpoint = aws.String(ep)
	})

	tableName := "CVData"

	tc := TableConfig{TableName: tableName, DBClient: db}

	exists, err := tc.doesTableExist()

	if err != nil {
		log.Fatalf("Unable to check if table exists, due to %v\n", err)
	}

	if !exists {
		// Creates the table on local DDB if it does not exist already
		_, err = tc.createTable()
		if err != nil {
			panic(err)
		}
	}

	// Populates the database with my data
	err = tc.populateDB()
	if err != nil {
		panic(err)
	}

	log.Println("Dynamo DB was setup successfully")

	return tc
}
