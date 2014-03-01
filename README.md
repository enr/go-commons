Go commons
==========

[![Build Status](https://travis-ci.org/enr/go-commons.png?branch=master)](https://travis-ci.org/enr/go-commons)

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

   Copyright 2014 go-commons contributors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
