INSERT INTO authors (name) VALUES ('JK. Rowling');
SELECT * FROM authors;

INSERT INTO books (title, author_id) VALUES ("Harry Potter and the Deathly Hallows", 1);
INSERT INTO author_to_books (book_id, author_id) VALUES (7,1);

SELECT
	b.*
FROM
	books AS b
	INNER JOIN author_books AS ab ON ab.book_id = b.id
WHERE
	ab.author_id = 1;

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL,
    surname VARCHAR(200) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
);

INSERT INTO users (name, surname) VALUES ('Amilcar', 'Cattaneo');
INSERT INTO users (name, surname) VALUES ('Lucas', 'Bley');

CREATE TABLE user_books (
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);

INSERT INTO user_books (user_id, book_id) VALUES (1, 7);

CREATE TABLE user_loans (
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    due_date DATETIME NOT NULL,
    user_id_lender INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    PRIMARY KEY (user_id, book_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (user_id_lender) REFERENCES users(id)
);

INSERT INTO user_loans (user_id, book_id, user_id_lender, due_date) VALUES (2, 7, 1, '2020-09-19');
SELECT
	ub.name AS Borrower,
	b.title AS Book,
	ul.name AS Lender,
	al.due_date AS DueDate
FROM
	user_loans AS al
	INNER JOIN books AS b ON al.book_id = b.id
	INNER JOIN users AS ul ON al.user_id_lender = ul.id
	INNER JOIN users AS ub ON al.user_id = ub.id
WHERE
	al.user_id = 2
	AND al.book_id = 7;

SELECT
    b.*
FROM
    user_books AS ub
    INNER JOIN books AS b ON ub.book_id = b.id
WHERE
    ub.user_id = 1;

UPDATE user_books SET available = TRUE;
DELETE FROM user_loans;

SELECT
    *
FROM
    user_books
WHERE
    user_id = 1;
SELECT * FROM user_loans;

SELECT
	a.*
FROM
	authors AS a
	INNER JOIN author_books AS ab ON ab.author_id = a.id
WHERE
	ab.book_id = 2;