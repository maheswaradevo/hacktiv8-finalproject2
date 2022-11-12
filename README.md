# My Gram

Repository for Hacktiv8 Final Project 2. Already deployed on railways `hacktiv8-fp2.up.railway.app/api/v1`

## Developer's Manual

### Migrations

- First you need to install the golang-migrate to do database migrations.

MacOS

```bash
brew install golang-migrate
```

Windows (use scoop)

```bash
scoop install migrate
```

To run a migrations

```bash
migrate -source file://./db/migrations -database "mysql://root:@tcp(localhost:3306)/my_gram" up
```

To rollback a migrations

```bash
migrate -source file://./db/migrations -database "mysql://root:@tcp(localhost:3306)/my_gram" down
```

### How To Run

```bash
make run
```

## Our Team

- Pande Putu Devo Punda Maheswara
- Hanif Fadillah Amrynudin
- I Putu Agus Arya Wiguna
