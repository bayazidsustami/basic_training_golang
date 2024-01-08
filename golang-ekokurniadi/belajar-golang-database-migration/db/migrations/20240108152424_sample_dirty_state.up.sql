CREATE TABLE correct (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    PRIMARY KEY(id)
) ENGINE = InnoDB;

CREATE TABLE wrong (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255), #use varchir to simulate wrong query
    PRIMARY KEY(id)
) ENGINE = InnoDB;