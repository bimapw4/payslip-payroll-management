-- +goose Up
-- +goose StatementBegin
-- Seeder untuk admin
INSERT INTO users (id, name, username, password, salary, is_admin, is_active, created_at, updated_at, created_by, updated_by)
VALUES
(gen_random_uuid(), 'Admin User', 'admin', '$2a$10$7P2Xq7l2Hu.VnBf0M6HguuMnyYmEwJXKIXlpiQuSRcku4J47oGhI2', 0, TRUE, TRUE, now(), now(), 'system', '');

DO $$
BEGIN
  FOR i IN 1..100 LOOP
    INSERT INTO users (id, name, username, password, salary, is_admin, is_active, created_at, updated_at, created_by, updated_by)
    VALUES (
      gen_random_uuid(),
      CONCAT('Employee ', i),
      CONCAT('employee', i),
      '$2a$10$7P2Xq7l2Hu.VnBf0M6HguuMnyYmEwJXKIXlpiQuSRcku4J47oGhI2',
      (random() * 5000000 + 3000000)::INT, -- gaji random antara 3jt - 8jt
      FALSE,
      TRUE,
      now(),
      now(),
      'system',
      ''
    );
  END LOOP;
END $$;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
