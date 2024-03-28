# Simples v2.0.0

An easy to use configuration loader that understands sections of settings in an `ini` style format:

``` ini
[SITE]
title = Sample Site

[NETWORK]
hostname = localhost
port     = 8080
```

## How to use it

An example follows below, but in brief:

- Create a file of settings - whilst the name is unimportant, it is similar enough to an `ini` file to be worth using that file extention (and a related plugin for your editor)
- Use sections to group together related settings
- Use `simples` to access by individual setting or a section at a time

## Code example

Install the package first:

``` sh
go get github.com/kcartlidge/simples/v2
```

Then use like this:

``` go
package main

import (
    "log"
    simples "github.com/kcartlidge/simples/v2"
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

Getting a value will always fall back to the default if a real value cannot be found (or if `GetNumber` cannot convert that value to a number).
Therefore there is no need to check for errors.

## Configuration file format

The file is expected to follow an extremely simple format and layout.
Keys for section and setting are case-insensitive.
Values are trimmed but otherwise untouched.

``` ini
# Without a section, these appear under the DEFAULT one.
MODE          = Production
PAGE_SIZE     = 10
PROJECT_TITLE = My Example Project

[DATABASE]
MONGODB       = my-long-winded-connection-string:27017
USERNAME      = site_user

# Your code can fetch an entire section at once.
# This will include a sequence, allowing the order here to be retained in your code.
[MENU]
Home           = /home.html
About Me       = /about.html
Privacy Policy = /privacy.html
```

- Entries are added to a `DEFAULT` section until the first named section is reached
- Lines enclosed in `[]` provide section names for all following lines until another section name overrides it
- Section and key names are *case-insensitive* for comparison/lookup purposes
- Everything to the left of the first `=` symbol is the key (and is trimmed)
- Everything to the right of the first `=` symbol is the value (and is trimmed)
    - This may also include further `=` symbols, which are then treated as part of the value in the same way as any other character
- Indentation is irrelevant, as is line-spacing
- Lines starting with a hash (`#`) symbol are comments; these are ignored
    - You can also use a semi-colon (`;`) to denote a comment
- Entries under a duplicated section name will be appended to the existing section

## Methods available

### CreateConfig

``` go
CreateConfig(filename string) (Config, error)
```

This reads in the given filename. It returns an error if the file could not be loaded. Any lines that are not key/value pairs are ignored.

The main return object is a `Config` which has the following methods available on it.

### GetString

``` go
GetString(section, key string, defaultValue string) string
```

This returns the value for the given section and key from the from the loaded configuration file (or the default value if missing).

### GetNumber

``` go
GetNumber(section, key string, defaultValue int) int
```

This works the same as `GetString` but expects to find a numeric value.
If the value is not convertible to a number, the default is returned.
Integers (whole numbers) only.

### GetSections

``` go
GetSections() []string
```

Returns a slice with the names of all known sections. Useful for keying into sections dynamically.

### GetSection

``` go
GetSection(section string) map[int]Entry
```

Returns a single section's entries, keyed by sequence (1-based).

Each entry has its own sequence within it, and these sequences match the ordering in the original `ini` file (as mentioned in the `MENU` section of the example above).
Items are keyed by their sequence, so they can be accessed like this:

``` go
// Iterates in a predictable sequence.
menu := c.GetSection("MENU")
count := len(menu)
for i := 1; i <= count; i++ {
  entry := menu[i]
  fmt.Println(i, entry.Key, "=", entry.Value)
}
```

In Go simply iterating over the map *does not* guarantee the order, whereas keying does.
If you *don't care* about the order you can use `range`:

``` go
// Warning: range ordering over map keys is not guaranteed in Go.
for _, v := range c.GetSection("MENU") {
  fmt.Println(v.Sequence, v.Key, "=", v.Value)
}
```

## Fetching and running tests

``` sh
go test
```

## Performance and caching

Configuration file settings *are cached*, and will always reflect the value *at launch*.

Copyright 2017-2024, **K Cartlidge** | License: **MIT**
