# mvdspawn

Parse player spawns from an MVD demo.

## Usage

```sh
$ mvdspawn testdata/20241210-2125_4on4_sk_vs_tot[dm3].mvd | head -n8
dm3;tot;ToT_fix;sng-mega
dm3;tot;ToT_slime;sng-ra
dm3;sk;gosciu;ya
dm3;sk;snapcase;lifts
dm3;sk;rokky;rl
dm3;tot;ToT_phren;ra-tunnel
dm3;tot;ToT_Javve;sng-mega
dm3;sk;kane;lifts
```

## Locs

`mvdspawn` comes with built-in locations for `dm2`, `dm3`, and `e1m2`. You can use
custom locations by specifying the `-locs-path` command-line option.

## Known issues

The application parses the `svc_playerinfo` message and checks whether the
`DF_DEAD` bit is set to determine if a player has died. While this method is not
completely foolproof, there is unfortunately no better way to determine this
without significantly increasing the code's complexity, which I prefer to
avoid at this point.
