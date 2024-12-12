INSERT INTO Products (name, price)
VALUES
    ('Laptop', 50000.00),
    ('Mouse', 2500.50),
    ('Keyboard', 4500.00);


UPDATE Products
SET price = 111100.00
WHERE name = 'Laptop';


DELETE FROM Products
WHERE id = 2;

SELECT * FROM Products;


SELECT * FROM Products
WHERE price > 50000;