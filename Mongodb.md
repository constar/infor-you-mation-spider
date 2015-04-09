# Mongo Database Record

## Usage

```
db.keyword.ensureIndex({"lastmodified":-1})
db.feeds.ensureIndex({"urlmd5":1},{"unique":true})
db.keyword.ensureIndex({"keyword":1, "feedid":1},{"unique":true})
```
