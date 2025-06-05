# Gator CLI for RSS Feed aggregation
Gator is a CLI application that allows users to register and login and follow RSS feeds. It allows specific RSS feeds to be gathered, and can aggregate posts to be stored.

## Installation
- Install gator by running ```go install```
- Gator uses a Postgres database with connection string saved in ```~\.gatorconfig.json```. Gator assumes this file exists with format:
```
{
    "db_url":"postgres://postgres:<username>:<password>:5432/gator_go?sslmode=disable"
}
```
  
## CLI usage
To run the CLI first ensure gator is installed, and then run:

``` gator <command> [args...] ```

The list of commands with their usage:
- ``` login <name_of_user> ```
  - sets the currently logged-in user to specified user if they have already registered with gator
- ``` register <name_of_user> ```
  - registers the specified user if they have not already registered with gator
- ``` reset ```
  - deletes all users from gator
- ``` users ```
  - lists all registered users of gator and indicates which is currently logged-in
- ``` agg <time_between_requests> ```
  - fetches posts from RSS feed urls at a specified rate
- ``` addfeed <name_of_user> <feed_url> ```
  - creates an RSS feed from a url for a registered user
- ``` feeds ```
  - lists all RSS feeds
- ``` follow <feed_url> ```
  - follows a specified RSS feed from a url for the logged-in user
- ``` following ```
  - lists all RSS feeds that are followed by the logged-in user
- ``` unfollow <feed_url> ```
  - unfollows an RSS feed for the logged-in user
- ``` browse [post_limit] ```
  - gets a specified number of posts for the logged-in user. If no post limit is provided, default is 2
  
  
## Developer notes
- Gator needs Go to be installed to build
- To create a new migration, add SQL migration in sql/schema directory and run from that directory:
  - ```goose postgres "postgres://<username>:<password>@localhost:5432/gator_go" up```
- To run a down migration run the following from sql/schema directory:
  - ```goose postgres "postgres://<username>:<password>@localhost:5432/gator_go" down```
- To create queries, add SQL to sql/queries directory and then run:
  - ```sqlc generate```
- To build and run for development, replace ```gator``` in the command above with ```go run .```