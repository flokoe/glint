INSERT INTO users (email, name) VALUES ('zaphod@pangalactic.gov', 'Zaphod');

INSERT INTO providers (type, url) VALUES ('openai-compatible', 'http://127.0.0.1:8080');

INSERT INTO llms (string, name, provider_id, context_size) VALUES ('gemma3/gemma3:1b', 'Gemma 3 1B', 1, 2048);
