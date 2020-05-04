// Code generated by goa v3.1.2, DO NOT EDIT.
//
// enduro HTTP client CLI support package
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	collectionc "github.com/artefactual-labs/enduro/internal/api/gen/http/collection/client"
	pipelinec "github.com/artefactual-labs/enduro/internal/api/gen/http/pipeline/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `collection (list|show|delete|cancel|retry|workflow|download|decide|bulk|bulk-status)
pipeline (list|show)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` collection list --name "Neque possimus et sunt." --original-id "Ex aut." --transfer-id "A129B534-C1FC-F09D-BF29-3DA5781E0ECB" --aip-id "4CCDE767-7648-444F-D09F-4B4FFE4EB36B" --pipeline-id "0C589E55-99C1-3ED8-809A-1463C91242B6" --earliest-created-time "1978-07-01T23:54:38Z" --latest-created-time "1989-06-26T13:51:36Z" --status "pending" --cursor "In cum et quia."` + "\n" +
		os.Args[0] + ` pipeline list --name "Est ut excepturi odio."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		collectionFlags = flag.NewFlagSet("collection", flag.ContinueOnError)

		collectionListFlags                   = flag.NewFlagSet("list", flag.ExitOnError)
		collectionListNameFlag                = collectionListFlags.String("name", "", "")
		collectionListOriginalIDFlag          = collectionListFlags.String("original-id", "", "")
		collectionListTransferIDFlag          = collectionListFlags.String("transfer-id", "", "")
		collectionListAipIDFlag               = collectionListFlags.String("aip-id", "", "")
		collectionListPipelineIDFlag          = collectionListFlags.String("pipeline-id", "", "")
		collectionListEarliestCreatedTimeFlag = collectionListFlags.String("earliest-created-time", "", "")
		collectionListLatestCreatedTimeFlag   = collectionListFlags.String("latest-created-time", "", "")
		collectionListStatusFlag              = collectionListFlags.String("status", "", "")
		collectionListCursorFlag              = collectionListFlags.String("cursor", "", "")

		collectionShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		collectionShowIDFlag = collectionShowFlags.String("id", "REQUIRED", "Identifier of collection to show")

		collectionDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		collectionDeleteIDFlag = collectionDeleteFlags.String("id", "REQUIRED", "Identifier of collection to delete")

		collectionCancelFlags  = flag.NewFlagSet("cancel", flag.ExitOnError)
		collectionCancelIDFlag = collectionCancelFlags.String("id", "REQUIRED", "Identifier of collection to remove")

		collectionRetryFlags  = flag.NewFlagSet("retry", flag.ExitOnError)
		collectionRetryIDFlag = collectionRetryFlags.String("id", "REQUIRED", "Identifier of collection to retry")

		collectionWorkflowFlags  = flag.NewFlagSet("workflow", flag.ExitOnError)
		collectionWorkflowIDFlag = collectionWorkflowFlags.String("id", "REQUIRED", "Identifier of collection to look up")

		collectionDownloadFlags  = flag.NewFlagSet("download", flag.ExitOnError)
		collectionDownloadIDFlag = collectionDownloadFlags.String("id", "REQUIRED", "Identifier of collection to look up")

		collectionDecideFlags    = flag.NewFlagSet("decide", flag.ExitOnError)
		collectionDecideBodyFlag = collectionDecideFlags.String("body", "REQUIRED", "")
		collectionDecideIDFlag   = collectionDecideFlags.String("id", "REQUIRED", "Identifier of collection to look up")

		collectionBulkFlags    = flag.NewFlagSet("bulk", flag.ExitOnError)
		collectionBulkBodyFlag = collectionBulkFlags.String("body", "REQUIRED", "")

		collectionBulkStatusFlags = flag.NewFlagSet("bulk-status", flag.ExitOnError)

		pipelineFlags = flag.NewFlagSet("pipeline", flag.ContinueOnError)

		pipelineListFlags    = flag.NewFlagSet("list", flag.ExitOnError)
		pipelineListNameFlag = pipelineListFlags.String("name", "", "")

		pipelineShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		pipelineShowIDFlag = pipelineShowFlags.String("id", "REQUIRED", "Identifier of pipeline to show")
	)
	collectionFlags.Usage = collectionUsage
	collectionListFlags.Usage = collectionListUsage
	collectionShowFlags.Usage = collectionShowUsage
	collectionDeleteFlags.Usage = collectionDeleteUsage
	collectionCancelFlags.Usage = collectionCancelUsage
	collectionRetryFlags.Usage = collectionRetryUsage
	collectionWorkflowFlags.Usage = collectionWorkflowUsage
	collectionDownloadFlags.Usage = collectionDownloadUsage
	collectionDecideFlags.Usage = collectionDecideUsage
	collectionBulkFlags.Usage = collectionBulkUsage
	collectionBulkStatusFlags.Usage = collectionBulkStatusUsage

	pipelineFlags.Usage = pipelineUsage
	pipelineListFlags.Usage = pipelineListUsage
	pipelineShowFlags.Usage = pipelineShowUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "collection":
			svcf = collectionFlags
		case "pipeline":
			svcf = pipelineFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "collection":
			switch epn {
			case "list":
				epf = collectionListFlags

			case "show":
				epf = collectionShowFlags

			case "delete":
				epf = collectionDeleteFlags

			case "cancel":
				epf = collectionCancelFlags

			case "retry":
				epf = collectionRetryFlags

			case "workflow":
				epf = collectionWorkflowFlags

			case "download":
				epf = collectionDownloadFlags

			case "decide":
				epf = collectionDecideFlags

			case "bulk":
				epf = collectionBulkFlags

			case "bulk-status":
				epf = collectionBulkStatusFlags

			}

		case "pipeline":
			switch epn {
			case "list":
				epf = pipelineListFlags

			case "show":
				epf = pipelineShowFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "collection":
			c := collectionc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = collectionc.BuildListPayload(*collectionListNameFlag, *collectionListOriginalIDFlag, *collectionListTransferIDFlag, *collectionListAipIDFlag, *collectionListPipelineIDFlag, *collectionListEarliestCreatedTimeFlag, *collectionListLatestCreatedTimeFlag, *collectionListStatusFlag, *collectionListCursorFlag)
			case "show":
				endpoint = c.Show()
				data, err = collectionc.BuildShowPayload(*collectionShowIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = collectionc.BuildDeletePayload(*collectionDeleteIDFlag)
			case "cancel":
				endpoint = c.Cancel()
				data, err = collectionc.BuildCancelPayload(*collectionCancelIDFlag)
			case "retry":
				endpoint = c.Retry()
				data, err = collectionc.BuildRetryPayload(*collectionRetryIDFlag)
			case "workflow":
				endpoint = c.Workflow()
				data, err = collectionc.BuildWorkflowPayload(*collectionWorkflowIDFlag)
			case "download":
				endpoint = c.Download()
				data, err = collectionc.BuildDownloadPayload(*collectionDownloadIDFlag)
			case "decide":
				endpoint = c.Decide()
				data, err = collectionc.BuildDecidePayload(*collectionDecideBodyFlag, *collectionDecideIDFlag)
			case "bulk":
				endpoint = c.Bulk()
				data, err = collectionc.BuildBulkPayload(*collectionBulkBodyFlag)
			case "bulk-status":
				endpoint = c.BulkStatus()
				data = nil
			}
		case "pipeline":
			c := pipelinec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = pipelinec.BuildListPayload(*pipelineListNameFlag)
			case "show":
				endpoint = c.Show()
				data, err = pipelinec.BuildShowPayload(*pipelineShowIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// collectionUsage displays the usage of the collection command and its
// subcommands.
func collectionUsage() {
	fmt.Fprintf(os.Stderr, `The collection service manages packages being transferred to Archivematica.
Usage:
    %s [globalflags] collection COMMAND [flags]

COMMAND:
    list: List all stored collections
    show: Show collection by ID
    delete: Delete collection by ID
    cancel: Cancel collection processing by ID
    retry: Retry collection processing by ID
    workflow: Retrieve workflow status by ID
    download: Download collection by ID
    decide: Make decision for a pending collection by ID
    bulk: Bulk operations (retry, cancel...).
    bulk-status: Retrieve status of current bulk operation.

Additional help:
    %s collection COMMAND --help
`, os.Args[0], os.Args[0])
}
func collectionListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection list -name STRING -original-id STRING -transfer-id STRING -aip-id STRING -pipeline-id STRING -earliest-created-time STRING -latest-created-time STRING -status STRING -cursor STRING

List all stored collections
    -name STRING: 
    -original-id STRING: 
    -transfer-id STRING: 
    -aip-id STRING: 
    -pipeline-id STRING: 
    -earliest-created-time STRING: 
    -latest-created-time STRING: 
    -status STRING: 
    -cursor STRING: 

Example:
    `+os.Args[0]+` collection list --name "Neque possimus et sunt." --original-id "Ex aut." --transfer-id "A129B534-C1FC-F09D-BF29-3DA5781E0ECB" --aip-id "4CCDE767-7648-444F-D09F-4B4FFE4EB36B" --pipeline-id "0C589E55-99C1-3ED8-809A-1463C91242B6" --earliest-created-time "1978-07-01T23:54:38Z" --latest-created-time "1989-06-26T13:51:36Z" --status "pending" --cursor "In cum et quia."
`, os.Args[0])
}

func collectionShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection show -id UINT

Show collection by ID
    -id UINT: Identifier of collection to show

Example:
    `+os.Args[0]+` collection show --id 8051351490510170373
`, os.Args[0])
}

func collectionDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection delete -id UINT

Delete collection by ID
    -id UINT: Identifier of collection to delete

Example:
    `+os.Args[0]+` collection delete --id 2802088767362527743
`, os.Args[0])
}

func collectionCancelUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection cancel -id UINT

Cancel collection processing by ID
    -id UINT: Identifier of collection to remove

Example:
    `+os.Args[0]+` collection cancel --id 9153290315792347469
`, os.Args[0])
}

func collectionRetryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection retry -id UINT

Retry collection processing by ID
    -id UINT: Identifier of collection to retry

Example:
    `+os.Args[0]+` collection retry --id 16485831467066386148
`, os.Args[0])
}

func collectionWorkflowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection workflow -id UINT

Retrieve workflow status by ID
    -id UINT: Identifier of collection to look up

Example:
    `+os.Args[0]+` collection workflow --id 6600600312516747591
`, os.Args[0])
}

func collectionDownloadUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection download -id UINT

Download collection by ID
    -id UINT: Identifier of collection to look up

Example:
    `+os.Args[0]+` collection download --id 8609244776313237956
`, os.Args[0])
}

func collectionDecideUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection decide -body JSON -id UINT

Make decision for a pending collection by ID
    -body JSON: 
    -id UINT: Identifier of collection to look up

Example:
    `+os.Args[0]+` collection decide --body '{
      "option": "Recusandae laudantium quidem consequatur ducimus excepturi perferendis."
   }' --id 14968097700807042938
`, os.Args[0])
}

func collectionBulkUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection bulk -body JSON

Bulk operations (retry, cancel...).
    -body JSON: 

Example:
    `+os.Args[0]+` collection bulk --body '{
      "operation": "cancel",
      "size": 4963086957530573892,
      "status": "unknown"
   }'
`, os.Args[0])
}

func collectionBulkStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection bulk-status

Retrieve status of current bulk operation.

Example:
    `+os.Args[0]+` collection bulk-status
`, os.Args[0])
}

// pipelineUsage displays the usage of the pipeline command and its subcommands.
func pipelineUsage() {
	fmt.Fprintf(os.Stderr, `The pipeline service manages Archivematica pipelines.
Usage:
    %s [globalflags] pipeline COMMAND [flags]

COMMAND:
    list: List all known pipelines
    show: Show pipeline by ID

Additional help:
    %s pipeline COMMAND --help
`, os.Args[0], os.Args[0])
}
func pipelineListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] pipeline list -name STRING

List all known pipelines
    -name STRING: 

Example:
    `+os.Args[0]+` pipeline list --name "Est ut excepturi odio."
`, os.Args[0])
}

func pipelineShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] pipeline show -id STRING

Show pipeline by ID
    -id STRING: Identifier of pipeline to show

Example:
    `+os.Args[0]+` pipeline show --id "5D93A99A-535C-03BB-87BD-865144549E25"
`, os.Args[0])
}
