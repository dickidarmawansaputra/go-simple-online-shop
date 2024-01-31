CREATE TABLE transactions
(
    id INT AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    product_id INT NOT NULL,
    product_price INT NOT NULL,
    amount INT NOT NULL,
    sub_total INT NOT NULL,
    platform_fee INT NOT NULL DEFAULT 0,
    grand_total INT NOT NULL,
    status INT NOT NULL,
    product_snapshot JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB;