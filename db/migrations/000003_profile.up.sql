CREATE TABLE profile (
    id_profile INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    avatar VARCHAR(255),
    bio TEXT,
    user_id INT UNIQUE NOT NULL,
    CONSTRAINT fk_user_profile FOREIGN KEY (user_id) REFERENCES users(id_users) ON DELETE CASCADE
);