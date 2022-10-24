create table public.groups
(
    id                     bigint not null
        constraint groups_pk
            primary key,
    package_size           integer default 15    not null,
    questions_types        integer ARRAY default ARRAY[1, 2]::integer[]
    not
    null,
    timer                  integer default 60    not null,
    next_question_on_timer bool    default false not null,
    earliest_year          integer
);

comment on table public.groups is 'Groups with settings';

comment on column public.groups.id is 'Group id from telegram';

comment on column public.groups.package_size is 'Size of the package';

comment on column public.groups.questions_types is ' Types of the questions from db.chgk.info database';

comment on column public.groups.earliest_year is 'Earliest year for download from DB';
