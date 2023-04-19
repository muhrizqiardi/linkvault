-- migrate:up
create extension if not exists "uuid-ossp";

create table if not exists public.users (
    id uuid default uuid_generate_v4() primary key, 
    email text not null unique,
    full_name text not null,
    password text not null
);

create table if not exists public.folders (
    id uuid default uuid_generate_v4() primary key,
    name text not null unique,
    owner_id uuid not null,

    foreign key (owner_id)
        references public.users (id)
);

create table if not exists public.tags (
    id uuid default uuid_generate_v4() primary key,
    name text not null unique,
    owner_id uuid not null,

    foreign key (owner_id)
        references public.users (id)
);

create table if not exists public.links (
    id uuid default uuid_generate_v4() primary key,
    url text not null unique,
    excerpt text not null,
    cover_url text not null,
    owner_id uuid not null,
    folder_id uuid not null,

    foreign key (owner_id)
        references public.users (id),
    foreign key (folder_id)
        references public.folders (id)
);

create table if not exists public.link_medias (
    id uuid default uuid_generate_v4() primary key,
    link_id uuid not null,
    mediaUrl text not null,
    owner_id uuid not null,

    foreign key (owner_id)
        references public.users (id),
    foreign key (link_id)
        references public.links (id)
);

create table if not exists public.files (
    id uuid default uuid_generate_v4() primary key,
    link_id uuid not null,
    file_url text not null,
    owner_id uuid not null,

    foreign key (owner_id)
        references public.users (id),
    foreign key (link_id)
        references public.links (id)
);

-- migrate:down
drop table public.files;
drop table public.link_medias;
drop table public.links;
drop table public.folders;
drop table public.tags;
drop table public.users;
