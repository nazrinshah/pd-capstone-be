CREATE DATABASE capstone;

CREATE TABLE dish_tab (
                          id MEDIUMINT NOT NULL AUTO_INCREMENT,
                          vendorid MEDIUMINT NOT NULL,
                          name CHAR(30) NOT NULL,
                          dishstatus MEDIUMINT NOT NULL,
                          price FLOAT(7,2) NOT NULL,
     description CHAR(255) NOT NULL,
     currency CHAR(255) NOT NULL,
     imagename CHAR(255) NOT NULL,
     PRIMARY KEY (id)
);

CREATE TABLE order_tab (
                           id MEDIUMINT NOT NULL AUTO_INCREMENT,
                           orderstatus MEDIUMINT NOT NULL,
                           items VARCHAR(1024) NOT NULL,
                           subtotal FLOAT(7,2) NOT NULL,
     platformfee FLOAT(7,2) NOT NULL,
     deliveryfee FLOAT(7,2) NOT NULL,
     PRIMARY KEY (id)
);

INSERT INTO dish_tab (vendorid, name, dishstatus, price, description, currency, imagename) VALUES (2, "Pearl Milk Tea", 1, 4.3, "Delicious boba with milk tea", "SGD", "fp-drink-gong-cha-pearl-milk-tea");
INSERT INTO dish_tab (vendorid, name, dishstatus, price, description, currency, imagename) VALUES (2, "Brown Sugar Milk Tea", 1, 7.3, "Healthier choice", "SGD", "fp-drink-gong-cha-brown-sugar-mlik-tea-with-pearl");
INSERT INTO dish_tab (vendorid, name, dishstatus, price, description, currency, imagename) VALUES (2, "Strawberry Milk Tea", 1, 6.3, "Fruity", "SGD", "fp-drink-gong-cha-strawberry-milk-tea");