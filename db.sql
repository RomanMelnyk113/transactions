--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.accounts (
    id text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    category text
);


ALTER TABLE public.accounts OWNER TO admin;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.transactions (
    id text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    account_id text,
    made_on timestamp with time zone,
    amount numeric,
    currency_code text,
    description text,
    details jsonb,
    "order" bigint,
    is_parent boolean,
    parent_id character varying(64)
);


ALTER TABLE public.transactions OWNER TO admin;

--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.accounts (id, created_at, updated_at, deleted_at, name, category) FROM stdin;
314706700999329909	\N	\N	\N	revolut	banks
314707590107891965	\N	\N	\N	paypal_gbp	payment_methods
314707589973674235	\N	\N	\N	paypal_eur	payment_methods
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.transactions (id, created_at, updated_at, deleted_at, account_id, made_on, amount, currency_code, description, details, "order", is_parent, parent_id) FROM stdin;
314707597087213849	\N	\N	\N	314707590107891965	2020-09-17 00:00:00+00	-1.86	GBP	To Euro	\N	1	f	314706840032118839
314707597087213848	\N	\N	\N	314707590107891965	2020-09-17 00:00:00+00	1.86	GBP	Credit Card	\N	0	f	314706840032118839
314706840032118839	\N	\N	\N	314706700999329909	2020-09-17 00:00:00+00	-2.1	EUR	Paypal *ny The New York Times Co.	\N	4	t	\N
314707596835555593	\N	\N	\N	314707589973674235	2020-09-17 00:00:00+00	-2.0	EUR	The New York Times Co.	\N	3	f	314706840032118839
314707596835555592	\N	\N	\N	314707589973674235	2020-09-17 00:00:00+00	2.0	EUR	From British Pound	\N	2	f	314706840032118839
\.


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: idx_accounts_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_accounts_deleted_at ON public.accounts USING btree (deleted_at);


--
-- Name: idx_transactions_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_transactions_deleted_at ON public.transactions USING btree (deleted_at);


--
-- Name: transactions fk_transactions_account; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_account FOREIGN KEY (account_id) REFERENCES public.accounts(id);


--
-- Name: transactions fk_transactions_parent; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_parent FOREIGN KEY (parent_id) REFERENCES public.transactions(id);


--
-- PostgreSQL database dump complete
--

