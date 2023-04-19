-- migrate:up
create extension if not exists "uuid-ossp";

create table if not exists public."User" (
    id uuid default uuid_generate_v4() primary key, 
    email text not null unique,
    fullName text not null,
    password text not null
);

create table if not exists public."Folder" (
    id uuid default uuid_generate_v4() primary key,
    name text not null unique,
    ownerId uuid not null,

    foreign key (ownerId)
        references public."User" (id)
);

create table if not exists public."Tag" (
    id uuid default uuid_generate_v4() primary key,
    name text not null unique,
    ownerId uuid not null,

    foreign key (ownerId)
        references public."User" (id)
);

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
drop table public."LinkMedia";
drop table public."Link";
drop table public."Folder";
drop table public."Tag";
drop table public."User";
