-- +goose Up
CREATE TYPE subscription_plan AS enum ('beta', 'basic', 'pro');
CREATE TYPE subscription_status AS enum ('active', 'canceled', 'unpaid');

CREATE TABLE IF NOT EXISTS user_subscriptions
(
  id               uuid                NOT NULL,
  created_at       timestamptz         NOT NULL,
  updated_at       timestamptz         NOT NULL,
  plan             subscription_plan   NOT NULL,
  status           subscription_status NOT NULL DEFAULT 'active',
  billing_try      smallint            NOT NULL DEFAULT 0,
  trial_started_at timestamptz         NOT NULL,
  trial_ends_at    timestamptz         NOT NULL,
  canceled_at      timestamptz,
  last_billed_at   timestamptz,

  CONSTRAINT fk_user_subscriptions_user
    FOREIGN KEY (id)
      REFERENCES users (id)
      ON UPDATE NO ACTION
      ON DELETE CASCADE
);

CREATE INDEX idx_user_subscriptions_status
  ON user_subscriptions (status);

INSERT INTO user_subscriptions (id, created_at, updated_at, plan, trial_started_at, trial_ends_at)
SELECT id, NOW(), NOW(), 'beta', NOW(), NOW() + INTERVAL '14 days'
FROM users
WHERE id NOT IN (SELECT id FROM user_subscriptions);

-- +goose Down
DROP TABLE IF EXISTS user_subscriptions;
DROP TYPE IF EXISTS subscription_plan;
DROP TYPE IF EXISTS subscription_status;
