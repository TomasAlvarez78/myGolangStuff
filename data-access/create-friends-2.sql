DROP TABLE IF EXISTS friendsUpgraded;
CREATE TABLE friendsUpgraded (
  id                INT AUTO_INCREMENT NOT NULL,
  personName        VARCHAR(128) NOT NULL,
  age               INT NOT NULL,
  career            VARCHAR(255) NOT NULL,
  birthdate         DATE NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO friendsUpgraded
  (personName, age, career, birthdate)
VALUES
  ('Tomas', 22 , 'Software Engineering','2000-06-10'),
  ('Emma', 19, 'Plants','2003-12-26'),
  ('Romanito', 22, 'Telecommunications Engineering','2000-08-14'),
  ('Facundito', 20 , 'Software Engineering','2002-10-13'),
  ('Chaqueno', 19 , 'Software Engineering','2002-07-17'),
  ('Benja', 20 , 'Software Engineering','2003-04-12'),
  ('Lucas', 22 , 'Software Engineering','2000-04-26');
  