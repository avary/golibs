Go libraries that I reuse across my projects. NOT SAFE FOR PUBLIC USE ☢️

## Libs


`barcode` is a fork of https://github.com/skerkour/golibs/barcode commit 65580ac6e377baeeebd5bb702e15795ead22e48d - License: MIT


`cobra` is a fork of https://github.com/spf13/cobra commit 212ea4078323771dc49b6f25a41d84efbaac3a4c to remove a lot of dependencies (`doc` directory deleted) - License: Apache 2.0

```shell
$ cd cobra && rm -rf doc go.mod go.sum && cd ..
```

`cors` is a fork of https://github.com/rs/cors commit fcebdb403f4d4585c705318c0e4d6d05a761a4ab - License: MIT

`cpuinfo` is a fork of https://github.com/klauspost/cpuid commit d685acd433f5dde4e315aa5b1eb8e72b9ecce117 - License: MIT

`cron` is a fork of https://github.com/robfig/cron commit bc59245fe10efaed9d51b56900192527ed733435 - License: MIT

`feeds` is a fork of https://github.com/jlelse/feeds commit 189f94254ad4b61b3b1908cf493f94390dc79870 - License: BSD 2-Clause

`goldmark-highlighting` is a fork of https://github.com/yuin/goldmark-highlighting commit 151362477c8778cdfd54adeea5cb3405231f3018 - License: MIT

`imaging` is a fork of https://github.com/disintegration/imaging commit d40f48ce0f098c53ab1fcd6e0e402da682262da5 - License: MIT

`migrate` is a fork of https://github.com/joncalhoun/migrate commit 34a9ee7d2b52f7b9f2261c2b6d9a8a0a83890a7f - License: MIT (see https://www.calhoun.io/database-migrations-in-go/)

`namesgenerator` is a fork of https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go commit 0f052eb4f56c05dcb8c444823ebde6ce0fac7197 - Licesne: Apache 2.0

`otp` is a fork of https://github.com/skerkour/golibs/otp commit c62dc589378ae5c364d36819ddfb03fe391635ad - License: Apache 2.0

`retry` is a fork of https://github.com/avast/retry-go commit 27363a141859f4031a7726c5bffcd670fb42d286 - License: MIT

`singleinstance` is a fork of https://github.com/snabb/sitemap commit ac70ad656cd7f36cc5623306806a3583707c9a56 - License: MIT

`sitemap` is a fork of https://github.com/postfinance/singlecommit 43dccf267e7c560d78380cbc074044ee27fb6e97 - License: MIT

`sysinfo` is a fork of https://github.com/skerkour/golibs/sysinfo commit 99e836ba64f229922382eecd8a346e6b3d1b560a - License: MIT

`toml` is a fork of https://github.com/skerkour/golibs/toml commit 1ba7f5b05951c7331f4eefea9bce295e2cb141f2 - License: MIT

`uuid` is a fork of https://github.com/google/uuid commit 44b5fee7c49cf3bcdf723f106b36d56ef13ccc88 - License: BSD-3 Clause

`validate` is a fork of https://github.com/asaskevich/govalidator commit f21760c49a8d602d863493de796926d2a5c1138d - License: MIT


## Maintenance

```bash
$ go get -u ./...
$ go mod tidy
```
