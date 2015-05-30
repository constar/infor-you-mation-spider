# Infor-You-Mation-Spider

## Usage

```
make
make run
```

## Mongo Database Record

```
db.keyword.ensureIndex({"lastmodified":-1})
db.feeds.ensureIndex({"urlmd5":1},{"unique":true})
db.keyword.ensureIndex({"keyword":1, "feedid":1},{"unique":true})
```

## Contact

```
i@yanyiwu.com
```
