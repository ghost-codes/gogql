CREATE TABLE "video" (
  "id" serial PRIMARY KEY,
  "title" varchar,
  "url" varchar,
  "author" int
);

CREATE TABLE "user" (
  "id" serial PRIMARY KEY,
  "name" varchar
);

ALTER TABLE "video" ADD FOREIGN KEY ("author") REFERENCES "user" ("id");
