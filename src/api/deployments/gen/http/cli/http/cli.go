// Code generated by goa v3.8.5, DO NOT EDIT.
//
// http HTTP client CLI support package
//
// Command:
// $ goa gen github.com/acul009/control-panel-2/src/api/deployments/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	deploymentsc "github.com/acul009/control-panel-2/src/api/deployments/gen/http/deployments/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `deployments (upsert|list|get|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` deployments upsert --body '{
      "containers": [
         {
            "image": "Quisquam sed.",
            "name": "Dignissimos et nemo sunt commodi iste.",
            "parameters": [
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               },
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               }
            ],
            "ports": [
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               }
            ],
            "services": [
               "Rem fugiat est.",
               "Quas enim magnam libero consectetur qui nesciunt.",
               "Quam libero."
            ]
         },
         {
            "image": "Quisquam sed.",
            "name": "Dignissimos et nemo sunt commodi iste.",
            "parameters": [
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               },
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               }
            ],
            "ports": [
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               }
            ],
            "services": [
               "Rem fugiat est.",
               "Quas enim magnam libero consectetur qui nesciunt.",
               "Quam libero."
            ]
         }
      ],
      "name": "Necessitatibus et autem placeat et.",
      "parameters": [
         {
            "name": "m",
            "value": "Voluptatem est eius alias est sit."
         },
         {
            "name": "m",
            "value": "Voluptatem est eius alias est sit."
         },
         {
            "name": "m",
            "value": "Voluptatem est eius alias est sit."
         }
      ]
   }'` + "\n" +
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
		deploymentsFlags = flag.NewFlagSet("deployments", flag.ContinueOnError)

		deploymentsUpsertFlags    = flag.NewFlagSet("upsert", flag.ExitOnError)
		deploymentsUpsertBodyFlag = deploymentsUpsertFlags.String("body", "REQUIRED", "")

		deploymentsListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		deploymentsGetFlags = flag.NewFlagSet("get", flag.ExitOnError)
		deploymentsGetPFlag = deploymentsGetFlags.String("p", "REQUIRED", "Name of the Deployment")

		deploymentsDeleteFlags = flag.NewFlagSet("delete", flag.ExitOnError)
		deploymentsDeletePFlag = deploymentsDeleteFlags.String("p", "REQUIRED", "Name of the Deployment")
	)
	deploymentsFlags.Usage = deploymentsUsage
	deploymentsUpsertFlags.Usage = deploymentsUpsertUsage
	deploymentsListFlags.Usage = deploymentsListUsage
	deploymentsGetFlags.Usage = deploymentsGetUsage
	deploymentsDeleteFlags.Usage = deploymentsDeleteUsage

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
		case "deployments":
			svcf = deploymentsFlags
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
		case "deployments":
			switch epn {
			case "upsert":
				epf = deploymentsUpsertFlags

			case "list":
				epf = deploymentsListFlags

			case "get":
				epf = deploymentsGetFlags

			case "delete":
				epf = deploymentsDeleteFlags

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
		case "deployments":
			c := deploymentsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "upsert":
				endpoint = c.Upsert()
				data, err = deploymentsc.BuildUpsertPayload(*deploymentsUpsertBodyFlag)
			case "list":
				endpoint = c.List()
				data = nil
			case "get":
				endpoint = c.Get()
				data = *deploymentsGetPFlag
			case "delete":
				endpoint = c.Delete()
				data = *deploymentsDeletePFlag
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// deploymentsUsage displays the usage of the deployments command and its
// subcommands.
func deploymentsUsage() {
	fmt.Fprintf(os.Stderr, `used to manage active deployments
Usage:
    %[1]s [globalflags] deployments COMMAND [flags]

COMMAND:
    upsert: Upsert implements upsert.
    list: List implements list.
    get: Get implements get.
    delete: Delete implements delete.

Additional help:
    %[1]s deployments COMMAND --help
`, os.Args[0])
}
func deploymentsUpsertUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] deployments upsert -body JSON

Upsert implements upsert.
    -body JSON: 

Example:
    %[1]s deployments upsert --body '{
      "containers": [
         {
            "image": "Quisquam sed.",
            "name": "Dignissimos et nemo sunt commodi iste.",
            "parameters": [
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               },
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               }
            ],
            "ports": [
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               }
            ],
            "services": [
               "Rem fugiat est.",
               "Quas enim magnam libero consectetur qui nesciunt.",
               "Quam libero."
            ]
         },
         {
            "image": "Quisquam sed.",
            "name": "Dignissimos et nemo sunt commodi iste.",
            "parameters": [
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               },
               {
                  "environment": [
                     "Voluptatem laboriosam perspiciatis est.",
                     "Est veritatis.",
                     "Quo aut recusandae.",
                     "Temporibus asperiores et voluptatem a consequatur ut."
                  ],
                  "files": [
                     "Voluptatem necessitatibus labore aliquam consequatur est.",
                     "Ullam voluptate esse.",
                     "Aut et totam voluptatem quia illo voluptate."
                  ],
                  "name": "z"
               }
            ],
            "ports": [
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               },
               {
                  "container": 63148,
                  "host": 21062,
                  "protocol": "tcp"
               }
            ],
            "services": [
               "Rem fugiat est.",
               "Quas enim magnam libero consectetur qui nesciunt.",
               "Quam libero."
            ]
         }
      ],
      "name": "Necessitatibus et autem placeat et.",
      "parameters": [
         {
            "name": "m",
            "value": "Voluptatem est eius alias est sit."
         },
         {
            "name": "m",
            "value": "Voluptatem est eius alias est sit."
         },
         {
            "name": "m",
            "value": "Voluptatem est eius alias est sit."
         }
      ]
   }'
`, os.Args[0])
}

func deploymentsListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] deployments list

List implements list.

Example:
    %[1]s deployments list
`, os.Args[0])
}

func deploymentsGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] deployments get -p STRING

Get implements get.
    -p STRING: Name of the Deployment

Example:
    %[1]s deployments get --p "Et cumque temporibus sit."
`, os.Args[0])
}

func deploymentsDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] deployments delete -p STRING

Delete implements delete.
    -p STRING: Name of the Deployment

Example:
    %[1]s deployments delete --p "Et aut ipsam in laborum distinctio veniam."
`, os.Args[0])
}
