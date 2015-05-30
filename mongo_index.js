use inforyoumation
db.feeds.ensureIndex({"urlmd5":1},{"unique":true})
db.keyword.ensureIndex({"keyword":1, "feedid":1},{"unique":true})
db.feeds.ensureIndex({"lastmodified":-1}, {"expireAfterSeconds": 3600})
db.keyword.ensureIndex({"lastmodified":-1}, {"expireAfterSeconds": 3600})
