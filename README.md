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
 null TextEditorComponent::domNode during visibility check           ┌────────────────────────────────────────────────────────────────────┐
  Avoid infinite recursion when bad values are passed to tz aware .. │  Add build script                                                  │
  Avoid distinct if a subquery has already materialized              │                                                                    │
  Add -enable-experimental-nested-generic-types frontend flag        │                                                                    │
  Add support for activating and deactivating package-specific key.. │                                                                    │
  Add a helper method mayHaveOpenedArchetypeOperands to SILInstruc.. │                                                                    │
  Add a typealias to avoid a build ordering dependency between pro.. │                                                                    │
  Add --main-process flag to run specs in the main process           │                                                                    │
  Add support for closure contexts to readMetadataFromInstance()     │                                                                    │
  Add a basic test for opening an editor in largeFileMode if >= 2M.. │                                                                    │
  Add support for allocators that require tensors with zero          │                                                                    │
  Add "event" parameter for "click" handler of MenuItem              │                                                                    │
  Add a design-decisions section to the CONTRIBUTING guide           │                                                                    │
  Add Throws flag and ThrowsLoc to AbstractFunctionDecl              │                                                                    │
  Add TODO about blinkFeatures -> enableBlinkFeatures                │                                                                    │
  Add validation test for projecting existentials                    │                                                                    │
  Add failing spec for Menu.buildFromTemplate                        │                                                                    │
  Add SkUserConfig.h with blank SkDebugf macro                       │                                                                    │
  Add support for launching HTML files directly                      │                                                                    │
  Add documentation for --proxy-bypass-list                          │                                                                    │
  Add assertions for no available bookmark                           │                                                                    │
  Add assert for role with app name in label                         │                                                                    │
  Add convenience API for demangling                                 │                                                                    │
  Add specs for moveSelectionLeft()                                  │                                                                    │
  Add comment about map key/values                                   │                                                                    │
  Add TypeLowering::hasFixedSize()                                   │                                                                    │
  Add File > Exit menu on Windows                                    │                                                                    │
  Add tests for pending pane items                                   │                                                                    │
  Add missing period in comment                                      │                                                                    │
  Add docs for app.getLocale()                                       │                                                                    │
  Add asserts for properties                                         │                                                                    │
  Add style.less examples                                            │                                                                    │
  Add overflow scrolling                                             │                                                                    │
  Add npm start script                                               │                                                                    │
  Add missing return                                                 │                                                                    │
> Add build script                                                   │                                                                    │
  38/125                                                             │                                                                    │
> Add
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
