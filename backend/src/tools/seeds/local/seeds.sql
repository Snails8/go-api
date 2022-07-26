DELETE FROM public.users;

COPY public.users (id, name, email, created_at, updated_at) FROM stdin;
1	管理者	admin@test.com	2022-05-12	2022-05-12
2	テスト	test@test.com	2022-05-12	2022-05-12
\.

SELECT pg_catalog.setval('public.users_id_seq', (SELECT MAX(id) FROM public.users), true);