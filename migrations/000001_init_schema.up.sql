DROP TABLE IF EXISTS brands cascade;
DROP TABLE IF EXISTS products cascade;
DROP TABLE IF EXISTS customers cascade;
DROP TABLE IF EXISTS transactions cascade;
DROP TABLE IF EXISTS orders cascade;

CREATE TABLE "brands" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "brand_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "transaction_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "products" ADD FOREIGN KEY ("brand_id") REFERENCES "brands" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE INDEX ON "orders" ("transaction_id");

CREATE INDEX ON "orders" ("product_id");

CREATE INDEX ON "transactions" ("customer_id");

CREATE INDEX ON "products" ("brand_id");

COMMENT ON COLUMN "products"."price" IS 'can be negative or positive';

COMMENT ON COLUMN "transactions"."amount" IS 'must be positive';