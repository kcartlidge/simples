# simples-config

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
keys might be in *AWS* environment variables (if you run
an *EC2* instance), whereas a standard list page size
could be in a config file for all environments and then
overridden by an environment variable if needed.

### Code example.

``` go
package main

import (
    simples "github.com/kcartlidge/simples-config"
)

func main() {
    c := simples.CreateConfig('.env')
    v := c.Get("PAGE_SIZE", "10")
    fmt.Println(v)
}
```

## Configuration file format.

The file is expected to follow an extremely simple
format and layout. Setting names ideally should be in
capitals (see the note about casing below), and values
are trimmed but otherwise untouched.

``` ini
MODE        = Production
DB_CONN     = my-long-winded-connection-string:27017
PAGE_SIZE   = 10

# Shows at the top of the web page.
BANNER      = Version 2 has been released!
```

Your setting name is first. This is followed by
one or more whitespace characters (which are
ignored), then an *equals* delimiter.
The remainder of the line forms the value, which
has leading/trailing whitespace removed and can
include *equals* too.

Lining up (as per the example above) is entirely
optional, as are comments - which are lines that
start with a hash symbol.

### A note about casing.

Requests via ```Get``` always treat the ```Key``` as
uppercase. Keys for entries in the *config* files are
also treated as uppercase, so lookups will succeed.
Environment variables are *not* treated as uppercase
by default, so unless your environment variables
are uppercased already they *will not* succeed.

Given this disjoint on casing, I suggest you simply
restrict yourself to uppercase for every key name
whether in a config file or the environment. As it
happens, that's what most developers do anyway.

## Performance and timing.

Environment variables are *not* cached; whenever a
setting is requested whose value derives from an
environment variable the *current* value is given.

Configuration file settings *are* cached, and will
always reflect the value at launch.

As file settings are cached, performance is not
impacted by where the setting comes from.

Copyright: **K Cartlidge** | License: **MIT**
