ALTER TABLE users
    ADD CONSTRAINT UC_EmailAndPhoneNumber UNIQUE (email, phone_number);

ALTER TABLE books
    ADD CONSTRAINT UC_BookISBN UNIQUE (isbn);

ALTER TABLE users
    ADD CONSTRAINT UC_EMAIL UNIQUE (email);

ALTER TABLE books
    ADD CONSTRAINT UC_BookTitleAuthor UNIQUE (title, author);

ALTER TABLE visitations
    ADD CONSTRAINT vst_fk_user 
    FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE visitations
    ADD CONSTRAINT vst_fk_book
    FOREIGN KEY (book_id) REFERENCES books (id);

ALTER TABLE reservations
    ADD CONSTRAINT rst_fk_user
    FOREIGN KEY (user_id) REFERENCES users (id);
    
ALTER TABLE reservations
    ADD CONSTRAINT rst_fk_book
    FOREIGN KEY (book_id) REFERENCES books (id);
    