Go commons
==========

Go libraries for common functionalities.

Environment
-----------

Library pertaining system and environment.

Import the library:

```Go
    import (
        "github.com/enr/go-commons/environment"
    )
```

Get environment variable:

```Go
    res := environment.GetenvEitherCase("TestGetenvEitherCase")

```

Get user home:

```Go
    home, err := environment.UserHome()

```

Lang
----

Library pertining common data structures.

Import the library:

```Go
    import (
        "github.com/enr/go-commons/lang"
    )
```

Find if the given string is present in a strings slice:

```Go
    SliceContainsString(list []string, a string)

```

Converts a interface to slice of strings:

```Go
    InterfaceToStringSlice(v interface{}, k string) ([]string, error) 

```


License
-------

Apache 2.0 - see LICENSE file.
