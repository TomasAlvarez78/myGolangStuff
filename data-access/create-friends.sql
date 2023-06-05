DROP TABLE IF EXISTS friends;
CREATE TABLE friends (
  id                INT AUTO_INCREMENT NOT NULL,
  personName        VARCHAR(128) NOT NULL,
  age               INT NOT NULL,
  career            VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO friends
  (personName, age, career)
VALUES
  ('Tomas', 22 , 'Software Engineering'),
  ('Emma', 19, 'Plants'),
  ('Romanito', 22, 'Telecommunications Engineering'),
  ('Facundito', 20 , 'Software Engineering'),
  ('Chaqueno', 19 , 'Software Engineering'),
  ('Benja', 20 , 'Software Engineering');