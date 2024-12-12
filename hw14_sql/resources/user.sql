INSERT INTO Users (name, email, password)
VALUES
    ('Alice', 'arubin@example.com', 'password123'),
    ('Bob', 'alexei@example.com', 'password456');

UPDATE Users
SET name = 'Nikolai Plinda', email = 'nplinda@example.com'
WHERE id = 1;

DELETE FROM Users
WHERE id = 2;

SELECT
    u.name,
    COUNT(DISTINCT o.id) AS total_orders,
    SUM(o.total_amount) AS total_spent,
    AVG(p.price) AS avg_product_price
FROM Users u
         LEFT JOIN Orders o ON u.id = o.user_id
         LEFT JOIN OrderProducts op ON o.id = op.order_id
         LEFT JOIN Products p ON op.product_id = p.id
WHERE u.id = 1
GROUP BY u.name;