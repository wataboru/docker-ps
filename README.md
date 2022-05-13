# docker-ps

docker-ps provides a fuzzy search function for `docker ps`.

![description](https://raw.githubusercontent.com/wiki/wataboru/git-fuzzy-find-commit-message/images/fcm_description.gif)

## Use as CLI

### Install

```
$ go get github.com/wataboru/docker-ps/cmd/fcm
```

Download from [releases](https://github.com/wataboru/git-fuzzy-find-commit-message/releases)

### Usage

1. Run
```
$ docker-ps
```

2. Choose contaienr
```
  b65cdae7874b   hogee:alipne   "bash"                   3 weeks ago         Exited (2) 3 weeks ago                                   hogee
  5232f50e8d12   fuga:alpine    "docker-entrypoint.s… "  11 days ago         Up About an hour               0.0.0.0:12345->3210/tcp   fuga-1
  349556d99c30   piyo:5.6       "docker-entrypoint.s… "  11 days ago         Up About an hour               0.0.0.0:10000->0123/tcp   piyo-1
  4d1c0ad7eda8   poyo:alpine    "bash"                   10 days ago         Exited (1) 20 days ago                                   poyo_run_81dd864f669c
  8d78bcf3693a   test1          "bash"                   19 days ago         Exited (1) 19 days ago                                   test1_run_59d0150a8b64
  ac59c93da4e8   test2_app      "bash"                   18 days ago         Exited (1) 18 days ago                                   test2_run_63ed7e08a01c
  94e703547a2d   test3_app      "bash"                   17 days ago         Exited (0) 17 days ago                                   test3_run_79a51946aca1
  b275a81b642e   test4_app      "bash"                   16 days ago         Exited (1) 16 days ago                                   test4_run_3bcd0ec0a301
  e9ca6ebb0723   test5_app      "bash"                   15 days ago         Exited (0) 15 days ago                                   test5_run_c075b8e0c8d1
  5bc6f6bbdff7   test6_app      "bash"                   14 days ago         Exited (0) 14 days ago                                   test6_run_ae6eb9b7635e
> 3e1b9dcb1231   test7_app      "bash"                   13 days ago         Exited (1) 13 days ago                                   test7_run_89fc025d56eb
  a092081a5fb9   test8_app      "bash"                   12 days ago         Exited (1) 12 days ago                                   test8_run_56918d43cfb5
  b5b4311601a0   test9_app      "bash"                   11 days ago         Exited (1) 11 days ago                                   test9_run_353288f2f896
  f3fdd457fda4   test0_app      "/bin/sh"                About an hour ago   Exited (0) About an hour ago                             test-app-1
  14/14
```

4. Container ID will be copied to the clipboard
```
```

### Version

```
$ docker-ps -v
```

## Use as a libary

- https://github.com/ktr0731/go-fuzzyfinder
- https://github.com/tcnksm/ghr

## Refer to the following for implementation

- https://github.com/skanehira/fk  
I referred to CI and Build as a whole.
