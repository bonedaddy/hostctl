---
title: "toggle"
weight: 100
---

## hostctl toggle

Change status of a profile on your hosts file.

### Synopsis


Alternates between on/off status of an existing profile.


```
hostctl toggle [flags]
```

### Options

```
  -h, --help            help for toggle
  -w, --wait duration   Toggles a profile for a specific amount of time (default -1ns)
```

### Options inherited from parent commands

```
  -c, --column strings     Columns to show on lists
      --host-file string   Hosts file path (default "/etc/hosts")
  -q, --quiet              Run command without output
      --raw                Output without table borders
```

##### SEE ALSO

* [hostctl](/docs/cli-usage/hostctl)	 - Your dev tool to manage /etc/hosts like a pro

*Auto generated by spf13/cobra*