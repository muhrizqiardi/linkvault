import * as z from 'zod';

export const userEntitySchema = z.object({
  id: z.string(),
  email: z.string().email(),
  full_name: z.string(),
  created_at: z.string().optional(),
  updated_at: z.string().optional(),
});
export interface UserEntity extends z.infer<typeof userEntitySchema> {}

export const folderEntitySchema = z.object({
  id: z.string(),
  name: z.string(),
  owner_id: z.string(),
  created_at: z.string().optional(),
  updated_at: z.string().optional(),
});
export interface FolderEntity extends z.infer<typeof folderEntitySchema> {}

export const tagEntitySchema = z.object({
  id: z.string(),
  name: z.string(),
  link_id: z.string(),
  owner_id: z.string(),
  created_at: z.string().optional(),
  updated_at: z.string().optional(),
});
export interface TagEntity extends z.infer<typeof tagEntitySchema> {}

export const linkEntitySchema = z.object({
  id: z.string(),
  url: z.string().url(),
  title: z.string(),
  excerpt: z.string(),
  cover_url: z.string().url().optional(),
  folder_id: z.string(),
  owner_id: z.string(),
  created_at: z.string().optional(),
  updated_at: z.string().optional(),
});
export interface LinkEntity extends z.infer<typeof linkEntitySchema> {}
