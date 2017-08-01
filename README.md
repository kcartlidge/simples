# Simples-Config v1.0.0

An easy to use configuration loader that understands sections of settings in an ```ini``` style format:

``` ini
[SITE]
title = Sample Site

[NETWORK]
hostname = localhost
port     = 8080
```

Previously this allowed overrides by environment variables. The environment is no longer checked.

## How to use it

An example follows below, but in brief:

1. Create a file of settings. Whilst the name is unimportant, it is similar enough to an ```ini``` file to be worth using that extention.

2. Use sections to group together related settings.

3. Use ```simples-config``` to acces by individual setting or a section at a time.

## Code example

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
    // Create a config object from an ini file.
    c, err := simples.CreateConfig("settings.ini")
    if err != nil {
        log.Fatalln(err.Error())
    }

    // Extract the values (with defaults).
    vs := c.GetString("NETWORK", "DomainName" "localhost")
    vn := c.GetNumber("NETWORK", "Port", 10)

    // Show the results.
    fmt.Println("Domain:", vs)
    fmt.Println("Port:", vn)
}
```

Getting a value will always fall back to the default if a real value cannot be found (or ```getNumber``` cannot convert that value to a number).
Therefore there is no need to check for errors.

## Configuration file format

The file is expected to follow an extremely simple
format and layout. Section and setting keys are always compared without regard to casing; values
are trimmed but otherwise untouched.

``` ini
# Without a section, these go under DEFAULT.
MODE          = Production
PAGE_SIZE     = 10
PROJECT_TITLE = My Example Project

[DATABASE]
MONGODB       = my-long-winded-connection-string:27017
USERNAME      = site_user

# You can fetch sections with a sequence, allowing
# the ordering here to be replicated in your code.
[MENU]
Home         = /home.html
About Me     = /about.html
```

* Entries are added to a ```DEFAULT``` section until an actual named section is reached.
* Lines enclosed in ```[]``` provide section names for all following lines until another section is declared.
* Section and key names are *not* treated as case-sensitive for comparison/lookup purposes.
* Everything to the left of the first ```=``` symbol is the key (though it *is* trimmed).
* Everything to the right of the first ```=``` symbol is the value (though it *is* trimmed). This may also include further ```=``` symbols, which are then treated exactly as per any other character.
* Indentation is irrelevant, as is line-spacing.
* Lines starting with a hash (```#```) symbol are comments; these are ignored.
* A duplicate section's entries will be appended to the existing section.

## Methods available

### CreateConfig

``` go
CreateConfig(filename string) (Config, error)
```

This reads in the given filename. It returns an error if the file could not be loaded. Any lines that are not key/value pairs are ignored.

The main return object is a ```Config``` which
has the following methods available on it.

### GetString

``` go
GetString(section, key string, defaultValue string) string
```

This returns the given section's key's value from the from the loaded configuration file, or the default value.

### GetNumber

``` go
GetNumber(section, key string, defaultValue int) int
```

This works the same as ```GetString``` but expects
to find a number value. If the value is not
convertible to a number, the default is
returned. *Note that by number I mean integer, so
whole numbers only.*

### GetSections

``` go
GetSections() []string
```

Returns a slice with the names of all known sections. Useful for keying into sections dynamically.

### GetSection

``` go
GetSection(section string)  map[int]Entry
```

Returns a single section's entries, keyed by sequence (1-based).

Each entry has it's own sequence within it, and these sequences match the ordering in the original ```ini``` file. As items are keyed by their sequence, they should be accessed like this:

``` go
cnt := len(m)
for i := 1; i <= cnt; i++ {
  e := m[i]
  fmt.Println(i, e.Key, e.Value)
}
```

In Go simply iterating over the map *does not* guarantee the order, whereas keying does.

## Fetching and running tests

Simply fetch the package and run as usual:

``` sh
go get github.com/kcartlidge/simples-config
cd $GOPATH/src/github.com/kcartlidge/simples-config
go test
```

## Performance and cacheing

Configuration file settings *are cached*, and will always reflect the *value at launch*.

Copyright: **K Cartlidge** | License: **MIT**
