-- ADMINISTRATORS password: admin123

INSERT INTO administrators (first_name, last_name, login, password, email, phone, create_date, modify_date, last_login) VALUES ('a', 'b', 'c', '240be518fabd2724ddb6f04eeb1da5967448d7e831c08c8fa822809f74c720a9', 'a@b.c', 123456789, '10-10-2025', '10-10-2025', '10-10-2025');

-- NAVIGATION

INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('Home', '/shop/main', 10, true, '10-10-2025', '10-10-2025', null, 1);
    
INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('Products', '/search/products', 20, true, '10-10-2025', '10-10-2025', null, 1);

INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('Search', '/search/products', 30, true, '10-10-2025', '10-10-2025', null, 1);

INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('Product', '/products/1', 40, true, '10-10-2025', '10-10-2025', 3, 1);

INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('Checkout', '/checkout', 50, true, '10-10-2025', '10-10-2025', null, 1);

INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('Contact', '/page/contact', 60, true, '10-10-2025', '10-10-2025', null, 1);

INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('Cart', '/cart', 70, true, '10-10-2025', '10-10-2025', null, 1);

INSERT INTO navigation (name, url, position, status, create_date, modify_date, navigation_navigation_id, administrators_administrator_id)
    VALUES ('About', '/page/about', 80, true, '10-10-2025', '10-10-2025', null, 1);


-- PAGES

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('main', 'shop', '', '10-10-2025', '10-10-2025', true, 1);

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('cart', 'cart', '', '10-10-2025', '10-10-2025', true, 1);

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('checkout', 'checkout', '', '10-10-2025', '10-10-2025', true, 1);

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('products', 'products', '{"Availability":true,"ShowType":true,"AllowQuantity":true,"ShowRelated":false,"ShowReviews":false}', '10-10-2025', '10-10-2025', true, 1);

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('contact', 'page', '', '10-10-2025', '10-10-2025', true, 1);

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('about', 'page', '', '10-10-2025', '10-10-2025', true, 1);

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('search', 'search', '', '10-10-2025', '10-10-2025', true, 1);

INSERT INTO pages (name, type, content, create_date, modify_date, status, administrators_administrator_id)
    VALUES ('profile', 'profile', '', '10-10-2025', '10-10-2025', true, 1);


-- PAYMENT DICTIONARY

INSERT INTO payment_dictionary (name, price) VALUES ('Blik', 5.50);
INSERT INTO payment_dictionary (name, price) VALUES ('Visa', 1.00);
INSERT INTO payment_dictionary (name, price) VALUES ('Personal', 0.00);


-- PRODUCT CATEGORY DICTIONARY

INSERT INTO product_category_dictionary (name) VALUES ('Żabeczki');
INSERT INTO product_category_dictionary (name) VALUES ('Terrarium');
INSERT INTO product_category_dictionary (name) VALUES ('Kotki');
INSERT INTO product_category_dictionary (name) VALUES ('Salamandry');

-- VAT PERCENTAGE DICTIONARY

INSERT INTO vat_percentage_dictionary (percentage) VALUES (12);
INSERT INTO vat_percentage_dictionary (percentage) VALUES (27);

-- DELIVERY DICTIONARY

INSERT INTO delivery_dictionary (name, price) VALUES ('Personal', 0.00);
INSERT INTO delivery_dictionary (name, price) VALUES ('DHL', 14.89);
INSERT INTO delivery_dictionary (name, price) VALUES ('Inpost', 19.99);

-- PRODUCTS (EXAMPLE)

INSERT INTO products (name, create_date, is_animal, price, availability, vat_percentage, product_code, description, product_category, product_metadata_product_metadata_id, modify_date)
    VALUES ('Żaba', '10-10-2025', true, 49.99, true, 12, 'ZB123', 'Little distinguish man', 'Żabeczki', null, '10-10-2025');

INSERT INTO products (name, create_date, is_animal, price, availability, vat_percentage, product_code, description, product_category, product_metadata_product_metadata_id, modify_date)
    VALUES ('Terrarium', '10-10-2025', false, 199.43, true, 12, 'TR123', 'Cool thingy for your froggies', 'Terrarium', null, '10-10-2025');

INSERT INTO products (name, create_date, is_animal, price, availability, vat_percentage, product_code, description, product_category, product_metadata_product_metadata_id, modify_date)
    VALUES ('A cat', '10-10-2025', true, 999, true, 27, 'CT123', 'Little distinguish man as well', 'Kotki', null, '10-10-2025');

INSERT INTO products (name, create_date, is_animal, price, availability, vat_percentage, product_code, description, product_category, product_metadata_product_metadata_id, modify_date)
    VALUES ('Salamandra bezplamista', '10-10-2025', true, 21.37, true, 12, 'ZB123', 'He tries his best', 'Salamandry', null, '10-10-2025');


-- MEDIAS

INSERT INTO medias (name, path, create_date, modify_date) VALUES ('frog.jpg', '../media/products/main/frog.jpg', '10-10-2025', '10-10-2025');
INSERT INTO medias (name, path, create_date, modify_date) VALUES ('frog.jpg', '../media/products/icons/frog.jpg', '10-10-2025', '10-10-2025');
INSERT INTO medias (name, path, create_date, modify_date) VALUES ('terrarium.jpg', '../media/products/main/terrarium.jpg', '10-10-2025', '10-10-2025');
INSERT INTO medias (name, path, create_date, modify_date) VALUES ('terrarium.jpg', '../media/products/icons/terrarium.jpg', '10-10-2025', '10-10-2025');
INSERT INTO medias (name, path, create_date, modify_date) VALUES ('cat.jpg', '../media/products/main/cat.jpg', '10-10-2025', '10-10-2025');
INSERT INTO medias (name, path, create_date, modify_date) VALUES ('cat.png', '../media/products/icons/cat.png', '10-10-2025', '10-10-2025');
INSERT INTO medias (name, path, create_date, modify_date) VALUES ('salamandra.jpg', '../media/products/main/salamandra.jpg', '10-10-2025', '10-10-2025');
INSERT INTO medias (name, path, create_date, modify_date) VALUES ('salamandra.jpg', '../media/products/icons/salamandra.jpg', '10-10-2025', '10-10-2025');


-- PRODUCT_MEDIAS

INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (1, 1);
INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (1, 2);
INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (2, 3);
INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (2, 4);
INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (3, 5);
INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (3, 6);
INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (4, 7);
INSERT INTO products_medias (products_product_id, medias_media_id) VALUES (4, 8);


-- CLIENTS

INSERT INTO clients (first_name, last_name, login, password, email, phone, status, newsletter, last_login, join_date, birthdate, create_date, modify_date)
    VALUES ('Barry', 'Benson', 'Bee', 'a@b.c', '65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5', 'a@b.c', null, false, false, '10-10-2025', '10-10-2025', null, '10-10-2025', '10-10-2025')
