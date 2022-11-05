create table public.groups
(
    id                     bigint not null
        constraint groups_pk
            primary key,
    package_size           bigint default 15    not null,
    questions_types        bigint ARRAY default ARRAY[1, 2]::bigint[]
    not
    null,
    timer                  bigint default 60    not null,
    next_question_on_timer bool    default false not null,
    earliest_year          integer
);

comment on table public.groups is 'groups with settings';

comment on column public.groups.id is 'group id from telegram';

comment on column public.groups.package_size is 'size of the package; IIGI';

comment on column public.groups.questions_types is 'types of the questions from db.chgk.info database IIGI'; -- IIGI - ignore in generated insert

comment on column public.groups.earliest_year is 'earliest year for download from DB';

comment on column public.groups.timer is 'IIGI';
