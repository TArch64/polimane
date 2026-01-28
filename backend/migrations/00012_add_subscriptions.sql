-- +goose Up
CREATE TYPE subscription_plan_id AS enum ('beta', 'basic', 'pro');
CREATE TYPE subscription_status AS enum ('active', 'canceled', 'unpaid');

CREATE TABLE IF NOT EXISTS user_subscriptions
(
  user_id          uuid                 NOT NULL,
  plan_id          subscription_plan_id NOT NULL,
  status           subscription_status  NOT NULL DEFAULT 'active',
  counters         jsonb                NOT NULL DEFAULT '{}'::jsonb,
  billing_try      smallint             NOT NULL DEFAULT 0,
  trial_started_at timestamptz          NOT NULL,
  trial_ends_at    timestamptz          NOT NULL,
  canceled_at      timestamptz,
  last_billed_at   timestamptz,
  PRIMARY KEY (user_id),

  CONSTRAINT fk_user_subscriptions_user
    FOREIGN KEY (user_id)
      REFERENCES users (id)
      ON UPDATE NO ACTION
      ON DELETE CASCADE
);

CREATE INDEX idx_user_subscriptions_status
  ON user_subscriptions (status);

INSERT INTO user_subscriptions (user_id, plan_id, trial_started_at, trial_ends_at, counters)
SELECT id,
       'beta',
       NOW(),
       NOW() + INTERVAL '14 days',
       JSON_BUILD_OBJECT(
         'schemasCreated',
         (SELECT COUNT(user_schemas.schema_id)
          FROM user_schemas
          WHERE user_id = users.id
            AND user_schemas.deleted_at IS NULL)
       )
FROM users
ON CONFLICT (user_id)
  DO UPDATE SET counters = excluded.counters;

-- +goose Down
DROP TABLE IF EXISTS user_subscriptions;
DROP TYPE IF EXISTS subscription_plan_id;
DROP TYPE IF EXISTS subscription_status;
