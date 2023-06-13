## Tech Stack

- User i/o
	- [cobra](https://github.com/spf13/cobra)
	- [pterm](https://github.com/pterm/pterm)
- Core system
	- golang
	- [viper](https://github.com/spf13/viper)
    - [go-money](https://github.com/Rhymond/go-money)
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
