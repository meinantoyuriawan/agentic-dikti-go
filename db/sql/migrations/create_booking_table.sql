CREATE TABLE counseling_bookings (
    bookingid      BIGSERIAL PRIMARY KEY,
    nama            VARCHAR(255) NOT NULL,
    nim             VARCHAR(50) NOT NULL,
    schedule     TIMESTAMP NOT NULL,
    universityid   INTEGER NOT NULL
);