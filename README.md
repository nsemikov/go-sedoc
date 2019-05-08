## Simple example

```go
package main

import (
    "encoding/xml"
    "fmt"
    "time"

    sedoc "github.com/stdatiks/go-sedoc"
)

const (
    MyAppName    = "MyService"
    MyAppVersion = "0.1.0"
)

func main() {
    // create api
    api := sedoc.New()
    api.Description = "My self documented API"
    // create commands
    infoCmd := sedoc.Command{
        Name: "info",
        Handler: func(c sedoc.Context) error {
            c.Response().Result = fmt.Sprintf("%s v%s", MyAppName, MyAppVersion)
            return nil
        },
    }
    // add examples to command (optionaly)
    infoCmd.Examples = sedoc.Examples{
        sedoc.Example{
            Name: "simple",
            Request: sedoc.ExampleRequest{
                Object: sedoc.Request{
                    Datetime: func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-10-16T09:58:03.487508407Z"); return t }(),
                    Command:  "info",
                },
            },
            Responses: sedoc.ExampleResponses{
                sedoc.ExampleResponse{
                    Name: "simple",
                    Object: sedoc.Response{
                        Datetime: func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-10-16T09:58:03.487508407Z"); return t }(),
                        Command:  "info",
                        Result:   "MyService v0.1.0",
                    },
                },
            },
        },
    }
    // add command to api
    api.AddCommand(infoCmd)
    // create request
    request := sedoc.NewRequest()
    request.Command = "help"
    // and execute it
    response := api.Execute(request)
    // ...
    if b, err := xml.MarshalIndent(response, "", "  "); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(b))
    }
}
```

This simple example will print to stdout something like this:
```xml
<response datetime="2018-11-10T16:51:30.41952529Z" command="help">
  <result description="My self documented API">
    <request_format required="1" count="7">
      <arg name="id" type="string" description="Request identifier (for debugging)"></arg>
      <arg name="datetime" type="string" description="Datetime string (ISO 8601)"></arg>
      <arg name="session" type="uuid" description="Session token uuid, formatted like &#34;01234567-89ab-cdef-0123-456789abcdef&#34;"></arg>
      <arg name="command" type="string" description="Command name string" required="true"></arg>
      <arg name="args" type="object" description="Extra request parameters, one-level object"></arg>
      <arg name="where" type="array" description="Search item(s) parameters, simple Array of one-level objects"></arg>
      <arg name="set" type="object" description="Item(s) data to set, one-level object"></arg>
    </request_format>
    <response_format required="1" count="7">
      <arg name="id" type="string" description="Request identifier (for debugging)"></arg>
      <arg name="datetime" type="string" description="Datetime string (ISO 8601)"></arg>
      <arg name="session" type="uuid" description="Session token uuid, formatted like &#34;01234567-89ab-cdef-0123-456789abcdef&#34;"></arg>
      <arg name="command" type="string" description="Command name string" required="true"></arg>
      <arg name="args" type="object" description="Extra request parameters, one-level object"></arg>
      <arg name="result" type="object" description="Result object. For XML maybe used another name"></arg>
      <arg name="error" type="object" description="Error object. Contains `code` and `desc` fields"></arg>
    </response_format>
    <commands count="2">
      <command name="help" description="Get list of commands">
        <examples count="1">
          <example name="simple help" description="simple help command usage example">
            <request>
              <request datetime="2018-10-16T09:58:03.487508407Z" command="help"></request>
            </request>
          </example>
        </examples>
      </command>
      <command name="info" description="">
        <examples count="1">
          <example name="simple">
            <request>
              <request datetime="2018-10-16T09:58:03.487508407Z" command="info"></request>
            </request>
            <responses count="1">
              <response name="simple">
                <response datetime="2018-10-16T09:58:03.487508407Z" command="info">
                  <result>MyService v0.1.0</result>
                </response>
              </response>
            </responses>
          </example>
        </examples>
      </command>
    </commands>
    <errors>
      <error code="1" desc="unknown error occurred"></error>
      <error code="2" desc="can&#39;t parse request"></error>
      <error code="3" desc="unknown command"></error>
      <error code="4" desc="invalid command argument parameter regexp"></error>
      <error code="5" desc="match command argument parameter regexp fails"></error>
      <error code="6" desc="require command argument parameter missing"></error>
      <error code="7" desc="unknown command argument parameter in request"></error>
      <error code="8" desc="invalid command argument parameter value"></error>
    </errors>
  </result>
</response>
```

You may use JSON or XML format over any transport (WebSocket, HTTP, TCP, UDP, etc)

See [godoc](http://godoc.org/github.com/stdatiks/go-sedoc) for details
