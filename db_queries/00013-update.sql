UPDATE users SET last_name='MF' WHERE id=2;

UPDATE users
SET last_name = 'Khan',
    updated_at = CURRENT_TIMESTAMP
WHERE id = 2;
