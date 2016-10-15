
CREATE TABLE IF NOT EXISTS events(
  Id INTEGER PRIMARY KEY,
  LibId INT,
  Req  TEXT,
  Name TEXT,
  Stime TEXT,
  Etime TEXT,
  `date` TEXT,
  `desc` TEXT);
