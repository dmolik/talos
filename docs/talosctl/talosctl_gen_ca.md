<!-- markdownlint-disable -->
## talosctl gen ca

Generates a self-signed X.509 certificate authority

### Synopsis

Generates a self-signed X.509 certificate authority

```
talosctl gen ca [flags]
```

### Options

```
  -h, --help                  help for ca
      --hours int             the hours from now on which the certificate validity period ends (default 87600)
      --organization string   X.509 distinguished name for the Organization
      --rsa                   generate in RSA format
```

### Options inherited from parent commands

```
      --context string       Context to be used in command
  -e, --endpoints strings    override default endpoints in Talos configuration
  -n, --nodes strings        target the specified nodes
      --talosconfig string   The path to the Talos configuration file (default "/home/user/.talos/config")
```

### SEE ALSO

* [talosctl gen](talosctl_gen.md)	 - Generate CAs, certificates, and private keys

