CREATE TABLE transactions
(
    id                SERIAL PRIMARY KEY,
    booking_id        INTEGER,
    payment_method_id INTEGER,
    virtual_rek       INTEGER,
    total_price       INTEGER NOT NULL,
    status            status_transaction,
    qr_code           VARCHAR(255),
    CONSTRAINT fk_transactions_booking FOREIGN KEY (booking_id) REFERENCES bookings (id),
    CONSTRAINT fk_transactions_payment FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id)
);