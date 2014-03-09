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
    actual := lang.SliceContainsString(haystack, needle)
```

Extract a first level field value from a JSON:

```Go
    jsonStr := `{"labels":[],"versions":["0.1","0.1.1","0.4","0.9"]}`
    jsonB := []byte(jsonStr)
    versions, err := lang.ExtractJsonFieldValue(jsonB, "versions")
```

Converts a JSON array (that is `[]interface {}{"foo", "bar"}`) to a slice of strings:

```Go
    jsonStr := `{"labels":[],"versions":["0.1","0.1.1","0.4","0.9"]}`
    jsonB := []byte(jsonStr)
    var b map[string]interface{}
    err := json.Unmarshal(jsonB, &b)
    if err != nil {
        t.Errorf("unexpected error thrown %s", err)
    }
    if versions, keyExists := b["versions"]; keyExists {
        versionsSlice, err := lang.JsonArrayToStringSlice(versions, "versions")
        // ...
    }
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
