-- Access the data from your HTTP Request software (Postman or Insomnia)
-- with this auth:
-- key: test 
-- token: password

INSERT INTO "administrators" ("id", "key", "token", "last_used") VALUES
(1, 'test', '$argon2id$v=19$m=65536,t=16,p=4$3a08c79fbf2222467a623df9a9ebf75802c65a4f9be36eb1df2f5d2052d53cb7$ce434bd38f7ba1fc1f2eb773afb8a1f7f2dad49140803ac6cb9d7256ce9826fb3b4afa1e2488da511c852fc6c33a76d5657eba6298a8e49d617b9972645b7106', '');

-- 10 jokes is enough right?

INSERT INTO "jokesbapak2" ("link", "creator") VALUES
('https://picsum.photos/id/1000/500/500', 1),
('https://picsum.photos/id/1001/500/500', 1),
('https://picsum.photos/id/1002/500/500', 1),
('https://picsum.photos/id/1003/500/500', 1),
('https://picsum.photos/id/1004/500/500', 1),
('https://picsum.photos/id/1005/500/500', 1),
('https://picsum.photos/id/1006/500/500', 1),
('https://picsum.photos/id/1010/500/500', 1),
('https://picsum.photos/id/1008/500/500', 1),
('https://picsum.photos/id/1009/500/500', 1);