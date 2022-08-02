DELETE FROM public.users;
DELETE FROM public.fishes;

COPY public.users (id, name, email, created_at, updated_at) FROM stdin;
1	管理者	admin@test.com	2022-05-12	2022-05-12
2	テスト	test@test.com	2022-05-12	2022-05-12
\.

COPY public.fishes (id, name, classification, type_id, created_at, updated_at) FROM stdin;
1	オイカワ	コイ科	1	2022-08-01	2022-08-01
2	カワムツ	コイ科	1	2022-08-01	2022-08-01
3	ドジョウ	コイ科	1	2022-08-01	2022-08-01
4	シマドジョウ	コイ科	1	2022-08-01	2022-08-01
5	オヤニラミ	スズキ目スズキ亜目ケツギョ科	1	2022-08-01	2022-08-01
6	カワアナゴ	 カワアナゴ科カワアナゴ亜科カワアナゴ属	1	2022-08-01	2022-08-01
7	ニホンウナギ	ウナギ科 ウナギ属	1	2022-08-01	2022-08-01
8	クロメダカ	ダツ目メダカ科メダカ属	1	2022-08-01	2022-08-01
9	タモロコ	コイ科	1	2022-08-01	2022-08-01
100	真鯛	a	2	2022-08-01	2022-08-01
101	黒鯛	a	2	2022-08-01	2022-08-01
102	ネコザメ	a	2	2022-08-01	2022-08-01
103	イワシ	a	2	2022-08-01	2022-08-01
104	アジ	a	2	2022-08-01	2022-08-01
105	スズキ	a	2	2022-08-01	2022-08-01
\.

SELECT pg_catalog.setval('public.users_id_seq', (SELECT MAX(id) FROM public.users), true);