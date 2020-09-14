# kreepr
[![Build Status](https://travis-ci.com/80-am/kreepr.svg?branch=master&status=started)](https://travis-ci.com/80-am/kreepr)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/80-am/kreepr)](https://golang.org/)
[![License: GPL-3.0](https://img.shields.io/github/license/80-am/kreepr)](https://opensource.org/licenses/GPL-3.0)

Kreeping social media for changes in trends and behavior.

## Getting Started
These instructions will get you up and running on your local machine.

```sql
CREATE DATABASE kreepr;
```

Import [db.sql](db.sql) using mysql or a db visualization program.
```bash
mysql -u username -p password kreepr < ./db.sql
```

Copy [secrets.yml.sample](secrets.yml.sample) into secrets.yml and fill in your secrets.

```yml
# db
user: "your db user"
password: "your db password"
schema: "/kreepr"

# twitter
key: "your twitter key"
secret: "your twitter secret" 
token: "your twitter token"
access_token: "your twitter access token"
access_secret: "your twitter access secret
```

### Subjects
Use `-add` argument to start kreeping someone.
```bash
add="twitterUserName"
```
Similar you can `-drop` your subject.
```bash
drop="twitterUserName"
```

### Daily Job
You can use crontab to update the history of your subjects.
```bash
0 0 * * * go run /PATH_TO_BUILT_APP/app
```