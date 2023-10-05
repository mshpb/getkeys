# getkeys

## Installation

As a library

```shell
go get github.com/mshpb/getkeys
```

## Usage

Add your application configuration to your `.env` file in the root of your project:

```shell
S3_BUCKET=YOURKEY
SECRET_KEY=YOURSECRETKEYGOESHERE
```

Then in your Go app you can do something like

```go
package main

import (
    "log"
    "os"

    "github.com/mshpb/getkeys"
)

func main() {
    // enter your .env file with the following
  err := getkeys.LoadEnv(".end")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  s3Bucket := os.Getenv("S3_BUCKET")
  secretKey := os.Getenv("SECRET_KEY")

  // now do something with s3 or whatever
}
```
