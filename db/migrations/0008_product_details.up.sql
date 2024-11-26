create SEQUENCE seq_product_details;

create table product_details(
    id bigint not null primary key default nextval('seq_product_details'),
    product_id uuid not null,
    state numeric(2) not null default 1,
    color integer,
    weight numeric,
    capacity numeric,
    two_dimensional_height numeric,
    two_dimensional_width numeric,
    three_dimensional_height numeric,
    three_dimensional_width numeric,
    three_dimensional_thick numeric,
    amount bigint not null,
    price numeric not null,
    discount_price numeric,
    discount_percent numeric,
    created_at timestamp not null default current_timestamp,
    created_by uuid,
    updated_at timestamp,
    updated_by uuid
);