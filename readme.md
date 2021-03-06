# Exile
Exile provides an API for generating certificates using multiple CAs with
Cloudflare [cfssl](https://github.com/cloudflare/cfssl).

# Demo
[![Demo](https://raw.githubusercontent.com/ehazlett/exile/master/demo.png)](https://asciinema.org/a/23977)

# Usage
Exile combines the configuration for cfssl signing along with a set of roots
to be used for signing.  There is an example config (`exile.conf.sample`) in
this repo for reference.

```
{
    "roots": {
        "primary": {
            "key": "test/certs/primary.key",
            "certificate": "test/certs/primary.pem"
        },
        "secondary": {
            "key": "test/certs/secondary.key",
            "certificate": "test/certs/secondary.pem"
        }
    },
    "signing": {
        "default": {
            "expiry": "8760h"
        },
        "profiles": {
            "client": {
                    "usages": [
                            "signing",
                            "key encipherment",
                            "client auth"
                    ],
                    "expiry": "8760h"
            },
            "server": {
                    "usages": [
                            "signing",
                            "key encipherment",
                            "server auth",
                            "client auth"
                    ],
                    "expiry": "8760h"
            },
            "intermediate": {
                    "usages": [
                            "signing",
                            "key encipherment",
                            "cert sign",
                            "crl sign"
                    ],
                    "is_ca": true,
                    "expiry": "8760h"
            }
        }
    }
}

```

In the above example, it defines two CAs (`primary` and `secondary`).  These
are then referenced during the CSR process by using a label.  You can also
specify the cfssl profile to be used for sigining.  Here is an example request
that is sent to the Exile API:

```
{
    "certificate_request":"-----BEGIN CERTIFICATE REQUEST-----\nMIIC1jCCAcACAQAwbjELMAkGA1UEBhMCVVMxDjAMBgNVBAoTBWxvY2FsMQ4wDAYD\nVQQLEwVsb2NhbDEVMBMGA1UEBxMMSW5kaWFuYXBvbGlzMRAwDgYDVQQIEwdJbmRp\nYW5hMRYwFAYDVQQDEw1ub2RlLTAwLmxvY2FsMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEA3jG8YXh3pcYfTKLtbLI4+a3NTi5bMbzPsuQnGRk8cqrkZGdI\nmRgJ8nuKXxKOT/MnxBTlVQwPYkcVRrh8VLBc40peKW8IvEGn1Ptke2MVLwwmPWRL\n+nFcLxuLXKCLAF+toe8uW5Ytb8u7muFB4uiEbap5P3ADV7sIImZoF/1AtWLDuRCB\nKwUK7PWd+JRGCzaT8B79ulNjK7BhWQ1wJ5bIx9jFFykY83zNyhwBSrYGJ02NMk+W\nUgY9R/kBvC25NW376jcFM1LuTDlzT7iMevk4nlrbd24dXEe8fZrJ3X01fcusPVSX\nO+14UMrJDWB8jdl2ix/14lw9JfiCpWd9SG6J0QIDAQABoCUwIwYJKoZIhvcNAQkO\nMRYwFDASBgNVHREECzAJggdub2RlLTAwMAsGCSqGSIb3DQEBCwOCAQEAdrrsYxL9\nffeXHj10StQneMVc8nFhxqmOhiK5r2PJ4r6pKuVv3J1Za7GP8S6/NJYJwzrKTw/E\nKDJT6MzLuLqEsoLIFiQZte4ExuUWpKWAH70pGFRjMagfOSHU0Vxc5d9POjkZIsd8\nz8PeWxm7adyBSZgXO9hdE3DF6ljvNLY2KixfYTSFVdZ3X7ctiW+64GBtO5CJAff7\nwFERxfwYjEF9P2ixtNpRuHGmBchYtRbIbF9yAcUfvg34zkfeyOrPOzqkYSfvLqEF\nRV/q1Etp0klgJDVMCdMPp85DIvfaK4/XTSUTGSGEjZwM/ielM0amck7Ssq0j3YGy\nyWG7lLJkh+8FZg==\n-----END CERTIFICATE REQUEST-----\n",
    "hosts": [
        "node-00"
    ],
    "profile": "server",
    "label": "primary"
}
```

This request contains a CSR, extra "hosts" to be added as alt names, a cfssl
profile name and the label to select which CA to be used for signing.

# Tests
To run the unit tests:

`godep go test -v ./...`

To run the integration tests:

`bash test.sh`
