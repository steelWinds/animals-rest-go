CREATE TABLE owner_units(
	  "id" SERIAL,
	  "name" TEXT NULL DEFAULT NULL,
	  "created_at" TIMESTAMP NULL DEFAULT NULL,
	  "updated_at" TIMESTAMP NULL DEFAULT NULL,
	  "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY ("id")
);
