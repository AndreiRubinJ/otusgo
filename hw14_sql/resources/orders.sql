INSERT INTO Orders (user_id, order_date, total_amount)
VALUES (1, '2023-12-10', 1245.00);


INSERT INTO OrderProducts (order_id, product_id, quantity)
VALUES
    (1, 1, 1), -- 1x Laptop
    (1, 3, 1); -- 1x Keyboard

DELETE FROM Orders
WHERE id = 1;

SELECT o.id AS order_id, o.order_date, o.total_amount
FROM Orders o
         JOIN Users u ON o.user_id = u.id
WHERE u.email = 'nplinda@example.com';