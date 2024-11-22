create SEQUENCE seq_products;
create SEQUENCE seq_files;
CREATE EXTENSION pg_trgm;

create table files(
    id bigint not null primary key default nextval('seq_files'),
    file_id uuid not null,
    name varchar(500),
    file_path varchar(256),
    size bigint,
    extension varchar,
    state numeric(2) not null default 1,
    created_at timestamp not null default current_timestamp,
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    is_main boolean not null default false,
    attach_id  uuid
);

comment on column files.attach_id is 'yuklangan rasm yoki video qaysi tablega (yani product yoki kategory) biriktirilganligini bildiradi';


create table products(
    id bigint not null primary key default nextval('seq_products'),
    product_id uuid not null UNIQUE,
    category_id uuid not null,
    sub_category_id uuid,
    created_at timestamp not null default current_timestamp,
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    state numeric(2) default 1,
    name_kr varchar(255),
    name_lt varchar(255) not null,
    name_ru varchar(255),
    short_info_kr varchar(500),
    short_info_lt varchar(500) not null,
    short_info_ru varchar(500),
    description_kr varchar(1000),
    description_lt varchar(1000) not null,
    description_ru varchar(1000),
    status varchar(20)
);

create table lists
(
    type_id                  integer not null,
    id                       bigint  not null,
    attributes               json,
    created_at               timestamp,
    created_by               bigint,
    date01                   timestamp,
    date02                   timestamp,
    date03                   timestamp,
    date04                   timestamp,
    date05                   timestamp,
    geo_polygon              json,
    int01                    integer,
    int02                    integer,
    int03                    integer,
    int04                    integer,
    int05                    integer,
    int06                    integer,
    int07                    integer,
    int08                    integer,
    int09                    integer,
    int10                    integer,
    key1                     varchar(255),
    key2                     varchar(255),
    key3                     varchar(255),
    long01                   varchar(2000),
    long02                   varchar(2000),
    long03                   varchar(2000),
    name1                    varchar(255),
    name2                    varchar(255),
    name3                    varchar(255),
    name4                    varchar(255),
    num                      integer,
    num01                    double precision,
    num02                    double precision,
    num03                    double precision,
    num04                    double precision,
    num05                    double precision,
    num06                    double precision,
    num07                    double precision,
    num08                    double precision,
    num09                    double precision,
    num10                    double precision,
    state                    integer default 1,
    tag                      integer,
    update_id                bigint,
    updated_at               timestamp,
    updated_by               bigint,
    val01                    varchar(1000),
    val02                    varchar(255),
    val03                    varchar(255),
    val04                    varchar(255),
    val05                    varchar(255),
    val06                    varchar(255),
    val07                    varchar(255),
    val08                    varchar(255),
    val09                    varchar(255),
    val10                    varchar(355),
    version                  integer,
    primary key (type_id, id)
);


create index lists_int01_type_id_index
    on lists (int01, type_id)
    where (state = 1);

create index lists_name1_idx
    on lists using gin (name1 gin_trgm_ops);

create index lists_type_id_idx
    on lists (type_id)
    where (state = 1);

create index lists_type_id_index
    on lists (type_id);

create index lists_type_id_name1_id_idx
    on lists (type_id) include (name1, id)
    where (state = 1);

create index lists_type_id_state_index
    on lists (type_id, state);

create index lists_id_idx
    on lists (id);

create index lists_state_type_id_idx
    on lists (state, type_id);


insert into lists(type_id, id, name1, state)
values 
(0, 0, 'Mundarija', 1),
(0, 5, 'Ranglar', 1);

INSERT INTO lists(type_id, id, name1, name2, name3, state)
VALUES
(5, 1, 'Oq', 'Оқ', 'белый', 1),
(5, 2, 'Qora', 'Қора', 'черный', 1),
(5, 3, 'Qizil', 'Қизил', 'красный', 1),
(5, 4, 'Yashil', 'Яшил', 'зеленый', 1),
(5, 5, 'Ko‘k', 'Кўк', 'синий', 1),
(5, 6, 'Sariq', 'Сариқ', 'желтый', 1),
(5, 7, 'Jigarrang', 'Жигарранг', 'коричневый', 1),
(5, 8, 'To‘q ko‘k', 'Тўқ кўк', 'темно-синий', 1),
(5, 9, 'Pushti', 'Пушти', 'розовый', 1),
(5, 10, 'Kulrang', 'Кулранг', 'серый', 1),
(5, 11, 'Moviy', 'Мовий', 'голубой', 1),
(5, 12, 'Olovrang', 'Оловранг', 'оранжевый', 1),
(5, 13, 'Binafsha', 'Бинафша', 'фиолетовый', 1),
(5, 14, 'Tillarang', 'Тилларанг', 'золотой', 1),
(5, 15, 'Kumushrang', 'Кумушранг', 'серебряный', 1);
