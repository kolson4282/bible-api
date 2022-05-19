CREATE TABLE characters (
            id INT NOT NULL AUTO_INCREMENT,
            name VARCHAR(255) NOT NULL,
            description VARCHAR(255) NOT NULL,
            gender VARCHAR(8) NULL,
            PRIMARY KEY (id)
);

LOAD DATA INFILE  '/var/lib/mysql-files/characters.csv' into table characters
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\r' 
IGNORE 1 LINES
(@val, name, description, gender);