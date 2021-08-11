# Overview

Sign up [here](https://keyspot.app)

KeySpot is a tool to help manage environment variables for individuals and teams of developers. The service stores environment variables for your project in a centralized place so you don't have to juggle different .env files for your environements and applications. Once you have signed in at [keyspot.app](https://keyspot.app), you can create new records, share them with your

## Installation

```go
import (
    "github.com/keyspot/gopackage"
)
```

## Usage

Sign in to [KeySpot](https://keyspot.app), and create a record. At the top of each record's page there is an accessKey. Copy the accessKey at the top of the page as you will be using this to access your environment variables in code.

### Set Environment Variables

You can directly add your secrets as environment variables in code:

```go
import (
    "github.com/keyspot/gopackage"
    "os"
)

keyspot.SetEnvironment("<accessKey>")

// access secrets with var secretValue string = os.Getenv("<secret-key>")
```

### Get Secrets in Code Directly

Get your secrets in code directly:

```go
import (
    "github.com/keyspot/gopackage"
)

var secrets map[string]string = keyspot.GetRecord("<accessKey>")
```

### Update Secets Programmatically

Update your secrets porgrammatically:

```go
import (
    "github.com/keyspot/gopackage"
)

var secrets map[string]string = keyspot.GetRecord("<accessKey>")
```
