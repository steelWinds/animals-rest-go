CREATE TABLE own_maps(
	  "id" SERIAL,
	  "animal_unit_id" INTEGER NULL DEFAULT NULL UNIQUE,
	  "owner_unit_id" INTEGER NULL DEFAULT NULL,
	  "created_at" TIMESTAMP NULL DEFAULT NULL,
	  "updated_at" TIMESTAMP NULL DEFAULT NULL,
	  "deleted_at" TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY ("id"),
	  CONSTRAINT "FK_own_maps_animal_units" FOREIGN KEY ("animal_unit_id") REFERENCES "animal_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
	  CONSTRAINT "FK_own_maps_owner_units" FOREIGN KEY ("owner_unit_id") REFERENCES "owner_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
