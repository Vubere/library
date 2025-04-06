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
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE;

ALTER TABLE borroweds
    ADD CONSTRAINT brw_fk_user
    FOREIGN KEY (user_id) REFERENCES users (id);
    
ALTER TABLE borroweds
    ADD CONSTRAINT brw_fk_book
    FOREIGN KEY (book_id) REFERENCES books (id)
    ON DELETE CASCADE;

ALTER TABLE book_reads
    ADD CONSTRAINT rd_fk_user
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE;
    
ALTER TABLE book_reads
    ADD CONSTRAINT rd_fk_book
    FOREIGN KEY (book_id) REFERENCES books (id)
    ON DELETE CASCADE;

ALTER TABLE book_reads
    ADD CONSTRAINT rd_fk_visitation
    FOREIGN KEY (visitation_id) REFERENCES visitations (id) 
    ON DELETE CASCADE;

