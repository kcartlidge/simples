# Simples-Config

An easy to use configuration loader that
follows simple priority rules for each setting.

1. Environment variables take priority.
2. Check a specified config file next.
3. Drop back to a coded default.

## How to use it.

The idea is that you should be exact about things
that must change for a named environment, but
also degrade to more general defaults whose values
are shared across multiple environments.

For example your database connection strings and secret
keys might be in environment variables in your cloud,
whereas a standard list page size could be in a config
file (for all environments) which is then overridden by
a more specific value in an environment variable if needed.

## Code example.

Install the package first:

``` sh
go get github.com/kcartlidge/simples-config
```

Then use like this:

``` go
package main

import (
    "log"
    simples "github.com/kcartlidge/simples-config"
)

func main() {
    // Create a config object from an env file.
    c, err := simples.CreateConfig(".env")
    if err != nil {
	    log.Fatalln(err.Error())
    }

    // Extract the values (with defaults).
    valueAsString := c.Get("PROJECT_TITLE", "Unnamed Project")
    valueAsNumber := c.GetNumber("PAGE_SIZE", 10)

    // Show the results.
    fmt.Println("Project:", valueAsString)
    fmt.Println("Page Size:", valueAsNumber)
}
```

Getting a value will always fall back to the default
if a real value cannot be found (or ```getNumber```
cannot convert that value to a number). Therefore
there is no need to check for errors.

## Configuration file format.

The file is expected to follow an extremely simple
format and layout. Setting names may as well be in
capitals as they are treated as such anyway; values
are trimmed but otherwise untouched.

``` ini
MODE          = Production
DB_CONN       = my-long-winded-connection-string:27017
PAGE_SIZE     = 10
PROJECT_TITLE = My Example Project

# Shows at the top of the web page.
BANNER        = Version 2 has been released!
```

Your setting name is first. This is followed by one
or more whitespace characters (which are ignored),
then an *equals* delimiter.

The remainder of the line forms the value, which
has leading/trailing whitespace removed and can
include *equals* too.

Lining up (as per the example above) is entirely
optional, as are comments - which are lines that
start with a hash symbol.

## Environment variables.

Standard stuff. If you provide one, it overrides
any loaded configuration file value with the same
name (case insensitive).

``` sh
export PROJECT_TITLE=Website
```

## Methods available.

### CreateConfig

``` go
CreateConfig(filename string) (Config, error)
```

This reads in the given filename and caches all
settings found. It returns an error if the
file could not be loaded. Any lines that are
not key/value pairs are ignored.

The main return object is a ```Config``` which
has the following methods available on it.

### Get

``` go
Get(key string, defaultValue string) string
```

This returns the given key's value from the
environment variable if it exists. Otherwise
it will return from the loaded configuration
file or drop back to returning the default.

### GetNumber

``` go
GetNumber(key string, defaultValue int) int
```

This works the same as ```Get``` but expects
to find a number value. If the value is not
convertible to a number, the default is
returned.

*Note that by number I mean an integer, so
whole numbers only.*

## Fetching and running tests.

Simply fetch the package and run as usual:

``` sh
go get github.com/kcartlidge/simples-config
go test
```

## Performance and cacheing.

Environment variables are *not cached*; whenever a
setting is requested whose value derives from an
environment variable the *current live value* is returned.

Configuration file settings *are cached*, and will
always reflect the *value at launch*.

As file settings are cached, performance is not
impacted by where the setting comes from.

Copyright: **K Cartlidge** | License: **MIT**
