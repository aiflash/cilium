// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes your Spot Fleet requests.
//
// Spot Fleet requests are deleted 48 hours after they are canceled and their
// instances are terminated.
func (c *Client) DescribeSpotFleetRequests(ctx context.Context, params *DescribeSpotFleetRequestsInput, optFns ...func(*Options)) (*DescribeSpotFleetRequestsOutput, error) {
	if params == nil {
		params = &DescribeSpotFleetRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpotFleetRequests", params, optFns, c.addOperationDescribeSpotFleetRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpotFleetRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotFleetRequests.
type DescribeSpotFleetRequestsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	MaxResults *int32

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// The IDs of the Spot Fleet requests.
	SpotFleetRequestIds []string

	noSmithyDocumentSerde
}

// Contains the output of DescribeSpotFleetRequests.
type DescribeSpotFleetRequestsOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about the configuration of your Spot Fleet.
	SpotFleetRequestConfigs []types.SpotFleetRequestConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSpotFleetRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotFleetRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotFleetRequests{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSpotFleetRequests"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotFleetRequests(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeSpotFleetRequestsPaginatorOptions is the paginator options for
// DescribeSpotFleetRequests
type DescribeSpotFleetRequestsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSpotFleetRequestsPaginator is a paginator for DescribeSpotFleetRequests
type DescribeSpotFleetRequestsPaginator struct {
	options   DescribeSpotFleetRequestsPaginatorOptions
	client    DescribeSpotFleetRequestsAPIClient
	params    *DescribeSpotFleetRequestsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSpotFleetRequestsPaginator returns a new
// DescribeSpotFleetRequestsPaginator
func NewDescribeSpotFleetRequestsPaginator(client DescribeSpotFleetRequestsAPIClient, params *DescribeSpotFleetRequestsInput, optFns ...func(*DescribeSpotFleetRequestsPaginatorOptions)) *DescribeSpotFleetRequestsPaginator {
	if params == nil {
		params = &DescribeSpotFleetRequestsInput{}
	}

	options := DescribeSpotFleetRequestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSpotFleetRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSpotFleetRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSpotFleetRequests page.
func (p *DescribeSpotFleetRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSpotFleetRequestsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeSpotFleetRequests(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeSpotFleetRequestsAPIClient is a client that implements the
// DescribeSpotFleetRequests operation.
type DescribeSpotFleetRequestsAPIClient interface {
	DescribeSpotFleetRequests(context.Context, *DescribeSpotFleetRequestsInput, ...func(*Options)) (*DescribeSpotFleetRequestsOutput, error)
}

var _ DescribeSpotFleetRequestsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeSpotFleetRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSpotFleetRequests",
	}
}
