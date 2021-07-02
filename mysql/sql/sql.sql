CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(35) NOT NULL,
    username VARCHAR(15) NOT NULL UNIQUE,
    email VARCHAR(60) NOT NULL UNIQUE,
    password CHAR(60) NOT NULL,
    status ENUM('ENABLED', 'DISABLED') DEFAULT 'ENABLED',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
ENGINE = INNODB
DEFAULT CHARSET = UTF8;

CREATE TABLE IF NOT EXISTS posts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    author_id CHAR(36) NOT NULL,
    title VARCHAR(100) NOT NULL,
    body TEXT NOT NULL,
    likes INT DEFAULT 0,
    dislikes INT DEFAULT 0,
    favorites INT DEFAULT 0,
    deleted CHAR(1) DEFAULT '0',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT posts_author_fk FOREIGN KEY(author_id)
    REFERENCES users(id)
)
ENGINE = INNODB
DEFAULT CHARSET = UTF8;