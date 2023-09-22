CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
    -- Add more columns as needed
);

INSERT INTO users ( email, password, phone)
VALUES
    ('user1@example.com', 'password', '123-456-7890'),
    ('user2@example.com', 'password', '987-654-3210'),
    ('user3@example.com', 'password', '555-555-5555');

-- Create the Order Table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    users_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT NOW(),
    status VARCHAR(255) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    shipping_address VARCHAR(255) NOT NULL,
    payment_method VARCHAR(255) NOT NULL
);

-- Insert dummy data into the Order Table
INSERT INTO orders (users_id, status, total_price, shipping_address, payment_method)
VALUES
    (1, 'Pending', 59.99, '123 Main St, City, Country', 'Credit Card'),
    (2, 'Shipped', 79.99, '456 Elm St, City, Country', 'PayPal'),
    (3, 'Delivered', 129.99, '789 Oak St, City, Country', 'Credit Card');

-- Create the OrderItem Table
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    orders_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL
);

-- Insert dummy data into the OrderItem Table
INSERT INTO order_items (orders_id, product_id, quantity, unit_price, subtotal)
VALUES
    (1, 1, 2, 29.99, 59.98),
    (1, 2, 1, 19.99, 19.99),
    (2, 3, 3, 9.99, 29.97),
    (3, 1, 5, 29.99, 149.95);


-- Create the products table
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    category VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    theme VARCHAR(255) NOT NULL
);

-- Insert dummy data into the products table
INSERT INTO product (category, name, description, price, theme)
VALUES
    ('Electronics', 'Smartphone', 'A high-end smartphone with advanced features', 599.99, 'Premium'),
    ('Clothing', 'T-shirt', 'Comfortable cotton t-shirt in various colors', 19.99, 'Casual'),
    ('Books', 'Programming Book', 'Learn programming with this bestseller', 39.99, 'Educational'),
    ('Electronics', 'Laptop', 'Powerful laptop for work and gaming', 999.99, 'Premium'),
    ('Home & Garden', 'Garden Tools Set', 'Essential tools for gardening', 49.99, 'Outdoor'),
    ('Toys', 'Board Game', 'Fun board game for the whole family', 29.99, 'Entertainment');

-- Create the Warehouses table with an "id" prefixed primary key
CREATE TABLE warehouse (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    is_active boolean DEFAULT true
);

-- Insert dummy data into the Warehouses table
INSERT INTO warehouse (name, location)
VALUES
    ('Warehouse A', 'Location A'),
    ('Warehouse B', 'Location B');

-- Create the ProductStock table with an "id" prefixed primary key
CREATE TABLE product_stock (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    warehouse_id INT NOT NULL,
    stock_quantity INT NOT NULL
);

-- Insert dummy data into the ProductStock table to represent initial stock in warehouses
INSERT INTO product_stock (product_id, warehouse_id, stock_quantity)
VALUES
    (1, 1, 50), -- 50 laptops in Warehouse A
    (1, 2, 50), -- 50 laptops in Warehouse A
    (2, 1, 100), -- 100 chairs in Warehouse A
    (3, 2, 75), -- 75 drills in Warehouse B
    (4, 2, 200); -- 200 t-shirts in Warehouse B;

-- Create the ProductTransfer table with an "id" prefixed primary key
CREATE TABLE product_transfer (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    from_warehouse_id INT NOT NULL,
    to_warehouse_id INT NOT NULL,
    quantity INT NOT NULL,
    transfer_date TIMESTAMP DEFAULT NOW()
);

-- Insert dummy data into the ProductTransfer table to represent product movement between warehouses
INSERT INTO product_transfer (product_id, from_warehouse_id, to_warehouse_id, quantity)
VALUES
    (1, 1, 2, 10), -- Transfer 10 laptops from Warehouse A to Warehouse B
    (3, 2, 1, 5); -- Transfer 5 drills from Warehouse B to Warehouse A;


-- Create the cart table with hold_duration column
CREATE TABLE cart (
    id SERIAL PRIMARY KEY,
    users_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    hold_duration INTERVAL DEFAULT INTERVAL '30 minutes', -- Default hold duration is 24 hours
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_paid bool DEFAULT FALSE,
    released_stock bool DEFAULT FALSE,
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);

-- Insert dummy data into the cart table
INSERT INTO cart (users_id, product_id, quantity)
VALUES
    (1, 1, 2), -- User 1 adds 2 units of Product 1 to their cart
    (1, 2, 3), -- User 1 adds 3 units of Product 2 to their cart
    (2, 3, 1), -- User 2 adds 1 unit of Product 3 to their cart
    (3, 1, 1); -- User 3 adds 1 unit of Product 1 to their cart

-- -- Release stocks that have exceeded the hold_duration
-- DELETE FROM cart
-- WHERE updated_at < NOW() - hold_duration;