# Simples-Config

An extremely easy to use configuration loader that
follows simple priority rules for each setting.

1. Environment variables take priority.
2. Check a specified config file next.
3. Drop back to a specified default.

## How to use it.

The idea is that you should be specific about things
that must change for your running environment, but
also degrade to more general defaults whose values
are shared across multiple environments.

For example your database connection strings and secret
keys might be in environment variables in your cloud,
whereas a standard list page size could be in a config
file (for all environments) which is then overridden by
an environment variable if needed.

### Code example.

``` go
package main

import (
    simples "github.com/kcartlidge/simples-config"
)

func main() {
    // Create a config object from an env file.
    c := simples.CreateConfig('.env')

    // Extract the values (with defaults).
    valueAsString := c.Get("PROJECT_TITLE", "Unnamed Project")
    valueAsNumber := c.GetNumber("PAGE_SIZE", 10)

    // Show the results.
    fmt.Println("Project:", valueAsString)
    fmt.Println("Page Size:", valueAsNumber)
}
```

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

## Performance and cacheing.

Environment variables are *not cached*; whenever a
setting is requested whose value derives from an
environment variable the *current live value* is returned.

Configuration file settings *are cached*, and will
always reflect the *value at launch*.

As file settings are cached, performance is not
impacted by where the setting comes from.

Copyright: **K Cartlidge** | License: **MIT**
