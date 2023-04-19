-- migrate:up
create extension if not exists "uuid-ossp";

-- migrate:up
create table if not exists public."User" (
    id uuid default uuid_generate_v4() primary key, 
    email text not null unique,
    fullName text not null,
    password text not null
);

-- migrate:down
drop table public."User";

-- migrate:up
create table if not exists public."Folder" (
    id uuid default uuid_generate_v4() primary key,
    name text not null unique,
    ownerId uuid not null,

    foreign key (ownerId)
        references public."User" (id)
);

-- migrate:down 
drop table public."Folder";

-- migrate:up
create table if not exists public."Tag" (
    id uuid default uuid_generate_v4() primary key,
    name text not null unique,
    ownerId uuid not null,

    foreign key (ownerId)
        references public."User" (id)
);

-- migrate:down
drop table public."Tag";

-- migrate:up
create table if not exists public."Link" (
    id uuid default uuid_generate_v4() primary key,
    url text not null unique,
    excerpt text not null,
    coverUrl text not null,
    ownerId uuid not null,
    folderId uuid not null,

    foreign key (ownerId)
        references public."User" (id),
    foreign key (folderId)
        references public."Folder" (id)
);

-- migrate:down
drop table public."Link";

-- migrate:up
create table if not exists public."LinkMedia" (
    id uuid default uuid_generate_v4() primary key,
    linkId uuid not null,
    mediaUrl text not null,
    ownerId uuid not null,

    foreign key (ownerId)
        references public."User" (id),
    foreign key (linkId)
        references public."Link" (id)
);

-- migrate:down
drop table public."LinkMedia";

-- migrate:up
create table if not exists public."File" (
    id uuid default uuid_generate_v4() primary key,
    linkId uuid not null,
    fileUrl text not null,
    ownerId uuid not null,

    foreign key (ownerId)
        references public."User" (id),
    foreign key (linkId)
        references public."Link" (id)
);

-- migrate:down
drop table public."File";

