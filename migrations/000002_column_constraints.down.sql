ALTER TABLE users
    DROP CONSTRAINT UC_EmailAndPhoneNumber;

ALTER TABLE books
    DROP CONSTRAINT UC_BookISBN;

ALTER TABLE books
    DROP CONSTRAINT UC_BookTitleAuthor;

ALTER TABLE visitations
    DROP CONSTRAINT vst_fk_user;

ALTER TABLE visitations
    DROP CONSTRAINT vst_fk_book;

ALTER TABLE reservations
    DROP CONSTRAINT rst_fk_user;
    
ALTER TABLE reservations
    DROP CONSTRAINT rst_fk_book;