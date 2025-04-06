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

ALTER TABLE borroweds
    DROP CONSTRAINT brw_fk_user;
    
ALTER TABLE borroweds
    DROP CONSTRAINT brw_fk_book;

ALTER TABLE book_reads
    DROP CONSTRAINT rd_fk_user;
    
ALTER TABLE book_reads
    DROP CONSTRAINT rd_fk_book;

ALTER TABLE book_reads
    DROP CONSTRAINT rd_fk_visitation;