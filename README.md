# Appliance Surge

The #1 Q&A site for all appliance repair needs!

## Getting Started

git clone the repository, then run docker-compose up in the root directory.

Use the built in gopher commands to create the database migrations.

./gopher migrate up

### Prerequisites

These tools must be installed to run the application
- [Docker](https://www.docker.com/)

These tools must be installed for development
- [Goose](https://github.com/pressly/goose)

### Installing

Here is how to get the application up and running

Clone the repository

    git clone https://github.com/Appliance-Surge/Appliance-Surge.git

Run docker-compose up in the root directory.
This should get the build the Go application, Postgres, and Adminer

    docker-compose up

Use the built-in gopher command to run the migrations

    ./gopher migrate up

## Gopher Commands

Run migrations

    ./gopher migrate up

Tear down migrations

    ./gopher migrate down

Rebuild migrations

    ./gopher migrate up --fresh

Create a migration

    ./gopher make:migration create_users_table

Run factories

    ./gopher factory

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code
of conduct, and the process for submitting pull requests to us.

## Versioning

We use [Semantic Versioning](http://semver.org/) for versioning. For the versions
available, see the [tags on this
repository](https://github.com/PurpleBooth/a-good-readme-template/tags).

## Authors

  - **Damion Voshall** - *Maintainer* -
    [Github](https://github.com/DamoFD)

See also the list of
[contributors](https://github.com/Appliance-Surge/Appliance-Surge/contributors)
who participated in this project.

## License

This project is licensed under the [LICENSE](LICENSE.md)
Creative Commons License - see the [LICENSE.md](LICENSE.md) file for
details
