# UUID CLI

**A versatile tool for generating UUIDs**

* native binary, zero dependencies
* easy to use in containers
* all UUID versions supported (thanks to [github.com/gofrs/uuid](https://github.com/gofrs/uuid))


## Reference

```
Usage of uuid-cli:
  -d string
    	Use organisation (org), UID (person) or GID (group) for V2 UUID
  -hex
    	Encode UUID as hex number (without hyphens)
  -ns string
    	Namespace UUID for V3/V5 UUID
  -v int
    	UUID version (1, 2, 3 or 4), in a nutshell:
    	  1: machine
    	  2: domain (org, person, group)
    	  3: namespaced (MD5)
    	  4: random
    	  5: namespaced (SHA-1)
    	  (default 4)
```

## Examples

### Version 4 (default)

    $ uuid-cli
    be040956-4caf-43b3-bfe7-8cf36fd9e573

### Version 1

    $ uuid-cli -v 1
    be040956-4caf-43b3-bfe7-8cf36fd9e573

### Version 5

    $ uuid-cli -v 5 -ns be040956-4caf-43b3-bfe7-8cf36fd9e573 my-name
    07b676c5-6b33-53a0-a3a0-4857ee8252e8

## Docker

    $ docker run networkteamcom/uuid-cli -v 1
    0a86245a-4411-11ea-9636-acde48001122

# License

[MIT](./LICENSE)