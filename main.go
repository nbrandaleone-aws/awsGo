package main

import (
  "fmt"
  "os"
  "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
  "github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

/*
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
*/

/* AWS Regions as specified in the Go SDK

    ApNortheast1RegionID = "ap-northeast-1" // Asia Pacific (Tokyo).
    ApNortheast2RegionID = "ap-northeast-2" // Asia Pacific (Seoul).
    ApSouth1RegionID     = "ap-south-1"     // Asia Pacific (Mumbai).
    ApSoutheast1RegionID = "ap-southeast-1" // Asia Pacific (Singapore).
    ApSoutheast2RegionID = "ap-southeast-2" // Asia Pacific (Sydney).
    CaCentral1RegionID   = "ca-central-1"   // Canada (Central).
    EuCentral1RegionID   = "eu-central-1"   // EU (Frankfurt).
    EuWest1RegionID      = "eu-west-1"      // EU (Ireland).
    EuWest2RegionID      = "eu-west-2"      // EU (London).
    EuWest3RegionID      = "eu-west-3"      // EU (Paris).
    SaEast1RegionID      = "sa-east-1"      // South America (Sao Paulo).
    UsEast1RegionID      = "us-east-1"      // US East (N. Virginia).
    UsEast2RegionID      = "us-east-2"      // US East (Ohio).
    UsWest1RegionID      = "us-west-1"      // US West (N. California).
    UsWest2RegionID      = "us-west-2"      // US West (Oregon).
*/

func LambdaHandler() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

  // Get the region and cluster name from env variables
  var region, cluster string 
  if r, ok := os.LookupEnv("region"); ok {
    region = r
  } else {
    region = os.Getenv("AWS_REGION")
  }
  fmt.Println(region)
  if c, ok := os.LookupEnv("cluster"); ok {
    cluster = c
  } else {
    panic("unable to determine which EKS cluster to use")
  }
  fmt.Println(cluster)
  
	// Set the AWS Region that the service clients should use
  // See https://docs.aws.amazon.com/sdk-for-go/api/aws/endpoints/#pkg-constants
	cfg.Region = endpoints.UsWest2RegionID
  svc := eks.New(cfg)
  	input := &eks.DescribeClusterInput{
  		Name: aws.String(cluster),
  	}

  req := svc.DescribeClusterRequest(input)
  result, err := req.Send()
  if err != nil {
  	if aerr, ok := err.(awserr.Error); ok {
  			switch aerr.Code() {
  			case eks.ErrCodeResourceNotFoundException:
  				fmt.Println(eks.ErrCodeResourceNotFoundException, aerr.Error())
  			case eks.ErrCodeClientException:
  				fmt.Println(eks.ErrCodeClientException, aerr.Error())
  			case eks.ErrCodeServerException:
  				fmt.Println(eks.ErrCodeServerException, aerr.Error())
  			case eks.ErrCodeServiceUnavailableException:
  				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
  			default:
  				fmt.Println(aerr.Error())
  			} // switch
  	} else {
  		// Print the error, cast err to awserr.Error to get the Code and
  		// Message from an error.
  		fmt.Println(err.Error())
  	}
  return
  }

  //log.Println(result)
  fmt.Print("\n=========================================\n")
  fmt.Println(*result.Cluster.Name)
  fmt.Println(*result.Cluster.Version)
  fmt.Println(*result.Cluster.Arn)
  fmt.Println(*result.Cluster.CertificateAuthority.Data)
  fmt.Println(*result.Cluster.Endpoint)
  fmt.Println(*result.Cluster.ResourcesVpcConfig)
  fmt.Println(result.Cluster.Status)
}

// TODO: Add context for future acceptance of triggers
func main() {
  lambda.Start(LambdaHandler)
}