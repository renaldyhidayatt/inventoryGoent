CREATE TABLE "categories" (
  "id" bigint NOT NULL,
  "name" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "category_productkeluar" (
  "category_id" bigint NOT NULL,
  "product_keluar_id" bigint NOT NULL
);

CREATE TABLE "customers" (
  "id" bigint NOT NULL,
  "name" "character varying" NOT NULL,
  "alamat" "character varying" NOT NULL,
  "telepon" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "product_category" (
  "product_id" bigint NOT NULL,
  "category_id" bigint NOT NULL
);

CREATE TABLE "product_keluars" (
  "id" bigint NOT NULL,
  "qty" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "product_masuk_supplier" (
  "product_masuk_id" bigint NOT NULL,
  "supplier_id" bigint NOT NULL
);

CREATE TABLE "product_masuks" (
  "id" bigint NOT NULL,
  "name" "character varying" NOT NULL,
  "qty" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "product_productkeluar" (
  "product_id" bigint NOT NULL,
  "product_keluar_id" bigint NOT NULL
);

CREATE TABLE "product_productmasuk" (
  "product_id" bigint NOT NULL,
  "product_masuk_id" bigint NOT NULL
);

CREATE TABLE "products" (
  "id" bigint NOT NULL,
  "name" "character varying" NOT NULL,
  "image" "character varying" NOT NULL,
  "qty" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "suppliers" (
  "id" bigint NOT NULL,
  "name" "character varying" NOT NULL,
  "alamat" "character varying" NOT NULL,
  "telepon" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "users" (
  "id" bigint NOT NULL,
  "firstname" "character varying" NOT NULL,
  "lastname" "character varying" NOT NULL,
  "email" "character varying" NOT NULL,
  "password" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

ALTER TABLE "category_productkeluar" ADD CONSTRAINT "category_productkeluar_category_id" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE;

ALTER TABLE "category_productkeluar" ADD CONSTRAINT "category_productkeluar_product_keluar_id" FOREIGN KEY ("product_keluar_id") REFERENCES "product_keluars" ("id") ON DELETE CASCADE;

ALTER TABLE "product_category" ADD CONSTRAINT "product_category_category_id" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE;

ALTER TABLE "product_category" ADD CONSTRAINT "product_category_product_id" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE "product_masuk_supplier" ADD CONSTRAINT "product_masuk_supplier_product_masuk_id" FOREIGN KEY ("product_masuk_id") REFERENCES "product_masuks" ("id") ON DELETE CASCADE;

ALTER TABLE "product_masuk_supplier" ADD CONSTRAINT "product_masuk_supplier_supplier_id" FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id") ON DELETE CASCADE;

ALTER TABLE "product_productkeluar" ADD CONSTRAINT "product_productkeluar_product_id" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE "product_productkeluar" ADD CONSTRAINT "product_productkeluar_product_keluar_id" FOREIGN KEY ("product_keluar_id") REFERENCES "product_keluars" ("id") ON DELETE CASCADE;

ALTER TABLE "product_productmasuk" ADD CONSTRAINT "product_productmasuk_product_id" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE "product_productmasuk" ADD CONSTRAINT "product_productmasuk_product_masuk_id" FOREIGN KEY ("product_masuk_id") REFERENCES "product_masuks" ("id") ON DELETE CASCADE;
