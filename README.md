## Local Test

Clone this repo and run

```
make setup-local
```

![show](https://github.com/LintaoAmons/undercontrol-go/assets/95092244/cebd9f42-3d69-481d-9382-cd66f7e801f4)

## Commands

### account

#### add

![add](https://github.com/LintaoAmons/undercontrol-go/assets/95092244/3f6b50ea-0e5c-4e8a-8251-af5264d8764f)

#### list

![image](https://github.com/LintaoAmons/undercontrol-go/assets/95092244/2f81b758-1d36-456d-9cfc-5e9a76c5047a)

## Tech Stack

- User i/o
  - [cobra](https://github.com/spf13/cobra)
  - [pterm](https://github.com/pterm/pterm)
- Core system
  - golang
  - [viper](https://github.com/spf13/viper)
    - [go-money](https://github.com/LintaoAmons/go-money)
- Persistency
  - entgo with sqlite

## TODO

- [ ] InitDBClient connect string to configuration
- [x] Show History of account
- [ ] docker-compose and dockerfile
  - [ ] extract env variable to allow docker-compose configuration
- [ ] Configuration: db
- [ ] Add Context to init dependencies
- [ ] Export data as encrypted file, restore by the file
