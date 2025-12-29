-- sqlite3 app.db -
PRAGMA foreign_keys = ON;

CREATE TABLE item_ (
	item_name_ TEXT PRIMARY KEY,
	unit_ TEXT NOT NULL,
	UNIQUE (item_name_)
);
CREATE TABLE recipe_ (
	recipe_name_ TEXT PRIMARY KEY,
	UNIQUE (recipe_name_)
);
CREATE TABLE recipe_item_ (
	recipe_name_ TEXT NOT NULL REFERENCES recipe_ ON DELETE CASCADE ON UPDATE CASCADE,
	item_name_ TEXT NOT NULL REFERENCES item_ ON DELETE CASCADE ON UPDATE CASCADE,
	item_type_ TEXT NOT NULL,
	quantity_ INTEGER NOT NULL,
	UNIQUE (recipe_name_, item_name_, item_type_)
);

INSERT INTO item_ (item_name_, unit_) VALUES
	('crude oil', 'mb'),
	('light oil', 'mb'),
	('heavy oil', 'mb'),
	('petroleum', 'mb'),
	('water', 'mb'),
	('steam', 'mb'),
	('coal', 'pcs'),
	('plastic', 'pcs');
INSERT INTO recipe_ (recipe_name_) VALUES
	('oil cracking'),
	('coal liquefaction'),
	('plastic forming');
INSERT INTO recipe_item_ (recipe_name_, item_name_, item_type_, quantity_) VALUES
	('oil cracking', 'crude oil', 'ingredient', 100),
	('oil cracking', 'water', 'ingredient', 50),
	('oil cracking', 'petroleum', 'product', 55),
	('oil cracking', 'light oil', 'product', 45),
	('oil cracking', 'heavy oil', 'product', 25),
	('coal liquefaction', 'coal', 'ingredient', 10),
	('coal liquefaction', 'steam', 'ingredient', 50),
	('coal liquefaction', 'heavy oil', 'ingredient', 25),
	('coal liquefaction', 'petroleum', 'product', 10),
	('coal liquefaction', 'light oil', 'product', 20),
	('coal liquefaction', 'heavy oil', 'product', 90),
	('plastic forming', 'petroleum', 'ingredient', 20),
	('plastic forming', 'coal', 'ingredient', 1),
	('plastic forming', 'plastic', 'product', 2);

UPDATE item_
SET item_name_ = 'plastic bar'
WHERE item_name_ = 'plastic';

SELECT * FROM recipe_item_;

SELECT recipe_name_
FROM recipe_item_
WHERE TRUE
AND item_type_ = 'product'
AND item_name_ = 'petroleum';

SELECT * FROM recipe_
INNER JOIN recipe_item_ USING (recipe_name_)
WHERE recipe_name_ = 'oil cracking';

