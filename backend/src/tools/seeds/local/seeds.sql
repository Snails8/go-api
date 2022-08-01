DELETE FROM public.users;
DELETE FROM public.fishes;

COPY public.users (id, name, email, created_at, updated_at) FROM stdin;
1	管理者	admin@test.com	2022-05-12	2022-05-12
2	テスト	test@test.com	2022-05-12	2022-05-12
\.

COPY public.fishes (id, name, type_id, created_at, updated_at) FROM stdin;
1	オイカワ	1	2022-08-01	2022-08-01
2	カワムツ	1	2022-08-01	2022-08-01
3	ドジョウ	1	2022-08-01	2022-08-01
4	シマドジョウ	1	2022-08-01	2022-08-01
5	オヤニラミ	1	2022-08-01	2022-08-01
6	カワアナゴ	1	2022-08-01	2022-08-01
7	ニホンウナギ	1	2022-08-01	2022-08-01
8	クロメダカ	1	2022-08-01	2022-08-01
9	タモロコ	1	2022-08-01	2022-08-01
100	真鯛	2	2022-08-01	2022-08-01
101	黒鯛	2	2022-08-01	2022-08-01
102	ネコザメ	2	2022-08-01	2022-08-01
103	イワシ	2	2022-08-01	2022-08-01
104	アジ	2	2022-08-01	2022-08-01
105	スズキ	2	2022-08-01	2022-08-01
\.

SELECT pg_catalog.setval('public.users_id_seq', (SELECT MAX(id) FROM public.users), true);