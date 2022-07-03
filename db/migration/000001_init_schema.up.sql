CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(32) NOT NULL,
  "username" varchar(16) UNIQUE NOT NULL,
  "email" varchar(32) NOT NULL,
  "password" varchar(60) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "user_levels" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint UNIQUE NOT NULL,
  "level" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" bigserial PRIMARY KEY,
  "title" varchar(32) NOT NULL,
  "author" varchar(32) NOT NULL,
  "year" char(4) NOT NULL,
  "pages" int NOT NULL,
  "synopsis" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "category_books" (
  "id" bigserial PRIMARY KEY,
  "book_id" bigint NOT NULL,
  "category_id" bigint NOT NULL
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(16) UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "read_confirmations" (
  "id" bigserial PRIMARY KEY,
  "book_list_id" bigint NOT NULL,
  "pages_read" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "book_lists" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "book_id" bigint NOT NULL,
  "is_read" bool NOT NULL DEFAULT false,
  "pages_read" int NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "user_levels" ("user_id");

CREATE INDEX ON "books" ("title");

CREATE INDEX ON "books" ("author");

CREATE INDEX ON "read_confirmations" ("book_list_id");

CREATE INDEX ON "book_lists" ("user_id");

ALTER TABLE "user_levels" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "category_books" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE;

ALTER TABLE "category_books" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE;

ALTER TABLE "read_confirmations" ADD FOREIGN KEY ("book_list_id") REFERENCES "book_lists" ("id") ON DELETE CASCADE;

ALTER TABLE "book_lists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "book_lists" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE;
