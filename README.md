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
	- Postgres
    - Database Migration
        - ✅ [golang-migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)
        - [goose](https://github.com/pressly/goose)
        - [atlas](https://atlasgo.io/)
        - Next time may try `atlas`'s declarative way of migration
    - [go-jet](https://github.com/go-jet/jet)

## TODO
- [ ] Export data as encrypted file, restore by the file
