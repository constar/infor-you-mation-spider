use inforyoumation;
//db.dropDatabase();
var expireSeconds = 3600*24*7;
db.feeds.ensureIndex({"urlmd5":1},{"unique":true});
db.keyword.ensureIndex({"keyword":1, "feedid":1},{"unique":true});
db.feeds.ensureIndex({"lastmodified":-1}, {"expireAfterSeconds": expireSeconds});
db.keyword.ensureIndex({"lastmodified":-1}, {"expireAfterSeconds": expireSeconds});
