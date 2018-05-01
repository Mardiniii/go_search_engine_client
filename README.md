# Go Search Engine Client

This is an experimental project to build a search engine, this repository contains a WEB client built on top of Go language. This code is built to consume the search engine indexes from this [repository](https://github.com/Mardiniii/go_search_engine_indexer).

Feel free to make any comment, pull request, code review, shared post, fork or feedback. Everything is welcome.

## License

This project is licensed under the **MIT License**.

## Usage

You can run all the components together the ElasticSearch server, the indexer service, two WEB app clients services and the load balancer with HAProxy using the docker-compose tool. Clone the project to your local a machine a run the commands below to start ElasticSearch server, the WEB client containers and the load balancer:

```bash
docker-compose up --build
```

Then, if you want to start crawling the internet, you should open a new tab and type the next command:

```bash
docker-compose run indexer index STARTING_URL
```

To stop the docker-compose containers just run:

```bash
docker-compose down
```

## Contributions
Feel free to make any comment, pull request, code review, shared post, fork or feedback. Everything is welcome.

## Authors

**Sebastian Zapata Mardini** - [GitHub profile](https://github.com/Mardiniii)
