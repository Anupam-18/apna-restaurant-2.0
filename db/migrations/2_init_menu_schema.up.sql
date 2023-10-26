CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE menus (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "category" VARCHAR NOT NULL,
    "menu_item_ids" UUID[],
    CONSTRAINT "unique_menu_category" UNIQUE ("category")
);

CREATE TABLE menuitems (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR NOT NULL,
    "price" INT NOT NULL,
    "image_url" VARCHAR NOT NULL,
    "menu_id" UUID NOT NULL,
    FOREIGN KEY ("menu_id") REFERENCES menus ("id"),
    CONSTRAINT "unique_menuitem_name" UNIQUE ("name")
);
