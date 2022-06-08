CREATE TABLE IF NOT EXISTS todo_items (
    id                INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title             VARCHAR(255) NOT NULL,
    activity_group_id INT,
    is_active         BOOLEAN DEFAULT 1 NOT NULL,
    priority          VARCHAR(255) NOT NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at        TIMESTAMP,
    FOREIGN KEY (activity_group_id) REFERENCES activity_groups(id)
);
