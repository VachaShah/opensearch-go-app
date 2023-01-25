package main

import (
	"context"
	// "net/http"
	// "io"
	"log"
	// "os"
	// "context"
	// "crypto/tls"
	// "fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
	opensearchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"

	// "net/http"
	"strings"
)

const IndexName = "go-test-index2"

//   const endpoint = "" // e.g. https://opensearch-domain.region.com managed service endpoint
const endpoint = "" //serverless endpoint

func main() {

	// Normal invocation of the go client

	// // Initialize the client with SSL/TLS enabled.
	// client, err := opensearch.NewClient(opensearch.Config{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // For testing only. Use certificate for validation.
	// 	},
	// 	Addresses: []string{"https://localhost:9200"},
	// 	Username:  "admin", // For testing only. Don't store credentials in code.
	// 	Password:  "admin",
	// })
	// if err != nil {
	// 	fmt.Println("cannot initialize", err)
	// 	os.Exit(1)
	// }

	// // Print OpenSearch version information on console.
	// fmt.Println(client.Info())

	// // Define index mapping.
	// mapping := strings.NewReader(`{
	//  'settings': {
	//    'index': {
	//         'number_of_shards': 4
	//         }
	//       }
	//  }`)

	// // Create an index with non-default settings.
	// createIndex := opensearchapi.IndicesCreateRequest{
	// 	Index: IndexName,
	// 	Body:  mapping,
	// }
	// createIndexResponse, err := createIndex.Do(context.Background(), client)
	// if err != nil {
	// 	fmt.Println("failed to create index ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(createIndexResponse)

	// // Add a document to the index.
	// document := strings.NewReader(`{
	//     "title": "Moneyball",
	//     "director": "Bennett Miller",
	//     "year": "2011"
	// }`)

	// docId := "1"
	// req := opensearchapi.IndexRequest{
	// 	Index:      IndexName,
	// 	DocumentID: docId,
	// 	Body:       document,
	// }
	// insertResponse, err := req.Do(context.Background(), client)
	// if err != nil {
	// 	fmt.Println("failed to insert document ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(insertResponse)

	// // Search for the document.
	// content := strings.NewReader(`{
	//    "size": 5,
	//    "query": {
	//        "multi_match": {
	//        "query": "miller",
	//        "fields": ["title^2", "director"]
	//        }
	//   }
	// }`)

	// search := opensearchapi.SearchRequest{
	// 	Body: content,
	// }

	// searchResponse, err := search.Do(context.Background(), client)
	// if err != nil {
	// 	fmt.Println("failed to search document ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(searchResponse)

	// // Delete the document.
	// delete := opensearchapi.DeleteRequest{
	// 	Index:      IndexName,
	// 	DocumentID: docId,
	// }

	// deleteResponse, err := delete.Do(context.Background(), client)
	// if err != nil {
	// 	fmt.Println("failed to delete document ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("deleting document")
	// fmt.Println(deleteResponse)

	// // Delete previously created index.
	// deleteIndex := opensearchapi.IndicesDeleteRequest{
	// 	Index: []string{IndexName},
	// }

	// deleteIndexResponse, err := deleteIndex.Do(context.Background(), client)
	// if err != nil {
	// 	fmt.Println("failed to delete index ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("deleting index", deleteIndexResponse)

	// // Invocation for managed service with sigv4

	// ctx := context.Background()

	// // cfg, err := config.LoadDefaultConfig(ctx)
	// // if err != nil {
	// // 	log.Fatal(err) // Do not log.fatal in a production ready app.
	// // }

	// awsCfg, err := config.LoadDefaultConfig(ctx,
	// 	config.WithRegion("us-west-2"),
	// 	config.WithCredentialsProvider(
	// 		getCredentialProvider("", "", ""),
	// 	),
	// )

	// // Create an AWS request Signer and load AWS configuration using default config folder or env vars.
	// // See https://docs.aws.amazon.com/opensearch-service/latest/developerguide/request-signing.html#request-signing-go
	// signer, err := requestsigner.NewSigner(awsCfg)
	// if err != nil {
	// 	log.Fatal(err) // Do not log.fatal in a production ready app.
	// }

	// // Create an opensearch client and use the request-signer
	// client, err := opensearch.NewClient(opensearch.Config{
	// 	Addresses: []string{endpoint},
	// 	Signer:    signer,
	// })
	// if err != nil {
	// 	log.Fatal("client creation err", err)
	// }

	// ping := opensearchapi.PingRequest{}

	// resp, err := ping.Do(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// if resp.IsError() {
	// 	log.Println("ping response status ", resp.Status())

	// 	respBody, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal("response body read err", err)
	// 	}

	// 	log.Fatal("ping resp body", respBody)
	// }

	// log.Println("PING OK")

	// // Define index mapping.
	// mapping := strings.NewReader(`{
	//  "settings": {
	//    "index": {
	//         "number_of_shards": 4
	//         }
	//       }
	//  }`)

	// // Create an index with non-default settings.
	// createIndex := opensearchapi.IndicesCreateRequest{
	// 	Index: IndexName,
	// 	Body:  mapping,
	// }
	// createIndexResponse, err := createIndex.Do(context.Background(), client)
	// if err != nil {
	// 	log.Println("failed to create index ", err)
	// 	log.Fatal("create response body read err", err)
	// }
	// log.Println(createIndexResponse)

	// // Add a document to the index.
	// document := strings.NewReader(`{
	//     "title": "Moneyball",
	//     "director": "Bennett Miller",
	//     "year": "2011"
	// }`)

	// docId := "1"
	// req := opensearchapi.IndexRequest{
	// 	Index:      IndexName,
	// 	DocumentID: docId,
	// 	Body:       document,
	// }
	// insertResponse, err := req.Do(context.Background(), client)
	// if err != nil {
	// 	log.Println("failed to insert document ", err)
	// 	log.Fatal("index response body read err", err)
	// }
	// log.Println(insertResponse)

	// // Search for the document.
	// content := strings.NewReader(`{
	//    "size": 5,
	//    "query": {
	//        "multi_match": {
	//        "query": "miller",
	//        "fields": ["title^2", "director"]
	//        }
	//   }
	// }`)

	// search := opensearchapi.SearchRequest{
	// 	Body: content,
	// }

	// searchResponse, err := search.Do(context.Background(), client)
	// if err != nil {
	// 	log.Println("failed to search document ", err)
	// 	log.Fatal("search response body read err", err)
	// }
	// log.Println(searchResponse)

	// // Delete the document.
	// delete := opensearchapi.DeleteRequest{
	// 	Index:      IndexName,
	// 	DocumentID: docId,
	// }

	// deleteResponse, err := delete.Do(context.Background(), client)
	// if err != nil {
	// 	log.Println("failed to delete document ", err)
	// 	log.Fatal("delete response body read err", err)
	// }
	// log.Println("deleting document")
	// log.Println(deleteResponse)

	// // Delete previously created index.
	// deleteIndex := opensearchapi.IndicesDeleteRequest{
	// 	Index: []string{IndexName},
	// }

	// deleteIndexResponse, err := deleteIndex.Do(context.Background(), client)
	// if err != nil {
	// 	log.Println("failed to delete index ", err)
	// 	log.Fatal("delete index response body read err", err)
	// }
	// log.Println("deleting index", deleteIndexResponse)

	// Invocation for serverless with service name aoss
	ctx := context.Background()

	// cfg, err := config.LoadDefaultConfig(ctx)
	// if err != nil {
	// 	log.Fatal(err) // Do not log.fatal in a production ready app.
	// }

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-west-2"),
		config.WithCredentialsProvider(
			getCredentialProvider("", "", ""),
		),
	)

	// Create an AWS request Signer and load AWS configuration using default config folder or env vars.
	// See https://docs.aws.amazon.com/opensearch-service/latest/developerguide/request-signing.html#request-signing-go
	signer, err := requestsigner.NewSignerWithService(awsCfg, "aoss")
	if err != nil {
		log.Fatal(err) // Do not log.fatal in a production ready app.
	}

	// Create an opensearch client and use the request-signer
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{endpoint},
		Signer:    signer,
	})
	if err != nil {
		log.Fatal("client creation err", err)
	}

	// ping := opensearchapi.PingRequest{}

	// resp, err := ping.Do(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// if resp.IsError() {
	// 	log.Println("ping response status ", resp.Status())

	// 	respBody, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal("response body read err", err)
	// 	}

	// 	log.Fatal("ping resp body", respBody)
	// }

	// log.Println("PING OK")

	// Define index mapping.
	mapping := strings.NewReader(`{
	 "settings": {
	   "index": {
	        "number_of_shards": 4
	        }
	      }
	 }`)
	

	// Create an index with non-default settings.
	createIndex := opensearchapi.IndicesCreateRequest{
		Index: IndexName,
		Body:  mapping,
		// Header: http.Header{"X-Amz-Content-Sha256": []string{"UNSIGNED-PAYLOAD"}},
	}
	// delete(createIndex.Header, "Content-Length")
	log.Println("Headers ", createIndex.Header)
	createIndexResponse, err := createIndex.Do(context.Background(), client)
	if err != nil {
		log.Println("Error ", err.Error())
		log.Println("failed to create index ", err)
		log.Fatal("create response body read err", err)
	}
	log.Println(createIndexResponse)

	// Delete previously created index.
	deleteIndex := opensearchapi.IndicesDeleteRequest{
		Index: []string{IndexName},
	}

	deleteIndexResponse, err := deleteIndex.Do(context.Background(), client)
	if err != nil {
		log.Println("failed to delete index ", err)
		log.Fatal("delete index response body read err", err)
	}
	log.Println("deleting index", deleteIndexResponse)

}

func getCredentialProvider(accessKey, secretAccessKey, token string) aws.CredentialsProviderFunc {
	return func(ctx context.Context) (aws.Credentials, error) {
		c := &aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretAccessKey,
			SessionToken:    token,
		}
		return *c, nil
	}
}
