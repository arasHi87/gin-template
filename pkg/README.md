# Arch

```
--- pkg
 |------- common
 |------- controller
 |------- middleware
 |------- model
 |------- router
 |------- setting
```

- common
  - put tool which use in other service, ex. jwt, logger, etc.
- controller
  - accept input come from router
- middleware
  - some middleware, ex. logger, auth, etc.
- router
  - put all router here
- setting
  - convert enviroment value to CONFIG struct here
